package Repositories

import (
	ps "blogapp/Infrastructure/password_services"

	"blogapp/Domain"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-playground/validator"
)

type userRepository struct {
	validator       *validator.Validate
	collection      Domain.Collection
	TokenRepository Domain.RefreshRepository
}

func NewUserRepository(_collection Domain.Collection, token_collection Domain.Collection) *userRepository {
	return &userRepository{
		validator:       validator.New(),
		collection:      _collection,
		TokenRepository: NewRefreshRepository(token_collection),
	}

}

// create user
func (as *userRepository) CreateUser(ctx context.Context, user *Domain.User) (Domain.OmitedUser, error, int) {

	// Check if user email is taken
	existingUserFilter := bson.D{{"email", user.Email}}
	existingUserCount, err := as.collection.CountDocuments(ctx, existingUserFilter)
	if err != nil {
		return Domain.OmitedUser{}, err, 500
	}
	if existingUserCount > 0 {
		return Domain.OmitedUser{}, errors.New("Email is already taken"), http.StatusBadRequest
	}
	// User registration logic
	hashedPassword, err := ps.GenerateFromPasswordCustom(user.Password)
	if err != nil {
		return Domain.OmitedUser{}, err, 500
	}
	user.Password = string(hashedPassword)
	insertResult, err := as.collection.InsertOne(ctx, user)
	if err != nil {
		return Domain.OmitedUser{}, err, 500
	}
	// Fetch the inserted task
	var fetched Domain.OmitedUser
	err = as.collection.FindOne(context.TODO(), bson.D{{"_id", insertResult.InsertedID.(primitive.ObjectID)}}).Decode(&fetched)
	if err != nil {
		fmt.Println(err)
		return Domain.OmitedUser{}, errors.New("User Not Created"), 500
	}
	if fetched.Email != user.Email {
		return Domain.OmitedUser{}, errors.New("User Not Created"), 500
	}
	fetched.Password = ""
	return fetched, nil, 200
}

// get all users
func (us *userRepository) GetUsers(ctx context.Context) ([]*Domain.OmitedUser, error, int) {
	// us.mu.RLock()
	// defer us.mu.RUnlock()

	// Pass these options to the Find method
	findOptions := options.Find()
	// findOptions.SetLimit(2)
	filter := bson.D{{}}

	// Here's an array in which you can store the decoded documents
	var results []*Domain.OmitedUser

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := us.collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal("error in finding users", err)
		log.Fatal(err)
		return []*Domain.OmitedUser{}, err, 0
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var elem Domain.OmitedUser
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("error in decoding user", err)
			fmt.Println(err.Error())
			// #handelthislater
			// should this say there was a decoding error and return?
			return []*Domain.OmitedUser{}, err, 500
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return []*Domain.OmitedUser{}, err, 500
	}

	// Close the cursor once finished
	cur.Close(ctx)

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results, nil, 200
}

// get user by id
func (us *userRepository) GetUsersById(ctx context.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {

	var filter bson.D
	filter = bson.D{{"_id", id}}
	var result Domain.OmitedUser
	err := us.collection.FindOne(ctx, filter).Decode(&result)
	// # handel this later
	if err != nil {
		return Domain.OmitedUser{}, errors.New("User not found"), http.StatusNotFound
	}
	if current_user.Role == "user" && result.ID != current_user.ID {
		return Domain.OmitedUser{}, errors.New("permission denied"), http.StatusForbidden

	}
	return result, nil, 200
}

// update user by id
func (us *userRepository) UpdateUsersById(ctx context.Context, id primitive.ObjectID, user Domain.User, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	if current_user.ID != id {
		return Domain.OmitedUser{}, errors.New("permission denied"), http.StatusForbidden
	}
	var NewUser Domain.OmitedUser
	statusCode := 200

	// Retrieve the existing user
	NewUser, err, statusCode := us.GetUsersById(ctx, id, current_user)
	if err != nil {
		return Domain.OmitedUser{}, err, 500
	}

	// Update only the specified fields
	if user.Email != "" {
		NewUser.Email = user.Email
	}
	if user.UserName != "" {
		NewUser.UserName = user.UserName
	}
	if user.Email != "" {
		NewUser.Email = user.Email
	}
	if user.Role != "" {
		NewUser.Role = user.Role
	}
	if user.ProfilePicture != "" {
		NewUser.ProfilePicture = user.ProfilePicture
	}
	if user.Bio != "" {
		NewUser.Bio = user.Bio
	}
	if !user.CreatedAt.IsZero() {
		NewUser.CreatedAt = user.CreatedAt
	}
	if !user.UpdatedAt.IsZero() {
		NewUser.UpdatedAt = user.UpdatedAt
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"email", NewUser.Email},
			{"userName", NewUser.UserName},
			{"role", NewUser.Role},
			{"profile_picture", NewUser.ProfilePicture},
			{"bio", NewUser.Bio},
			{"created_at", NewUser.CreatedAt},
			{"updated_at", NewUser.UpdatedAt},
		}},
	}

	updateResult, err := us.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		statusCode = 500
		return Domain.OmitedUser{}, err, statusCode
	}
	if updateResult.ModifiedCount == 0 {
		statusCode = 400
		return Domain.OmitedUser{}, errors.New("user does not exist"), statusCode
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return NewUser, nil, statusCode
}

// delete user by id
func (us *userRepository) DeleteUsersById(ctx context.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (error, int) {

	filter := bson.D{{"_id", id}}
	if current_user.Role == "user" && current_user.ID != id {
		return errors.New("permission denied"), http.StatusForbidden
	}

	deleteResult, err := us.collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return err, 500
	}
	if deleteResult.DeletedCount == 0 {
		return errors.New("User does not exist"), http.StatusNotFound
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	if current_user.ID == id {
		// delete the refresh token if you are deleting you self
		err, statusCode := us.TokenRepository.DeleteToken(ctx, id)
		if err != nil {
			return err, statusCode
		}
	}
	return nil, 200
}

func (us *userRepository) PromoteUser(ctx context.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	if current_user.Role != "admin" || current_user.ID == id {
		return Domain.OmitedUser{}, errors.New("permission denied"), http.StatusForbidden
	}
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"role", "admin"},
		}},
	}
	updateResult, err := us.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return Domain.OmitedUser{}, err, 500
	}
	if updateResult.ModifiedCount == 0 {
		return Domain.OmitedUser{}, errors.New("user does not exist"), 400
	}
	var result Domain.OmitedUser
	err = us.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return Domain.OmitedUser{}, errors.New("User Not Found"), 500
	}
	return result, nil, 200
}

func (us *userRepository) DemoteUser(ctx context.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	if current_user.Role != "admin" || current_user.ID == id {
		return Domain.OmitedUser{}, errors.New("permission denied"), http.StatusForbidden
	}
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"role", "user"},
		}},
	}
	updateResult, err := us.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return Domain.OmitedUser{}, err, 500
	}
	if updateResult.ModifiedCount == 0 {
		return Domain.OmitedUser{}, errors.New("user does not exist"), 400
	}
	var result Domain.OmitedUser
	err = us.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return Domain.OmitedUser{}, errors.New("User Not Found"), 500
	}
	return result, nil, 200
}

func (us *userRepository) ChangePassByEmail(ctx context.Context, email string, password string) (Domain.OmitedUser, error, int) {
	statusCode := 200
	filter := bson.D{{"email", email}}
	update := bson.D{
		{"$set", bson.D{
			{"password", password},
		}},
	}
	updateResult, err := us.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		statusCode = 500
		return Domain.OmitedUser{}, err, statusCode
	}
	if updateResult.ModifiedCount == 0 {
		statusCode = 400
		fmt.Println("user does not exist:", email)
		return Domain.OmitedUser{}, errors.New("user does not exist"), statusCode
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return Domain.OmitedUser{}, nil, statusCode
}

// find by email
func (us *userRepository) FindByEmail(ctx context.Context, email string) (Domain.OmitedUser, error, int) {
	filter := bson.D{{"email", email}}
	var result Domain.OmitedUser
	err := us.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return Domain.OmitedUser{}, err, 500
	}
	return result, nil, 200
}

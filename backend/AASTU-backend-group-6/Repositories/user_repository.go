package repositories

import (
	domain "blogs/Domain"
	"blogs/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

// CreateUser implements domain.UserRepository.
func (ur *userRepository) CreateUser(c context.Context, user domain.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

// DeleteUser implements domain.UserRepository.
func (u *userRepository) DeleteUser(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindUserByEmail implements domain.UserRepository.
func (ur *userRepository) FindUserByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

// FindUserByID implements domain.UserRepository.
func (ur *userRepository) FindUserByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{} , err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)

	if err != nil {
		return domain.User{} , err
	}
	
	return user, nil
	
}

// FindUserByUsername implements domain.UserRepository.
func (ur *userRepository) FindUserByUsername(c context.Context, username string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"username": username}).Decode(&user)
	return user, err
}

// ForgotPassword implements domain.UserRepository.
func (u *userRepository) ForgotPassword(ctx context.Context, email string, token string) error {
	panic("unimplemented")
}

// UpdateUser implements domain.UserRepository.
func (ur *userRepository) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	collection := ur.database.Collection(ur.collection) 

	// changing the id to primitive object type
	
	filter := bson.M{"_id": user.ID}       
	update := bson.M{
		"$set": bson.M{
			"full_name":         user.Full_Name,
			"username":          user.Username,
			"password":          user.Password,
			"profile_image_url": user.Profile_image_url,
			"contact":           user.Contact,
			"bio":               user.Bio,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.User{}, err
	}

	// After updating, fetch the updated user from the database
	var updatedUser domain.User
	err = collection.FindOne(ctx, filter).Decode(&updatedUser)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}


func (ur *userRepository) AllUsers(c context.Context) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []domain.User

	err = cursor.All(c, &users)
	if users == nil {
		return []domain.User{}, err
	}

	return users, err
}


func (ur *userRepository) PromoteandDemoteUser(c context.Context , id string , role string) (error) {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": idHex}
	update := bson.M{
		"$set": bson.M{
			"role": role,
		},
	}

	_, err = collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}
package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRespository(collection *mongo.Collection) domain.UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (userRepo *UserRepository) CountByUsername(ctx context.Context, username string) (int, domain.Error) {
	filter := bson.D{{Key: "username", Value: username}}
	count, errCount := userRepo.collection.CountDocuments(ctx, filter)
	if errCount != nil {
		return 0, domain.CustomError{Message: fmt.Sprintln("error counting the user by username."), Code: http.StatusInternalServerError}
	}
	return int(count), nil
}

func (userRepo *UserRepository) CountByEmail(ctx context.Context, email string) (int, domain.Error) {
	filter := bson.D{{Key: "email", Value: email}}
	count, errCount := userRepo.collection.CountDocuments(ctx, filter)
	if errCount != nil {
		return 0, domain.CustomError{Message: fmt.Sprintln("error counting the user by email."), Code: http.StatusInternalServerError}
	}
	return int(count), nil
}

func (userRepo *UserRepository) CheckExistence(ctx context.Context, id string) (int, domain.Error) {
	userID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return 0, domain.CustomError{Message: fmt.Sprintf("error parsing the user id. %v \n", errIDParse.Error()), Code: http.StatusInternalServerError}
	}
	filter := bson.D{{Key: "_id", Value: userID}}
	count, errCount := userRepo.collection.CountDocuments(ctx, filter)
	if errCount != nil {
		return 0, domain.CustomError{Message: fmt.Sprintln("error counting the user."), Code: http.StatusInternalServerError}
	}
	return int(count), nil
}

func (userRepo *UserRepository) FindById(ctx context.Context, id string) (*domain.User, domain.Error) {
	userID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return &domain.User{}, domain.CustomError{Message: fmt.Sprintf("error parsing the user id. %v \n", errIDParse.Error()), Code: http.StatusInternalServerError}
	}
	filter := bson.D{{Key: "_id", Value: userID}}
	var fetchedUser domain.User
	errFetchUser := userRepo.collection.FindOne(ctx, filter).Decode(&fetchedUser)
	if errFetchUser != nil {
		if errors.Is(errFetchUser, mongo.ErrNoDocuments) {
			return &domain.User{}, domain.CustomError{Message: fmt.Sprintf("user not found. %v \n", errFetchUser.Error()), Code: http.StatusNotFound}
		}
		return &domain.User{}, domain.CustomError{Message: fmt.Sprintf("error fetching the user. %v \n", errFetchUser.Error()), Code: http.StatusInternalServerError}
	}
	return &fetchedUser, nil
}

func (userRepo *UserRepository) FindByEmail(cxt context.Context, email string) (*domain.User, domain.Error) {
	filter := bson.D{{Key: "email", Value: email}}
	var fetchedUser domain.User
	errFetchUser := userRepo.collection.FindOne(cxt, filter).Decode(&fetchedUser)
	if errFetchUser != nil {
		if errors.Is(errFetchUser, mongo.ErrNoDocuments) {
			return &domain.User{}, domain.CustomError{Message: fmt.Sprintf("user not found. %v \n", errFetchUser.Error()), Code: http.StatusNotFound}
		}
		return &domain.User{}, domain.CustomError{Message: fmt.Sprintf("error fetching the user. %v \n", errFetchUser.Error()), Code: http.StatusInternalServerError}
	}
	return &fetchedUser, nil
}

func (userRepo *UserRepository) FindByUsername(cxt context.Context, username string) (*domain.User, domain.Error) {
	filter := bson.D{{Key: "username", Value: username}}

	var fetchedUser domain.User
	errFetchUser := userRepo.collection.FindOne(cxt, filter).Decode(&fetchedUser)

	if errFetchUser != nil {
		if errors.Is(errFetchUser, mongo.ErrNoDocuments) {
			return &domain.User{}, domain.CustomError{Message: fmt.Sprintf("user not found. %v \n", errFetchUser.Error()), Code: http.StatusNotFound}
		}
		return &domain.User{}, domain.CustomError{Message: fmt.Sprintf("error fetching the user. %v \n", errFetchUser.Error()), Code: http.StatusInternalServerError}
	}
	return &fetchedUser, nil
}
func (userRepo *UserRepository) FindAll(cxt context.Context) ([]domain.User, domain.Error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "username", Value: 1}})
	var fetchedUsers []domain.User
	cursor, errFetchUsers := userRepo.collection.Find(cxt, filter, opts)
	if errFetchUsers != nil {
		return []domain.User{}, domain.CustomError{Message: fmt.Sprintf("error fetching the users. %v \n", errFetchUsers.Error()), Code: http.StatusInternalServerError}
	}

	errCusrsor := cursor.All(cxt, &fetchedUsers)
	if errCusrsor != nil {
		return []domain.User{}, domain.CustomError{Message: fmt.Sprintf("error decoding the users. %v \n", errCusrsor.Error()), Code: http.StatusInternalServerError}
	}
	return fetchedUsers, nil
}

func (userRepo *UserRepository) Create(cxt context.Context, user *domain.User) (*domain.User, domain.Error) {
	insertResult, errInsert := userRepo.collection.InsertOne(cxt, user)
	if errInsert != nil {
		if mongo.IsDuplicateKeyError(errInsert) {
			return &domain.User{}, domain.CustomError{Message: fmt.Sprintf("user already exists. %v \n", errInsert.Error()), Code: http.StatusConflict}
		}
		return &domain.User{}, domain.CustomError{Message: fmt.Sprintf("error inserting the user. %v \n", errInsert.Error()), Code: http.StatusInternalServerError}
	}
	returnedID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return &domain.User{}, domain.CustomError{Message: fmt.Sprintln("error getting the user id"), Code: http.StatusInternalServerError}
	}
	user.ID = returnedID
	return user, nil
}

func (userRepo *UserRepository) UpdateProfile(cxt context.Context, id string, user map[string]interface{}) domain.Error {
	updateID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return domain.CustomError{Message: fmt.Sprintf("error parsing the user id. %v \n", errIDParse.Error()), Code: http.StatusInternalServerError}
	}

	updateDoc := bson.M{}
	if user["username"] != nil {
		updateDoc["username"] = user["username"]
	}
	if user["email"] != nil {
		updateDoc["email"] = user["email"]
	}
	updateDoc["bio"] = user["bio"]
	updateDoc["updated_at"] = time.Now()

	filter := bson.D{{Key: "_id", Value: updateID}}
	update := bson.D{{Key: "$set", Value: updateDoc}}
	opts := options.Update().SetUpsert(false)
	updateResult, errUpdate := userRepo.collection.UpdateOne(cxt, filter, update, opts)

	if errUpdate != nil {
		return domain.CustomError{Message: fmt.Sprintf("error updating the user. %v \n", errUpdate.Error()), Code: http.StatusInternalServerError}
	}
	if updateResult.MatchedCount == 0 {
		return domain.CustomError{Message: fmt.Sprintln("user not found."), Code: http.StatusNotFound}
	}

	return nil
}

func (userRepo *UserRepository) UpdatePassword(cxt context.Context, id, password string) domain.Error {
	updateID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return domain.CustomError{Message: fmt.Sprintln("error parsing the user id."), Code: http.StatusInternalServerError}
	}
	filter := bson.D{{Key: "_id", Value: updateID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: password}}}}
	opts := options.Update().SetUpsert(false)
	updateResult, errUpdate := userRepo.collection.UpdateOne(cxt, filter, update, opts)
	if errUpdate != nil {
		return domain.CustomError{Message: fmt.Sprintf("error updating the user. %v \n", errUpdate.Error()), Code: http.StatusInternalServerError}
	}
	if updateResult.MatchedCount == 0 {
		return domain.CustomError{Message: fmt.Sprintln("user not found"), Code: http.StatusNotFound}
	}
	return nil

}

func (userRepo *UserRepository) UpdateRole(cxt context.Context, id, role string) domain.Error {
	updateID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return domain.CustomError{Message: fmt.Sprintln("error parsing the user id."), Code: http.StatusInternalServerError}
	}
	filter := bson.D{{Key: "_id", Value: updateID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "role", Value: role}}}}
	opts := options.Update().SetUpsert(false)
	updateResult, errUpdate := userRepo.collection.UpdateOne(cxt, filter, update, opts)
	if errUpdate != nil {
		return domain.CustomError{Message: fmt.Sprintf("error updating the user. %v \n", errUpdate.Error()), Code: http.StatusInternalServerError}
	}
	if updateResult.MatchedCount == 0 {
		return domain.CustomError{Message: fmt.Sprintln("user not found"), Code: http.StatusNotFound}
	}
	return nil

}

func (userRepo *UserRepository) Delete(cxt context.Context, id string) domain.Error {
	deleteID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return domain.CustomError{Message: fmt.Sprintf("error parsing the user id. %v \n", errIDParse.Error()), Code: http.StatusInternalServerError}
	}

	filter := bson.D{{Key: "_id", Value: deleteID}}
	deleteResult, errDelete := userRepo.collection.DeleteOne(cxt, filter)

	if errDelete != nil {
		return domain.CustomError{Message: fmt.Sprintf("error updating the user. %v \n", errDelete.Error()), Code: http.StatusInternalServerError}
	}
	if deleteResult.DeletedCount == 0 {
		return domain.CustomError{Message: fmt.Sprintln("user not found"), Code: http.StatusNotFound}
	}
	return nil
}

func (userRepo *UserRepository) UploadProfilePicture(cxt context.Context, picture domain.Photo, id string) domain.Error {
	updateID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return domain.CustomError{Message: fmt.Sprintln("error parsing the user id."), Code: http.StatusInternalServerError}
	}
	filter := bson.D{{Key: "_id", Value: updateID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "profile_picture", Value: picture}}}}
	opts := options.Update().SetUpsert(false)
	updateResult, errUpdate := userRepo.collection.UpdateOne(cxt, filter, update, opts)
	if errUpdate != nil {
		return domain.CustomError{Message: fmt.Sprintf("error updating the user profile_picture. %v \n", errUpdate.Error()), Code: http.StatusInternalServerError}
	}
	if updateResult.MatchedCount == 0 {
		return domain.CustomError{Message: fmt.Sprintln("user not found"), Code: http.StatusNotFound}
	}
	return nil
}

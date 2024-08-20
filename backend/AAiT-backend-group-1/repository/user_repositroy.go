package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

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

func (userRepo *UserRepository) FindById(ctx context.Context, id string) (*domain.User, domain.Error) {
	userID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return &domain.User{}, &domain.CustomError{Message: fmt.Sprintf("error parsing the user id. %v \n", errIDParse.Error()), Code: http.StatusInternalServerError}
	}
	filter := bson.D{{"_id", userID}}
	var fetchedUser domain.User
	errFetchUser := userRepo.collection.FindOne(ctx, filter).Decode(&fetchedUser)
	if errFetchUser != nil {
		if errors.Is(errFetchUser, mongo.ErrNoDocuments) {
			return &domain.User{}, &domain.CustomError{Message: fmt.Sprintf("user not found. %v \n", errFetchUser.Error()), Code: http.StatusNotFound}
		}
		return &domain.User{}, &domain.CustomError{Message: fmt.Sprintf("error fetching the user. %v \n", errFetchUser.Error()), Code: http.StatusInternalServerError}
	}
	return &fetchedUser, nil
}

func (userRepo *UserRepository) FindByEmail(cxt context.Context, email string) (*domain.User, domain.Error) {
	filter := bson.D{{"email", email}}
	var fetchedUser domain.User
	errFetchUser := userRepo.collection.FindOne(cxt, filter).Decode(&fetchedUser)
	if errFetchUser != nil {
		if errors.Is(errFetchUser, mongo.ErrNoDocuments) {
			return &domain.User{}, &domain.CustomError{Message: fmt.Sprintf("user not found. %v \n", errFetchUser.Error()), Code: http.StatusNotFound}
		}
		return &domain.User{}, &domain.CustomError{Message: fmt.Sprintf("error fetching the user. %v \n", errFetchUser.Error()), Code: http.StatusInternalServerError}
	}
	return &fetchedUser, nil
}
func (userRepo *UserRepository) FindByUsername(cxt context.Context, username string) (*domain.User, domain.Error) {
	filter := bson.D{{"username", username}}
	var fetchedUser domain.User
	errFetchUser := userRepo.collection.FindOne(cxt, filter).Decode(&fetchedUser)
	if errFetchUser != nil {
		if errors.Is(errFetchUser, mongo.ErrNoDocuments) {
			return &domain.User{}, &domain.CustomError{Message: fmt.Sprintf("user not found. %v \n", errFetchUser.Error()), Code: http.StatusNotFound}
		}
		return &domain.User{}, &domain.CustomError{Message: fmt.Sprintf("error fetching the user. %v \n", errFetchUser.Error()), Code: http.StatusInternalServerError}
	}
	return &fetchedUser, nil
}
func (userRepo *UserRepository) FindAll(cxt context.Context) ([]domain.User, domain.Error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"username", 1}})
	var fetchedUsers []domain.User
	cursor, errFetchUsers := userRepo.collection.Find(cxt, filter, opts)
	if errFetchUsers != nil {
		return []domain.User{}, &domain.CustomError{Message: fmt.Sprintf("error fetching the users. %v \n", errFetchUsers.Error()), Code: http.StatusInternalServerError}
	}

	errCusrsor := cursor.All(cxt, &fetchedUsers)
	if errCusrsor != nil {
		return []domain.User{}, &domain.CustomError{Message: fmt.Sprintf("error decoding the users. %v \n", errCusrsor.Error()), Code: http.StatusInternalServerError}
	}
	return fetchedUsers, nil
}

func (userRepo *UserRepository) Create(cxt context.Context, user *domain.User) (*domain.User, domain.Error) {
	insertResult, errInsert := userRepo.collection.InsertOne(cxt, user)
	if errInsert != nil {
		if mongo.IsDuplicateKeyError(errInsert) {
			return &domain.User{}, &domain.CustomError{Message: fmt.Sprintf("user already exists. %v \n", errInsert.Error()), Code: http.StatusConflict}
		}
		return &domain.User{}, &domain.CustomError{Message: fmt.Sprintf("error inserting the user. %v \n", errInsert.Error()), Code: http.StatusInternalServerError}
	}
	returnedID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return &domain.User{}, &domain.CustomError{Message: fmt.Sprintf("error getting the user id. %v \n", errInsert.Error()), Code: http.StatusInternalServerError}
	}
	user.ID = returnedID
	return user, nil
}

func (userRepo *UserRepository) Update(cxt context.Context, id string, user *domain.User) domain.Error {
	updateID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return &domain.CustomError{Message: fmt.Sprintf("error parsing the user id. %v \n", errIDParse.Error()), Code: http.StatusInternalServerError}
	}

	filter := bson.D{{"_id", updateID}}
	updateDoc := bson.D{{"$set", user}}
	opts := options.Update().SetUpsert(false)
	updateResult, errUpdate := userRepo.collection.UpdateOne(cxt, filter, updateDoc, opts)

	if errUpdate != nil {
		return &domain.CustomError{Message: fmt.Sprintf("error updating the user. %v \n", errUpdate.Error()), Code: http.StatusInternalServerError}
	}
	if updateResult.MatchedCount == 0 {
		return &domain.CustomError{Message: fmt.Sprintf("user not found. %v \n", errUpdate.Error()), Code: http.StatusNotFound}
	}

	return nil
}
func (userRepo *UserRepository) Delete(cxt context.Context, id string) domain.Error {
	deleteID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return &domain.CustomError{Message: fmt.Sprintf("error parsing the user id. %v \n", errIDParse.Error()), Code: http.StatusInternalServerError}
	}

	filter := bson.D{{"_id", deleteID}}
	deleteResult, errDelete := userRepo.collection.DeleteOne(cxt, filter)

	if errDelete != nil {
		return &domain.CustomError{Message: fmt.Sprintf("error updating the user. %v \n", errDelete.Error()), Code: http.StatusInternalServerError}
	}
	if deleteResult.DeletedCount == 0 {
		return &domain.CustomError{Message: fmt.Sprintf("user not found. %v \n", errDelete.Error()), Code: http.StatusNotFound}
	}
	return nil
}

package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRespository(collection *mongo.Collection) domain.UserRepository {
	return UserRepository{
		collection: collection,
	}
}

func (userRepo *UserRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	userID := uuid.String(id)
	if userID == "" {
		return domain.User{}, domain.UserError{Message: fmt.Sprintf("invalid uuid format, got %v \n", id), Code: http.StatusBadRequest}
	}
	filter := bson.D{{"id": userID}}
	var fetchedUser domain.User
	err := userRepo.collection.FindOne(ctx, filter).Decode(&fetchedUser)
	if err != nil {
		return domain.User{}, domain.UserError{Message: fmt.Sprintf("an error occured during fetching user"), Code: http.StatusInternalServerError}
	}
	return fetchedUser, nil
}

func (userRepo *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	filter := bson.D{{"email": email}}
	var fetchedUser domain.User
	err := userRepo.collection.FindOne(ctx, filter).Decode(&fetchedUser)
	if err != nil {
		return domain.User, domain.UserError{Message: fmt.Sprintf("an error occured during fetching user, %v \n", errFetching.Error()), Code: http.StatusInternalServerError}
	}
	return fetchedUser, nil
}

func (userRepo *UserRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	filter := bson.D{{"username": username}}
	var fetchedUser domain.User
	err := userRepo.collection.FindOne(ctx, filter).Decode(&fetchedUser)
	if err != nil {
		return domain.User, domain.UserError{Message: fmt.Sprintf("an error occured during fetching user, %v \n", errFetching.Error()), Code: http.StatusInternalServerError}
	}
	return fetchedUser, nil
}

func (userRepo *UserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"username", 1}})
	cursor, errFetching := userRepo.collection.Find(ctx, filter, opts)
	if errFetching != nil {
		return []domain.User, domain.UserError{Message: fmt.Sprintf("an error occured during fetching users, %v \n", errFetching.Error()), Code: http.StatusInternalServerError}
	}

	var fetchedUsers []domain.User
	errCursor := cursor.All(&fetchedUsers)
	if errCursor != nil {
		return []domain.User, domain.UserError{Message: fmt.Sprintf("an error occured during fetching users, %v \n", errFetching.Error()), Code: http.StatusInternalServerError}
	}

	return fetchedUsers, nil
}

func (userRepo *UserRepository) Create(cxt context.Context, user *domain.User) (*domain.User, error) {
	insertedResult, errInserting := userRepo.collection.InsertOne(cxt, user)
	if errInserting != nil {
		return domain.User{}, domain.UserError{Message: fmt.Sprintf("an error occured during inserting user, %v \n", errInserting.Error()), Code: http.StatusInternalServerError}
	}
	return user, nil
}
func (userRepo *UserRepository) Update(cxt context.Context, user *domain.User) (*domain.User, error) {
	userID := uuid.String(user.ID)
	if userID == "" {
		return domain.User{}, domain.UserError{Message: fmt.Sprintf("invalid uuid format, got %v \n", user.ID), Code: http.StatusBadRequest}
	}
	filter := bson.D{{"id": userID}}
	update := bson.D{{"$set", bson.D{{"username", user.Username}, {"email", user.Email}, {"password", user.Password}}}}

	var updatedUser domain.User
	errUpdating := userRepo.collection.FindOneAndUpdate(cxt, filter, update).Decode(&updatedUser)
	if errUpdating != nil {
		if errors.Is(errUpdating, mongo.ErrNoDocuments) {
			return domain.User{}, domain.UserError{Message: fmt.Sprintf("user with id %v not found to update\n", userID), Code: http.StatusNotFound}
		}
		return domain.User{}, domain.UserError{Message: fmt.Sprintf("an error occured during updating user, %v \n", err.Error()), Code: http.StatusInternalServerError}
	}
	return updatedUser, nil
}
func (userRepo *UserRepository) Delete(cxt context.Context, id uuid.UUID) error {
	userID := uuid.String(id)
	if userID == "" {
		return domain.UserError{Message: fmt.Sprintf("invalid uuid format, got %v \n", id), Code: http.StatusBadRequest}
	}
	filter := bson.D{{"id": userID}}
	_, errDeleting := userRepo.collection.DeleteOne(cxt, filter)
	if errDeleting != nil {
		return domain.UserError{Message: fmt.Sprintf("an error occured during deleting user, %v \n", errDeleting.Error()), Code: http.StatusInternalServerError}
	}
	return nil
}

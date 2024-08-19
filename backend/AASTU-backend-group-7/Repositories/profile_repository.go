package Repositories

import (
	"blogapp/Domain"
	"blogapp/Infrastructure/password_services"
	"context"
	"errors"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-playground/validator"
)

type profileRepository struct {
	validator       *validator.Validate
	collection      Domain.Collection
	TokenRepository Domain.RefreshRepository
}

func NewProfileRepository(_collection Domain.Collection, token_collection Domain.Collection) *profileRepository {
	return &profileRepository{
		validator: validator.New(),

		collection:      _collection,
		TokenRepository: NewRefreshRepository(token_collection),
	}

}

// get user by id
func (ps *profileRepository) GetProfile(ctx context.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	if current_user.ID != id {
		return Domain.OmitedUser{}, errors.New("permission denied"), http.StatusForbidden
	}

	var filter bson.D
	filter = bson.D{{"_id", id}}
	var result Domain.OmitedUser
	err := ps.collection.FindOne(ctx, filter).Decode(&result)
	// # handel this later
	if err != nil {
		return Domain.OmitedUser{}, errors.New("User not found"), http.StatusNotFound
	}
	return result, nil, 200
}

// update user by id
func (ps *profileRepository) UpdateProfile(ctx context.Context, id primitive.ObjectID, user Domain.User, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	if current_user.ID != id {
		return Domain.OmitedUser{}, errors.New("permission denied"), http.StatusForbidden
	}
	var NewUser Domain.OmitedUser
	statusCode := 200

	// Retrieve the existing user
	NewUser, err, statusCode := ps.GetProfile(ctx, id, current_user)
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
	if user.Password != "" {
		err = password_services.CheckPasswordStrength(user.Password)
		if err != nil {
			return Domain.OmitedUser{}, err, 400
		}
		newpass, er := password_services.GenerateFromPasswordCustom(user.Password)
		if er != nil {
			return Domain.OmitedUser{}, er, 500
		}
		NewUser.Password = newpass
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

	updateResult, err := ps.collection.UpdateOne(ctx, filter, update)
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
func (ps *profileRepository) DeleteProfile(ctx context.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (error, int) {

	filter := bson.D{{"_id", id}}
	if current_user.ID != id {
		return errors.New("permission denied"), http.StatusForbidden
	}

	deleteResult, err := ps.collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return err, 500
	}
	if deleteResult.DeletedCount == 0 {
		return errors.New("User does not exist"), http.StatusNotFound
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	// delete the refresh token
	err, statusCode := ps.TokenRepository.DeleteToken(ctx, id)
	if err != nil {
		return err, statusCode
	}
	return nil, 200

}

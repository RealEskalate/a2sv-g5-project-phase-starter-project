package Repositories

import (
	"blogapp/Domain"
	"blogapp/Infrastructure/password_services"
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-playground/validator"
)

type profileRepository struct {
	validator       *validator.Validate
	userCollection  Domain.Collection
	blogCollection  Domain.Collection
	TokenRepository Domain.RefreshRepository
	mu              sync.RWMutex
}

func NewProfileRepository(userCollection Domain.Collection, token_collection Domain.Collection, blogCollection Domain.Collection) *profileRepository {
	return &profileRepository{
		validator: validator.New(),
		blogCollection: blogCollection,
		userCollection:  userCollection,
		TokenRepository: NewRefreshRepository(token_collection),
		mu:              sync.RWMutex{},
	}

}

// get user by id
func (ps *profileRepository) GetProfile(ctx context.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	if current_user.ID != id {
		return Domain.OmitedUser{}, errors.New("permission denied"), http.StatusForbidden
	}

	var filter bson.D
	filter = bson.D{{"_id", id}}
	var result Domain.OmitedUser
	err := ps.userCollection.FindOne(ctx, filter).Decode(&result)
	// # handel this later
	if err != nil {
		return Domain.OmitedUser{}, errors.New("User not found"), http.StatusNotFound
	}
	return result, nil, 200
}

// update user by id
func (ps *profileRepository) UpdateProfile(ctx context.Context, id primitive.ObjectID, user Domain.User, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
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
	if user.Name != "" {
		NewUser.Name = user.Name
	}
	NewUser.UpdatedAt = time.Now()

	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"email", NewUser.Email},
			{"username", NewUser.UserName},
			{"name", NewUser.Name},
			{"role", NewUser.Role},
			{"profile_picture", NewUser.ProfilePicture},
			{"bio", NewUser.Bio},
			{"updatedat", NewUser.UpdatedAt},
		}},
	}

	updateResult, err := ps.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		statusCode = 500
		return Domain.OmitedUser{}, err, statusCode
	}
	if updateResult.ModifiedCount == 0 {
		statusCode = 400
		return Domain.OmitedUser{}, errors.New("user does not exist or no changes"), statusCode
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return NewUser, nil, statusCode
}

// DeleteProfile removes a user profile, updates posts, and deletes the user's refresh token
func (ps *profileRepository) DeleteProfile(ctx context.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (error, int) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	// Check if the current user has permission to delete the profile
	
	fmt.Println(ps.userCollection)
	fmt.Println("-------------")
	fmt.Println(ps.blogCollection)
	if current_user.ID != id {
		return errors.New("permission denied"), http.StatusForbidden
	}

	fakeID, err:= primitive.ObjectIDFromHex("000000000000000000000000")
	if err != nil {
		return errors.New("internal server error"), http.StatusInternalServerError
	}
	// Update all posts where the authorid matches the deleted user's id
	updateFilter := bson.D{{"authorid", id}}
	update := bson.D{
		{"$set", bson.D{{"authorid", fakeID}}},
	}
	updateResult, err := ps.blogCollection.UpdateMany(ctx, updateFilter, update)
	if err != nil {
		fmt.Println("Error updating posts:", err)
		return errors.New("internal server error"), http.StatusInternalServerError
	}

	// Attempt to delete the user profile
	filter := bson.D{{"_id", id}}
	deleteResult, err := ps.userCollection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println("Error deleting user profile:", err)
		return errors.New("internal server error"), http.StatusInternalServerError
	}
	if deleteResult.DeletedCount == 0 {
		return errors.New("user does not exist"), http.StatusNotFound
	}

	// create a fake id
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Delete the refresh token associated with the user
	err, statusCode := ps.TokenRepository.DeleteToken(ctx, id)
	if err != nil {
		return err, statusCode
	}

	return nil, http.StatusOK
}

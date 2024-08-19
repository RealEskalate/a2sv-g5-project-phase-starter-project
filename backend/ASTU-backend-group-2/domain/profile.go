package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	FirstName  string             `json:"first_name" bson:"first_name" binding:"required,min=3,max=30"`
	LastName   string             `json:"last_name" bson:"last_name" binding:"max=30"`
	Email      string             `json:"email" bson:"email" binding:"required,email"`
	Bio        string             `json:"bio" bson:"bio"`
	ProfileImg string             `json:"profile_img" bson:"profile_img"`
	IsOwner    bool               `json:"is_owner" bson:"is_owner"`
	Role       string             `json:"role" bson:"role"` //may make only tobe admin or user
}
type ProfileUpdate struct {
	FirstName  string `json:"first_name" bson:"first_name"`
	LastName   string `json:"last_name" bson:"last_name"`
	Email      string `json:"email" bson:"email"`
	Bio        string `json:"bio" bson:"bio"`
	ProfileImg string `json:"profile_img" bson:"profile_img"`
}

type ProfileUsecase interface {
	GetUserProfile(c context.Context, userID string) (*User, error)                 // Retrieves a user's profile by ID
	UpdateUserProfile(c context.Context, userID string, updatedProfile *User) error // Updates a user's profile
	DeleteUserProfile(c context.Context, userID string) error                       // Deletes a user's profile
	PromoteUserToAdmin(c context.Context, userID string) error                      // Promotes a user to admin
	DemoteAdminToUser(c context.Context, userID string) error                       // Demotes an admin to a regular user
}



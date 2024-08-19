package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	FirstName  string             `json:"first_name" bson:"first_name" binding:"required,min=3,max=30"`
	LastName   string             `json:"last_name" bson:"last_name" binding:"max=30"`
	Email      string             `json:"email" bson:"email" binding:"required,email"`
	Active     bool               `json:"is_active" bson:"is_active"`
	Bio        string             `json:"bio" bson:"bio"`
	ProfileImg string             `json:"profile_img" bson:"profile_img"`
	Password   string             `json:"password" bson:"password" binding:"required,min=4,max=30,StrongPassword"`
	IsOwner    bool               `json:"is_owner" bson:"is_owner"`
	Role       string             `json:"role" bson:"role"` //may make only tobe admin or user
	Tokens     []string           `json:"tokens" bson:"tokens"`
}

// this structure defined for data sent as a response
type UserOut struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	FirstName  string             `json:"first_name" bson:"first_name"`
	LastName   string             `json:"last_name" bson:"last_name"`
	Email      string             `json:"email" bson:"email"`
	Bio        string             `json:"bio" bson:"bio"`
	ProfileImg string             `json:"profile_img" bson:"profile_img"`
	IsOwner    bool               `json:"is_owner" bson:"is_owner"`
	Role       string             `json:"role" bson:"role"` //may make only tobe admin or user
}
type UserUpdate struct {
	FirstName  string `json:"first_name" bson:"first_name"`
	LastName   string `json:"last_name" bson:"last_name"`
	Bio        string `json:"bio" bson:"bio"`
	ProfileImg string `json:"profile_img" bson:"profile_img"`
}

// user knows the password and wants to update
type UpdatePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

// user forgot the password and wants to reset
// reset passowrd token will be exreacted from the url /reset-password/:user_id/:<reset password token>
type ResetPassword struct {
	NewPassword string `json:"new_password" binding:"required"`
}

type UserRepository interface {
	GetByID(c context.Context, userID string) (*UserOut, error)                       // Retrieves a user's profile by ID
	UpdateProfile(c context.Context, userID string, updatedProfile *UserUpdate) error // Updates a user's profile
	PromoteToAdmin(c context.Context, userID string) error                            // Promotes a user to admin
	DemoteToUser(c context.Context, userID string) error                              // Demotes an admin to a regular user
	Delete(c context.Context, userID string) error                                    // Deletes a user's profile
}
type UserUsecase interface {
	GetUserProfile(c context.Context, userID string) (*User, error)                 // Retrieves a user's profile by ID
	UpdateUserProfile(c context.Context, userID string, updatedProfile *User) error // Updates a user's profile
	DeleteUserProfile(c context.Context, userID string) error                       // Deletes a user's profile
	PromoteUserToAdmin(c context.Context, userID string) error                      // Promotes a user to admin
	DemoteAdminToUser(c context.Context, userID string) error                       // Demotes an admin to a regular user
}

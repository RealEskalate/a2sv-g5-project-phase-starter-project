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



type UserRepository interface {
	GetByID(c context.Context, userID string) (*User, error)          // Retrieves a user's profile by ID
	Update(c context.Context, userID string, updatedProfile *User) error // Updates a user's profile
	Delete(c context.Context, userID string) error                      // Deletes a user's profile
	PromoteToAdmin(c context.Context, userID string) error              // Promotes a user to admin
	DemoteToUser(c context.Context, userID string) error                // Demotes an admin to a regular user
}

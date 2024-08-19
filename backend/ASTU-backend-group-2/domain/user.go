package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
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

type UserUsecase interface {
	CreateUser(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUser(c context.Context, userID string) (*User, error)
	UpdateUser(c context.Context, userID string, updatedUser *User) error
	DeleteUser(c context.Context, userID string) error
	IsUserActive(c context.Context, userID string) (bool, error)

	ResetUserPassword(c context.Context, userID string, resetPassword *ResetPassword) error
	UpdateUserPassword(c context.Context, userID string, updatePassword *UpdatePassword) error

	PromoteUserToAdmin(c context.Context, userID string) error
	DemoteAdminToUser(c context.Context, userID string) error
}

type UserRepository interface {
	CreateUser(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUser(c context.Context, userID string) (*User, error)
	UpdateUser(c context.Context, userID string, updatedUser *User) error
	DeleteUser(c context.Context, userID string) error
	IsUserActive(c context.Context, userID string) (bool, error)

	ResetUserPassword(c context.Context, userID string, resetPassword *ResetPassword) error
	UpdateUserPassword(c context.Context, userID string, updatePassword *UpdatePassword) error

	PromoteUserToAdmin(c context.Context, userID string) error
	DemoteAdminToUser(c context.Context, userID string) error
}

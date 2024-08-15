package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Role is a type for user roles
type UserRepository interface {
	GetUserByID(c context.Context, id primitive.ObjectID) (*User, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserByUsername(c context.Context, username string) (*User, error)
	CreateUser(c context.Context, user *User) error
	UpdateUser(c context.Context, user *User) error
	DeleteUser(c context.Context, id primitive.ObjectID) error
}

// SignupRepository is an interface that contains the CreateUser method

type UserUsecase interface {
	GetUserByID(c context.Context, id primitive.ObjectID) (*User, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserByUsername(c context.Context, username string) (*User, error)
	CreateUser(c context.Context, user *CreateUser) error
	PromoteUser(c context.Context, id primitive.ObjectID) (*Privilage, error)
	DemoteUser(c context.Context, id primitive.ObjectID) (*Privilage, error)
	UpdateUser(c context.Context, profile *Profile) (*ProfileResponse, error)
	DeleteUser(c context.Context, id primitive.ObjectID) error
}

type SignupUsecase interface {
	RegisterUser(c context.Context, user *AuthSignup) (*primitive.ObjectID, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserByUsername(c context.Context, username string) (*User, error)
	CreateAccessToken(user *AuthSignup, secret string, expiry int) (string, error)
	CreateRefreshToken(user *AuthSignup, secret string, expiry int) (string, error)
}

type ProfileUsecase interface {
	CreateProfile(c context.Context, profile *Profile) error
	GetProfileByID(c context.Context, id primitive.ObjectID) (*ProfileResponse, error)
	UpdateProfile(c context.Context, profile *Profile) error
	DeleteProfile(c context.Context, id primitive.ObjectID) error
}

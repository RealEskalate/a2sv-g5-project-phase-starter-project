package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	UserCollection = "users"
)
type User struct{
    ID                      primitive.ObjectID `bson:"_id" json:"_id"`
    Name                    string             `bson:"name" json:"name" validate:"required,min=2,max=100"`
    Username                string             `bson:"username" json:"username" validate:"required,min=2,max=100"`
    Password                string             `bson:"password" json:"password" validate:"required,min=8,max=32"`
    Email                   string             `bson:"email" json:"email" validate:"email,required"`
    User_type               string             `bson:"user_type" json:"user_type"`
    Is_active               bool               `bson:"is_active" json:"is_active"`
    VerificationCode        string             `bson:"verification_code" json:"verification_code"`
    VerificationCodeExpiry  time.Time          `bson:"verification_code_expiry" json:"verification_code_expiry"`
    PWRecoveryToken         string             `bson:"pw_recovery_token" json:"pw_recovery_token"`
    PWRecoveryTokenExpiry   time.Time          `bson:"pw_recovery_token_expiry" json:"pw_recovery_token_expiry"`
    Token                   string             `bson:"token" json:"token"`
    Refresh_token           string             `bson:"refresh_token" json:"refresh_token"`
    Created_at              time.Time          `bson:"created_at" json:"created_at"`
    Updated_at              time.Time          `bson:"updated_at" json:"updated_at"`
}


type UserUsecase interface{
	CreateUser(c context.Context, user *User) error
	GetUserByEmail(c context.Context, id string) (*User, error)
	GetUserByUsername(c context.Context, username string) (*User, error)
	GetUserByID(c context.Context, id string) (*User, error)
	GetUsers(c context.Context, ) ([]*User, error)
	UpdateUser(c context.Context, user *User) error
	DeleteUser(c context.Context, id string) error
}

type UserRepository interface{
	CreateUser(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserByUsername(c context.Context, username string) (*User, error)
	GetUserByID(c context.Context, id string) (*User, error)
	GetUsers(c context.Context, ) ([]*User, error)
	UpdateUser(c context.Context, user *User) error
	DeleteUser(c context.Context, id string) error
}
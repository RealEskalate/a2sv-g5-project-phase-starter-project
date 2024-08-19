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
	ID						primitive.ObjectID		`bson:"_id"`
	Name		    		string			`json:"name" validate:"required,min=2,max=100"`
	Username				string			`json:"username" validate:"required,min=2,max=100"`
	Password				string			`json:"password" validate:"required,min=6"`
	Email					string			`json:"email" validate:"email,required"`
	User_type				string			`json:"user_type"`
	Is_active       		bool            `json:"is_active"`
	VerificationCode 		string          `json:"verification_code"`
	VerificationCodeExpiry  time.Time       `json:"verification_code_expiry"`
	Token 					string			`json:"token"`
	Refresh_token			string			`json:"refresh_token"`
	Created_at				time.Time		`json:"created_at"`
	Updated_at				time.Time		`json:"updated_at"`
	User_id					string			`json:"user_id"`
}

type UserUsecase interface{
	CreateUser(c context.Context, user *User) error
	GetByEmail(c context.Context, id string) (*User, error)
	GetByID(c context.Context, id string) (*User, error)
	GetUsers(c context.Context, ) ([]*User, error)
	UpdateUser(c context.Context, user *User, user_id string) error
	DeleteUser(c context.Context, id string) error
	LoginUser(c context.Context, user *User) (*User, error)
}

type UserRepository interface{
	CreateUser(c context.Context, user *User) error
	GetByEmail(c context.Context, email string) (User, error)
	GetByUsername(c context.Context, username string) (User, error)
	GetByID(c context.Context, id string) (User, error)
	GetUsers(c context.Context, ) ([]User, error)
	UpdateUser(c context.Context, user *User) error
	DeleteUser(c context.Context, id string) error
	LoginUser(c context.Context, user *User) (User, error)
}
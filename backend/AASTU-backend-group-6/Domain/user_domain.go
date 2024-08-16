package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Full_Name         string             `json:"full_name"`
	Email             string             `json:"email" validate:"required,email"`
	Username          string             `json:"username" validate:"required"`
	Password          string             `json:"password" validate:"required"`
	Profile_image_url string             `json:"profile_image" `
	GoogleID          string             `json:"googleId"`
	Posts             []Post             `json:"posts"`
	RefreshToken      string             `json:"refreshToken" validate:"required"`
	AccessToken       string             `json:"accessToken" validate:"required"`
	Contact           string             `json:"contact"`
	Bio               string             `json:"bio"`
	Role              string             `json:"roles" validate:"required"`
	Comments          []Comment          `json:"comments"`
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	FindUserByEmail(email string) (User, error)
	FindUserByUsername(username string) (User, error)
	FindUserByID(id string) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id string) error
	ForgotPassword(email string, token string) error
}

type SignupRepository interface {
	Create(User) (User, error)
	FindUserByEmail(c context.Context , email string) (User, error)
}

type SignupUseCase interface {
	Create(c context.Context , user User) interface{}
}

type UserUseCase interface {
	CreateUser(user User) interface{}
	FindUserByEmail(email string) interface{}
	FindUserByUsername(username string) interface{}
	FindUserByID(id string) interface{}
	UpdateUser(user User) interface{}
	DeleteUser(id string) interface{}
	ForgotPassword(email string, token string) interface{}
}

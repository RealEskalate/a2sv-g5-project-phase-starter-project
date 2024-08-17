package domain

import (
	"context"
	"time"

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
	Posts             []Blog             `json:"posts"`
	RefreshToken      string             `json:"refreshToken" validate:"required"`
	AccessToken       string             `json:"accessToken" validate:"required"`
	Contact           string             `json:"contact"`
	Bio               string             `json:"bio"`
	Role              string             `json:"roles" validate:"required"`
	Comments          []Comment          `json:"comments"`
	Verified 		 bool 				`json:"verified"`
	OTP 			 string 			`json:"otp"`
	ExpiresAt 		 time.Time			`json:"expires_at"` 
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
	Create(c context.Context , user User) (User, error)
	FindUserByEmail(c context.Context , email string) (User, error)
	SetOTP(c context.Context , email string , otp string) (error)
	VerifyUser(c context.Context , user User) (User, error)
}

type SignupUseCase interface {
	Create(c context.Context , user User) interface{}
	VerifyOTP(c context.Context , otp OtpToken) interface{}
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

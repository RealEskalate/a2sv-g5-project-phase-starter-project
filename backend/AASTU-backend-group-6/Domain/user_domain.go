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
	CreateUser(ctx context.Context, user User) (error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	FindUserByUsername(ctx context.Context, username string) (User, error)
	FindUserByID(ctx context.Context, id string) (User, error)
	UpdateUser(ctx context.Context, user User) (User, error)
	DeleteUser(ctx context.Context, id string) error
	ForgotPassword(ctx context.Context, email string, token string) error
	AllUsers(c context.Context) ([]User, error) 
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
	CreateUser(ctx context.Context, user User) interface{}
	FindUserByEmail(ctx context.Context, email string) interface{}
	FindUserByUsername(ctx context.Context, username string) interface{}
	FindUserByID(ctx context.Context, id string) interface{}
	UpdateUser(ctx context.Context, user User) interface{}
	DeleteUser(ctx context.Context, id string) interface{}
	ForgotPassword(ctx context.Context, email string, token string) interface{}
}

package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserName     string             `json:"username,omitempty"`
	Email        string             `json:"email,omitempty"`
	Imageuri     string             `json:"imageuri,omitempty"`
	Bio          string             `json:"bio,omitempty"`
	Contact      string             `json:"contact,omitempty"`
	Password     string             `json:"password,omitempty"`
	IsAdmin      bool               `json:"isadmin,omitempty"`
	JoinedAt     time.Time          `json:"joinedat,omitempty"`
	RefreshToken string             `json:"refreshtoken,omitempty"`
	IsVerified   bool               `bson:"isverified,omitempty" json:"isverified,omitempty"`
	Oauth        bool               `json:"oauth,omitempty"`
}

type RestRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"password"`
}

type UserUsecase interface {
	UpdateUserDetails(c context.Context, user *User) *AppError
	RegisterUser(c context.Context, user *User) *AppError
	LoginUser(c context.Context, user User) (string, *AppError)
	ForgotPassword(c context.Context, email string) *AppError
	LogoutUser(c context.Context, uid string) *AppError
	PromoteDemoteUser(c context.Context, userid string, isAdmin bool) *AppError
	ResetPassword(c context.Context, token string, newPassword string) *AppError
	VerifyUserEmail(c context.Context, token string) *AppError
}

type UserRepository interface {
	UpdateUserDetails(user *User) *AppError
	RegisterUser(user *User) *AppError
	LoginUser(user User) (string, *AppError)
	ForgotPassword(email string) *AppError
	LogoutUser(uid string) *AppError
	PromoteDemoteUser(userid string, isAdmin bool) *AppError
	ResetPassword(token string, newPassword string) *AppError
	VerifyUserEmail(token string) *AppError
}

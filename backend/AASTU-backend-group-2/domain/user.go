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
	Password     string             `json:"password,omitempty"`
	IsAdmin      bool               `json:"isadmin,omitempty"`
	JoinedAt     time.Time          `json:"joinedat,omitempty"`
	RefreshToken string             `json:"refreshtoken,omitempty"`
}

type RestRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"password"`
}

type UserUsecase interface {
	RegisterUser(c context.Context, user *User) error
	LoginUser(c context.Context, user User) (string, error)
	ForgotPassword(c context.Context, email string) error
	LogoutUser(c context.Context, uid string) error
	PromoteDemoteUser(c context.Context, userid string, isAdmin bool) error
	ResetPassword(c context.Context, token string, newPassword string) error
}

type UserRepository interface {
	RegisterUser(user *User) error
	LoginUser(user User) (string, error)
	ForgotPassword(email string) error
	LogoutUser(uid string) error
	PromoteDemoteUser(userid string, isAdmin bool) error
	ResetPassword(token string, newPassword string) error
}

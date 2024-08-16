package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserName     string             `json:"username"`
	Email        string             `json:"email"`
	Password     string             `json:"password,omitempty"`
	IsAdmin      bool               `json:"isadmin,omitempty"`
	JoinedAt     time.Time          `json:"joinedat,omitempty"`
	RefreshToken string             `json:"refreshtoken,omitempty"`
}

type UserUsecase interface {
	RegisterUser(c context.Context, user User) error
	LoginUser(c context.Context, user User) (string, error)
	ForgotPassword(c context.Context, email string) error
	LogoutUser(c context.Context) error
	PromoteDemoteUser(c context.Context, userid string) error
}

type UserRepository interface {
	RegisterUser(user User) error
	LoginUser(user User) (string, error)
	ForgotPassword(email string) error
	LogoutUser() error
	PromoteDemoteUser(userid string) error
}

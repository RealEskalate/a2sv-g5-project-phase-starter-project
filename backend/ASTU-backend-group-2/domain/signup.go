package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserIn struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	FirstName  string             `json:"first_name" bson:"first_name" binding:"required,min=3,max=30"`
	LastName   string             `json:"last_name" bson:"last_name" binding:"max=30"`
	Email      string             `json:"email" bson:"email" binding:"required,email"`
	Password   string             `json:"password" bson:"password" binding:"required,min=4,max=30,StrongPassword"`
	Bio        string             `json:"bio" bson:"bio"`
	ProfileImg string             `json:"profile_img" bson:"profile_img"`
	IsOwner    bool               `json:"-" bson:"is_owner"`
	Role       string             `json:"-" bson:"role"` //may make only tobe admin or user
	Tokens     []string           `json:"tokens" bson:"tokens"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
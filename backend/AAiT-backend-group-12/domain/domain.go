package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

/*
Defines the names of the collections in the DB
*/
const (
	CollectionUsers = "users"
	CollectionBlogs = "blogs"
)

type Response gin.H

type User struct {
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	RefreshToken string    `json:"refresh_token"`
}

type UserRepositoryInterface interface {
	CreateUser(c context.Context, user *User) CodedError
	FindUser(c context.Context, user *User) (User, CodedError)
	SetRefreshToken(c context.Context, user *User, newRefreshToken string) CodedError
}

type UserUsecaseInterface interface {
	Signup(c context.Context, user *User) CodedError
	Login(c context.Context, user *User) (string, string, CodedError)
}

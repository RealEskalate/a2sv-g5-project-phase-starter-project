package domain

import (
	"context"

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
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	CreatedAt    string `json:"created_at"`
	RefreshToken string `json:"refresh_token"`
}

type UserRepositoryInterface interface {
	CreateUser(c context.Context, user *User) CodedError
}

type UserUsecaseInterface interface {
	Signup(c context.Context, user *User) CodedError
}

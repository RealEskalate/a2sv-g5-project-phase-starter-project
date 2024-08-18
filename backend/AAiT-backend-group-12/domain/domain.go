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

type Blog struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	UserID     string    `json:"user_id"`
	Tags       []string  `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ViewCount  uint      `json:"view_count"`
	LikedBy    []User    `json:"liked_by"`
	DislikedBy []User    `json:"disliked_by"`
	Comments   []Comment `json:"comment"`
}


// Comment represents a comment entity in the domain.
type Comment struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ViewCount uint      `json:"view_count"`
	Comments  []Comment `json:"comment"`
}

type BlogRepositoryInterface interface {
	FindBlogPostByID(ctx context.Context, id string) (*Blog, error)
	InsertBlogPost(ctx context.Context, blog *Blog) error
	UpdateBlogPost(ctx context.Context, id string, blog *Blog) error
	DeleteBlogPost(ctx context.Context, id string) error
}

type BlogUseCaseInterface interface {
	GetBlogPost(ctx context.Context, id string) (*Blog, error)
	CreateBlogPost(ctx context.Context, blog *Blog) error
	EditBlogPost(ctx context.Context, id string, blog *Blog) error
	DeleteBlogPost(ctx context.Context, id string) error
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

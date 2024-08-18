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
	Username   string    `json:"username"`
	Tags       []string  `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ViewCount  uint      `json:"view_count"`
	LikedBy    []string  `json:"liked_by"`
	DislikedBy []string  `json:"disliked_by"`
	Comments   []Comment `json:"comment"`
}

// Comment represents a comment entity in the domain.
type Comment struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ViewCount uint      `json:"view_count"`
	Comments  []Comment `json:"comment"`
}

type BlogFilterOptions struct {
	Title         string // Search by title
	Author        string // Search by author name
	Tags          []string
	DateFrom      time.Time
	DateTo        time.Time
	SortBy        string // Sort by criteria: date, like count, dislike count, view count
	SortDirection string // Sort direction: asc, desc
	Page          int    // Pagination: Page number
	PostsPerPage  int    // Pagination: Posts per page
	MinLikes      int    // Filter by minimum likes
	MinDislikes   int    // Filter by minimum dislikes
	MinComments   int    // Filter by minimum comments
	MinViewCount  int    // Filter by minimum view count
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

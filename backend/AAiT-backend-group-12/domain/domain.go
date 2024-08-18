package domain

import (
	"context"
	"time"
)

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

// User represents a user entity in the domain.
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
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

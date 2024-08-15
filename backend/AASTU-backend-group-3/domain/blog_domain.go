package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

)

type Tag struct {
	Name string `bson:"name"`
}

type Blog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Tags      []Tag              `bson:"tags,omitempty"` // Embedded Tags
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type BlogUsecase interface {
	GetBlogs(ctx context.Context) ([]Blog, error)
	GetBlog(ctx context.Context, blogID string) (Blog, error)
	SearchBlogs(ctx context.Context, query string) ([]Blog, error)
	FilterBlogs(ctx context.Context, filters map[string]interface{}) ([]Blog, error)
	CreateBlog(ctx context.Context, blog Blog) error
	UpdateBlog(ctx context.Context, blog Blog) error
	DeleteBlog(ctx context.Context, blogID string) error
	UpdateBlogVisibility(ctx context.Context, blogID string, visibility bool) error
}



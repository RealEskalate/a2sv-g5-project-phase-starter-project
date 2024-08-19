package domain

import (
	"context"
	"time"
)

type Blog struct {
	ID           string    `bson:"_id" json:"id"`
	Title        string    `bson:"title" json:"title"`
	Content      string    `bson:"content" json:"content"`
	Author       string    `bson:"author" json:"author"`
	Tags         []string  `bson:"tags" json:"tags"`
	ViewCount    int       `bson:"view_count" json:"view_count"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
	Comments     []Comment `json:"comments,omitempty"`
	LikeCount    int       `json:"like_count,omitempty"`
	DislikeCount int       `json:"dislike_count,omitempty"`
}

type RequestBlog struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Author  string   `json:"author"`
	Tags    []string `json:"tags"`
}

type BlogRepository interface {
	FindAll(ctx context.Context, page int, pageSize int, sortBy string, sortOrder string) ([]Blog, int, error)
	FindByID(ctx context.Context, id string) (*Blog, error)
	Save(ctx context.Context, blog *Blog) error
	Update(ctx context.Context, blog *Blog) error
	Delete(ctx context.Context, id string) error
}

type BlogUsecase interface {
	GetAllBlogs() ([]Blog, error)
	GetBlogByID(id string) (*Blog, error)
	CreateBlog(blog *Blog) error
	UpdateBlog(blog *Blog) error
	DeleteBlog(id string) error
}

package domain

import (
	"context"
	"errors"
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
	Tags    []string `json:"tags"`
}

type BlogRepository interface {
	FindAll(ctx context.Context, page int, pageSize int, sortBy string, sortOrder string) ([]Blog, int, error)
	FindByID(ctx context.Context, id string) (*Blog, error)
	Save(ctx context.Context, blog *Blog) error
	Update(ctx context.Context, blog *Blog) error
	Delete(ctx context.Context, id string) error
	Filter(ctx context.Context,tags []string, startDate,endDate,sortBy string) ([]Blog,error) 
}

type BlogUseCase interface {
    GetAllBlogs(ctx context.Context, page, pageSize int, sortBy, sortOrder string) ([]Blog, int, error)
    GetBlogByID(ctx context.Context, id string) (*Blog, error)
    CreateBlog(ctx context.Context, req *RequestBlog, author string) error
    UpdateBlog(ctx context.Context, req *RequestBlog, author,id string) error
    DeleteBlog(ctx context.Context, author,id string) error
	FilterBlogs(ctx context.Context,tags []string,startDate,endDate,sortBy string) ([]Blog,error)
}



// Validate BlogRequest field
func (r *RequestBlog) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.Content == "" {
		return errors.New("content is required")
	}
	
	if len(r.Tags) == 0 {
		return errors.New("tags cannot be empty")
	}
	return nil
}
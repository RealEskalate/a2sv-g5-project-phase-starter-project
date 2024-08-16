package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title   string             `bson:"title,omitempty" json:"title,omitempty"`
	Content string             `bson:"content,omitempty" json:"content,omitempty"`
	UserID  primitive.ObjectID `bson:"userid,omitempty" json:"userid,omitempty"`
	Tags    []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	Date    time.Time          `bson:"date,omitempty" json:"date,omitempty"`
}

type BlogUsecase interface {
	CreateBlog(c context.Context, blog Blog) error
	SearchBlog(c context.Context, postName string, authorName string) ([]Blog, error)
	FilterBlog(c context.Context, tag []string, date time.Time) ([]Blog, error)
	RetrieveBlog(c context.Context, page int) ([]Blog, error)
	UpdateBlog(c context.Context, updatedblog Blog, blogID string) error
	DeleteBlog(c context.Context, blogID string) error
}

type BlogRepository interface {
	CreateBlog(blog Blog) error
	RetrieveBlog(pgnum int) ([]Blog, error)
	UpdateBlog(updatedblog Blog, blogID string) error
	DeleteBlog(blogID string) error
	SearchBlog(postName string, authorName string) ([]Blog, error)
	FilterBlog(tag []string, date time.Time) ([]Blog, error)
}

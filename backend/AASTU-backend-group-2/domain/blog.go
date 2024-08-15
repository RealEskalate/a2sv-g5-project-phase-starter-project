package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Title   string             `bson:"title,omitempty" json:"title,omitempty"`
	Content string             `bson:"content,omitempty" json:"content,omitempty"`
	UserID  primitive.ObjectID `bson:"userid,omitempty" json:"-"`
	Tags    []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	Date    time.Time          `bson:"due_date,omitempty" json:"due_date,omitempty"`
}

type BlogUsecase interface {
	CreateBlog(c context.Context, blog Blog) error
	RetrieveBlog(c context.Context, page int) ([]Blog, error)
	UpdateBlog(c context.Context, updatedblog Blog) error
	DeleteBlog(c context.Context, blogID primitive.ObjectID) error
	SearchBlog(c context.Context) ([]Blog, error)
	FilterBlog(c context.Context) ([]Blog, error)
}

type BlogRepository interface {
	CreateBlog(blog Blog) error
	RetrieveBlog(pgnum int) ([]Blog, error)
	UpdateBlog(updatedblog Blog) error
	DeleteBlog(blogID primitive.ObjectID) error
	SearchBlog(postName string, authorName string) ([]Blog, error)
	FilterBlog(tag string, date time.Time) ([]Blog, error)
}

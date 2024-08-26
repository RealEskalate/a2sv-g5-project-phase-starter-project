package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title    string             `json:"title,omitempty"`
	Content  string             `json:"content,omitempty"`
	Imageuri string             `json:"imageuri,omitempty"`
	UserID   primitive.ObjectID `json:"userid,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
	Date     time.Time          `json:"date,omitempty"`
	Likes    int                `json:"likes,omitempty"`
	DisLikes int                `json:"dislikes,omitempty"`
	Comments int                `json:"comments,omitempty"`
}

type BlogUsecase interface {
	CreateBlog(c context.Context, blog *Blog) error
	SearchBlog(c context.Context, postName string, authorName string) ([]Blog, error)
	FilterBlog(c context.Context, tag []string, date time.Time) ([]Blog, error)
	RetrieveBlog(c context.Context, page int, sortby string, dir string) ([]Blog, int, error)
	UpdateBlog(c context.Context, updatedblog Blog, blogID string, isadmin bool, userid string) error
	DeleteBlog(c context.Context, blogID string, isadmin bool, userid string) error
}

type BlogRepository interface {
	CreateBlog(blog *Blog) error
	RetrieveBlog(pgnum int, sortby string, dir string) ([]Blog, int, error)
	UpdateBlog(updatedblog Blog, blogID string, isadmin bool, userid string) error
	DeleteBlog(blogID string, isadmin bool, userid string) error
	SearchBlog(postName string, authorName string) ([]Blog, error)
	FilterBlog(tag []string, date time.Time) ([]Blog, error)
}

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
	CreateBlog(c context.Context, blog *Blog) *AppError
	SearchBlog(c context.Context, postName string, authorName string) ([]Blog, *AppError)
	FilterBlog(c context.Context, tag []string, date time.Time) ([]Blog, *AppError)
	RetrieveBlog(c context.Context, page int, sortby string, dir string) ([]Blog, int, *AppError)
	UpdateBlog(c context.Context, updatedblog Blog, blogID string, isadmin bool, userid string) *AppError
	DeleteBlog(c context.Context, blogID string, isadmin bool, userid string) *AppError
}

type BlogRepository interface {
	CreateBlog(blog *Blog) *AppError
	RetrieveBlog(pgnum int, sortby string, dir string) ([]Blog, int, *AppError)
	UpdateBlog(updatedblog Blog, blogID string, isadmin bool, userid string) *AppError
	DeleteBlog(blogID string, isadmin bool, userid string) *AppError
	SearchBlog(postName string, authorName string) ([]Blog, *AppError)
	FilterBlog(tag []string, date time.Time) ([]Blog, *AppError)
}

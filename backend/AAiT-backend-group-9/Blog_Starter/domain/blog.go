package domain

import (
	"context"
	"time"
)

const (
	CollectionBlog = "blogs"
)

type Blog struct {
	BlogID  string `json:"blog_id" bson:"_id"`
	UserID  string `json:"user_id" bson:"user_id"`
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	Tags    []string `json:"tags" bson:"tags"`
	Author string `json:"author" bson:"author"`
	CreatedAt   time.Time `json:"createtimestamp" bson:"createtimestamp"`
	UpdatedAt    time.Time `json:"updatetimestamp" bson:"updatetimestamp"`

	AverageRating float64 `json:"average_rating" bson:"average_rating"` // do some math and modify when rating is called
	TotalRating int `json:"total_rating" bson:"total_rating"` // add when rating is called
	ViewCount int `json:"view_count" bson:"view_count"`  //add when getbyid is used and blog is viewed
	LikeCount int `json:"like_count" bson:"like_count"`  // add when like is called and substarct when it is unliked
	CommentCount int `json:"comment_count" bson:"comment_count"` // add when comment is called and substarct when it is deleted
	
}

type BlogCreate struct{
	UserID  string `json:"user_id" bson:"user_id" binding:"required"`
	Title   string `json:"title" bson:"title" binding:"required"`
	Content string `json:"content" bson:"content" binding:"required"`
	Tags    []string `json:"tags" bson:"tags" binding:"required"`
}

type BlogUpdate struct{
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	Tags    []string `json:"tags" bson:"tags"`
	UpdatedAt    time.Time `json:"updatetimestamp" bson:"updatetimestamp"`
	
}

type BlogFilter struct {
    Tags    []string
    Date    time.Time
    SortByPopularity bool
}


type BlogRepository interface {
	CreateBlog(c context.Context, blog *Blog) (*Blog, error)
	GetBlogByID(c context.Context, blogID string) (*Blog, error) // add view count when this is called
	GetAllBlog(c context.Context) ([]*Blog, error)
	UpdateBlog(c context.Context, blog *BlogUpdate, blogID string) (*Blog, error)
	DeleteBlog(c context.Context, blogID string) error
	SearchBlogs(c context.Context, title string, author string) ([]*Blog, error)
	FilterBlogs(c context.Context, filters *BlogFilter) ([]*Blog, error)
	
	IncrementViewCount(c context.Context, blogID string) error
	UpdateRating(c context.Context, blogID string, rating float64) error
	UpdateLikeCount(c context.Context, blogID string, increment bool) error
	UpdateCommentCount(c context.Context, blogID string, increment bool) error

}

type BlogUseCase interface{
	CreateBlog(c context.Context, blog *BlogCreate) (*Blog, error)
	GetBlogByID(c context.Context, blogID string) (*Blog, error)
	GetAllBlog(c context.Context) ([]*Blog, error)
	UpdateBlog(c context.Context, blog *BlogUpdate, blogID string) (*Blog, error)
	DeleteBlog(c context.Context, blogID string) error
	SearchBlogs(c context.Context, title string, author string) ([]*Blog, error)

}


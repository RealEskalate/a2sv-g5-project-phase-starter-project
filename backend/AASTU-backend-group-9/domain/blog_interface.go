package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogService interface {
	CreateBlog(blog *Blog) (*Blog, error)
	GetBlogByID(id primitive.ObjectID) (*Blog, error)
	GetAllBlogs(page, pageSize int, sortBy string) ([]Blog, error)
	UpdateBlog(blog *Blog) (*Blog, error)
	DeleteBlog(id primitive.ObjectID, userID primitive.ObjectID) error
	SearchBlogs(title string, author string) ([]Blog, error)
	FilterBlogs(tags []string, date time.Time, sortBy string) ([]Blog, error)
	TrackPopularity(blogID primitive.ObjectID, action string) error // e.g., "like", "dislike", "view"
}

type CommentService interface {
	AddComment(blogID primitive.ObjectID, comment *Comment) (*Comment, error)
	DeleteComment(blogID, commentID primitive.ObjectID, userID primitive.ObjectID) error
	GetCommentsByBlogID(blogID primitive.ObjectID) ([]Comment, error)
}


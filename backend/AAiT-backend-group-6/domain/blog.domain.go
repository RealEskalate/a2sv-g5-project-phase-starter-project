package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBlogs = "blogs"
)

type Blog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Author    string             `json:"author" bson:"author"`
	Content   string             `json:"content" bson:"content"`
	Title     string             `json:"title" bson:"title"`
	Tags      []string           `json:"tags" bson:"tags"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type Like struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogID  primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	IsLiked bool               `json:"is_liked" bson:"is_liked"`
}

type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	AuthorID  primitive.ObjectID `json:"author" bson:"author"`
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type BlogUseCase interface {
	CreateBlog(c context.Context, blog *Blog) (Blog, error)
	GetBlog(c context.Context, id string) (*Blog, error)

	GetBlogs(c context.Context, pagination *Pagination) ([]*Blog, error)
	UpdateBlog(c context.Context, blog *Blog, blog_id string) error
	DeleteBlog(c context.Context, id string) error
	LikeBlog(c context.Context, blogID string, userID string) error
	UnlikeBlog(c context.Context, blogID string, userID string) error
	CommentBlog(c context.Context, blogID string, comment *Comment) error
}

type BlogRepository interface {
	CreateBlog(c context.Context, blog *Blog) (Blog, error)
	GetBlog(c context.Context, id string) (*Blog, error)
	GetBlogs(c context.Context, pagination *Pagination) ([]*Blog, error)
	UpdateBlog(c context.Context, blog *Blog) error
	DeleteBlog(c context.Context, id string) error
	LikeBlog(c context.Context, blogID string, userID string) error
	UnlikeBlog(c context.Context, blogID string, userID string) error
	CommentBlog(c context.Context, blogID string, comment *Comment) error
}

type CommentUseCase interface {
	CreateComment(c context.Context, comment *Comment) error
	GetComment(c context.Context, id string) (*Comment, error)
	UpdateComment(c context.Context, comment *Comment) error
	DeleteComment(c context.Context, id string) error
}

type CommentRepository interface {
	CreateComment(c context.Context, comment *Comment) error
	GetComment(c context.Context, id string) (*Comment, error)
	UpdateComment(c context.Context, comment *Comment) error
	DeleteComment(c context.Context, id string) error
}

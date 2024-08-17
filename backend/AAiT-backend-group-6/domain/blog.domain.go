package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)



type  Blog struct{
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Author primitive.ObjectID `json:"author" bson:"author"`
	Content string `json:"content" bson:"content"`
	Title string `json:"title" bson:"title"`
	Tags []string `json:"tags" bson:"tags"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Like struct{
	ID primitive.ObjectID `json:"id" bson:"_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogID primitive.ObjectID `json:"blog_id" bson:"blog_id"`
    IsLiked bool `json:"is_liked" bson:"is_liked"`
}


type  Comment struct{
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Author User `json:"author" bson:"author"`
	BlogID primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content string `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Pagination struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
}


type BlogUseCase interface {
	CreateBlog(blog *Blog) error
	GetBlog(id string) (*Blog, error)
	GetBlogs(pagination *Pagination) ([]*Blog, error)
	UpdateBlog(blog *Blog,blog_id string) error
	DeleteBlog(id string) error
	LikeBlog(blogID string, userID string) error
	UnlikeBlog(blogID string, userID string) error
	CommentBlog(blogID string, comment *Comment) error
}

type BlogRepository interface {
	CreateBlog(blog *Blog) error
	GetBlog(id string) (*Blog, error)
	GetBlogs() ([]*Blog, error)
	UpdateBlog(blog *Blog) error
	DeleteBlog(id string) error
	LikeBlog(blogID string, userID string) error
	UnlikeBlog(blogID string, userID string) error
	CommentBlog(blogID string, comment *Comment) error
}

type CommentUseCase interface {
	CreateComment(comment *Comment) error
	GetComment(id string) (*Comment, error)
	UpdateComment(comment *Comment) error
	DeleteComment(id string) error
}

type CommentRepository interface {
	CreateComment(comment *Comment) error
	GetComment(id string) (*Comment, error)
	UpdateComment(comment *Comment) error
	DeleteComment(id string) error
}


package domain

import "time"

type Comment struct {
	ID        string    `bson:"_id" json:"id"`
	BlogID    string    `bson:"blog_id" json:"blog_id"`
	Content   string    `bson:"content" json:"content"`
	Author    string    `bson:"author" json:"author"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type CommentRepository interface {
	FindAll() ([]Comment, error)
	FindByID(id string) (*Comment, error)
	FindByBlogID(blogID string) ([]Comment, error)
	Save(comment *Comment) error
	Update(comment *Comment) error
	Delete(id string) error
}

type CommentUsecase interface {
	GetAllComments() ([]Comment, error)
	GetCommentByID(id string) (*Comment, error)
	GetCommentsByBlogID(blogID string) ([]Comment, error)
	CreateComment(comment *Comment) error
	UpdateComment(comment *Comment) error
	DeleteComment(id string) error
}
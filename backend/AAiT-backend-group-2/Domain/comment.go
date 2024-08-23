package domain

import (
	"context"
	"time"
)

type Comment struct {
	ID        string    `bson:"_id" json:"id"`
	BlogID    string    `bson:"blog_id" json:"blog_id"`
	Content   string    `bson:"content" json:"content"`
	Author    string    `bson:"author" json:"author"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type CommentRepository interface {
	FindAll(ctx context.Context) ([]Comment, error)
	FindByID(ctx context.Context, id string) (*Comment, error)
	FindByBlogID(ctx context.Context, blogID string) ([]Comment, error)
	Save(ctx context.Context, comment *Comment) error
	Update(ctx context.Context, comment *Comment) error
	Delete(ctx context.Context, id string) error
}

type CommentUsecase interface {
	GetAllComments(ctx context.Context) ([]Comment, error)
	GetCommentByID(ctx context.Context, id string) (*Comment, error)
	GetCommentsByBlogID(ctx context.Context, blogID string) ([]Comment, error)
	CreateComment(ctx context.Context, comment *Comment) error
	UpdateComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, id string) error
}

package domain

import (
	"context"
	"time"
)

// Comment represents a comment made by a user on a post.
type Comment struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	UserID    string    `json:"userId" bson:"userId"`
	PostID    string    `json:"postId" bson:"postId"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

// CommentRepository defines the methods that a comment repository must implement.
type CommentRepository interface {
	GetByID(id string) (*Comment, error)
	Create(comment *Comment) error
	Update(comment *Comment) error
	Delete(id string) error
}

type CommentUsecase interface {
	GetByID(ctx context.Context, id string) (*Comment, error)
	Create(ctx context.Context, comment *Comment) error
	Update(ctx context.Context, comment *Comment) error
	Delete(ctx context.Context, id string) error
}

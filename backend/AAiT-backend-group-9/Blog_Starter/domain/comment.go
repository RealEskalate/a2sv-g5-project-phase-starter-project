package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionComment = "comments"
)

type Comment struct {
	CommentID primitive.ObjectID `json:"comment_id" bson:"_id"`
	UserID    string             `json:"user_id" bson:"user_id"`
	BlogID    string             `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"createtimestamp" bson:"createtimestamp"`
	UpdatedAt time.Time          `json:"updatetimestamp" bson:"updatetimestamp"`
}

type CommentRequest struct {
	CommentID string `json:"comment_id"`
	UserID    string `json:"user_id"`
	BlogID    string `json:"blog_id"`
	Content   string `json:"content"`
}

type CommentRepository interface {
	GetComments(c context.Context, userID string, blogID string) ([]*Comment, error)
	GetCommentByID(c context.Context, commentID string) (*Comment, error)
	Create(c context.Context, comment *Comment) (*Comment, error)
	Delete(c context.Context, commentID string) (*Comment, error)
	Update(c context.Context, content string, commentID string) (*Comment, error)
}

type CommentUseCase interface {
	GetComments(c context.Context, userID string, blogID string) ([]*Comment, error)
	GetCommentByID(c context.Context, commentID string) (*Comment, error)
	Create(c context.Context, comment *CommentRequest) (*Comment, error)
	Delete(c context.Context, commentID string) (*Comment, error)
	Update(c context.Context, content string, commentID string) (*Comment, error)
}

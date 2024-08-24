package entities

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/forms"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionComment = "comments"
)

type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CommentUpdate struct {
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	UpdatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type CommentRepository interface {
	GetComments(c context.Context, blogID string, limit int64, page int64) ([]Comment, mongopagination.PaginationData, error)
	CreateComment(c context.Context, comment *Comment) (Comment, error)
	GetComment(c context.Context, commentID string) (Comment, error)
	UpdateComment(c context.Context, commentID string, updatedComment *CommentUpdate) (Comment, error)
	DeleteComment(c context.Context, commentID string) error
}

type CommentUsecase interface {
	GetComments(c context.Context, blogID string, limit int64, page int64) ([]Comment, mongopagination.PaginationData, error)
	CreateComment(c context.Context, userID string, blogID string, comment *forms.CommentForm) (Comment, error)
	GetComment(c context.Context, commentID string) (Comment, error)
	UpdateComment(c context.Context, commentID string, updatedComment *CommentUpdate) (Comment, error)
	DeleteComment(c context.Context, commentID string) error
}

type CommentErrors struct {
	Err error
}

package repository_interface


import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"AAIT-backend-group-3/internal/domain/models"
)

type CommentRepositoryInterface interface {
	CreateComment(ctx context.Context, comment *models.Comment) error
	GetCommentByID(ctx context.Context, commentID primitive.ObjectID) (*models.Comment, error)
	EditComment(ctx context.Context,commentID primitive.ObjectID, newComment *models.Comment ) error
	DeleteComment(ctx context.Context, commentID primitive.ObjectID) error
}
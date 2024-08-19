package repository_interface


import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"AAIT-backend-group-3/internal/domain/models"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *models.Comment) error
	EditComment(ctx context.Context,commentID primitive.ObjectID, newComment *models.Comment ) error
	DeleteComment(ctx context.Context, commentID primitive.ObjectID) error
}
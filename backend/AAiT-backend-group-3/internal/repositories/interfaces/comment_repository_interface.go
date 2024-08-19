package repository_interface

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"AAIT-backend-group-3/internal/domain/models"
)
type CommentRepositoryInterface interface {
	CreateComment(comment *models.Comment) error
	GetCommentByID(commentID primitive.ObjectID) (*models.Comment, error)
	EditComment(commentID primitive.ObjectID, newComment *models.Comment ) error
	DeleteComment(commentID primitive.ObjectID) error
	GetCommentsByIDList(commentIDs []primitive.ObjectID) ([]*models.Comment, error)
	GetCommentByAuthorID(authorID primitive.ObjectID) ([]*models.Comment, error)
}
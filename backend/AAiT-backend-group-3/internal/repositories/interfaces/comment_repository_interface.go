package repository_interface

import (
	"AAIT-backend-group-3/internal/domain/models"
)
type CommentRepositoryInterface interface {
	CreateComment(comment *models.Comment, blogId string) (string, error)
	GetCommentByID(commentID string) (*models.Comment, error)
	EditComment(commentID string, newComment *models.Comment ) error
	DeleteComment(commentID string) error
	GetCommentsByIDList(commentIDs []string) ([]*models.Comment, error)
	GetCommentByAuthorID(authorID string) ([]*models.Comment, error)
	DeleteCommentByID(commentID string) error
}

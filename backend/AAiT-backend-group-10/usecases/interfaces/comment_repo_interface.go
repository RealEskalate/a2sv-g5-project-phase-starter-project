package interfaces

import (
	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

type CommentRepositoryInterface interface {
	GetCommentByID(commentID uuid.UUID) (domain.Comment, error)
	GetComments(blogID uuid.UUID) ([]domain.Comment, error)
	GetCommentsCount(blogID uuid.UUID) (int, error)
	AddComment(comment domain.Comment) error
	UpdateComment(updatedComment domain.Comment) error
	DelelteComment(commentID uuid.UUID) error
}

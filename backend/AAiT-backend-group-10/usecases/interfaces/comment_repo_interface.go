package interfaces

import (
	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

type CommentRepositoryInterface interface {
	GetComments(blogID uuid.UUID) ([]domain.Comment, error)
	GetCommentsCount(blogID uuid.UUID) (int, error)
	AddComment(comment domain.Comment) error
	UpdateComment(commentID uuid.UUID, updatedComment domain.Comment) error
	DelelteComment(commentID uuid.UUID) error
}

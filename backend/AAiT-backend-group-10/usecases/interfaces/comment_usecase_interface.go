package interfaces

import (
	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

type CommentUsecaseInterface interface {
	GetComments(blogID uuid.UUID) ([]domain.Comment, error)
	AddComment(comment domain.Comment) error
	UpdateComment(updatedComment domain.Comment) error
	DelelteComment(blogID uuid.UUID, userID uuid.UUID) error
}

package interfaces

import (
	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

type CommentRepositoryInterface interface {
	GetCommentByID(commentID uuid.UUID) (domain.Comment, *domain.CustomError)
	GetComments(blogID uuid.UUID) ([]domain.Comment, *domain.CustomError)
	GetCommentsCount(blogID uuid.UUID) (int, *domain.CustomError)
	AddComment(comment domain.Comment) *domain.CustomError
	UpdateComment(updatedComment domain.Comment) *domain.CustomError
	DelelteComment(commentID uuid.UUID) *domain.CustomError
}

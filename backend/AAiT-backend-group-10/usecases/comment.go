package usecases

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type CommentUsecase struct {
	CommentRepository interfaces.CommentRepositoryInterface
}

func NewCommentUsecase(cr interfaces.CommentRepositoryInterface) interfaces.CommentUsecaseInterface {
	return &CommentUsecase{
		CommentRepository: cr,
	}
}

// AddComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) AddComment(comment domain.Comment) error {
	return cu.CommentRepository.AddComment(comment)
}

// DelelteComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) DelelteComment(blogID uuid.UUID, userID uuid.UUID) error {
	return cu.CommentRepository.DelelteComment(blogID, userID)
}

// GetComments implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) GetComments(blogID uuid.UUID) ([]domain.Comment, error) {
	return cu.CommentRepository.GetComments(blogID)
}

// UpdateComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) UpdateComment(updatedComment domain.Comment) error {
	return cu.CommentRepository.UpdateComment(updatedComment)
}

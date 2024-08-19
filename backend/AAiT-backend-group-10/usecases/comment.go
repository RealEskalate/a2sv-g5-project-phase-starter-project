package usecases

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type CommentUsecase struct {
	CommentRepository interfaces.CommentRepositoryInterface
}

type CommentUsecaseInterface interface {
	GetComments(blogID uuid.UUID) ([]domain.Comment, error)
	GetCommentsCount(blogID uuid.UUID) (int, error)
	AddComment(comment domain.Comment) error
	UpdateComment(commentID uuid.UUID, updatedComment domain.Comment) error
	DelelteComment(commentID uuid.UUID) error
}

func NewCommentUsecase(cr interfaces.CommentRepositoryInterface) CommentUsecaseInterface {
	return &CommentUsecase{
		CommentRepository: cr,
	}
}

// AddComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) AddComment(comment domain.Comment) error {
	return cu.CommentRepository.AddComment(comment)
}

// DelelteComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) DelelteComment(commentID uuid.UUID) error {
	return cu.CommentRepository.DelelteComment(commentID)
}

// GetComments implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) GetComments(blogID uuid.UUID) ([]domain.Comment, error) {
	return cu.CommentRepository.GetComments(blogID)
}

func (cu *CommentUsecase) GetCommentsCount(blogID uuid.UUID) (int, error) {
	return cu.CommentRepository.GetCommentsCount(blogID)
}

// UpdateComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) UpdateComment(commentID uuid.UUID, updatedComment domain.Comment) error {
	return cu.CommentRepository.UpdateComment(commentID, updatedComment)
}

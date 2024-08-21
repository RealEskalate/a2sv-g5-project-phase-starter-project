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
	GetCommentByID(commentID uuid.UUID) (domain.Comment, *domain.CustomError)
	GetComments(blogID uuid.UUID) ([]domain.Comment, *domain.CustomError)
	GetCommentsCount(blogID uuid.UUID) (int, *domain.CustomError)
	AddComment(comment domain.Comment) *domain.CustomError
	UpdateComment(requester_id uuid.UUID, updatedComment domain.Comment) *domain.CustomError
	DelelteComment(commentID uuid.UUID, requesterID uuid.UUID, isAdmin bool) *domain.CustomError
}

func NewCommentUsecase(cr interfaces.CommentRepositoryInterface) CommentUsecaseInterface {
	return &CommentUsecase{
		CommentRepository: cr,
	}
}

// GetCommentByID implements CommentUsecaseInterface.
func (cu *CommentUsecase) GetCommentByID(commentID uuid.UUID) (domain.Comment, *domain.CustomError) {
	return cu.CommentRepository.GetCommentByID(commentID)
}

// AddComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) AddComment(comment domain.Comment) *domain.CustomError {
	return cu.CommentRepository.AddComment(comment)
}

// DelelteComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) DelelteComment(commentID uuid.UUID, requesterID uuid.UUID, isAdmin bool) *domain.CustomError {
	originalComment, err := cu.CommentRepository.GetCommentByID(commentID)
	if err != nil {
		return err
	}
	if originalComment.UserID != requesterID && !isAdmin{
		return domain.ErrUnAuthorized
	}
	return cu.CommentRepository.DelelteComment(commentID)
}

// GetComments implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) GetComments(blogID uuid.UUID) ([]domain.Comment, *domain.CustomError) {
	return cu.CommentRepository.GetComments(blogID)
}

func (cu *CommentUsecase) GetCommentsCount(blogID uuid.UUID) (int, *domain.CustomError) {
	return cu.CommentRepository.GetCommentsCount(blogID)
}

// UpdateComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) UpdateComment(requester_id uuid.UUID, updatedComment domain.Comment) *domain.CustomError {
	originalComment, err := cu.CommentRepository.GetCommentByID(updatedComment.ID)
	if err != nil {
		return err
	}
	if originalComment.UserID != requester_id{
		return domain.ErrUnAuthorized
	}
	return cu.CommentRepository.UpdateComment(updatedComment)
}

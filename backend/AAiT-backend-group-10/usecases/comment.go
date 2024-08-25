package usecases

import (
	"encoding/json"
	"fmt"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type CommentUsecaseInterface interface {
	GetCommentByID(commentID uuid.UUID) (*dto.CommentDto, *domain.CustomError)
	GetComments(blogID uuid.UUID) ([]*dto.CommentDto, *domain.CustomError)
	AddComment(comment domain.Comment) *domain.CustomError
	UpdateComment(requester_id uuid.UUID, updatedComment domain.Comment) *domain.CustomError
	DeleteComment(commentID uuid.UUID, requesterID uuid.UUID, isAdmin bool) *domain.CustomError
}

type CommentUsecase struct {
	CommentRepository interfaces.CommentRepositoryInterface
	userRepo          interfaces.IUserRepository
	CacheRepo         interfaces.CacheRepoInterface
}

func NewCommentUsecase(cr interfaces.CommentRepositoryInterface, ur interfaces.IUserRepository, cacheRepo interfaces.CacheRepoInterface) *CommentUsecase {
	return &CommentUsecase{
		CommentRepository: cr,
		userRepo:          ur,
		CacheRepo:         cacheRepo,
	}
}

// GetCommentByID implements CommentUsecaseInterface.
func (cu *CommentUsecase) GetCommentByID(commentID uuid.UUID) (*dto.CommentDto, *domain.CustomError) {
	comment, err := cu.CommentRepository.GetCommentByID(commentID)
	if err != nil {
		return nil, err
	}
	commenter, err := cu.userRepo.GetUserByID(comment.CommenterID)
	if err != nil {
		return nil, err
	}
	return dto.NewCommentDto(comment, commenter.FullName), nil
}

// AddComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) AddComment(comment domain.Comment) *domain.CustomError {
	comment.ID = uuid.New()
	err := cu.CommentRepository.AddComment(comment)
	if err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("Comments:%s", comment.BlogID)
	_ = cu.CacheRepo.Delete(cacheKey)

	return nil
}

// DeleteComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) DeleteComment(commentID uuid.UUID, requesterID uuid.UUID, isAdmin bool) *domain.CustomError {
	originalComment, err := cu.CommentRepository.GetCommentByID(commentID)
	if err != nil {
		return err
	}
	if originalComment.CommenterID != requesterID && !isAdmin {
		return domain.ErrUnAuthorized
	}
	err = cu.CommentRepository.DeleteComment(commentID)
	if err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("Comments:%s", originalComment.BlogID)
	_ = cu.CacheRepo.Delete(cacheKey)

	return nil
}

// GetComments implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) GetComments(blogID uuid.UUID) ([]*dto.CommentDto, *domain.CustomError) {
	cacheKey := fmt.Sprintf("Comments:%s", blogID)
	cachedComments, err := cu.CacheRepo.Get(cacheKey)
	if err == nil && cachedComments != "" {
		var commentDtos []*dto.CommentDto
		err := json.Unmarshal([]byte(cachedComments), &commentDtos)
		if err == nil {
			return commentDtos, nil
		}
	}

	comments, err := cu.CommentRepository.GetComments(blogID)
	if err != nil {
		return nil, err
	}

	changedComments := make([]*dto.CommentDto, len(comments))
	for i, comment := range comments {
		commenter, err := cu.userRepo.GetUserByID(comment.CommenterID)
		if err != nil {
			return nil, err
		}
		changedComments[i] = dto.NewCommentDto(comment, commenter.FullName)
	}

	commentsJson, cerr := json.Marshal(changedComments)
	if cerr == nil {
		_ = cu.CacheRepo.Set(cacheKey, string(commentsJson), 0)
	}

	return changedComments, nil
}

// UpdateComment implements interfaces.CommentUsecaseInterface.
func (cu *CommentUsecase) UpdateComment(requester_id uuid.UUID, updatedComment domain.Comment) *domain.CustomError {
	originalComment, err := cu.CommentRepository.GetCommentByID(updatedComment.ID)
	if err != nil {
		return err
	}
	if originalComment.CommenterID != requester_id {
		return domain.ErrUnAuthorized
	}
	err = cu.CommentRepository.UpdateComment(updatedComment)
	if err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("Comments:%s", updatedComment.BlogID)
	_ = cu.CacheRepo.Delete(cacheKey)
	return nil
}

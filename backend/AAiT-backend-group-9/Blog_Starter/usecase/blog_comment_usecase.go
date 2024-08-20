package usecase

import (
	"Blog_Starter/domain"
	"context"
	"time"
)

type BlogCommentUseCase struct {
	BlogRepository    domain.BlogRepository
	CommentRepository domain.CommentRepository
}

func NewCommentUseCase(commentRepository domain.CommentRepository, blogRepository domain.BlogRepository) domain.CommentUseCase {
	return &BlogCommentUseCase{
		BlogRepository:    blogRepository,
		CommentRepository: commentRepository,
	}
}

// Create implements domain.CommentUseCase.
func (bc *BlogCommentUseCase) Create(ctx context.Context, comment *domain.CommentRequest) (*domain.Comment, error) {
	formattedComment := &domain.Comment {
		UserID: comment.UserID,
		BlogID: comment.BlogID,
		Content: comment.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	createdComment, err := bc.CommentRepository.Create(ctx, formattedComment)
	if err != nil {
		return nil, err
	}
	err = bc.BlogRepository.UpdateCommentCount(ctx,  createdComment.BlogID, true)
	return createdComment, err
}

// Delete implements domain.CommentUseCase.
func (bc *BlogCommentUseCase) Delete(ctx context.Context, commentID string) (*domain.Comment, error) {
	deletedComment, err := bc.CommentRepository.Delete(ctx, commentID)
	if err != nil {
		return nil, err
	}
	err = bc.BlogRepository.UpdateCommentCount(ctx,  deletedComment.BlogID, false)
	return deletedComment, err
}

// GetCommentByID implements domain.CommentUseCase.
func (bc *BlogCommentUseCase) GetCommentByID(ctx context.Context, commentID string) (*domain.Comment, error) {
	returnedComment, err := bc.CommentRepository.GetCommentByID(ctx, commentID)
	return returnedComment, err
}

// GetComments implements domain.CommentUseCase.
func (bc *BlogCommentUseCase) GetComments(ctx context.Context, userID string, blogID string) ([]*domain.Comment, error) {
	returnedComments, err := bc.CommentRepository.GetComments(ctx, userID, blogID)
	return returnedComments, err
}

// Update implements domain.CommentUseCase.
func (bc *BlogCommentUseCase) Update(ctx context.Context, content string, commentID string) (*domain.Comment, error) {
	updatedComment, err := bc.CommentRepository.Update(ctx,content,commentID)
	return updatedComment, err
}

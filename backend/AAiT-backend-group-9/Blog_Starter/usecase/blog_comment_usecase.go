package usecase

import (
	"Blog_Starter/domain"
	"context"
	"errors"
	"time"
)

type BlogCommentUseCase struct {
	BlogRepository    domain.BlogRepository
	CommentRepository domain.CommentRepository
	timeout					time.Duration
}

func NewCommentUseCase(commentRepository domain.CommentRepository, blogRepository domain.BlogRepository, timeout time.Duration) domain.CommentUseCase {
	return &BlogCommentUseCase{
		BlogRepository:    blogRepository,
		CommentRepository: commentRepository,
		timeout : timeout,
	}
}

// Create implements domain.CommentUseCase.
func (bc *BlogCommentUseCase) Create(c context.Context, comment *domain.CommentRequest) (*domain.Comment, error) {

	ctx, cancel := context.WithTimeout(c, bc.timeout)
	defer cancel()

	if len(comment.Content) < 10 {
		return nil, errors.New("comment content too short")
	}

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
func (bc *BlogCommentUseCase) Delete(c context.Context, commentID string) (*domain.Comment, error) {

	ctx, cancel := context.WithTimeout(c, bc.timeout)
	defer cancel()

	deletedComment, err := bc.CommentRepository.Delete(ctx, commentID)
	if err != nil {
		return nil, err
	}
	err = bc.BlogRepository.UpdateCommentCount(ctx,  deletedComment.BlogID, false)
	return deletedComment, err
}

// GetCommentByID implements domain.CommentUseCase.
func (bc *BlogCommentUseCase) GetCommentByID(c context.Context, commentID string) (*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(c, bc.timeout)
	defer cancel()

	returnedComment, err := bc.CommentRepository.GetCommentByID(ctx, commentID)
	return returnedComment, err
}

// GetComments implements domain.CommentUseCase.
func (bc *BlogCommentUseCase) GetComments(c context.Context, userID string, blogID string) ([]*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(c, bc.timeout)
	defer cancel()

	returnedComments, err := bc.CommentRepository.GetComments(ctx, userID, blogID)
	return returnedComments, err
}

// Update implements domain.CommentUseCase.
func (bc *BlogCommentUseCase) Update(c context.Context, content string, commentID string) (*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(c, bc.timeout)
	defer cancel()

	if len(content) < 10 {
		return nil, errors.New("content too short")
	}
	updatedComment, err := bc.CommentRepository.Update(ctx,content,commentID)
	return updatedComment, err
}

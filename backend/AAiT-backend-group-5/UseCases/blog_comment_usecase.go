package usecases

import (
	// config "github.com/aait.backend.g5.main/backend/Config"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

import (
	"context"
)

type CommentUsecase struct {
	repository interfaces.BlogCommentRepository
	blogRepo   interfaces.BlogRepository
	cacheService interfaces.RedisCache
}

func NewCommentUsecase(
	repository interfaces.BlogCommentRepository,
	blogRepo interfaces.BlogRepository,
	cacheService interfaces.RedisCache,
) interfaces.BlogCommentUsecase {
	return &CommentUsecase{
		repository: repository,
		blogRepo:   blogRepo,
		cacheService: cacheService,
	}
}

func (uc *CommentUsecase) AddComment(ctx context.Context, comment models.Comment) *models.ErrorResponse {
	_, err := uc.blogRepo.GetBlog(ctx, comment.BlogID)
	if err != nil {
		return err
	}

	if err := uc.repository.AddComment(ctx, comment); err != nil {
		return err
	}

	return nil
}

func (uc *CommentUsecase) GetComments(ctx context.Context, blogID string) ([]models.Comment, *models.ErrorResponse) {
	_, err := uc.blogRepo.GetBlog(ctx, blogID)
	if err != nil {
		return nil, err
	}

	comments, tErr := uc.repository.GetComments(ctx, blogID)
	if tErr != nil {
		return nil, tErr
	}

	return comments, nil
}

func (uc *CommentUsecase) GetComment(ctx context.Context, commentID string) (*models.Comment, *models.ErrorResponse) {
	comment, err := uc.repository.GetComment(ctx, commentID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (uc *CommentUsecase) UpdateComment(ctx context.Context, commentID string, userID string, comment dtos.CommentUpdateRequest) *models.ErrorResponse {
	existComment, err := uc.repository.GetComment(ctx, commentID)
	if err != nil {
		return err
	}

	if existComment.UserID != userID {
		return models.Unauthorized("You are not authorized to update this comment")
	}

	return uc.repository.UpdateComment(ctx, commentID, comment)
}

func (uc *CommentUsecase) DeleteComment(ctx context.Context, userID string, commentID string) *models.ErrorResponse {
	existComment, err := uc.repository.GetComment(ctx, commentID)
	if err != nil {
		return err
	}

	if existComment.UserID != userID {
		return models.Unauthorized("You are not authorized to delete this comment")
	}

	if err := uc.repository.DeleteComment(ctx, commentID); err != nil {
		return err
	}

	return nil
}

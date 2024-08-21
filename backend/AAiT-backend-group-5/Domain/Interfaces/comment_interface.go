package interfaces

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/gin-gonic/gin"
)

type BlogCommentRepository interface {
	AddComment(ctx context.Context, comment models.Comment) *models.ErrorResponse
	GetComments(ctx context.Context, blogID string) ([]models.Comment, *models.ErrorResponse)
	GetComment(ctx context.Context, commentID string) (*models.Comment, *models.ErrorResponse)
	UpdateComment(ctx context.Context, commentID string, comment dtos.CommentUpdateRequest) *models.ErrorResponse
	DeleteComment(ctx context.Context, commentID string) *models.ErrorResponse
	DeleteComments(ctx context.Context, blogID string) *models.ErrorResponse
}

type BlogCommentUsecase interface {
	AddComment(ctx context.Context, comment models.Comment) *models.ErrorResponse
	GetComments(ctx context.Context, blogID string) ([]models.Comment, *models.ErrorResponse)
	GetComment(ctx context.Context, commentID string) (*models.Comment, *models.ErrorResponse)
	UpdateComment(ctx context.Context, commentID string, userID string, comment dtos.CommentUpdateRequest) *models.ErrorResponse
	DeleteComment(ctx context.Context, commentID string, userID string) *models.ErrorResponse
}

type BlogCommentController interface {
	AddCommentController(ctx *gin.Context)
	GetCommentsController(ctx *gin.Context)
	GetCommentController(ctx *gin.Context)
	UpdateCommentController(ctx *gin.Context)
	DeleteCommentController(ctx *gin.Context)
}

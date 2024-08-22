package usecases

import (
	"blogapp/Domain"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentUsecase struct {
	commentRepository Domain.CommentRepository
	contextTimeout    time.Duration
}

func NewCommentUseCase(repo Domain.CommentRepository) *CommentUsecase {
	return &CommentUsecase{
		commentRepository: repo,
		contextTimeout:    time.Second * 10,
	}
}

func (uc *CommentUsecase) CommentOnPost(c *gin.Context, comment *Domain.Comment, objID primitive.ObjectID) (error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.commentRepository.CommentOnPost(ctx, comment, objID)
}

func (uc *CommentUsecase) GetCommentByID(c *gin.Context, id primitive.ObjectID) (*Domain.Comment, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.commentRepository.GetCommentByID(ctx, id)
}

func (uc *CommentUsecase) EditComment(c *gin.Context, id primitive.ObjectID, comment *Domain.Comment) (error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.commentRepository.EditComment(ctx, id, comment)
}

func (uc *CommentUsecase) GetUserComments(c *gin.Context, authorID primitive.ObjectID) ([]*Domain.Comment, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.commentRepository.GetUserComments(ctx, authorID)
}

func (uc *CommentUsecase) DeleteComment(c *gin.Context, id primitive.ObjectID) (error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.commentRepository.DeleteComment(ctx, id)
}

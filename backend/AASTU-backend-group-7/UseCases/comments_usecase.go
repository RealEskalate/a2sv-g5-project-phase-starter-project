package usecases

import (
	"blogapp/Domain"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentUsecase struct {
	commentRepository Domain.CommentRepository
	contextTimeout    time.Duration
}

func NewCommentUseCase(repo Domain.CommentRepository) *commentUsecase {
	return &commentUsecase{
		commentRepository: repo,
		contextTimeout:    time.Second * 10,
	}
}

func (uc *commentUsecase) CommentOnPost(c *gin.Context, comment *Domain.Comment, objID primitive.ObjectID) (error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.commentRepository.CommentOnPost(ctx, comment,objID)
}

func (uc *commentUsecase) GetCommentByID(c *gin.Context, id primitive.ObjectID) (*Domain.Comment, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.commentRepository.GetCommentByID(ctx, id)
}

func (uc *commentUsecase) EditComment(c *gin.Context, id primitive.ObjectID, comment *Domain.Comment) (error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.commentRepository.EditComment(ctx, id, comment)
}

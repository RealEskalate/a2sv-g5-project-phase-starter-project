package usecases

import (
	"blogapp/Domain"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type blogUsecase struct {
	BlogRepository Domain.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUseCase(repo Domain.BlogRepository) *blogUsecase {
	return &blogUsecase{
		BlogRepository: repo,
		contextTimeout: time.Second * 10,
	}
}

func (uc *blogUsecase) CreateBlog(c *gin.Context, post *Domain.Post) (error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.BlogRepository.CreateBlog(ctx,post)
}

func (uc *blogUsecase) GetPostBySlug(c *gin.Context, slug string) ([]*Domain.Post, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.BlogRepository.GetPostBySlug(ctx,slug)
}

func (uc *blogUsecase) GetPostByID(c *gin.Context, id primitive.ObjectID) (*Domain.Post, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.BlogRepository.GetPostByID(ctx,id)
}

func (uc *blogUsecase) GetPostByAuthorID(c *gin.Context, authorID primitive.ObjectID) ([]*Domain.Post, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.BlogRepository.GetPostByAuthorID(ctx,authorID)
}

func (uc *blogUsecase) UpdatePostByID(c *gin.Context, id primitive.ObjectID, post *Domain.Post) (error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.BlogRepository.UpdatePostByID(ctx,id,post)
}


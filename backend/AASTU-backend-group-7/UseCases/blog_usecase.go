package usecases

import (
	"blogapp/Domain"
	"context"
	"time"

	"github.com/gin-gonic/gin"
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

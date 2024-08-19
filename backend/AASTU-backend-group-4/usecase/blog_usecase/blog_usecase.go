package blog_usecase

import (
	"time"

	"blog-api/domain"
)

type BlogUsecase struct {
	repo           domain.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(BlogRepository domain.BlogRepository, timeout time.Duration) domain.BlogUsecase {
	return &BlogUsecase{
		repo:           BlogRepository,
		contextTimeout: timeout,
	}
}

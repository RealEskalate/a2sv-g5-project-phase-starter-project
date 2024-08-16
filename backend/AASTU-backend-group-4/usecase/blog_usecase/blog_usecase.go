package blog_usecase

import (
	"time"

	"blog-api/domain/blog"
)

type BlogUsecase struct {
	repo           blog.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(BlogRepository blog.BlogRepository, timeout time.Duration) blog.BlogUsecase {
	return &BlogUsecase{
		repo:           BlogRepository,
		contextTimeout: timeout,
	}
}

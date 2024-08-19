package blog_usecase

import (
	domain "blog-api/domain/blog"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (bu *BlogUsecase) GetBlog(ctx context.Context, blogID primitive.ObjectID) (*domain.Blog, error) {
	return bu.repo.GetBlog(ctx, blogID)
}

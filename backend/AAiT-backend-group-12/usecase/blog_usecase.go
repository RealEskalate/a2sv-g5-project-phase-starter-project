package usecase

import (
	"blog_api/domain"
	"context"
	"time"
)

type BlogUseCase struct {
	blogRepo       domain.BlogRepositoryInterface
	contextTimeOut time.Duration
}


func NewBlogUseCase(repo domain.BlogRepositoryInterface, t time.Duration) *BlogUseCase {
	return &BlogUseCase{
		blogRepo:       repo,
		contextTimeOut: t,
	}
}

// CreateBlogPost implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) CreateBlogPost(ctx context.Context, blog *domain.Blog) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	err := b.blogRepo.InsertBlogPost(ctx, blog)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBlogPost implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) DeleteBlogPost(ctx context.Context, blogId string) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	err := b.blogRepo.DeleteBlogPost(ctx, blogId)
	if err != nil{
		return err
	}
	return nil
}

// EditBlogPost implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) EditBlogPost(ctx context.Context, blogId string, blog *domain.Blog) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	err := b.blogRepo.UpdateBlogPost(ctx, blogId, blog)
	if err != nil{
		return err
	}
	return nil
}

// GetBlogPost implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) GetBlogPost(ctx context.Context, blogId string) (*domain.Blog, error) {
	panic("implemented by robel")
}
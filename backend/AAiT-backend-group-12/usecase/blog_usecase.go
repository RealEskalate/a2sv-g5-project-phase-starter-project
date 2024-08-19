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
func (b *BlogUseCase) CreateBlogPost(ctx context.Context, blog *domain.Blog) domain.CodedError {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	blog.CreatedAt = time.Now()
	defer cancel()

	err := b.blogRepo.InsertBlogPost(ctx, blog)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBlogPost implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) DeleteBlogPost(ctx context.Context, blogId string) domain.CodedError {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	err := b.blogRepo.DeleteBlogPost(ctx, blogId)
	if err != nil{
		return err
	}
	return nil
}

// EditBlogPost implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) EditBlogPost(ctx context.Context, blogId string, blog *domain.Blog) domain.CodedError {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	blog.UpdatedAt = time.Now()

	err := b.blogRepo.UpdateBlogPost(ctx, blogId, blog)
	if err != nil{
		return err
	}
	return nil
}

// Fetches all blogs
func (b *BlogUseCase) GetBlogPosts(ctx context.Context, filters domain.BlogFilterOptions) ([]domain.Blog, int, domain.CodedError) {
	context, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	// Set default pagination if not provided
	if filters.Page <= 0 {
		filters.Page = 1
	}
	if filters.PostsPerPage <= 0 {
		filters.PostsPerPage = 10 // Default to 10 posts per page
	}

	// Set default sorting if not provided
	if filters.SortBy == "" {
		filters.SortBy = "created_at" // Default sort by creation date
		filters.SortDirection = "desc"
	}

	return b.blogRepo.FetchBlogPosts(context, filters)
}

// FetchBlogPostByID retrieves a single blog post by its ID and increments its view count.
func (b *BlogUseCase) GetBlogPostByID(ctx context.Context, blogID string) (*domain.Blog, domain.CodedError) {
	context, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	return b.blogRepo.FetchBlogPostByID(context, blogID)
}

func (b *BlogUseCase) TrackBlogPopularity(ctx context.Context, blogId string, action string, username string) domain.CodedError {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	return b.blogRepo.TrackBlogPopularity(ctx, blogId, action, username)
}
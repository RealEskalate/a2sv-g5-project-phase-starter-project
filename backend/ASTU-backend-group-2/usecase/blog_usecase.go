package usecase

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
)

type blogUsecase struct {
	blogRepository domain.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(blogRepository domain.BlogRepository, timeout time.Duration) domain.BlogUsecase {
	return &blogUsecase{
		blogRepository: blogRepository,
		contextTimeout: timeout,
	}
}

func (b *blogUsecase) GetByTags(c context.Context, tags []string) ([]domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, err := b.blogRepository.GetByTags(ctx, tags)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (b *blogUsecase) GetAllBlogs(c context.Context) ([]domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, err := b.blogRepository.GetAllBlogs(ctx)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (b *blogUsecase) GetBlogByID(c context.Context, blogID string) (domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blog, err := b.blogRepository.GetBlogByID(ctx, blogID)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (b *blogUsecase) GetByPopularity(c context.Context) ([]domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, err := b.blogRepository.GetByPopularity(ctx)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (b *blogUsecase) Search(c context.Context, searchTerm string) ([]domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, err := b.blogRepository.Search(ctx, searchTerm)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (b *blogUsecase) CreateBlog(c context.Context, newBlog *domain.Blog) (domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blog, err := b.blogRepository.CreateBlog(ctx, newBlog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (b *blogUsecase) UpdateBlog(c context.Context, blogID string, updatedBlog *domain.Blog) (domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blog, err := b.blogRepository.UpdateBlog(ctx, blogID, updatedBlog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (b *blogUsecase) DeleteBlog(c context.Context, blogID string) error {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	err := b.blogRepository.DeleteBlog(ctx, blogID)
	if err != nil {
		return err
	}

	return nil
}

func (b *blogUsecase) SortByDate(c context.Context) ([]domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, err := b.blogRepository.SortByDate(ctx)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

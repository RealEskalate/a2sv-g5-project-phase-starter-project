package usecase

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	mongopagination "github.com/gobeam/mongo-go-pagination"
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

func (b *blogUsecase) GetByTags(c context.Context, tags []string, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.GetByTags(ctx, tags, limit, page)
	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
}

func (b *blogUsecase) GetAllBlogs(c context.Context, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.GetAllBlogs(ctx, limit, page)

	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
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

func (b *blogUsecase) GetByPopularity(c context.Context, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.GetByPopularity(ctx, limit, page)
	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
}

func (b *blogUsecase) Search(c context.Context, searchTerm string, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.Search(ctx, searchTerm, limit, page)
	if err != nil {
		return nil, meta, err
	}

	return blogs, meta, nil
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

func (b *blogUsecase) UpdateBlog(c context.Context, blogID string, updatedBlog *domain.BlogUpdate) (domain.Blog, error) {
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

func (b *blogUsecase) SortByDate(c context.Context, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.SortByDate(ctx, limit, page)
	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
}

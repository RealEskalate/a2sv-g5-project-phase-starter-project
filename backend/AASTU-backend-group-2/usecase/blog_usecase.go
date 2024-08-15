package usecase

import (
	"blog_g2/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUsecase struct {
	BlogRepo       domain.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(blogrepo domain.BlogRepository, timeout time.Duration) domain.BlogUsecase {
	return &BlogUsecase{
		BlogRepo:       blogrepo,
		contextTimeout: timeout,
	}

}

func (br *BlogUsecase) CreateBlog(c context.Context, blog domain.Blog) error {
	return nil
}

func (br *BlogUsecase) RetrieveBlog(c context.Context) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *BlogUsecase) UpdateBlog(c context.Context, updatedblog domain.Blog) error {
	return nil
}

func (br *BlogUsecase) DeleteBlog(c context.Context, blogID primitive.ObjectID) error {
	return nil
}

func (br *BlogUsecase) SearchBlog(c context.Context, name string, author string) ([]domain.Blog, error) {
	_, cancel := context.WithTimeout(c, br.contextTimeout)
	defer cancel()
	return br.BlogRepo.SearchBlog(name, author)
}

func (br *BlogUsecase) FilterBlog(c context.Context, tags []string, date time.Time) ([]domain.Blog, error) {
	_, cancel := context.WithTimeout(c, br.contextTimeout)
	defer cancel()
	return br.BlogRepo.FilterBlog(tags, date)
}

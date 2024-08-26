package usecase

import (
	"blog_g2/domain"
	"context"

	"time"
)

type BlogUsecase struct {
	BlogRepo       domain.BlogRepository
	contextTimeout time.Duration
	Aiservice      domain.AIService
}

func NewBlogUsecase(blogrepo domain.BlogRepository, aiserv domain.AIService, timeout time.Duration) domain.BlogUsecase {
	return &BlogUsecase{
		BlogRepo:       blogrepo,
		contextTimeout: timeout,
		Aiservice:      aiserv,
	}

}

func (br *BlogUsecase) CreateBlog(c context.Context, blog *domain.Blog) error {
	_, cancel := context.WithTimeout(c, br.contextTimeout)
	defer cancel()
	if err := br.Aiservice.Validate_Blog(blog.Content); err != nil {
		return err
	}

	return br.BlogRepo.CreateBlog(blog)
}

func (br *BlogUsecase) RetrieveBlog(c context.Context, page int, sortby string, dire string) ([]domain.Blog, error) {
	_, cancel := context.WithTimeout(c, br.contextTimeout)
	defer cancel()

	return br.BlogRepo.RetrieveBlog(page, sortby, dire)
}

func (br *BlogUsecase) UpdateBlog(c context.Context, updatedblog domain.Blog, blogID string, isadmin bool, userid string) error {
	_, cancel := context.WithTimeout(c, br.contextTimeout)
	defer cancel()
	return br.BlogRepo.UpdateBlog(updatedblog, blogID, isadmin, userid)
}

func (br *BlogUsecase) DeleteBlog(c context.Context, blogID string, isadmin bool, userid string) error {
	_, cancel := context.WithTimeout(c, br.contextTimeout)
	defer cancel()
	return br.BlogRepo.DeleteBlog(blogID, isadmin, userid)

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

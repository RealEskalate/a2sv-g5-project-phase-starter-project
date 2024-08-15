package repositories

import (
	"blog_g2/domain"
	"context"
)

type BlogUsecase struct {
	BlogRepo domain.BlogRepository
}

func NewBlogUsecase(blogrepo domain.BlogRepository) domain.BlogUsecase {
	return &BlogUsecase{
		BlogRepo: blogrepo,
	}

}

func (br *BlogUsecase) CreateBlog(c context.Context, blog domain.Blog) error {
	return nil
}

func (br *BlogUsecase) RetrieveBlog(c context.Context) ([]domain.Blog, error) {

	return []domain.Blog{}, nil
}

func (br *BlogUsecase) UpdateBlog(c context.Context, updatedblog domain.Blog, blogID string) error {
	err := br.BlogRepo.UpdateBlog(updatedblog, blogID)
	if err != nil {
		return err
	}

	return nil
}

func (br *BlogUsecase) DeleteBlog(c context.Context, blogID string) error {
	err := br.BlogRepo.DeleteBlog(blogID)
	if err != nil {
		return err
	}

	return nil
}

func (br *BlogUsecase) SearchBlog(c context.Context) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *BlogUsecase) FilterBlog(c context.Context) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

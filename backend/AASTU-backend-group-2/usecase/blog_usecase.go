package repositories

import (
	"blog_g2/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	br.BlogRepo.CreateBlog(blog)
	return nil
}

func (br *BlogUsecase) RetrieveBlog(c context.Context, page int) ([]domain.Blog, error) {
	return br.BlogRepo.RetrieveBlog(page)
}

func (br *BlogUsecase) UpdateBlog(c context.Context, updatedblog domain.Blog) error {
	return nil
}

func (br *BlogUsecase) DeleteBlog(c context.Context, blogID primitive.ObjectID) error {
	return nil
}

func (br *BlogUsecase) SearchBlog(c context.Context) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *BlogUsecase) FilterBlog(c context.Context) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

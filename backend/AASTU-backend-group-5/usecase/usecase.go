package usecase

import (
	"github.com/RealEskalate/blogpost/domain"
)

type BlogUseCase struct {
	BlogRepo domain.BlogRepository
}

func (uc *BlogUseCase) CreateBlog(blog domain.Blog) ([]domain.Blog, error) {
	return uc.BlogRepo.CreateBlog(blog)
}
func (uc *BlogUseCase) GetOneBlog(id string) ([]domain.Blog, error) {
	return uc.BlogRepo.GetOneBlog(id)
}
func (uc *BlogUseCase) UpdateBlog(id string, blog domain.Blog) ([]domain.Blog, error) {
	return uc.BlogRepo.UpdateBlog(id)
}
func (uc *BlogUseCase) DeleteBlog(id string) error {
	return uc.BlogRepo.DeleteBlog(id)
}
func (uc *BlogUseCase) FilterBlog(map[string]string) ([]domain.Blog, error) {
	return uc.BlogRepo.FilterBlog(map[string]string)
}

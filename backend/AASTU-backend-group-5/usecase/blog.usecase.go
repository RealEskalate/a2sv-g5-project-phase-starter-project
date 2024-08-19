package usecase

import (
	"github.com/RealEskalate/blogpost/domain"
)

type BlogUseCase struct {
	BlogRepo domain.Blog_Rerpository_interface
}

func (uc *BlogUseCase) CreateBlog(blog domain.Blog) (*domain.Blog, error) {
	createdBlog, err := uc.BlogRepo.CreateBlogDocunent(blog)
	if err != nil {
		return nil, err
	}
	return &createdBlog, nil
}
func (uc *BlogUseCase) GetOneBlog(id string) ([]domain.Blog, error) {
	return uc.BlogRepo.GetOneBlogDocunent(id)
}
func (uc *BlogUseCase) UpdateBlog(id string, updatedBlog domain.Blog) (domain.Blog, error) {
	return uc.BlogRepo.UpdateBlogDocunent(id, updatedBlog)
}

func (uc *BlogUseCase) DeleteBlog(id string) error {
	return uc.BlogRepo.DeleteBlogDocunent(id)
}
func (uc *BlogUseCase) FilterBlog(ans map[string]string) ([]domain.Blog, error) {
	return uc.BlogRepo.FilterBlogDocunent(ans)
}

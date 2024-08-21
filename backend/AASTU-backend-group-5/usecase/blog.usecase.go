package usecase

import (
	"time"

	"github.com/RealEskalate/blogpost/domain"
)

type BlogUseCase struct {
	BlogRepo domain.Blog_Repository_interface
}

func NewBlogUsecase(repo domain.Blog_Repository_interface) *BlogUseCase {
	return &BlogUseCase{
		BlogRepo: repo,
	}
}

func (uc *BlogUseCase) CreateBlog(iblog domain.PostBlog) (domain.Blog, error) {
	blog := domain.Blog{
		Title:     iblog.Title,
		Content:   iblog.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Tag:       iblog.Tag,
	}
	createdBlog, err := uc.BlogRepo.CreateBlogDocunent(blog)
	if err != nil {
		return domain.Blog{}, err
	}
	return createdBlog, nil
}

func (uc *BlogUseCase) GetBlogs(limit int, page_number int) ([]domain.Blog, error) {
	blogs, err := uc.BlogRepo.GetBlogDocunents(page_number, limit)
	if err != nil {
		return []domain.Blog{}, err
	}
	return blogs, nil
}
func (uc *BlogUseCase) GetOneBlog(id string) (domain.Blog, error) {
	blogs, err := uc.BlogRepo.GetOneBlogDocunent(id)
	if err != nil {
		return domain.Blog{}, err
	}
	return blogs, nil
}
func (uc *BlogUseCase) UpdateBlog(id string, updatedBlog domain.Blog) (domain.Blog, error) {
	blog, err := uc.BlogRepo.UpdateBlogDocunent(id, updatedBlog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (uc *BlogUseCase) DeleteBlog(id string, user_id string) error {
	return uc.BlogRepo.DeleteBlogDocument(id, user_id)
}
func (uc *BlogUseCase) FilterBlog(ans map[string]string) ([]domain.Blog, error) {
	blogs, err := uc.BlogRepo.FilterBlogDocunent(ans)
	if err != nil {
		return []domain.Blog{}, err
	}
	return blogs, nil
}

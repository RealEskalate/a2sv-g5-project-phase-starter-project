package usecase

import (
	"AAiT-backend-group-8/Domain"
	"time"
)

type BlogUseCase struct {
	repo Domain.IBlogRepository
}

func NewBlogUseCase(repo Domain.IBlogRepository) *BlogUseCase {
	return &BlogUseCase{repo: repo}
}

func (uc *BlogUseCase) GetAllBlogs(page int, pageSize int, sortBy string) (*[]Domain.Blog, error) {
	blogs, err := uc.repo.FindAll(page, pageSize, sortBy)

	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (uc *BlogUseCase) CreateBlog(blog *Domain.Blog) error {
	blog.CreatedAt = time.Now()
	err := uc.repo.Create(blog)

	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUseCase) GetBlogByID(ID string) (*Domain.Blog, error) {
	blog, err := uc.repo.FindByID(ID)

	if err != nil {
		return nil, err
	}

	return blog, nil

}

func (uc *BlogUseCase) UpdateBlogViewCount(id string) error {
	err := uc.repo.UpdateViewCount(id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUseCase) DeleteBlog(ID string) error {

	err := uc.repo.Delete(ID)

	if err != nil {
		return err
	}

	return nil
}

func (uc *BlogUseCase) UpdateBlog(blog *Domain.Blog) error {
	blog.LastUpdated = time.Now()
	err := uc.repo.Update(blog)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUseCase) UpdateBlogCommentCount(id string, inc bool) error {
	err := uc.repo.UpdateCommentCount(id, inc)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUseCase) UpdateBlogLikeCount(id string, inc bool) error {
	err := uc.repo.UpdateLikeCount(id, inc)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUseCase) SearchBlog(criteria *Domain.SearchCriteria) (*[]Domain.Blog, error) {

	blogs, err := uc.repo.Search(criteria)
	if err != nil {
		return nil, err
	}
	return blogs, nil

}

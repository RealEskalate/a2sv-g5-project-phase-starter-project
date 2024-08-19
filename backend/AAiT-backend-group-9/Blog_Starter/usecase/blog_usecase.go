package usecase

import (
	"Blog_Starter/domain"
	"context"
	"errors"
)

type BlogUseCase struct {
	blogRepo domain.BlogRepository
}

func NewBlogUseCase(blogRepo domain.BlogRepository) domain.BlogUseCase {
	return &BlogUseCase{
		blogRepo: blogRepo,
	}
}

func (uc *BlogUseCase) CreateBlog(c context.Context, blog *domain.BlogCreate) (*domain.Blog, error) {
	// implementation
	blogModel := &domain.Blog{
		UserID:  blog.UserID,
		Title:   blog.Title,
		Content: blog.Content,
		Tags:    blog.Tags,
	}
	return uc.blogRepo.CreateBlog(c, blogModel)
}

func (uc *BlogUseCase) GetBlogByID(c context.Context, blogID string) (*domain.Blog, error) {
	// implementation
	return uc.blogRepo.GetBlogByID(c, blogID)
}

func (uc *BlogUseCase) GetAllBlog(c context.Context) ([]*domain.Blog, error) {
	// implementation
	return uc.blogRepo.GetAllBlog(c)
}

func (uc *BlogUseCase) UpdateBlog(c context.Context, blog *domain.BlogUpdate, blogID string) (*domain.Blog, error) {
	// implementation i want only to change the title, content and tags
	_, err := uc.blogRepo.GetBlogByID(c, blogID)
	if err != nil {
		return nil, errors.New("blog not found")
	}

	updatedBlog := &domain.BlogUpdate{
		Title:   blog.Title,
		Content: blog.Content,
		Tags:    blog.Tags,
	}

	return uc.blogRepo.UpdateBlog(c, updatedBlog, blogID)
}

func (uc *BlogUseCase) DeleteBlog(c context.Context, blogID string) error {
	// check whether the blog exists or not
	_, err := uc.blogRepo.GetBlogByID(c, blogID)
	if err != nil {
		return errors.New("blog not found")
	}
	// implementation
	return uc.blogRepo.DeleteBlog(c, blogID)
}

// SearchBlogs implements domain.BlogUseCase.

func (uc *BlogUseCase) SearchBlogs(c context.Context, title string, author string) ([]*domain.Blog, error) {
	panic("unimplemented")
}

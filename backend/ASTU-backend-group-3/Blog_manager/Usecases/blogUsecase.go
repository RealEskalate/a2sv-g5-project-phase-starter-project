package Usecases

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Repository"
)

type BlogUsecase interface {
	CreateBlog(blog *Domain.Blog) (*Domain.Blog, error)
	RetrieveBlogs(page, pageSize int, sortBy string) ([]Domain.Blog, int64, error)
	DeleteBlogByID(id string) error
	SearchBlogs(title string, author string, tags []string) ([]*Domain.Blog, error)
}

type blogUsecase struct {
	blogRepo Repository.BlogRepository
}

func NewBlogUsecase(br Repository.BlogRepository) BlogUsecase {
	return &blogUsecase{
		blogRepo: br,
	}
}

func (uc *blogUsecase) CreateBlog(blog *Domain.Blog) (*Domain.Blog, error) {
	// Validate the blog details
	if err := blog.Validate(); err != nil {
		return nil, err
	}

	// Insert the blog into the database
	createdBlog, err := uc.blogRepo.Save(blog)
	if err != nil {
		return nil, err
	}

	return createdBlog, nil
}

func (bu *blogUsecase) RetrieveBlogs(page, pageSize int, sortBy string) ([]Domain.Blog, int64, error) {
	blogs, totalPosts, err := bu.blogRepo.RetrieveBlogs(page, pageSize, sortBy)
	if err != nil {
		return nil, 0, err
	}

	return blogs, totalPosts, nil
}
func (uc *blogUsecase) DeleteBlogByID(id string) error {
	return uc.blogRepo.DeleteBlogByID(id)
}

func (uc *blogUsecase) SearchBlogs(title string, author string, tags []string) ([]*Domain.Blog, error) {
	return uc.blogRepo.SearchBlogs(title, author, tags)
}

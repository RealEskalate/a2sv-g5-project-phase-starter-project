package Usecases

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Repository"
)

type BlogUsecase interface {
	CreateBlog(blog *Domain.Blog) (*Domain.Blog, error)
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

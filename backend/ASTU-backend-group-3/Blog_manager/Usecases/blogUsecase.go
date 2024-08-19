package Usecases

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Repository"
	"errors"
	"time"
)

type BlogUsecase interface {
	CreateBlog(blog *Domain.Blog) (*Domain.Blog, error)
	UpdateBlog(blogID string, input Domain.UpdateBlogInput, author string) (*Domain.Blog, error)
	RetrieveBlogs(page, pageSize int, sortBy string) ([]Domain.Blog, int64, error)
	DeleteBlogByID(id string) error
	SearchBlogs(title string, author string, tags []string) ([]Domain.Blog, error)
	FindByID(id string) (*Domain.Blog, error)
	IncrementViewCount(id string) error
	ToggleLike(blogID, username string) error
	ToggleDislike(blogID, username string) error
	AddComment(blogID string, comment Domain.Comment) error
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

func (uc *blogUsecase) SearchBlogs(title string, author string, tags []string) ([]Domain.Blog, error) {
	return uc.blogRepo.SearchBlogs(title, author, tags)
}

func (uc *blogUsecase) UpdateBlog(blogID string, input Domain.UpdateBlogInput, author string) (*Domain.Blog, error) {
	// Retrieve the existing blog
	existingBlog, err := uc.blogRepo.FindByID(blogID)
	if err != nil {
		return nil, err
	}

	// Check if the current user is the author
	if existingBlog.Author != author {
		return nil, errors.New("unauthorized: user is not the author of the blog")
	}

	// Update the existing blog with the new details
	existingBlog.Title = input.Title
	existingBlog.Content = input.Content
	existingBlog.Tags = input.Tags
	existingBlog.UpdatedAt = time.Now().Format(time.RFC3339)

	// Save the updated blog to the database
	updatedBlog, err := uc.blogRepo.Save(existingBlog)
	if err != nil {
		return nil, err
	}

	return updatedBlog, nil
}

func (uc *blogUsecase) FindByID(id string) (*Domain.Blog, error) {
	return uc.blogRepo.FindByID(id)
}

// Increment the view count of a blog post
func (u *blogUsecase) IncrementViewCount(id string) error {
	return u.blogRepo.IncrementViewCount(id)
}

// Toggle like on a blog post
func (u *blogUsecase) ToggleLike(blogID, username string) error {
	return u.blogRepo.ToggleLike(blogID, username)
}

// Toggle dislike on a blog post
func (u *blogUsecase) ToggleDislike(blogID, username string) error {
	return u.blogRepo.ToggleDislike(blogID, username)
}

// Add a comment to a blog post
func (u *blogUsecase) AddComment(blogID string, comment Domain.Comment) error {
	return u.blogRepo.AddComment(blogID, comment)
}

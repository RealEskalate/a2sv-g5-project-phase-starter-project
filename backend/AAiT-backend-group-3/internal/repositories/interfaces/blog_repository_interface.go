package repository_interface

import (
	"AAIT-backend-group-3/internal/domain/models"
)

type BlogRepositoryInterface interface {
	CreateBlog(blog *models.Blog, authorId string) (string, error)
	GetBlogByID(blogID string) (*models.Blog, error)
	GetBlogs(filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error)
	UpdateBlog(blogID string, newBlog *models.Blog ) error
	DeleteBlog(blogID string) error
	AddCommentToTheList(blogID string, commentID string) error
	DeleteCommentFromTheList(blogID string, commentID string) error
	GetBlogsByAuthorID(authorID string) ([]*models.Blog, error)
	GetBlogsByPopularity(limit int) ([]*models.Blog, error)
	LikeBlog(blogID string, userID string) error
	ViewBlog(blogID string) error
	
}

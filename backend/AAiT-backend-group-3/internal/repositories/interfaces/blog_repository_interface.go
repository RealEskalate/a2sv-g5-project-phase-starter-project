package repository_interface

import (
	"AAIT-backend-group-3/internal/domain/models"
)

type BlogRepositoryInterface interface {
	CreateBlog(blog *models.Blog, authorId string) error
	GetBlogByID(blogID string) (*models.Blog, error)
	GetBlogs(filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error)
	EditBlog(logID string, newBlog *models.Blog ) error
	DeleteBlog(blogID string) error
	AddCommentToTheList(blogID string, commentID string) error
}
package repository_interface

import (
	"AAIT-backend-group-3/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogRepositoryInterface interface {
	CreateBlog(blog *models.Blog, authorId string) (string, error)
	GetBlogByID(blogID string) (*models.Blog, error)
	GetBlogs(filter primitive.M, page int, limit int) ([]*models.Blog, error)
	UpdateBlog(blogID string, newBlog *models.Blog ) error
	DeleteBlog(blogID string) error
	AddCommentToTheList(blogID string, commentID string) error
	DeleteCommentFromTheList(blogID string, commentID string) error
	GetBlogsByAuthorID(authorID string) ([]*models.Blog, error)
	GetBlogsByPopularity(limit int) ([]*models.Blog, error)
	AddLike(blogID string, userID string) error
	RemoveLike(blogID string, userID string) error
	ViewBlog(blogID string) error
	
}

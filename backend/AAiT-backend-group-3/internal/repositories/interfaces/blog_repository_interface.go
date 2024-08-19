package repository_interface


import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"AAIT-backend-group-3/internal/domain/models"
)


type BlogRepositoryInterface interface {
	CreateBlog(ctx context.Context, blog *models.Blog) error
	GetBlogByID(ctx context.Context, blogID primitive.ObjectID) (*models.Blog, error)
	GetBlogs(ctx context.Context, filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error)
	EditBlog(ctx context.Context,blogID primitive.ObjectID, newBlog *models.Blog ) error
	DeleteBlog(ctx context.Context, blogID primitive.ObjectID) error
	AddCommentToTheList(ctx context.Context, blogID primitive.ObjectID, comment *models.Comment) error
}
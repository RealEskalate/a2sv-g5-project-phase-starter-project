package blog

import (
	"blogApp/internal/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogUseCase defines the business logic operations for managing blogs
type BlogUseCase interface {
	// Blog operations
	CreateBlog(ctx context.Context, blog *domain.Blog) error
	GetBlogByID(ctx context.Context, id primitive.ObjectID) (*domain.Blog, error)
	UpdateBlog(ctx context.Context, id primitive.ObjectID, blog *domain.Blog) error
	DeleteBlog(ctx context.Context, id primitive.ObjectID) error
	GetAllBlogs(ctx context.Context) ([]*domain.Blog, error)
	FilterBlogs(ctx context.Context, filter domain.BlogFilter) ([]*domain.Blog, error)
	PaginateBlogs(ctx context.Context, filter domain.BlogFilter, page, pageSize int) ([]*domain.Blog, error)

	// Tag operations
	AddTagToBlog(ctx context.Context, blogID primitive.ObjectID, tag domain.BlogTag) error
	RemoveTagFromBlog(ctx context.Context, blogID primitive.ObjectID, tagID primitive.ObjectID) error

	// Comment operations
	AddComment(ctx context.Context, comment *domain.Comment) error
	GetCommentsByBlogID(ctx context.Context, blogID primitive.ObjectID) ([]*domain.Comment, error)

	// Like operations
	AddLike(ctx context.Context, like *domain.Like) error
	GetLikesByBlogID(ctx context.Context, blogID primitive.ObjectID) ([]*domain.Like, error)

	// View operations
	AddView(ctx context.Context, view *domain.View) error
	GetViewsByBlogID(ctx context.Context, blogID primitive.ObjectID) ([]*domain.View, error)

	// Real-time updates(will be implemented later)
	WatchBlogs(ctx context.Context, pipeline []primitive.M) (<-chan domain.Blog, error)
	WatchBlogByID(ctx context.Context, id primitive.ObjectID) (<-chan domain.Blog, error)
}

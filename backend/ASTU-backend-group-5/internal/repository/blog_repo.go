package repository

import (
	"blogApp/internal/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogRepository interface {
	// Blog operations
	CreateBlog(ctx context.Context, blog *domain.Blog) error
	GetBlogByID(ctx context.Context, id primitive.ObjectID) (*domain.Blog, error)
	UpdateBlog(ctx context.Context, id primitive.ObjectID, blog *domain.Blog) error
	DeleteBlog(ctx context.Context, id primitive.ObjectID) error
	GetAllBlogs(ctx context.Context) ([]*domain.Blog, error)
	FilterBlogs(ctx context.Context, filter BlogFilter) ([]*domain.Blog, error)
	PaginateBlogs(ctx context.Context, filter BlogFilter, page, pageSize int) ([]*domain.Blog, error)

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

	// Real-time updates
	WatchBlogs(ctx context.Context, pipeline []primitive.M) (<-chan domain.Blog, error)
	WatchBlogByID(ctx context.Context, id primitive.ObjectID) (<-chan domain.Blog, error)
}

// BlogFilter represents criteria to filter blogs
type BlogFilter struct {
	AuthorID  *primitive.ObjectID // Filter by Author ID
	Tags      []string            // Filter by Tags
	Title     *string             // Filter by Title (exact or partial match)
	DateRange *DateRange          // Filter by Creation Date Range
	Content   *string             // Filter by Content (exact or partial match)
}

// DateRange represents a time range for filtering
type DateRange struct {
	From time.Time // Start date for the range
	To   time.Time // End date for the range
}

package blog

import (
	"blogApp/internal/domain"
	"context"
)

// BlogUseCase defines the business logic operations for managing blogs
type BlogUseCase interface {
	// Blog operations
	CreateBlog(ctx context.Context, blog *domain.Blog) error
	GetBlogByID(ctx context.Context, id string) (*domain.Blog, error)
	UpdateBlog(ctx context.Context, id string, blog *domain.Blog) error
	DeleteBlog(ctx context.Context, id string) error
	GetAllBlogs(ctx context.Context) ([]*domain.Blog, error)
	FilterBlogs(ctx context.Context, filter domain.BlogFilter) ([]*domain.Blog, error)
	PaginateBlogs(ctx context.Context, filter domain.BlogFilter, page, pageSize int) ([]*domain.Blog, error)

	// Tag operations
	AddTagToBlog(ctx context.Context, blogID string, tag domain.BlogTag) error
	RemoveTagFromBlog(ctx context.Context, blogID string, tagID string) error

	// Comment operations
	AddComment(ctx context.Context, comment *domain.Comment) error
	GetCommentsByBlogID(ctx context.Context, blogID string) ([]*domain.Comment, error)

	// Like operations
	AddLike(ctx context.Context, like *domain.Like) error
	GetLikesByBlogID(ctx context.Context, blogID string) ([]*domain.Like, error)

	// View operations
	AddView(ctx context.Context, view *domain.View) error
	GetViewsByBlogID(ctx context.Context, blogID string) ([]*domain.View, error)

	// Tag operations
	GetAllTags(ctx context.Context) ([]*domain.BlogTag, error)
	CreateTag(ctx context.Context, tag *domain.BlogTag) error
	UpdateTag(ctx context.Context, id string, tag *domain.BlogTag) error
	DeleteTag(ctx context.Context, id string) error
	GetTagByID(ctx context.Context, id string) (*domain.BlogTag, error)
}

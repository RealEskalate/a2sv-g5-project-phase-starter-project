package blog

import (
	"blogApp/internal/domain"
	"context"
)

// BlogUseCase defines the business logic operations for managing blogs
type BlogUseCase interface {
	// Blog operations
	CreateBlog(ctx context.Context, blog *domain.Blog, authorId string) error
	GetBlogByID(ctx context.Context, blogId, userId string) (*domain.GetSingleBlogDTO, error)
	UpdateBlog(ctx context.Context, id string, blog *domain.Blog, authorId string) error
	DeleteBlog(ctx context.Context, id, userId string, userRole string) error
	GetAllBlogs(ctx context.Context) ([]*domain.Blog, error)
	// FilterBlogs(ctx context.Context, filter domain.BlogFilter) ([]*domain.Blog, error)
	PaginateBlogs(ctx context.Context, filter domain.BlogFilter, page, pageSize int) ([]*domain.Blog, error)

	// Tag operations
	AddTagToBlog(ctx context.Context, blogID string, tag domain.BlogTag) error
	RemoveTagFromBlog(ctx context.Context, blogID string, tagID string) error

	// Comment operations
	AddComment(ctx context.Context, comment *domain.Comment, userId string) error
	GetCommentsByBlogID(ctx context.Context, blogID string) ([]*domain.Comment, error)

	// Like operations
	AddLike(ctx context.Context, like *domain.Like, userId string) error
	GetLikesByBlogID(ctx context.Context, blogID string) ([]*domain.Like, error)

	// View operations
	AddView(ctx context.Context, view *domain.View, userId string) error
	GetViewsByBlogID(ctx context.Context, blogID string) ([]*domain.View, error)

	// Tag operations
	GetAllTags(ctx context.Context) ([]*domain.BlogTag, error)
	CreateTag(ctx context.Context, tag *domain.BlogTag) error
	UpdateTag(ctx context.Context, id string, tag *domain.BlogTag) error
	DeleteTag(ctx context.Context, id string) error
	GetTagByID(ctx context.Context, id string) (*domain.BlogTag, error)

	DeleteComment(ctx context.Context, comment, userId string, userRole string) error
	RemoveLike(ctx context.Context, likeId, userId string, userRole string) error
	SearchBlogs(ctx context.Context, filter domain.BlogFilter, page int, pageSize int, orderBy []string) ([]*domain.GetBlogDTO, int, error)
	GetUserBlogs(ctx context.Context, userId string, page int, pageSize int) ([]*domain.Blog, error)
}

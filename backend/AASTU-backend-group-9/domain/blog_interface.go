package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogRepository defines the methods to interact with the data layer for blogs.
type BlogRepository interface {
	CreateBlog(ctx context.Context, blog *Blog) error
	GetBlogByID(ctx context.Context, id primitive.ObjectID) (*Blog, error)
	GetAllBlogs(ctx context.Context, page int, limit int, sortBy string) ([]*Blog, error)
	UpdateBlog(ctx context.Context, blog *Blog) error
	DeleteBlog(ctx context.Context, id primitive.ObjectID) error
	SearchBlogs(ctx context.Context, query string, filters *BlogFilters) ([]*Blog, error)
	FilterBlogs(ctx context.Context, filters *BlogFilters) ([]*Blog, error)
	IncrementPopularity(ctx context.Context, id primitive.ObjectID, metric string) error
}

// BlogUsecase defines the business logic methods for blogs.
type BlogUsecase interface {
	CreateBlog(ctx context.Context, blog *BlogCreationRequest) (*BlogResponse, error)
	GetBlogByID(ctx context.Context, id primitive.ObjectID) (*BlogResponse, error)
	GetAllBlogs(ctx context.Context, page int, limit int, sortBy string) ([]*BlogResponse, error)
	UpdateBlog(ctx context.Context, id primitive.ObjectID, blog *BlogUpdateRequest) (*BlogResponse, error)
	DeleteBlog(ctx context.Context, id primitive.ObjectID) error
	SearchBlogs(ctx context.Context, query string, filters *BlogFilters) ([]*BlogResponse, error)
	FilterBlogs(ctx context.Context, filters *BlogFilters) ([]*BlogResponse, error)
	TrackPopularity(ctx context.Context, id primitive.ObjectID, action *PopularityAction) error
}

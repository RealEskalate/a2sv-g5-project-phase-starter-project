package mongodb

import (
	"context"

	"github.com/RealEskalate/astu-backend-g4/backend/ASTU-backend-group-4/blog"
	"github.com/RealEskalate/astu-backend-g4/backend/ASTU-backend-group-4/pkg/infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogStorage struct {
	db *mongo.Database
}

func NewBlogStorage(db *mongo.Database) *BlogStorage {
	return &BlogStorage{db: db}
}

// CreateBlog implements BlogRepository.
func (b BlogStorage) CreateBlog(ctx context.Context, blog blog.Blog) (string, error) {
	panic("unimplemented")
}

// CreateComment implements BlogRepository.
func (b BlogStorage) CreateComment(ctx context.Context, comment blog.Comment) error {
	panic("unimplemented")
}

// DeleteBlog implements BlogRepository.
func (b BlogStorage) DeleteBlog(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DeleteComment implements BlogRepository.
func (b BlogStorage) DeleteComment(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DislikeBlog implements BlogRepository.
func (b BlogStorage) DislikeBlog(ctx context.Context, dislike blog.Dislike) error {
	panic("unimplemented")
}

// GetBlogByID implements BlogRepository.
func (b BlogStorage) GetBlogByID(ctx context.Context, id string) (blog.Blog, error) {
	panic("unimplemented")
}

// GetBlogs implements BlogRepository.
func (b BlogStorage) GetBlogs(ctx context.Context, filterOptions []blog.FilterOption, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[blog.Blog], error) {
	panic("unimplemented")
}

// GetCommentsByBlogID implements BlogRepository.
func (b BlogStorage) GetCommentsByBlogID(ctx context.Context, blogID string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[blog.Comment], error) {
	panic("unimplemented")
}

// LikeBlog implements BlogRepository.
func (b BlogStorage) LikeBlog(ctx context.Context, like blog.Like) error {
	panic("unimplemented")
}

// SearchBlogs implements BlogRepository.
func (b BlogStorage) SearchBlogs(ctx context.Context, query string) (infrastructure.PaginationResponse[blog.Blog], error) {
	panic("unimplemented")
}

// UpdateBlog implements BlogRepository.
func (b BlogStorage) UpdateBlog(ctx context.Context, id string, blog blog.Blog) error {
	panic("unimplemented")
}

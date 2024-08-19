package blog

import (
	"context"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
)

type BlogUseCaseImpl struct {
	blogRepository BlogRepository
}

func NewBlogUseCaseImpl(blogRepository BlogRepository) BlogUseCase {
	return &BlogUseCaseImpl{
		blogRepository: blogRepository,
	}
}

// CreateBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) CreateBlog(ctx context.Context, blog Blog) (Blog, error) {
	blogId, err := b.blogRepository.CreateBlog(ctx, blog)
	if err != nil {
		return Blog{}, err
	}

	newBlog, err := b.blogRepository.GetBlogByID(ctx, blogId)
	if err != nil {
		return Blog{}, err
	}

	return newBlog, nil
}

// CreateComment implements BlogUseCase.
func (b *BlogUseCaseImpl) CreateComment(ctx context.Context, comment Comment) error {
	panic("unimplemented")
}

// DeleteBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) DeleteBlog(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DeleteComment implements BlogUseCase.
func (b *BlogUseCaseImpl) DeleteComment(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DislikeBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) DislikeBlog(ctx context.Context, dislike Dislike) error {
	panic("unimplemented")
}

// GetBlogByID implements BlogUseCase.
func (b *BlogUseCaseImpl) GetBlogByID(ctx context.Context, id string) (Blog, error) {
	panic("unimplemented")
}

// GetBlogs implements BlogUseCase.
func (b *BlogUseCaseImpl) GetBlogs(ctx context.Context, filterQuery FilterQuery, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Blog], error) {
	panic("unimplemented")
}

// GetCommentsByBlogID implements BlogUseCase.
func (b *BlogUseCaseImpl) GetCommentsByBlogID(ctx context.Context, blogID string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Comment], error) {
	panic("unimplemented")
}

// LikeBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) LikeBlog(ctx context.Context, like Like) error {
	panic("unimplemented")
}

// SearchBlogs implements BlogUseCase.
func (b *BlogUseCaseImpl) SearchBlogs(ctx context.Context, query string) (infrastructure.PaginationResponse[Blog], error) {
	panic("unimplemented")
}

// UpdateBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) UpdateBlog(ctx context.Context, id string, blog Blog) (Blog, error) {
	panic("unimplemented")
}

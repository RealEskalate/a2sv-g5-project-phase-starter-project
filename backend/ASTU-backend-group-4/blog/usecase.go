package blog

import "github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"

type BlogUseCaseImpl struct {
	blogRepository BlogRepository
}

func NewBlogUseCaseImpl(blogRepository BlogRepository) BlogUseCase {
	return &BlogUseCaseImpl{
		blogRepository: blogRepository,
	}
}

// CreateBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) CreateBlog(blog Blog) (Blog, error) {
	panic("unimplemented")
}

// CreateComment implements BlogUseCase.
func (b *BlogUseCaseImpl) CreateComment(comment Comment) error {
	panic("unimplemented")
}

// DeleteBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) DeleteBlog(id string) error {
	panic("unimplemented")
}

// DeleteComment implements BlogUseCase.
func (b *BlogUseCaseImpl) DeleteComment(id string) error {
	panic("unimplemented")
}

// DislikeBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) DislikeBlog(dislike Dislike) error {
	panic("unimplemented")
}

// GetBlogByID implements BlogUseCase.
func (b *BlogUseCaseImpl) GetBlogByID(id string) (Blog, error) {
	panic("unimplemented")
}

// GetBlogs implements BlogUseCase.
func (b *BlogUseCaseImpl) GetBlogs(filterQuery FilterQuery, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Blog], error) {
	panic("unimplemented")
}

// GetCommentsByBlogID implements BlogUseCase.
func (b *BlogUseCaseImpl) GetCommentsByBlogID(blogID string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Comment], error) {
	panic("unimplemented")
}

// LikeBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) LikeBlog(like Like) error {
	panic("unimplemented")
}

// SearchBlogs implements BlogUseCase.
func (b *BlogUseCaseImpl) SearchBlogs(query string) (infrastructure.PaginationResponse[Blog], error) {
	panic("unimplemented")
}

// UpdateBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) UpdateBlog(id string, blog Blog) (Blog, error) {
	panic("unimplemented")
}

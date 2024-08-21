package blog

import (
	"context"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
)

type BlogUseCaseImpl struct {
	blogRepository BlogRepository
	authRepository auth.AuthRepository
}

func NewBlogUseCaseImpl(blogRepository BlogRepository) BlogUseCase {
	return &BlogUseCaseImpl{
		blogRepository: blogRepository,
	}
}

// CreateBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) CreateBlog(ctx context.Context, authorID string, blog CreateBlogRequest) (Blog, error) {
	author, err := b.authRepository.GetUserByUsername(ctx, authorID) // TODO: Change to GetUserByID
	if err != nil {
		return Blog{}, err
	}

	var newBlog Blog
	newBlog.AuthorID = author.ID
	newBlog.Title = blog.Title
	newBlog.Content = blog.Content
	newBlog.Tags = blog.Tags
	newBlog.ViewsCount = 0
	newBlog.CommentsCount = 0
	newBlog.LikesCount = 0
	newBlog.DislikesCount = 0
	newBlog.Popularity = 0

	blogId, err := b.blogRepository.CreateBlog(ctx, newBlog)
	if err != nil {
		return Blog{}, err
	}

	newBlog.ID = blogId
	return newBlog, nil
}

// CreateComment implements BlogUseCase.
func (b *BlogUseCaseImpl) CreateComment(ctx context.Context, userID, blogID string, comment CreateCommentRequest) error {
	user, err := b.authRepository.GetUserByUsername(ctx, userID) // TODO: Change to GetUserByID
	if err != nil {
		return err
	}
	blog, err := b.blogRepository.GetBlogByID(ctx, blogID)
	if err != nil {
		return err
	}

	var newComment Comment
	newComment.AuthorID = user.ID
	newComment.BlogID = blogID
	newComment.Content = comment.Content

	_, err = b.blogRepository.CreateComment(ctx, newComment)
	if err != nil {
		return err
	}

	blog.IncrementCommentsCount()
	b.blogRepository.UpdateBlog(ctx, blogID, blog)

	return nil
}

// DeleteBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) DeleteBlog(ctx context.Context, id, userID string) error {
	user, err := b.authRepository.GetUserByUsername(ctx, userID) // TODO: Change to GetUserByID
	if err != nil {
		return err
	}

	blog, err := b.blogRepository.GetBlogByID(ctx, id)
	if err != nil {
		return err
	}

	if blog.AuthorID != user.ID {
		return ErrBlogNotFound
	}

	err = b.blogRepository.DeleteBlog(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteComment implements BlogUseCase.
func (b *BlogUseCaseImpl) DeleteComment(ctx context.Context, id, userID string) error {
	user, err := b.authRepository.GetUserByUsername(ctx, userID) // TODO: Change to GetUserByID
	if err != nil {
		return err
	}

	comment, err := b.blogRepository.GetCommentByID(ctx, id)
	if err != nil {
		return err
	}

	if comment.AuthorID != user.ID {
		return ErrCommentNotFound
	}

	err = b.blogRepository.DeleteComment(ctx, id)
	if err != nil {
		return err
	}

	blog, err := b.blogRepository.GetBlogByID(ctx, comment.BlogID)
	if err != nil {
		return err
	}

	blog.DecrementCommentsCount()
	b.blogRepository.UpdateBlog(ctx, comment.BlogID, blog)

	return nil
}

// DislikeBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) DislikeBlog(ctx context.Context, userID string, blogID string) error {
	panic("unimplemented")
}

// GetBlogByID implements BlogUseCase.
func (b *BlogUseCaseImpl) GetBlogByID(ctx context.Context, id string) (Blog, error) {
	blog, err := b.blogRepository.GetBlogByID(ctx, id)
	if err != nil {
		return Blog{}, err
	}

	blog.IncrementViewsCount()
	blog.CalculatePopularity()
	b.blogRepository.UpdateBlog(ctx, id, blog)

	return blog, nil
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
func (b *BlogUseCaseImpl) LikeBlog(ctx context.Context, userID string, blogID string) error {
	panic("unimplemented")
}

// SearchBlogs implements BlogUseCase.
func (b *BlogUseCaseImpl) SearchBlogs(ctx context.Context, query string) (infrastructure.PaginationResponse[Blog], error) {
	panic("unimplemented")
}

// UpdateBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) UpdateBlog(ctx context.Context, id string, blog UpdateBlogRequest) (Blog, error) {
	panic("unimplemented")
}

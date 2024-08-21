package blog

import (
	"context"
	"time"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
)

type BlogUseCaseImpl struct {
	blogRepository BlogRepository
	authRepository auth.AuthRepository
}

func NewBlogUseCaseImpl(blogRepository BlogRepository, authRepository auth.AuthRepository) BlogUseCase {
	return &BlogUseCaseImpl{
		blogRepository: blogRepository,
		authRepository: authRepository,
	}
}

// CreateBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) CreateBlog(ctx context.Context, authorID string, blog CreateBlogRequest) (Blog, error) {
	author, err := b.authRepository.GetUserByID(ctx, authorID)
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
	newBlog.CreatedAt = time.Now()
	newBlog.UpdatedAt = time.Now()

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
	user, err := b.authRepository.GetUserByUsername(ctx, userID) // TODO: Change to GetUserByID
	if err != nil {
		return err
	}

	blog, err := b.blogRepository.GetBlogByID(ctx, blogID)
	if err != nil {
		return err
	}

	err = b.blogRepository.UnlikeBlog(ctx, Like{
		BlogID: blogID,
		UserID: user.ID,
	})

	if err == nil {
		blog.DecrementLikesCount()
	}

	err = b.blogRepository.DislikeBlog(ctx, Dislike{
		BlogID: blogID,
		UserID: user.ID,
	})

	if err != nil {
		return err
	}

	blog.IncrementDislikesCount()
	blog.CalculatePopularity()

	err = b.blogRepository.UpdateBlog(ctx, blogID, blog)
	if err != nil {
		return err
	}

	return nil
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
	blogs, err := b.blogRepository.GetBlogs(ctx, filterQuery, pagination)
	if err != nil {
		return infrastructure.PaginationResponse[Blog]{}, err
	}

	return blogs, nil
}

// GetCommentsByBlogID implements BlogUseCase.
func (b *BlogUseCaseImpl) GetCommentsByBlogID(ctx context.Context, blogID string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Comment], error) {
	comments, err := b.blogRepository.GetCommentsByBlogID(ctx, blogID, pagination)
	if err != nil {
		return infrastructure.PaginationResponse[Comment]{}, err
	}

	return comments, nil
}

// LikeBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) LikeBlog(ctx context.Context, userID string, blogID string) error {
	user, err := b.authRepository.GetUserByUsername(ctx, userID) // TODO: Change to GetUserByID
	if err != nil {
		return err
	}

	blog, err := b.blogRepository.GetBlogByID(ctx, blogID)
	if err != nil {
		return err
	}

	err = b.blogRepository.UndislikeBlog(ctx, Dislike{
		BlogID: blogID,
		UserID: user.ID,
	})

	if err == nil {
		blog.DecrementDislikesCount()
	}

	err = b.blogRepository.LikeBlog(ctx, Like{
		BlogID: blogID,
		UserID: user.ID,
	})

	if err != nil {
		return err
	}

	blog.IncrementLikesCount()
	blog.CalculatePopularity()

	b.blogRepository.UpdateBlog(ctx, blogID, blog)

	return nil
}

// SearchBlogs implements BlogUseCase.
func (b *BlogUseCaseImpl) SearchBlogs(ctx context.Context, query string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Blog], error) {
	blogs, err := b.blogRepository.SearchBlogs(ctx, query, pagination)
	if err != nil {
		return infrastructure.PaginationResponse[Blog]{}, err
	}

	return blogs, nil
}

// UpdateBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) UpdateBlog(ctx context.Context, id, userID string, blog UpdateBlogRequest) (Blog, error) {
	user, err := b.authRepository.GetUserByUsername(ctx, userID) // TODO: Change to GetUserByID
	if err != nil {
		return Blog{}, err
	}

	oldBlog, err := b.blogRepository.GetBlogByID(ctx, id)
	if err != nil {
		return Blog{}, err
	}

	if oldBlog.AuthorID != user.ID || !user.IsAdmin {
		return Blog{}, ErrBlogNotFound
	}

	var updatedBlog Blog
	updatedBlog.ID = id
	updatedBlog.AuthorID = oldBlog.AuthorID
	updatedBlog.Title = blog.Title
	updatedBlog.Content = blog.Content
	updatedBlog.Tags = blog.Tags
	updatedBlog.ViewsCount = oldBlog.ViewsCount
	updatedBlog.CommentsCount = oldBlog.CommentsCount
	updatedBlog.LikesCount = oldBlog.LikesCount
	updatedBlog.DislikesCount = oldBlog.DislikesCount
	updatedBlog.Popularity = oldBlog.Popularity

	err = b.blogRepository.UpdateBlog(ctx, id, updatedBlog)
	if err != nil {
		return Blog{}, err
	}

	return updatedBlog, nil
}

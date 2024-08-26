package blog

import (
	"context"
	"time"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"github.com/go-playground/validator/v10"
)

type BlogUseCaseImpl struct {
	blogRepository BlogRepository
	authRepository auth.AuthRepository
}

// UnDislikeBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) UnDislikeBlog(ctx context.Context, userID string, blogID string) error {
	return b.blogRepository.UndislikeBlog(ctx, Dislike{UserID: userID, BlogID: blogID})
}

// UnLikeBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) UnLikeBlog(ctx context.Context, userID string, blogID string) error {
	return b.blogRepository.UnlikeBlog(ctx, Like{UserID: userID, BlogID: blogID})
}

func NewBlogUseCaseImpl(blogRepository BlogRepository, authRepository auth.AuthRepository) BlogUseCase {
	return &BlogUseCaseImpl{
		blogRepository: blogRepository,
		authRepository: authRepository,
	}
}

// CreateBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) CreateBlog(ctx context.Context, authorID string, blog CreateBlogRequest) (Blog, error) {
	validate := validator.New()
	err := infrastructure.Validate(validate, blog)
	if err != nil {
		return Blog{}, err
	}

	var newBlog Blog
	newBlog.AuthorID = authorID
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
	validate := validator.New()
	err := infrastructure.Validate(validate, comment)
	if err != nil {
		return err
	}

	user, err := b.authRepository.GetUserByID(ctx, userID)
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
	newComment.CreatedAt = time.Now()
	newComment.UpdatedAt = time.Now()

	_, err = b.blogRepository.CreateComment(ctx, newComment)
	if err != nil {
		return err
	}

	blog.IncrementCommentsCount()
	blog.CalculatePopularity()
	b.blogRepository.UpdateBlog(ctx, blogID, blog)

	return nil
}

// DeleteBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) DeleteBlog(ctx context.Context, id, userID string) error {
	user, err := b.authRepository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	blog, err := b.blogRepository.GetBlogByID(ctx, id)
	if err != nil {
		return err
	}

	if blog.AuthorID != user.ID && !user.IsAdmin {
		return ErrBlogNotFound
	}

	err = b.blogRepository.DeleteBlog(ctx, id)
	if err != nil {
		return err
	}

	b.blogRepository.DeleteCommentsByBlogID(ctx, id)
	b.blogRepository.DeleteLikesByBlogID(ctx, id)
	b.blogRepository.DeleteDislikesByBlogID(ctx, id)

	return nil
}

// DeleteComment implements BlogUseCase.
func (b *BlogUseCaseImpl) DeleteComment(ctx context.Context, id, userID string) error {
	user, err := b.authRepository.GetUserByID(ctx, userID)
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
	user, err := b.authRepository.GetUserByID(ctx, userID)
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
func (b *BlogUseCaseImpl) GetBlogs(ctx context.Context, filterQuery FilterQuery, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[BlogSummary], error) {
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	if pagination.Page == 0 {
		pagination.Page = 1
	}

	if filterQuery.PopularityTo == 0 {
		filterQuery.PopularityTo = 100
	}
	if filterQuery.PopularityFrom == 0 {
		filterQuery.PopularityFrom = 0
	}

	if filterQuery.PopularityFrom > filterQuery.PopularityTo {
		return infrastructure.PaginationResponse[BlogSummary]{}, ErrInvalidPopularityRange
	}

	blogs, err := b.blogRepository.GetBlogs(ctx, filterQuery, pagination)
	if err != nil {
		return infrastructure.PaginationResponse[BlogSummary]{}, err
	}

	return blogs, nil
}

// GetCommentsByBlogID implements BlogUseCase.
func (b *BlogUseCaseImpl) GetCommentsByBlogID(ctx context.Context, blogID string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Comment], error) {
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	if pagination.Page == 0 {
		pagination.Page = 1
	}

	comments, err := b.blogRepository.GetCommentsByBlogID(ctx, blogID, pagination)
	if err != nil {
		return infrastructure.PaginationResponse[Comment]{}, err
	}

	return comments, nil
}

// LikeBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) LikeBlog(ctx context.Context, userID string, blogID string) error {
	user, err := b.authRepository.GetUserByID(ctx, userID)
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
func (b *BlogUseCaseImpl) SearchBlogs(ctx context.Context, query string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[BlogSummary], error) {
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	if pagination.Page == 0 {
		pagination.Page = 1
	}

	blogs, err := b.blogRepository.SearchBlogs(ctx, query, pagination)
	if err != nil {
		return infrastructure.PaginationResponse[BlogSummary]{}, err
	}

	return blogs, nil
}

// UpdateBlog implements BlogUseCase.
func (b *BlogUseCaseImpl) UpdateBlog(ctx context.Context, id, userID string, blog UpdateBlogRequest) (Blog, error) {
	validate := validator.New()
	err := infrastructure.Validate(validate, blog)
	if err != nil {
		return Blog{}, err
	}

	user, err := b.authRepository.GetUserByID(ctx, userID)
	if err != nil {
		return Blog{}, err
	}

	oldBlog, err := b.blogRepository.GetBlogByID(ctx, id)
	if err != nil {
		return Blog{}, err
	}

	if oldBlog.AuthorID != user.ID && !user.IsAdmin {
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

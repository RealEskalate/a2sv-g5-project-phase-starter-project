package usecase

import (
	"AAiT-backend-group-6/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type blogUseCase struct {
	blogRepo domain.BlogRepository
}

func NewBlogUseCase(blogRepo domain.BlogRepository) domain.BlogUseCase {
	return &blogUseCase{
		blogRepo: blogRepo,
	}
}

func (uc *blogUseCase) CreateBlog(c context.Context, blog *domain.Blog) (domain.Blog, error) {
	return uc.blogRepo.CreateBlog(c, blog)
}

func (uc *blogUseCase) GetBlog(c context.Context, id string) (*domain.Blog, error) {
	return uc.blogRepo.GetBlog(c, id)
}

func (uc *blogUseCase) GetBlogs(c context.Context, pagination *domain.Pagination) ([]*domain.Blog, error) {
	return uc.blogRepo.GetBlogs(c, pagination)
}

func (uc *blogUseCase) UpdateBlog(c context.Context, blog *domain.Blog, blogID string) error {
	// Ensure the blog ID is set
	blog.ID, _ = primitive.ObjectIDFromHex(blogID)
	return uc.blogRepo.UpdateBlog(c, blog)
}

func (uc *blogUseCase) DeleteBlog(c context.Context, id string) error {
	return uc.blogRepo.DeleteBlog(c, id)
}

func (uc *blogUseCase) LikeBlog(c context.Context, blogID string, userID string) error {
	return uc.blogRepo.LikeBlog(c, blogID, userID)
}

func (uc *blogUseCase) UnlikeBlog(c context.Context, blogID string, userID string) error {
	return uc.blogRepo.UnlikeBlog(c, blogID, userID)
}

func (uc *blogUseCase) CommentBlog(c context.Context, blogID string, comment *domain.Comment) error {
	return uc.blogRepo.CommentBlog(c, blogID, comment)
}

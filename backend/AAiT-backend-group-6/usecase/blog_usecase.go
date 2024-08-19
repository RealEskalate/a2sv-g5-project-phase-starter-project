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

func (uc *blogUseCase) CreateBlog(blog *domain.Blog) error {
	return uc.blogRepo.CreateBlog(context.Background(), blog)
}

func (uc *blogUseCase) GetBlog(id string) (*domain.Blog, error) {
	return uc.blogRepo.GetBlog(context.Background(), id)
}

func (uc *blogUseCase) GetBlogs(pagination *domain.Pagination) ([]*domain.Blog, error) {
	return uc.blogRepo.GetBlogs(context.Background(), pagination)
}

func (uc *blogUseCase) UpdateBlog(blog *domain.Blog, blogID string) error {
	// Ensure the blog ID is set
	blog.ID, _ = primitive.ObjectIDFromHex(blogID)
	return uc.blogRepo.UpdateBlog(context.Background(), blog)
}

func (uc *blogUseCase) DeleteBlog(id string) error {
	return uc.blogRepo.DeleteBlog(context.Background(), id)
}

func (uc *blogUseCase) LikeBlog(blogID string, userID string) error {
	return uc.blogRepo.LikeBlog(context.Background(), blogID, userID)
}

func (uc *blogUseCase) UnlikeBlog(blogID string, userID string) error {
	return uc.blogRepo.UnlikeBlog(context.Background(), blogID, userID)
}

func (uc *blogUseCase) CommentBlog(blogID string, comment *domain.Comment) error {
	return uc.blogRepo.CommentBlog(context.Background(), blogID, comment)
}

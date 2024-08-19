package blog_usecase

import (
	"blog-api/domain"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *BlogUsecase) CreateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog,error) {
	// Validate required fields
    if blog.Title == "" || blog.Content == "" || blog.Author == "" || len(blog.Tags) == 0 {
        return nil, errors.New("all fields are required")
    }

    // Set the blog ID, timestamps, and author ID
    blog.ID = primitive.NewObjectID()
    blog.CreatedAt = time.Now()
    blog.UpdatedAt = time.Now()

    // Call the repository to save the blog post
    err := u.repo.CreateBlog(context.Background(), blog)
    if err != nil {
        return nil, err
    }

    return blog, nil
}
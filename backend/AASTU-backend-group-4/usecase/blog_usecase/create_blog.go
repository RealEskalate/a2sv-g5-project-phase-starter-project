package blog_usecase

import (
	"blog-api/domain"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) CreateBlog(ctx context.Context, data *domain.BlogRequest, authorID primitive.ObjectID) (*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	if data.Title == "" || data.Content == "" || data.Author == "" || len(data.Tags) == 0 {
		return nil, errors.New("tiltle, content, author name and list of tags are required")
	}

	var blog domain.Blog

	blog.ID = primitive.NewObjectID()
	blog.Title = data.Title
	blog.Author = data.Author
	blog.Content = data.Content
	blog.AuthorID = authorID
	blog.Tags = data.Tags
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()

	err := bu.blogRepo.CreateBlog(ctx, &blog)
	if err != nil {
		return nil, err
	}

	return &blog, nil
}

package blog

import (
	"blogApp/internal/domain"
	"context"
	"errors"
	"fmt"
	"log"
)

// GetAllBlogs retrieves all blogs
func (u *blogUseCase) GetAllBlogs(ctx context.Context) ([]*domain.Blog, error) {
	blogs, err := u.repo.GetAllBlogs(ctx)
	if err != nil {
		log.Printf("Error retrieving all blogs: %v", err)
		return nil, fmt.Errorf("failed to retrieve blogs: %w", err)
	}
	return blogs, nil
}

// PaginateBlogs paginates the blogs
func (u *blogUseCase) PaginateBlogs(ctx context.Context, filter domain.BlogFilter, page, pageSize int) ([]*domain.Blog, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, errors.New("invalid pagination parameters")
	}
	blogs, err := u.repo.PaginateBlogs(ctx, filter, page, pageSize)
	if err != nil {
		log.Printf("Error paginating blogs: %v", err)
		return nil, fmt.Errorf("failed to paginate blogs: %w", err)
	}
	return blogs, nil
}

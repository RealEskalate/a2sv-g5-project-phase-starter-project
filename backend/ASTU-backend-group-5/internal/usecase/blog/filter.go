package blog

import (
	"blogApp/internal/domain"
	"context"
	"fmt"
	"log"
)

// FilterBlogs filters blogs based on certain criteria
func (u *blogUseCase) FilterBlogs(ctx context.Context, filter domain.BlogFilter) ([]*domain.Blog, error) {
	blogs, err := u.repo.FilterBlogs(ctx, filter)
	if err != nil {
		log.Printf("Error filtering blogs: %v", err)
		return nil, fmt.Errorf("failed to filter blogs: %w", err)
	}
	return blogs, nil
}

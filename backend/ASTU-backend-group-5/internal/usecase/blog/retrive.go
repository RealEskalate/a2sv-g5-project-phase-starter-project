package blog

import (
	"blogApp/internal/domain"
	"context"
	"fmt"
	"log"
)

func (u *blogUseCase) GetBlogByID(ctx context.Context, id string) (*domain.Blog, error) {
	blog, err := u.repo.GetBlogByID(ctx, id)
	if err != nil {
		log.Printf("Error retrieving blog by ID %s: %v", id, err)
		return nil, fmt.Errorf("failed to retrieve blog by ID: %w", err)
	}
	return blog, nil
}

// GetCommentsByBlogID retrieves comments by blog ID
func (u *blogUseCase) GetCommentsByBlogID(ctx context.Context, blogID string) ([]*domain.Comment, error) {
	comments, err := u.repo.GetCommentsByBlogID(ctx, blogID)
	if err != nil {
		log.Printf("Error retrieving comments for blog with ID %s: %v", blogID, err)
		return nil, fmt.Errorf("failed to retrieve comments: %w", err)
	}
	return comments, nil
}

// GetLikesByBlogID retrieves likes by blog ID
func (u *blogUseCase) GetLikesByBlogID(ctx context.Context, blogID string) ([]*domain.Like, error) {

	likes, err := u.repo.GetLikesByBlogID(ctx, blogID)
	if err != nil {
		log.Printf("Error retrieving likes for blog with ID %s: %v", blogID, err)
		return nil, fmt.Errorf("failed to retrieve likes: %w", err)
	}
	return likes, nil
}

// GetViewsByBlogID retrieves views by blog ID
func (u *blogUseCase) GetViewsByBlogID(ctx context.Context, blogID string) ([]*domain.View, error) {
	views, err := u.repo.GetViewsByBlogID(ctx, blogID)
	if err != nil {
		log.Printf("Error retrieving views for blog with ID %s: %v", blogID, err)
		return nil, fmt.Errorf("failed to retrieve views: %w", err)
	}
	return views, nil
}

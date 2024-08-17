package blog

import (
	"blogApp/internal/domain"
	"context"
	"fmt"
	"log"
)

// UpdateBlog updates an existing blog
func (u *blogUseCase) UpdateBlog(ctx context.Context, id string, blog *domain.Blog) error {
	err := u.repo.UpdateBlog(ctx, id, blog)
	if err != nil {
		log.Printf("Error updating blog with ID %s: %v", id, err)
		return fmt.Errorf("failed to update blog: %w", err)
	}
	return nil
}

// Tag operations

// AddTagToBlog adds a tag to a blog
func (u *blogUseCase) AddTagToBlog(ctx context.Context, blogID string, tag domain.BlogTag) error {
	err := u.repo.AddTagToBlog(ctx, blogID, tag)
	if err != nil {
		log.Printf("Error adding tag to blog with ID %s: %v", blogID, err)
		return fmt.Errorf("failed to add tag to blog: %w", err)
	}
	return nil
}

// RemoveTagFromBlog removes a tag from a blog
func (u *blogUseCase) RemoveTagFromBlog(ctx context.Context, blogID string, tagID string) error {
	err := u.repo.RemoveTagFromBlog(ctx, blogID, tagID)
	if err != nil {
		log.Printf("Error removing tag from blog with ID %s: %v", blogID, err)
		return fmt.Errorf("failed to remove tag from blog: %w", err)
	}
	return nil
}

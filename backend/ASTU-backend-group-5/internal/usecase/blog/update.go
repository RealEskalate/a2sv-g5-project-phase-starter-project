package blog

import (
	"blogApp/internal/domain"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateBlog updates an existing blog
func (u *blogUseCase) UpdateBlog(ctx context.Context, id string, blog *domain.Blog, userID string) error {
	if blog == nil {
		return errors.New("blog cannot be nil")
	}
	if blog.Author.Hex() != userID {
		return errors.New("you are not authorized to update this blog")
	}

	blog.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
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

package blog

import (
	"context"
	"fmt"
	"log"
)

// DeleteBlog deletes a blog by its ID
func (u *blogUseCase) DeleteBlog(ctx context.Context, id string, userId string) error {
	blog, err := u.repo.GetBlogByID(ctx, id)
	if err != nil {
		log.Printf("Error retrieving blog with ID %s: %v", id, err)
		return fmt.Errorf("failed to retrieve blog: %w", err)
	}

	if blog.Author.Hex() != userId {
		return fmt.Errorf("you are not authorized to delete this blog")
	}

	err = u.repo.DeleteBlog(ctx, id)
	if err != nil {
		log.Printf("Error deleting blog with ID %s: %v", id, err)
		return fmt.Errorf("failed to delete blog: %w", err)
	}
	return nil
}

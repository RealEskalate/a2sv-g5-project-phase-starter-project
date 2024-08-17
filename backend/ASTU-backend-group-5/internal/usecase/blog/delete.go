package blog

import (
	"context"
	"fmt"
	"log"
)

// DeleteBlog deletes a blog by its ID
func (u *blogUseCase) DeleteBlog(ctx context.Context, id string) error {
	err := u.repo.DeleteBlog(ctx, id)
	if err != nil {
		log.Printf("Error deleting blog with ID %s: %v", id, err)
		return fmt.Errorf("failed to delete blog: %w", err)
	}
	return nil
}

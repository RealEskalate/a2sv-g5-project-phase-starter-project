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
	if blog == nil {
		log.Printf("Blog with ID %s not found", id)
		return fmt.Errorf("blog not found")
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

func (u *blogUseCase) DeleteComment(ctx context.Context, commentId, userId string) error {
	comment, err := u.repo.GetCommentById(ctx, commentId)
	if err != nil {
		log.Printf("Error retrieving comment with ID %s: %v", commentId, err)
		return fmt.Errorf("failed to retrieve comment: %w", err)
	}

	if comment.UserID.Hex() != userId {
		return fmt.Errorf("you are not authorized to delete this comment")
	}
	err = u.repo.DeleteComment(ctx, comment.ID)
	if err != nil {
		log.Printf("Error deleting comment with ID %s: %v", comment.ID, err)
		return fmt.Errorf("failed to delete comment: %w", err)
	}
	return nil

}

func (u *blogUseCase) RemoveLike(ctx context.Context, likeId, userId string) error {
	like, err := u.repo.GetLikeById(ctx, likeId)
	if err != nil {
		log.Printf("Error retrieving like with ID %s: %v", likeId, err)
		return fmt.Errorf("failed to retrieve like: %w", err)
	}
	if like == nil {
		log.Printf("Like with ID %s not found", likeId)
		return fmt.Errorf("like not found")
	}

	if like.UserID.Hex() != userId {
		return fmt.Errorf("you are not authorized to delete this like")
	}
	err = u.repo.RemoveLike(ctx, like.ID)
	if err != nil {
		log.Printf("Error deleting like with ID %s: %v", like.ID, err)
		return fmt.Errorf("failed to delete like: %w", err)
	}
	return nil

}

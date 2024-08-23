package blog

import (
	"blogApp/internal/domain"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *blogUseCase) GetBlogByID(ctx context.Context, blogId, userId string) (*domain.GetSingleBlogDTO, error) {
	blog, err := u.repo.GetBlogByID(ctx, blogId)
	if err != nil {
		log.Printf("Error retrieving blog by ID %s: %v", blogId, err)
		return nil, fmt.Errorf("failed to retrieve blog by ID: %w", err)
	}

	if userId != "" {
		viewd, err := u.repo.HasUserViewedBlog(ctx, userId, blogId)

		if err != nil {
			log.Printf("Error checking if user has viewed blog with ID %s: %v", blogId, err)
			return nil, fmt.Errorf("failed to check if user has viewed blog: %w", err)
		}
		if !viewd {
			BlogIdObj, _ := primitive.ObjectIDFromHex(blogId)
			UserIdObj, _ := primitive.ObjectIDFromHex(userId)
			view := &domain.View{
				ID:     primitive.NewObjectID(),
				BlogID: BlogIdObj,
				UserID: UserIdObj,
			}
			err = u.repo.AddView(ctx, view)
			if err != nil {
				log.Printf("Error creating view for blog with ID %s: %v", blogId, err)
				return nil, fmt.Errorf("failed to create view: %w", err)
			}
		}
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

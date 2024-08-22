package usecases

import (
	domain "aait-backend-group4/Domain"
	"context"
	"time"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type likeUsecase struct {
	blogRepository  domain.BlogRepository
	likeRepository  domain.LikeRepository
	contextTimeouts time.Duration
}

// NewLikeUsecase initializes and returns a new likeUsecase instance
// It takes blog and like repositories along with a timeout duration
// to manage operations related to liking and disliking blog posts.
func NewLikeUsecase(blogRepository domain.BlogRepository, likeRepository domain.LikeRepository, timeouts time.Duration) domain.LikeUsecase {

	return &likeUsecase{
		blogRepository:  blogRepository,
		likeRepository:  likeRepository,
		contextTimeouts: timeouts,
	}

}

// Like allows a user to like a blog post
// It first adds the like to the like repository, then updates the blog's
// like count by incrementing it in the blog repository.
func (lu *likeUsecase) Like(ctx context.Context, userID string, blogID string) error {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	// Attempt to register the like action in the like repository
	err := lu.likeRepository.Like(ctx, userID, blogID)
	if err != nil {
		return err
	}

	// Increment the blog's like count in the blog repository
	return lu.blogRepository.UpdateFeedback(ctx, blogID, lu.blogRepository.IncrementLikes)
}

// Dislike allows a user to dislike a blog post
// It first adds the dislike to the like repository, then updates the blog's
// dislike count by incrementing it in the blog repository.
func (lu *likeUsecase) Dislike(ctx context.Context, userID string, blogID string) error {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	// Attempt to register the dislike action in the like repository
	err := lu.likeRepository.Dislike(ctx, userID, blogID)
	if err != nil {
		return err
	}

	// Increment the blog's dislike count in the blog repository
	return lu.blogRepository.UpdateFeedback(ctx, blogID, lu.blogRepository.IncrementDislike)
}

// RemoveLike allows a user to remove their like from a blog post
// It removes the like from the like repository and then decrements
// the blog's like count in the blog repository.
func (lu *likeUsecase) RemoveLike(ctx context.Context, likeID, blogID string) error {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	// Attempt to remove the like from the like repository
	err := lu.likeRepository.RemoveLike(ctx, likeID)
	if err != nil {
		return err
	}

	// Decrement the blog's like count in the blog repository
	return lu.blogRepository.UpdateFeedback(ctx, blogID, lu.blogRepository.DecrementLikes)
}

// RemoveDislike allows a user to remove their dislike from a blog post
// It removes the dislike from the like repository and then decrements
// the blog's dislike count in the blog repository.
func (lu *likeUsecase) RemoveDislike(ctx context.Context, dislikeID, blogID string) error {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	// Attempt to remove the dislike from the like repository
	err := lu.likeRepository.RemoveDislike(ctx, dislikeID)
	if err != nil {
		return err
	}

	// Decrement the blog's dislike count in the blog repository
	return lu.blogRepository.UpdateFeedback(ctx, blogID, lu.blogRepository.DecrementDislikes)
}

// GetLikesByUser retrieves a paginated list of likes made by a specific user
// It fetches the likes from the like repository based on the user ID,
// with support for pagination using limit and page parameters.
func (lu *likeUsecase) GetLikesByUser(ctx context.Context, userID string, limit, page int) ([]domain.Like, error) {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	// Calculate the offset for pagination
	offset := (page - 1) * limit

	// Retrieve the list of likes from the like repository
	likes, err := lu.likeRepository.GetLikesByUser(ctx, userID, limit, offset)
	if err != nil {
		return []domain.Like{}, err
	}

	return likes, nil
}

// GetLikesByBlog retrieves a paginated list of likes for a specific blog post
// It fetches the likes from the like repository based on the blog ID,
// with support for pagination using limit and page parameters.
func (lu *likeUsecase) GetLikesByBlog(ctx context.Context, blogID string, limit, page int) ([]domain.Like, error) {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	// Calculate the offset for pagination
	offset := (page - 1) * limit

	// Retrieve the list of likes from the like repository
	likes, err := lu.likeRepository.GetLikesByBlog(ctx, blogID, limit, offset)
	if err != nil {
		return []domain.Like{}, err
	}

	return likes, nil
}

// GetStatus checks if a user has liked or disliked a specific blog post
// It converts the userID and blogID to ObjectID, then queries the like repository
// to check the status and retrieves the corresponding like or dislike ID.
func (lu *likeUsecase) GetStatus(ctx context.Context, userID string, blogID string) (bool, string, error) {
    userIDObj, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return false, "", errors.New("invalid user ID format")
    }

    blogIDObj, err := primitive.ObjectIDFromHex(blogID)
    if err != nil {
        return false, "", errors.New("invalid blog ID format")
    }

    status, id, err := lu.likeRepository.GetStatus(ctx, userIDObj, blogIDObj)
    if err != nil {
        return false, "", err
    }

    return status, id, nil
}

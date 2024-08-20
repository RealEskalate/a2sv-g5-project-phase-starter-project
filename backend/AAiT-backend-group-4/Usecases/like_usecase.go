package usecases

import (
	domain "aait-backend-group4/Domain"
	"context"
	"time"
)

type likeUsecase struct {
	blogRepository  domain.BlogRepository
	likeRepository  domain.LikeReposiotory
	contextTimeouts time.Duration
}

func NewLikesUsecase(blogrepository domain.BlogRepository, likeRepository domain.LikeReposiotory, timeouts time.Duration) domain.LikeUsecase {
	return &likeUsecase{
		blogRepositry:   blogrepository,
		likeRepository:  likeRepository,
		contextTimeouts: timeouts,
	}

}

// Like increments the like count for a blog post
func (lu *likeUsecase) Like(ctx context.Context, userID string, blogID string) error {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	// Update the Like status in the like repository
	err := lu.likeRepository.Like(ctx, userID, blogID)
	if err != nil {
		return err
	}

	return lu.blogRepository.UpdateFeedback(ctx, blogID, lu.blogRepository.IncrementLikes)
}

func (lu *likeUsecase) Dislike(ctx context.Context, userID string, blogID string) error {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	// Update the Dislike status in the like repository
	err := lu.likeRepository.Dislike(ctx, userID, blogID)
	if err != nil {
		return err
	}

	return lu.blogRepository.UpdateFeedback(ctx, blogID, lu.blogRepository.IncrementDislike)
}

func (lu *likeUsecase) RemoveLike(ctx context.Context, likeID string) error {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	like, err := lu.likeReposiotory.GetLikeByID(ctx, likeID)
	if err != nil {
		return err
	}

	err := lu.likeReposiotory.RemoveLike(ctx, likeID)
	if err != nil {
		return err
	}

	blogID := like.BlogID.Hex()

	return lu.blogRepository.UpdateFeedback(ctx, blogID, blogRepository.DecrementLikes)
}

func (lu *likeUsecase) RemoveDislike(ctx context.Context, dislikeID string) error {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	err := lu.likeReposiotory.RemoveDislike(ctx, dislikeID)
	if err != nil {
		return err
	}

	dislike, err := lu.likeReposiotory.GetLikeByID(ctx, dislikeID)
	if err != nil {
		return err
	}
	blogID := dislike.BlogID.Hex()

	return lu.blogRepository.UpdateFeedback(ctx, blogID, blogRepository.DecrementDislikes)
}

func (lu *likeUsecase) GetLikesByUser(ctx context.Context, userID string, limit, page int) ([]domain.Like, error) {

	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	likes, err := lu.likeRepository.GetLikesByUser(ctx, userID, limit, offset)
	if err != nil {
		return []domian.Like{}, err
	}

	return likes, nil
}

func (lu *likeUsecase) GetLikesByBlog(ctx context.Context, blogID string, limit, page int) ([]domain.Like, error) {

	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeouts)
	defer cancel()

	offset := (page - 1) * limit

	likes, err := lu.likeRepository.GetLikesByBlog(ctx, blogID, limit, offset)
	if err != nil {
		return []domian.Like{}, err
	}

	return likes, nil
}

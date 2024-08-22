package usecase


import (
	"Blog_Starter/domain"
	"context"
	"time"
)

type BlogRatingUseCase struct {
	RatingRepository 		domain.BlogRatingRepository
	BlogRepository			domain.BlogRepository
	timeout					time.Duration
}

func NewBlogRatingUseCase(ratingRepository domain.BlogRatingRepository, blogRepository domain.BlogRepository, timeout time.Duration) domain.BlogRatingUseCase {
	return &BlogRatingUseCase {
		RatingRepository : ratingRepository,
		BlogRepository : blogRepository,
		timeout : timeout,
	}
}

// GetRatingByBlogID implements domain.BlogRatingRepository.
func (bu *BlogRatingUseCase) GetRatingByBlogID(c context.Context, blogID string) ([]*domain.BlogRating, error) {
	ctx, cancel := context.WithTimeout(c, bu.timeout)
	defer cancel()
	foundRatings, err := bu.RatingRepository.GetRatingByBlogID(ctx, blogID)
	return foundRatings, err
}
// GetRatingByID implements domain.BlogRatingRepository.
func (bu *BlogRatingUseCase) GetRatingByID(c context.Context, ratingID string) (*domain.BlogRating, error) {
	ctx, cancel := context.WithTimeout(c, bu.timeout)
	defer cancel()
	foundRating, err := bu.RatingRepository.GetRatingByID(ctx, ratingID)
	return foundRating, err
}

// InsertRating implements domain.BlogRatingRepository.
func (bu *BlogRatingUseCase) InsertRating(c context.Context, rating *domain.BlogRatingRequest) (*domain.BlogRating, error) {
	ctx, cancel := context.WithTimeout(c, bu.timeout)
	defer cancel()

	formattedRating := domain.BlogRating {
		BlogID : rating.BlogID,
		UserID : rating.UserID,
		Rating : rating.Rating,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}
	insertedRating, err := bu.RatingRepository.InsertRating(ctx, &formattedRating)
	if err != nil {return nil, err}
	err = bu.BlogRepository.InsertRating(ctx, insertedRating)
	return insertedRating, err
}

// UpdateRating implements domain.BlogRatingRepository.
func (bu *BlogRatingUseCase) UpdateRating(c context.Context, rating int, ratingID string) (*domain.BlogRating, error) {
	ctx, cancel := context.WithTimeout(c, bu.timeout)
	defer cancel()
	updatedRating,prevRating,err := bu.RatingRepository.UpdateRating(ctx, rating, ratingID)
	if err != nil {return nil, err}
	err = bu.BlogRepository.UpdateRating(ctx, updatedRating, prevRating)
	return updatedRating, err
}

func (bu *BlogRatingUseCase) DeleteRating(c context.Context, ratingID string) (*domain.BlogRating, error) {
	ctx, cancel := context.WithTimeout(c, bu.timeout)
	defer cancel()
	deletedRating, err := bu.RatingRepository.DeleteRating(ctx, ratingID)
	if err != nil {return nil, err}
	err = bu.BlogRepository.DeleteRating(ctx, deletedRating)
	return deletedRating, err
}
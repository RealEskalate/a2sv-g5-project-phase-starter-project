package usecase


import (
	"Blog_Starter/domain"
	"context"
	"time"
)

type BlogRatingUseCase struct {
	RatingRepository 		domain.BlogRatingRepository
	BlogRepository			domain.BlogRepository
}

func NewBlogRatingUseCase(ratingRepository domain.BlogRatingRepository, blogRepository domain.BlogRepository) domain.BlogRatingUseCase {
	return &BlogRatingUseCase {
		RatingRepository : ratingRepository,
		BlogRepository : blogRepository,
	}
}

// GetRatingByBlogID implements domain.BlogRatingRepository.
func (bu *BlogRatingUseCase) GetRatingByBlogID(ctx context.Context, blogID string) ([]*domain.BlogRating, error) {
	foundRatings, err := bu.RatingRepository.GetRatingByBlogID(ctx, blogID)
	return foundRatings, err
}
// GetRatingByID implements domain.BlogRatingRepository.
func (bu *BlogRatingUseCase) GetRatingByID(ctx context.Context, ratingID string) (*domain.BlogRating, error) {
	foundRating, err := bu.RatingRepository.GetRatingByID(ctx, ratingID)
	return foundRating, err
}

// InsertRating implements domain.BlogRatingRepository.
func (bu *BlogRatingUseCase) InsertRating(ctx context.Context, rating *domain.BlogRatingRequest) (*domain.BlogRating, error) {

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
func (bu *BlogRatingUseCase) UpdateRating(ctx context.Context, rating int, ratingID string) (*domain.BlogRating, error) {
	updatedRating,prevRating,err := bu.RatingRepository.UpdateRating(ctx, rating, ratingID)
	if err != nil {return nil, err}
	err = bu.BlogRepository.UpdateRating(ctx, updatedRating, prevRating)
	return updatedRating, err
}

func (bu *BlogRatingUseCase) DeleteRating(ctx context.Context, ratingID string) (*domain.BlogRating, error) {
	deletedRating, err := bu.RatingRepository.DeleteRating(ctx, ratingID)
	if err != nil {return nil, err}
	err = bu.BlogRepository.DeleteRating(ctx, deletedRating)
	return deletedRating, err
}
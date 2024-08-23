package blog_usecase

import (
	"blog-api/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) CreateLike(ctx context.Context, like *domain.LikeRequest) (*domain.Like, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	var newLike domain.Like

	newLike.ID = primitive.NewObjectID()
	newLike.BlogID = like.BlogID
	newLike.UserID = like.UserID
	newLike.LikedAt = time.Now()

	err := bu.likeRepo.AddLike(ctx, newLike)
	if err != nil {
		return nil, err
	}

	return &newLike, nil
}

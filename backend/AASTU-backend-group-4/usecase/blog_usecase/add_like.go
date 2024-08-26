package blog_usecase

import (
	"blog-api/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) AddLike(ctx context.Context, userID primitive.ObjectID, blogID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	var newLike domain.Like

	newLike.ID = primitive.NewObjectID()
	newLike.BlogID = blogID
	newLike.UserID = userID
	newLike.LikedAt = time.Now()

	err := bu.likeRepo.AddLike(ctx, newLike)
	if err != nil {
		return err
	}

	return nil
}

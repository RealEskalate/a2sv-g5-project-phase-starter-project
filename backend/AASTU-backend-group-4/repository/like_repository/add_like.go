package like_repository

import (
	"blog-api/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

func (lr *LikeRepository) AddLike(ctx context.Context, like domain.Like) error {
	result := lr.collection.FindOne(ctx, bson.M{"blog_id": like.BlogID, "user_id": like.UserID})
	if result == nil {
		return errors.New("you have already liked this blog")
	}
	_, err := lr.collection.InsertOne(ctx, like)
	return err
}

package like_repository

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lr *LikeRepository) GetLikeByID(ctx context.Context, likeID primitive.ObjectID) (domain.Like, error) {
	var like domain.Like
	err := lr.collection.FindOne(ctx, bson.M{"_id": likeID}).Decode(&like)
	if err != nil {
		return domain.Like{}, err
	}
	return like, nil
}

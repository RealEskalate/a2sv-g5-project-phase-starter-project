package like_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lr *LikeRepository) GetLikesCount(ctx context.Context, blogID primitive.ObjectID) (int, error) {
	count, err := lr.collection.CountDocuments(context.TODO(), bson.M{"blog_id": blogID})
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

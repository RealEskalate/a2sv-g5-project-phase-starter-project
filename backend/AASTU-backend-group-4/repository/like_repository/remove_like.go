package like_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (br LikeRepository) RemoveLike(ctx context.Context, likeID primitive.ObjectID) error {
	filter := bson.M{"_id": likeID}

	_, err := br.collection.DeleteOne(ctx, filter)
	return err
}

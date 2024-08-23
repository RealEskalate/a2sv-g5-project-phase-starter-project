package like_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lr *LikeRepository) RemoveBlogLikes(ctx context.Context, blogID primitive.ObjectID) error {
	filter := bson.M{"blog_id": blogID}

	_, err := lr.collection.DeleteMany(ctx, filter)
	return err
}

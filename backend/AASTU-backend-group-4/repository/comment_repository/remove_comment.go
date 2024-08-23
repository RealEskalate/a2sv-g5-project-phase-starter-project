package comment_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (br CommentRepository) RemoveComment(ctx context.Context, commentID primitive.ObjectID) error {
	filter := bson.M{"_id": commentID}

	_, err := br.collection.DeleteOne(ctx, filter)
	return err
}

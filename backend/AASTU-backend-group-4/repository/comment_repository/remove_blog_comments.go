package comment_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (cr *CommentRepository) RemoveBlogComments(ctx context.Context, blogID primitive.ObjectID) error {
	filter := bson.M{"blog_id": blogID}

	_, err := cr.collection.DeleteMany(ctx, filter)
	return err
}

package comment_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (cr *CommentRepository) GetCommentsCount(ctx context.Context, blogID primitive.ObjectID) (int, error) {
	count, err := cr.collection.CountDocuments(context.TODO(), bson.M{"blog_id": blogID})
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

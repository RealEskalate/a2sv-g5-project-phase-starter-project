package blog_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (br *BlogRepository) DeleteBlog(ctx context.Context, blogID primitive.ObjectID) error {
	filter := bson.M{"_id": blogID}

	_, err := br.collection.DeleteOne(ctx, filter)
	return err
}

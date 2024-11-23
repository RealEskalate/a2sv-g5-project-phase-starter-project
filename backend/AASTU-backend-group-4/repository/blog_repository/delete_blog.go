package blog_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (br *BlogRepository) DeleteBlog(ctx context.Context, userID, blogID primitive.ObjectID, isAdmin bool) error {
	filter := bson.M{"_id": blogID}

	existingBlog, err := br.GetBlogByID(ctx, blogID)
	if err != nil {
		return errors.New("blog with that ID doesn't exist")
	}
	if !isAdmin && existingBlog.AuthorID != userID {
		return errors.New("you do not have permission to delete this blog post")
	}
	_, err = br.collection.DeleteOne(ctx, filter)

	return err
}

package Repositories

import (
	"blogapp/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentRepository struct {
	postcollection    Domain.Collection
	commentcollection Domain.Collection
}

func NewCommentRepository(blogcollection Domain.BlogCollections) *commentRepository {
	return &commentRepository{
		postcollection:    blogcollection.Posts,
		commentcollection: blogcollection.Comments,
	}
}

func (cr *commentRepository) CommentOnPost(ctx context.Context, comment *Domain.Comment, objID primitive.ObjectID) (error, int) {
	_, err := cr.commentcollection.InsertOne(ctx, comment)
	if err != nil {
		return err, 500
	}
	// add comment to post collection in field which is an array of comments pointers
	filter := bson.D{{"_id", objID}}
	update := bson.D{{"$push", bson.D{{"comments", comment}}}}
	_, err = cr.postcollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err, 500
	}

	return nil, 200
}

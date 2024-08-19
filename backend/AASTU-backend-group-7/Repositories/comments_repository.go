package Repositories

import (
	"blogapp/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentRepository struct {
	postcollection    Domain.Collection
	commentCollection Domain.Collection
}

func NewCommentRepository(blogcollection Domain.BlogCollections) *commentRepository {
	return &commentRepository{
		postcollection:    blogcollection.Posts,
		commentCollection: blogcollection.Comments,
	}
}

func (cr *commentRepository) CommentOnPost(ctx context.Context, comment *Domain.Comment, objID primitive.ObjectID) (error, int) {
	
	_, err := cr.commentCollection.InsertOne(ctx, comment)
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

func (cr *commentRepository) GetCommentByID(ctx context.Context, id primitive.ObjectID) (*Domain.Comment, error, int) {
	var comment *Domain.Comment
	filter := bson.D{{"_id", id}}
	err := cr.commentCollection.FindOne(ctx, filter).Decode(&comment)
	if err != nil {
		return nil, err, 500
	}
	return comment, nil, 200
}

func (cr *commentRepository) EditComment(ctx context.Context, id primitive.ObjectID, comment *Domain.Comment) (error, int) {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", comment}}
	_, err := cr.commentCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err, 500
	}
	return nil, 200
}
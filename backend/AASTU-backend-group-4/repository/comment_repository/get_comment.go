package comment_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"blog-api/domain"
)

func (cr *CommentRepository) GetCommentByID(ctx context.Context, id primitive.ObjectID) (*domain.Comment, error) {
	var comment domain.Comment
	err := cr.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

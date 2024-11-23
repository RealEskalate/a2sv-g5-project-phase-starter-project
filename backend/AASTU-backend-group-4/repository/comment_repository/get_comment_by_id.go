package comment_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"blog-api/domain"
)

func (cr *CommentRepository) GetCommentByID(ctx context.Context, commentID primitive.ObjectID) (domain.Comment, error) {
	var comment domain.Comment
	err := cr.collection.FindOne(ctx, bson.M{"_id": commentID}).Decode(&comment)
	if err != nil {
		return domain.Comment{}, err
	}
	return comment, nil
}

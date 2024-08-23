package comment_repository

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (cr *CommentRepository) GetBlogComments(ctx context.Context, blogID primitive.ObjectID) ([]domain.Comment, error) {
	var comments []domain.Comment

	cursor, err := cr.collection.Find(ctx, bson.M{"blog_id": blogID})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

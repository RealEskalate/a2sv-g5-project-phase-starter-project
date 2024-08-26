package comment_repository

import (
	"blog-api/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (cr *CommentRepository) RemoveComment(ctx context.Context, userID, blogID primitive.ObjectID, isAdmin bool) error {
	var comment domain.Comment

	filter := bson.M{"blog_id": blogID, "user_id": userID}
	cr.collection.FindOne(ctx, filter).Decode(&comment)
	if !isAdmin && comment.UserID != userID {
		return errors.New("only the user or an admin can remove this comment")
	}

	_, err := cr.collection.DeleteOne(ctx, filter)
	return err
}

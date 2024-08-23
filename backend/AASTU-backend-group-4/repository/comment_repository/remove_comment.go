package comment_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (cr *CommentRepository) RemoveComment(ctx context.Context, userID, commentID primitive.ObjectID, isAdmin bool) error {
	filter := bson.M{"_id": commentID}

	existingComment, err := cr.GetCommentByID(ctx, commentID)
	if err != nil {
		return errors.New("comment with that ID doesn't exist")
	}

	if !isAdmin && existingComment.UserID != userID {
		return errors.New("only the user or admin can remove this like")
	}

	_, err = cr.collection.DeleteOne(ctx, filter)
	return err
}

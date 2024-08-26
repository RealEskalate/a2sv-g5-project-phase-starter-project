package like_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lr *LikeRepository) RemoveLike(ctx context.Context, userID, likeID primitive.ObjectID, isAdmin bool) error {
	filter := bson.M{"_id": likeID}

	existingLike, err := lr.GetLikeByID(ctx, likeID)
	if err != nil {
		return errors.New("you haven't liked this blog yet")
	}

	if !isAdmin && existingLike.UserID != userID {
		return errors.New("only the user or an admin can remove this like")
	}

	_, err = lr.collection.DeleteOne(ctx, filter)
	return err
}

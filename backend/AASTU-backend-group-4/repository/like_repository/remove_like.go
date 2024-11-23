package like_repository

import (
	"blog-api/domain"
	"context"

	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lr *LikeRepository) RemoveLike(ctx context.Context, userID, blogID primitive.ObjectID, isAdmin bool) error {
	var like domain.Like
	filter := bson.M{"blog_id": blogID, "user_id": userID}
	lr.collection.FindOne(ctx, filter).Decode(&like)
	if !isAdmin && like.UserID != userID {
		return errors.New("only the user or an admin can remove this like")
	}

	_, err := lr.collection.DeleteOne(ctx, filter)
	return err
}

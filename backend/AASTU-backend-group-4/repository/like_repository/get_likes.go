package like_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"blog-api/domain"
)

func (lr *LikeRepository) GetLikes(ctx context.Context, blogID, userID primitive.ObjectID) ([]domain.Like, error) {
	var likes []domain.Like
	filter := bson.M{}
	if blogID != primitive.NilObjectID {
		filter["blog_id"] = blogID
	}
	if userID != primitive.NilObjectID {
		filter["user_id"] = userID
	}
	cursor, err := lr.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &likes)
	if err != nil {
		return nil, err
	}
	return likes, nil
}

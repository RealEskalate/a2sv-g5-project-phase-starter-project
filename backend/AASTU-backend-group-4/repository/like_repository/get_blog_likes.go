package like_repository

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lc *LikeRepository) GetBlogLikes(ctx context.Context, blogID primitive.ObjectID) ([]domain.Like, error) {
	var likes []domain.Like

	cursor, err := lc.collection.Find(ctx, bson.M{"blog_id": blogID})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &likes)
	if err != nil {
		return nil, err
	}
	return likes, nil
}

package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
    ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    BlogID primitive.ObjectID `bson:"blog_id" json:"blog_id"`
    UserID primitive.ObjectID `bson:"user_id" json:"user_id"`
}

type LikeRepository interface {
    GetLikesCount(ctx context.Context, blogID string) (int, error)

    // CreateLike(ctx context.Context, like Like) error
    // DeleteLike(ctx context.Context, likeID string) error
    // IsLikedByUser(ctx context.Context, blogID, userID string) (bool, error)
}
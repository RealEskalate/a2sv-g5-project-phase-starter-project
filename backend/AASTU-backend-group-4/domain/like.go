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
	GetLikesCount(ctx context.Context, blogID primitive.ObjectID) (int, error)
	GetLikes(ctx context.Context, blogID primitive.ObjectID) ([]Like, error)
	CreateLike(ctx context.Context, like Like) error
	DeleteLike(ctx context.Context, likeID string) error
}

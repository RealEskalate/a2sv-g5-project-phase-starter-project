package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	PostID primitive.ObjectID `bson:"post_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id,omitempty"`
}

type DisLike struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	PostID primitive.ObjectID `bson:"post_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id,omitempty"`
}

type LikeRepository interface {
	CreateLike(like Like) error
	GetLikesByPostID(postID primitive.ObjectID) ([]Like, error)
	DeleteLike(likeID primitive.ObjectID) error
}

type LikeUsecase interface {
	CreateLike(ctx context.Context, like Like) error
	GetLikesByPostID(ctx context.Context, postID primitive.ObjectID) ([]Like, error)
	DeleteLike(ctx context.Context, likeID primitive.ObjectID) error
}

type DisLikeRepository interface {
	CreateLike(like Like) error
	GetLikesByPostID(postID primitive.ObjectID) ([]Like, error)
	DeleteLike(likeID primitive.ObjectID) error
}

type DisLikeUsecase interface {
	CreateDisLike(ctx context.Context, dislike DisLike) error
	GetDisLikesByPostID(ctx context.Context, postID primitive.ObjectID) ([]DisLike, error)
	DeleteDisLike(ctx context.Context, dislikeID primitive.ObjectID) error
}

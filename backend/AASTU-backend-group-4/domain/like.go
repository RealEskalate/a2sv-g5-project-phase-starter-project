package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BlogID  primitive.ObjectID `bson:"blog_id" json:"blog_id"`
	UserID  primitive.ObjectID `bson:"user_id" json:"user_id"`
	LikedAt time.Time          `bson:"created_at" json:"created_at"`
}

type LikeRepository interface {
	AddLike(ctx context.Context, like Like) error                                                             //
	GetLikesCount(ctx context.Context, blogID primitive.ObjectID) (int, error)                                //
	GetLikes(ctx context.Context, userID, blogID primitive.ObjectID) ([]Like, error)                          //
	GetBlogLikes(ctx context.Context, blogID primitive.ObjectID) ([]Like, error)                              //
	GetLikeByID(ctx context.Context, likeID primitive.ObjectID) (*Like, error)                                //
	RemoveLike(ctx context.Context, userID primitive.ObjectID, likeID primitive.ObjectID, isAdmin bool) error //
	RemoveBlogLikes(ctx context.Context, blogID primitive.ObjectID) error                                     //
}

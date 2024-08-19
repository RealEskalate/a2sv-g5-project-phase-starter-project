package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID primitive.ObjectID `bson:"user_id" json:"user_id"`
	BlogID primitive.ObjectID `bson:"blog_id" json:"blog_id"`
	Status *bool              `bson:"status" json:"status"`
}

type LikeReposiotory interface {
	Create(c context.Context, userID string, blogID string, status bool) error
	Like(c context.Context, userID string, blogID string) error
	Dislike(c context.Context, userId string, blogID string) error
	RemoveLike(c context.Context, id string) error
	RemoveDislike(c context.Context, id string) error
	GetLikesByUser(c context.Context, userID string) (likes []Like, err error)
	GetLikesByBlog(c context.Context, blogID string) (likes []Like, err error)
}

type LikeUsecase interface {
	Like(c context.Context, userID string, blogID string) error
	Dislike(c context.Context, userID string, blogID string) error
	GetLikesByUser(c context.Context, userID string) error
	GetLikesByBlog(c context.Context, blogID string) error
	GetLikeByID(c context.Context, likeID string) (Like, error)
}

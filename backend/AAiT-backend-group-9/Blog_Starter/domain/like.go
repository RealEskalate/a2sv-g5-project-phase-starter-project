package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionLike = "likes"
)

type Like struct {
	LikeID    primitive.ObjectID `json:"like_id" bson:"_id"`
	UserID    string             `json:"user_id" bson:"user_id"`
	BlogID    string             `json:"blog_id" bson:"blog_id"`
	CreatedAt time.Time          `json:"createtimestamp" bson:"createtimestamp"`
}

type LikeRepository interface {
	LikeBlog(c context.Context, like *Like) (*Like, error) // add like count but check if user liked the post already
	GetByID(c context.Context, userID string, blogID string) (*Like, error)
	UnlikeBlog(c context.Context, likeID string) (*Like, error) // subtract like count but check if user has liked the post first
	DeleteLikeByBlogID(c context.Context, blogID string) error
}

type LikeUseCase interface {
	LikeBlog(c context.Context, like *Like) (*Like, error)
	UnlikeBlog(c context.Context, likeID string) (*Like, error)
	GetByID(c context.Context, userID string, blogID string) (*Like, error)
}

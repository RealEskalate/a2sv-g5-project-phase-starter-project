package domain

import (
	"context"
	"time"
)

const (
	CollectionLike = "Likes"
)

type Like struct {
	LikeID string `json:"like_id" bson:"_id"`
	UserID string `json:"user_id" bson:"user_id"`
	BlogID string `json:"blog_id" bson:"blog_id"`
	CreatedAt   time.Time `json:"createtimestamp" bson:"createtimestamp"`

}

type LikeRepository interface {
	LikeBlog(c context.Context, like *Like) (*Like, error) // add like count but check if user liked the post already
	GetByID(c context.Context, userID string, blogID string) (*Like, error)
	UnlikeBlog(c context.Context, likeID string) (*Like, error) // subtract like count but check if user has liked the post first
}

type LikeUseCase interface {
	LikeBlog(c context.Context, like *Like) (*Like, error)
	UnlikeBlog(c context.Context, likeID string) (*Like, error)
	GetByID(c context.Context, userID string, blogID string) (*Like, error)

}
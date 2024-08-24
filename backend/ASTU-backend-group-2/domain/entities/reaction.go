package entities

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionReaction = "reactions"
)

// user reaction to the blog if liked or disliked
type Reaction struct {
	BlogID   primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	Liked    bool               `json:"liked" bson:"liked"`
	Disliked bool               `json:"disliked" bson:"disliked"`
	Date     time.Time          `json:"date" bson:"date"`
}

type ReactionRepository interface {
	Like(c context.Context, blogID, userID string) error
	Dislike(c context.Context, blogID, userID string) error
	RemoveLike(c context.Context, blogID, userID string) error
	RemoveDislike(c context.Context, blogID, userID string) error
	IsPostLiked(c context.Context, blogID, userID string) (bool, error)
	IsPostDisliked(c context.Context, blogID, userID string) (bool, error)
}

type ReactionUsecase interface {
	ToggleLike(c context.Context, blogID, userID string) error
	ToggleDislike(c context.Context, blogID, userID string) error
}

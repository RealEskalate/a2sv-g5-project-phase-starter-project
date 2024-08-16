package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Blog struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"id,omitempty"`
	Title        string             `json:"title " bson:"title"`
	Content      string             `json:"content" bson:"content"`
	Tag          string             `json:"tag" bson:"tag"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	Owner        User               `json:"owner" bson:"owner"`
	CommentCount int                `json:"comment_count" bson:"comment_count"`
	LikeCount    int                `json:"like_count" bson:"like_count"`
	DisLikeCount int                `json:"dislike_count" bson:"dislike_count"`
	ViewCount    int                `json:"view_count" bson:"view_count"`
}

package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reply struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    string             `json:"userId" bson:"userId"`
	CommentID primitive.ObjectID `json:"commentId" bson:"commentId"`
	Content   string             `json:"content" bson:"content"`
	LikesCount int               `json:"likesCount" bson:"likesCount"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
package forms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentForm struct {
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

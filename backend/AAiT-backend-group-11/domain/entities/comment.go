package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	BlogPostID primitive.ObjectID `bson:"blogPostId"`
	AuthorID   primitive.ObjectID `bson:"authorId"`
	Content    string             `bson:"content"`
	CreatedAt  time.Time          `bson:"createdAt"`
}
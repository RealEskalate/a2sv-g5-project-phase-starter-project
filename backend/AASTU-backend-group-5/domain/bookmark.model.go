package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bookmark struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogID primitive.ObjectID `json:"blog_id" bson:"blog_id"`
}

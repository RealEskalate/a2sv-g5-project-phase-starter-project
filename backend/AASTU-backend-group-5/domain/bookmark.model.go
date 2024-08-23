package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bookmark struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`
	BlogID primitive.ObjectID `bson:"blog_id"`
}

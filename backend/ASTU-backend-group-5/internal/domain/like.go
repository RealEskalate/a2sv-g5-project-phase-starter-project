package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Like struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`                // Unique identifier for the blog
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`       // ID of user
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`       // ID of blog
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"` // Timestamp for when the like was created
}

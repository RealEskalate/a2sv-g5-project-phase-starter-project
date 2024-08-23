package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`                // Unique identifier for the blog
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`       // ID of commenter
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`       // ID of blog
	Content   string             `json:"content" bson:"content"`       // Content of the comment
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"` // Timestamp for when the comment was created
	ReplyToId primitive.ObjectID `json:"reply_to" bson:"reply_to"`     // ID of comment

}

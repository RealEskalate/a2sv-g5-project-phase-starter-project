package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Blog represents a blog post with flexible content.
type Blog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`                // Unique identifier for the blog
	Author    primitive.ObjectID `json:"ownerID" bson:"ownerID"`       // ID of the blog author
	Title     string             `json:"title" bson:"title"`           // Title of the blog
	Content   []interface{}      `json:"content" bson:"content"`       // Array of any type of content
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"` // Timestamp for when the blog was created
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"` // Timestamp for when the blog was last updated
	Tags      []BlogTag          `json:"tags" bson:"tags"`             // Tags for categorizing the blog
}

type BlogTag struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`    // Unique identifier for the blog
	Name string             `json:"name" bson:"name"` // Name of the blog
}

type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`                // Unique identifier for the blog
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`       // ID of commenter
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`       // ID of blog
	Content   string             `json:"content" bson:"content"`       // Content of the comment
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"` // Timestamp for when the comment was created
}

type Like struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`                // Unique identifier for the blog
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`       // ID of user
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`       // ID of blog
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"` // Timestamp for when the like was created
}

type View struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`                // Unique identifier for the blog
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`       // ID of user
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`       // ID of blog
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"` // Timestamp for when the view was created
}

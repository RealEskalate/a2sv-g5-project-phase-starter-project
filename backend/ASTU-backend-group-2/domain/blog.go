package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// this structure will be used when responding for blog request
type Blog struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	AuthorID     primitive.ObjectID `json:"author_id" bson:"author_id"`
	Title        string             `json:"title" bson:"title" binding:"required"`
	Tags         []string           `json:"tags" bson:"tags"`
	Content      string             `json:"content" bson:"content" binding:"required"`
	ViewCount    int                `json:"view_count" bson:"view_count"`
	LikeCount    int                `json:"like_count" bson:"like_count"`
	DislikeCount int                `json:"dislike_count" bson:"dislike_count"`
	Comments     []Comment          `json:"comments" bson:"comments"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// defines the structure for the blogs that will be  received from the request when creating and updating
type BlogIn struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	AuthorID  primitive.ObjectID `json:"author_id" bson:"author_id"`
	Title     string             `json:"title" bson:"title" binding:"required"`
	Tags      []string           `json:"tags" bson:"tags"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// this structure when creating and updating comments
type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

// user reaction to the blog if liked or disliked
type Reaction struct {
	BlogID   primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	Liked    bool               `json:"liked" bson:"liked"`
	Disliked bool               `json:"disliked" bson:"disliked"`
	Date     time.Time          `json:"date" bson:"date"`
}

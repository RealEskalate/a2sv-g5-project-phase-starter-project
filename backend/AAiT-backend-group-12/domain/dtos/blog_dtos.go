package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogDTO represents the data structure stored in MongoDB for a blog post.
type BlogDTO struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title"`
	Content    string             `bson:"content"`
	Username   string             `bson:"username"`
	Tags       []string           `bson:"tags"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	ViewCount  uint               `bson:"view_count"`
	LikedBy    []string           `bson:"liked_by"`
	DislikedBy []string           `bson:"disliked_by"`
	Comments   []CommentDTO       `bson:"comments"`
}

type CommentDTO struct {
	ID        primitive.ObjectID `bson:"_id"`
	Content   string             `bson:"content"`
	Username  string             `bson:"user_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	ViewCount uint               `bson:"view_count"`
	Comments  []CommentDTO       `bson:"comments"`
}

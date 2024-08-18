package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogDTO represents the data structure stored in MongoDB for a blog post.
type BlogDTO struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	Title      string               `bson:"title"`
	Content    string               `bson:"content"`
	UserID     primitive.ObjectID   `bson:"user_id"`
	Tags       []string             `bson:"tags"`
	CreatedAt  time.Time            `bson:"created_at"`
	UpdatedAt  time.Time            `bson:"updated_at"`
	ViewCount  uint                 `bson:"view_count"`
	LikedBy    []primitive.ObjectID `bson:"liked_by"`
	DislikedBy []primitive.ObjectID `bson:"disliked_by"`
	Comments   []CommentDTO         `bson:"comments"`
}


type UserDTO struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserName  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Role      string             `bson:"role"`
	CreatedAt time.Time          `bson:"created_at"`
}


type CommentDTO struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Content   string             `bson:"content"`
	UserID    primitive.ObjectID `bson:"user_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	ViewCount uint               `bson:"view_count"`
	Comments  []CommentDTO       `bson:"comments"`
}

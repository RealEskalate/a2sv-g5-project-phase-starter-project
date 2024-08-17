
package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	BlogID    primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title" binding:"required"`
	Content   string             `bson:"content" binding:"required"`
	Author    User               `bson:"author" binding:"required"`
	Tags      []string           `bson:"tags"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Views     int                `bson:"views"`
	Likes     int                `bson:"likes"`
	Dislikes  int                `bson:"dislikes"`
	Comments  []Comment          `bson:"comments"`
}

type Comment struct {
	CommentID primitive.ObjectID `bson:"_id,omitempty"`
	Author    User               `bson:"author" binding:"required"`
	Content   string             `bson:"content" binding:"required"`
	CreatedAt time.Time          `bson:"created_at"`
}

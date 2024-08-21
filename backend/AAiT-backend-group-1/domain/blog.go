package domain

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title        string             `bson:"title" json:"title"`
	Content      string             `bson:"content" json:"content"`
	AuthorID     primitive.ObjectID `bson:"author_id" json:"author_id"`
	Tags         []string           `bson:"tags" json:"tags"`
	ViewCount    int                `bson:"view_count" json:"view_count"`
	Likes        []string           `bson:"likes" json:"likes"`
	Dislikes     []string           `bson:"dislikes" json:"dislikes"`
	Comments     []string           `bson:"comments" json:"comments"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}
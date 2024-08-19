package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Blog struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	AuthorID  primitive.ObjectID `bson:"author_id" json:"author_id"`
	Views     int                `bson:"views" json:"views"`
	Likes     int                `bson:"likes" json:"likes"`
	Slug      string             `bson:"slug" json:"slug"`
	Comments  []primitive.ObjectID          `bson:"comments" json:"comments"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Tags      []string           `bson:"tags" json:"tags"`
}
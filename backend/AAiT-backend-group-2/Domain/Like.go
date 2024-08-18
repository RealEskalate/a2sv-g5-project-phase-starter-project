package domain

import "time"

type Like struct {
	ID        string    `bson:"_id" json:"id"`
	BlogID    string    `bson:"blog_id" json:"blog_id"`
	Author    string    `bson:"author" json:"author"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
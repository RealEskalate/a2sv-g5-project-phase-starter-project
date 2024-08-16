package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID         primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Author     string             `json:"author" validate:"required"`
	Title      string             `json:"title" validate:"required,min=1,max=255"`
	Content    string             `json:"content" validate:"required,min=9"`
	Tags       []string           `json:"tags"`
	Creater_id primitive.ObjectID `json:"creater_id" validate:"required"`
	CreatedAt  time.Time          `json:"createdAt" validate:"required"`
	UpdatedAt  time.Time          `json:"updatedAt" validate:"required"`
	Comments   []Comment          `json:"comments"`
	Post_image string             `json:"post_image"`
}

type Comment struct {
	ID                 primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Post_ID            primitive.ObjectID `json:"post_id"`
	Commentor_ID       primitive.ObjectID `json:"commentor_id" validate:"required"`
	Commentor_username string             `json:"commentor_username" validate:"required"`
	Content            string             `json:"content" validate:"required,min=1,max=255"`
}

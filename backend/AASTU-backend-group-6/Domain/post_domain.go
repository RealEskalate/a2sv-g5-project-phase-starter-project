package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID         primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Author     string             `json:"author"`
	Title      string             `json:"title"`
	Content    string             `json:"content"`
	Tags       []string           `json:"tags"`
	Creater_id primitive.ObjectID `json:"creater_id"`
	CreatedAt  time.Time          `json:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt"`
	Comments   []Comment          `json:"comments"`
	Post_image string             `json:"post_image"`
}

type Comment struct {
	ID                 primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Post_ID            primitive.ObjectID `json:"post_id"`
	Commentor_ID       primitive.ObjectID `json:"commentor_id"`
	Commentor_username string             `json:"commentor_username"`
	Content            string             `json:"content"`
}

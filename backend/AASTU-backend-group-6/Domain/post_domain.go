package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Author      string             `json:"author"`
	Title       string             `json:"title"`
	Content	 string             `json:"content"`
	Tags	 []string           `json:"tags"`
	UserID	  string             `json:"userID"`
	CreatedAt  time.Time		  `json:"createdAt"`
	UpdatedAt  time.Time		  `json:"updatedAt"`
	Comments   []Comment         `json:"comments"`

}

type Comment struct {
	ID          primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Commentor      string             `json:"Commentor"`
	Content	 string             `json:"content"`
}


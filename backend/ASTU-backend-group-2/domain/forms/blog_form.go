package forms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogForm struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Title   string             `json:"title" bson:"title" binding:"required,max=100,min=5,alphanum"`
	Tags    []string           `json:"tags" bson:"tags"`
	Content string             `json:"content" bson:"content" binding:"required,max=1000,min=5"`
}

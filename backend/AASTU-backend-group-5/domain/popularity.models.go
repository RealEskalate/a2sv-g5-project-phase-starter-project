package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Like struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	PostID primitive.ObjectID `json:"post_id" bson:"post_id"`
}

type DisLike struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	PostID primitive.ObjectID `json:"post_id" bson:"post_id"`
}

type Comment struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id"`
	PostID  primitive.ObjectID `json:"post_id" bson:"post_id"`
	Content string             `json:"comment" bson:"comment"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Session struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       string             `bson:"user_id" json:"user_id"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
	AccessToken  string             `bson:"access_token" json:"access_token"`
}

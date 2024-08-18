package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthLogin struct {
	UserID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username  		string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password"`
}

type LoginResponse struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	AcessToken   string             `json:"access_token"`
	RefreshToken string             `json:"refresh_token"`
}
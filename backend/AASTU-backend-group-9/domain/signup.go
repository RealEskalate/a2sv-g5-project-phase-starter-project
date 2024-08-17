package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// CreateUser is a struct that contains the username, email, password and role of a user
type CreateUser struct {
	Username  		string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	Role            string `json:"role" binding:"required"`
}

// AuthSignup is a struct that contains the email and password of a user
type AuthSignup struct {
	Username  		string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password"`
}

type SignUpResponse struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	AcessToken      string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
}


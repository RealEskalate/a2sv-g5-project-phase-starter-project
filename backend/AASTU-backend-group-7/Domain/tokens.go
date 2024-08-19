package Domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	UserID        primitive.ObjectID `json:"user_id" bson:"_id,omitempty"`
	Refresh_token string             `json:"refresh_token"`
}

type ResetToken struct {
	Email       string `json:"email"`
	Reset_token string `json:"reset_token"`
}

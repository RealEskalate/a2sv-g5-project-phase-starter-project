package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	TokenID      primitive.ObjectID `json:"token_id" bson:"token_id"`
	AccessToken  string             `json:"access_token" bson:"access_token"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token"`
	Username     string             `json:"username" bson:"username"`
	ExpiresAt    time.Time          `json:"expires_at" bson:"expires_at"` // Add this field to track expiration time
}

func (t *Token) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}

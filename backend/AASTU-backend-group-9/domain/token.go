package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const(
	TokenCollection = "tokens"
)

type Token struct {
     ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    UserID       primitive.ObjectID `bson:"user_id"`
    AccessToken  string             `bson:"access_token" json:"access_token"`
    RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
    ExpiresAt    time.Time          `bson:"expires_at"`
    CreatedAt    time.Time          `bson:"created_at"`
}

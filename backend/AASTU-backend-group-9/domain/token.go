package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id"`
	RefreshToken string             `bson:"refresh_token"`
	ExpiresAt    time.Time          `bson:"expires_at"`
	CreatedAt    time.Time          `bson:"created_at"`
}

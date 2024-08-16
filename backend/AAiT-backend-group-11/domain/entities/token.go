package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       primitive.ObjectID `bson:"userId"`
	AccessToken  string             `bson:"accessToken"`
	RefreshToken string             `bson:"refreshToken"`
	CreatedAt    time.Time          `bson:"createdAt"`
	ExpiresAt    time.Time          `bson:"expiresAt"`
}

type PasswordResetToken struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userId"`
	Token     string             `bson:"token"`
	CreatedAt time.Time          `bson:"createdAt"`
	ExpiresAt time.Time          `bson:"expiresAt"`
}
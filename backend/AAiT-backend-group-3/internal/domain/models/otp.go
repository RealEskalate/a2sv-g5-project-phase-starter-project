package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OtpEntry struct {
	ID        primitive.ObjectID        `bson:"_id,omitempty" json:"_id"`
	OTP       string    				`bson:"otp" json:"otp"`
	UserID    string   					`bson:"user_id" json:"user_id"`
	ExpiresAt time.Time 				`bson:"expires_at" json:"expires_at"`
}
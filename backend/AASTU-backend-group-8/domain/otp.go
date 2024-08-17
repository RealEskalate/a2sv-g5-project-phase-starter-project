package domain

import(
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OTP struct {
	UserID    primitive.ObjectID `bson:"user_id"`
	OTP       string             `bson:"otp"`
	ExpiresAt time.Time          `bson:"expires_at"`
}

type OTPRequest struct {
	OTP string `json:"otp"`
}
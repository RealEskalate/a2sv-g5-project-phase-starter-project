package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OTP struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Email      string             `bson:"email"`
	Code       string             `bson:"code"`
	Expiration time.Time          `bson:"expiration"`
	IsValid    bool     		  `bson:"is_valid"`
}

type ResendOTPRequest struct {
	Email string `json:"email"`
}

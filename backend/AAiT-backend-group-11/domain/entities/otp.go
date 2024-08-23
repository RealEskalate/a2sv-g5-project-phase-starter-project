package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OTP struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Email      string             `bson:"email, required, email"`
	Code       string             `bson:"code, min=5"`
	Expiration time.Time          `bson:"expiration"`
	IsValid    bool     		  `bson:"is_valid"`
}

type ResendOTPRequest struct {
	Email string `json:"email required email"`
}

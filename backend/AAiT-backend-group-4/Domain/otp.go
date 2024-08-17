package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserOTPRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
}

type UserOTPVerification struct {
	ID         primitive.ObjectID `bson:"_id"`
	User_ID    string             `json:"user_id" bson:"user_id"`
	Email      string             `json:"email" bson:"email"`
	OTP        string             `json:"otp" bson:"otp"`
	Created_At time.Time          `json:"created_at" bson:"created_at"`
	Expires_At time.Time          `json:"expires_at" bson:"expires_at"`
}

type OTPVerificationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

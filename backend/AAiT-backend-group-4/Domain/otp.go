package domain

import (
	"context"
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

type OTPUsecase interface {
	GenerateOTP(user *UserOTPRequest) (otp UserOTPVerification, err error)
	VerifyOTP(user *UserOTPVerification) (resp OTPVerificationResponse, err error)
}

type OTPRepository interface {
	CreateOTP(c context.Context, otp *UserOTPVerification) error
	GetOTPByEmail(c context.Context, email string) (otp UserOTPVerification, err error)
	DeleteOTPByEmail(c context.Context, email string) error
}

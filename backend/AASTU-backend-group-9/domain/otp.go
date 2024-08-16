package domain

import "time"

const (
	CollectionOTP = "otp"
)

type OTP struct {
    Value      string    `bson:"value"`
    Username   string    `bson:"username"`
    Email      string    `bson:"email"`
    Password   string    `bson:"password"`
    CreatedAt  time.Time `bson:"created_at"`
    ExpiresAt  time.Time `bson:"expires_at"`
}

type OTPRequest struct {
    Email    string `json:"email" binding:"required"`
    Value    string `json:"value" binding:"required"`
}

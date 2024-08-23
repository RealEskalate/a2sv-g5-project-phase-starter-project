package domain

import (
	"context"
	"time"
)

type UnverifiedUser struct {
	Email     string
	UserToken string
	OTP       string
	ExpiresAt time.Time
}
type UnverifiedUserResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Otp   string `json:"otp"`
}

type UnverifiedUserRepository interface {
	StoreUnverifiedUser(ctx context.Context, uv UnverifiedUser) error
	FindUnverifiedUser(ctx context.Context, email string) (UnverifiedUser, error)
	DeleteUnverifiedUser(ctx context.Context, email string) error
	UpdateOTP(ctx context.Context, email string, otp string, expiry time.Time) (UnverifiedUserResponse, error)
}

package domain

import (
	"context"
)

// TODO: do email and password validation here
type UserSignUp struct {
	Username string `json:"username" bson:"username" binding:"required" `
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type VerifyEmailRequest struct {
	Email string `form:"email" binding:"required,email"`
	OTP   string `form:"otp" binding:"required"`
}

type ResendOTPRequest struct {
	Email string `form:"email" binding:"required,email"`
}

type FederatedSignupRequest struct {
	Provider string `json:"provider" binding:"required"`
	Token    string `json:"token" binding:"required"`

}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	CreateUser(c context.Context, user *UserSignUp) (*User, error) // do validation here
	VerifyEmail(c context.Context, req *VerifyEmailRequest) (*UserResponse, error)
	ResendOTP(c context.Context, req *ResendOTPRequest) error
	HandleFederatedSignup(c context.Context, token string) (*User, error)
	CreateTokens(c context.Context, user *User) (*TokenResponse, error)
}

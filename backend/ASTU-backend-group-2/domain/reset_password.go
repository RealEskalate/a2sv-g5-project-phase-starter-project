package domain

import (
	"context"
	"time"
)

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email" binding:"required"`
	Code     string `json:"code" binding:"required"`
	NewPassword  string `json:"password" binding:"required"`
}
type OtpSave struct{
	Email    string `json:"email" binding:"required"`
	Code     string `json:"code" binding:"required"`
	ExpiresAt time.Time `json:"expiresat" sql:"expiresat"`
}
type ResetPasswordUsecase interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	ResetPassword(c context.Context, userID string ,resetPassword *ResetPasswordRequest) error
	SaveOtp(c context.Context, otp *OtpSave) error
	DeleteOtp(c context.Context, email string) error
	GetOTPByEmail(c context.Context, email string) (*OtpSave, error)
	

}

type ResetPasswordRepository interface {
	GetUserByEmail(c context.Context, email string) (*User, error)
	ResetPassword(c context.Context, userID string ,resetPassword *ResetPasswordRequest) error
	SaveOtp(c context.Context, otp *OtpSave) error
	GetOTPByEmail(c context.Context, email string) (*OtpSave, error)
	DeleteOtp(c context.Context, email string) error
}

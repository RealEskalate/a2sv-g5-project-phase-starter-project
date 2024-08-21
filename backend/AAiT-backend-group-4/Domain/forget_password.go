package domain

import "context"


type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}


type ForgotPasswordResponse struct {
	Message string `json:"message"`
}

type ResetPasswordRequest struct {
    Email         string `json:"email"`
    OTP           string `json:"otp"`
    NewPassword   string `json:"new_password"`
}


type ResetPasswordResponse struct {
	Message string `json:"message"`
}

type ForgotPasswordUsecase interface {
	ForgotPassword(ctx context.Context, request ForgotPasswordRequest) (ForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, request ResetPasswordRequest) (ResetPasswordResponse, error)
}

type EmailService interface {
	SendPasswordResetEmail(email, token string) error
}

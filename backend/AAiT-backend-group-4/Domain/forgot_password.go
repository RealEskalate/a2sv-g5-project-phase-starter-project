package domain

import "context"

type ForgotPasswordRequest struct {
	Confirmation string `json:"confirmation" validate:"required"`
	New_Password string `json:"new_password" validate:"required"`
}

type ForgotPasswordResponse struct {
	Message       string `json:"message"`
	User_ID       string `json:"user_id"`
	Access_Token  string `json:"access_token"`
	Refresh_Token string `json:"refresh_token"`
}

type ForgotPasswordUsecase interface {
	ForgotPassword(c context.Context, email string, key string) (resp OTPVerificationResponse, err error)
	VerifyChangePassword(c context.Context, email string, request ForgotPasswordRequest) (resp ForgotPasswordResponse, err error)
	GetByEmail(c context.Context, email string) (user User, err error)
	GetByUsername(c context.Context, userName string) (user User, err error)
}

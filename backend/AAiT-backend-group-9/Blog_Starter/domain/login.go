package domain

import "context"

type UserLogin struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type LoginResponse struct {
	UserID       string `json:"id"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type ChangePasswordRequest struct {
	OTP      string `form:"otp" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginUsecase interface {
	Login(c context.Context, user *UserLogin) (*LoginResponse, error)
	UpdatePassword(c context.Context, req ChangePasswordRequest, userID string) error
}


type LoginRepository interface{
	Login(c context.Context, user *UserLogin ) (*LoginResponse, error) // check things here later
	UpdatePassword(c context.Context, req ChangePasswordRequest, userID string) (error)
}


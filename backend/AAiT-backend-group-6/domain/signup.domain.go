package domain

import (
	"context"
)

type SignupRequest struct {
	Name     	string `json:"name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Email    	string `json:"email" binding:"required,email"`
	Password 	string `json:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type VerifyEmailRequest struct{
	Email    	        string `form:"email" binding:"required,email"`
	Verification_code 	string `json:"verification_code"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	VerifyEmail(c context.Context, email string, code string) error
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
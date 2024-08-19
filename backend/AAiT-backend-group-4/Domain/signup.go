package domain

import "context"

type SignupRequest struct {
	First_Name string  `json:"first_name" validate:"required,min=2,max=100"`
	Last_Name  string  `json:"last_name" validate:"required,min=2,max=100"`
	User_Name  string  `json:"user_name" validate:"required,min=5"`
	Email      string  `json:"email" validate:"required,email"`
	Password   string  `json:"password" validate:"required,min=6"`
	Phone      *string `json:"phone"`
	Bio        *string `json:"bio"`
}

type SignupRespnse struct {
	Message       string `json:"message"`
	User_ID       string `json:"user_id"`
	Access_Token  string `json:"access_token"`
	Refresh_Token string `json:"refresh_token"`
}

type SignupUsecase interface {
	Signup(c context.Context, user *SignupRequest) (resp OTPVerificationResponse, err error)
	GetByEmail(c context.Context, email string) (user User, err error)
	GetByUsername(c context.Context, userName string) (user User, err error)
	CreateAllTokens(user *User, accessSecret string, refreshSecret string,
		accessExpiry int, refreshExpiry int) (accessToken string, refreshToken string, err error)
}

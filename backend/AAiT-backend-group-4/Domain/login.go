package domain

import "context"

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginUsecase interface {
	GetByIDentifier(c context.Context, identifier string) (user User, err error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (accessToken string, err error)
}

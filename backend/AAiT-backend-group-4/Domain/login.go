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
	CreateAllTokens(user *User, accessSecret string, refreshSecret string,
		accessExpiry int, refreshExpiry int) (accessToken string, refreshToken string, err error)
}

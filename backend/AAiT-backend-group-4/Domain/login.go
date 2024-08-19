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
	LoginWithIdentifier(c context.Context, identifier string) (accesToken string, refreshToken string, err error)
	CreateAllTokens(c context.Context, user *User, accessSecret string, refreshSecret string,
		accessExpiry int, refreshExpiry int) (accessToken string, refreshToken string, err error)
}

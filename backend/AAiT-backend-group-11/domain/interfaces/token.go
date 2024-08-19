package interfaces

import "backend-starter-project/domain/entities"


type RefreshTokenRepository interface{
	CreateRefreshToken(refreshToken *entities.RefreshToken) (*entities.RefreshToken, error)
	FindRefreshTokenByUserId(userId string) (*entities.RefreshToken, error)
	DeleteRefreshTokenByUserId(userId string) error
}

type TokenService interface {
	VerifyAccessToken(token string) (*entities.User, error)
	GenerateAccessToken(userId string) (string, error)
	GenerateRefreshToken(token string) (string, error)
	VerifyRefreshToken(token string) (string, error)
	InvalidateAccessToken(token string) (string, error)
	InvalidateRefreshToken(token string) (string, error)
}

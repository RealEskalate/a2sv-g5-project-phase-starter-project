package interfaces

import "backend-starter-project/domain/entities"

type TokenService interface {
	VerifyAccessToken(token string) (*entities.User, error)
	GenerateAccessToken(userId string) (*entities.Token, error)
	GenerateRefreshToken(token string) (*entities.Token, error)
	VerifyRefreshToken(token string) (*entities.Token, error)
	InvalidateAccessToken(token string) (*entities.Token, error)
	InvalidateRefreshToken(token string) (*entities.Token, error)
}
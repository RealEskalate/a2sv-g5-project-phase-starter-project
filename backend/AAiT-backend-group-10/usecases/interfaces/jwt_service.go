package interfaces

import (
	"aait.backend.g10/domain"
	"github.com/golang-jwt/jwt/v4"
)

type IJwtService interface {
	GenerateToken(user *domain.User) (string, string, *domain.CustomError)
	ValidateToken(token string) (*jwt.Token, *domain.CustomError)
	GenerateResetToken(email string, code int64) (string, *domain.CustomError)
	CheckToken(authPart string) (*jwt.Token, *domain.CustomError)
	FindClaim(token *jwt.Token) (jwt.MapClaims, bool)
	GenerateActivationToken(email string) (string, *domain.CustomError)
}

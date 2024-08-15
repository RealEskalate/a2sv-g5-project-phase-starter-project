package infrastructure

import (
	domain "task/Domain"

	"github.com/dgrijalva/jwt-go"
)

// JWTService represents a service for generating and validating JWT tokens.
type JWTService struct {
	SecretKey []byte
}

// NewJWTService creates a new instance of JWTService with the given secret key.
func NewJWTService(secretKey string) *JWTService {
	return &JWTService{
		SecretKey: []byte(secretKey),
	}
}

// GenerateToken generates a new JWT token with the given claims.
func (s *JWTService) GenerateToken(claims *domain.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.SecretKey)
}

// ValidateToken validates the given JWT token and returns the claims if the token is valid.
func (s *JWTService) ValidateToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}

// AuthMiddleware is a middleware function that validates the JWT token in the request header.

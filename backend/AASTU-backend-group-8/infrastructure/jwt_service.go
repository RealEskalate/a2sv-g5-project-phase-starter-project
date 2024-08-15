package infrastructure

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// JWTService interface
type JWTService interface {
	GenerateToken(userID string, isAdmin bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	UserID  string `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService creates a new JWTService
func NewJWTService(secretKey, issuer string) JWTService {
	return &jwtService{
		secretKey: secretKey,
		issuer:    issuer,
	}
}

// GenerateToken generates a new JWT token
func (j *jwtService) GenerateToken(userID string, isAdmin bool) (string, error) {
	claims := &jwtCustomClaims{
		UserID:  userID,
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// ValidateToken validates the given JWT token
func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
}

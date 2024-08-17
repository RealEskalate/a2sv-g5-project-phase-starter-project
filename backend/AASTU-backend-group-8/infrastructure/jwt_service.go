package infrastructure

import (
	"time"

	"meleket/domain"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtService struct {
	secretKey        string
	refreshSecretKey string
}

// NewJWTService creates a new JWTService
func NewJWTService(secretKey, refreshSecretKey string) *JwtService {
	return &JwtService{
		secretKey:        secretKey,
		refreshSecretKey: refreshSecretKey,
	}
}

// GenerateToken generates a new JWT token
func (j *JwtService) GenerateToken(userID primitive.ObjectID, username, role string) (string, error) {
	claims := &domain.Claims{
		ID:       userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // Shorter expiry time
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// GenerateRefreshToken generates a new refresh JWT token
func (j *JwtService) GenerateRefreshToken(userID primitive.ObjectID, username, role string) (string, error) {
	claims := &domain.Claims{
		ID:       userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // Longer expiry time for refresh token
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.refreshSecretKey))
}

// ValidateToken validates the given JWT token
func (j *JwtService) ValidateToken(tokenString string) (*jwt.Token, *domain.Claims, error) {
	claims := &domain.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// Check that the signing method is what we expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, nil, err
	}

	return token, claims, nil
}

// ValidateRefreshToken validates the given refresh JWT token
func (j *JwtService) ValidateRefreshToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Check that the signing method is what we expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.refreshSecretKey), nil
	})
}

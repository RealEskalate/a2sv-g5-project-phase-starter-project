package infrastracture

import (
	"group3-blogApi/config"
	"group3-blogApi/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

// TokenGenerator implementation
type TokenGeneratorImpl struct{}

// NewTokenGenerator creates a new TokenGenerator instance
func NewTokenGenerator() domain.TokenGenerator {
	return &TokenGeneratorImpl{}
}

// GenerateToken generates an access token for the user
func (tg *TokenGeneratorImpl) GenerateToken(user domain.User) (string, error) {
	accessTokenSecret := []byte(config.EnvConfigs.JwtSecret)
	accessTokenExpiryHour := config.EnvConfigs.AccessTokenExpiryHour

	claims := domain.JwtCustomClaims{
		Authorized:  true,
		UserID:      user.ID.Hex(),
		Role:        user.Role,
		Username:    user.Username,
		IsActivated: user.IsActive,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(accessTokenExpiryHour)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(accessTokenSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}

// GenerateRefreshToken generates a refresh token for the user
func (tg *TokenGeneratorImpl) GenerateRefreshToken(user domain.User) (string, error) {
	refreshTokenSecret := []byte(config.EnvConfigs.JwtRefreshSecret)
	refreshTokenExpiryHour := config.EnvConfigs.RefreshTokenExpiryHour

	claims := domain.JwtCustomClaims{
		Authorized:  true,
		UserID:      user.ID.Hex(),
		Role:        user.Role,
		Username:    user.Username,
		IsActivated: user.IsActive,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(refreshTokenExpiryHour)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(refreshTokenSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}

// RefreshToken parses and verifies a refresh token and returns the user ID
func (tg *TokenGeneratorImpl) RefreshToken(tokenString string) (string, error) {
	refreshTokenSecret := []byte(config.EnvConfigs.JwtRefreshSecret)

	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenSecret, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return "", err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", err
	}

	return userID, nil
}

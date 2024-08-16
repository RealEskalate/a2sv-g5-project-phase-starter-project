package config

import (
	"blogs/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// Define your secret key and token expiration times
var (
	accessSecretKay        = []byte("my-access-secret-key")
	refreshSecretKay       = []byte("my-refresh-secret-key")
	passwordResetSecretKay = []byte("my-password-reset-secret-key")

	accessTokenExpiry        = 15 * time.Minute    // 15 minutes
	refreshTokenExpiry       = 30 * 24 * time.Hour // 30 days
	passwordResetTokenExpiry = 1 * time.Hour       // 1 hour
)

func GenerateToken(username, role, tokenType string) (string, *domain.Token, error) {
	var expiry time.Duration
	var secretKey []byte

	switch tokenType {
	case "access":
		expiry = accessTokenExpiry
		secretKey = accessSecretKay
	case "refresh":
		expiry = refreshTokenExpiry
		secretKey = refreshSecretKay
	case "password-reset":
		expiry = passwordResetTokenExpiry
		secretKey = passwordResetSecretKay
	default:
		return "", nil, errors.New("invalid token type")
	}

	expireTime := time.Now().Add(expiry).Unix()
	token := &domain.Token{
		Username:  username,
		ExpiresAt: expireTime,
	}

	claims := domain.Claim{
		Username: username,
		Role:     role,
		Type:     tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiry).Unix(),
		},
	}

	claimsToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := claimsToken.SignedString(secretKey)
	if err != nil {
		return "", nil, err
	}

	return signedToken, token, nil
}

func ValidateToken(tokenString string, tokenType string) (*domain.Claim, error) {
	var secretKey []byte

	switch tokenType {
	case "access":
		secretKey = accessSecretKay
	case "refresh":
		secretKey = refreshSecretKay
	case "password-reset":
		secretKey = passwordResetSecretKay
	default:
		return nil, errors.New("invalid token type")
	}

	token, err := jwt.ParseWithClaims(tokenString, &domain.Claim{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signature method matches the one used
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*domain.Claim)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

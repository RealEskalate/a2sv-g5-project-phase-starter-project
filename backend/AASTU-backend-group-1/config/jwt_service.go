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
	registerSecretKay      = []byte("my-register-secret-key")

	accessTokenExpiry        = 15 * time.Minute    // 15 minutes
	refreshTokenExpiry       = 30 * 24 * time.Hour // 30 days
	passwordResetTokenExpiry = 1 * time.Hour       // 1 hour
	registerTokenExpiry      = 24 * time.Hour      // 24 hours
)

func GenerateToken(claims domain.Claims, tokenType string) (string, *domain.Token, error) {
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
	case "register":
		expiry = registerTokenExpiry
		secretKey = registerSecretKay
	default:
		return "", nil, errors.New("invalid token type")
	}

	expireTime := time.Now().Add(expiry).Unix()
	token := &domain.Token{
		Username:  claims.GetUsername(),
		ExpiresAt: expireTime,
	}

	claims.SetExpiresAt(expireTime)

	claimsToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := claimsToken.SignedString(secretKey)
	if err != nil {
		return "", nil, err
	}

	return signedToken, token, nil
}

func ValidateToken(tokenString string, tokenType string) (domain.Claims, error) {
	var secretKey []byte
	var claims domain.Claims

	switch tokenType {
	case "access":
		secretKey = accessSecretKay
		claims = &domain.LoginClaims{}
	case "refresh":
		secretKey = refreshSecretKay
		claims = &domain.LoginClaims{}
	case "password-reset":
		secretKey = passwordResetSecretKay
		claims = &domain.PasswordResetClaims{}
	case "register":
		secretKey = registerSecretKay
		claims = &domain.RegisterClaims{}
	default:
		return nil, errors.New("invalid token type")
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signature method matches the one used
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(domain.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

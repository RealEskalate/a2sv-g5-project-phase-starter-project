package config

import (
	"blogs/domain"
	"errors"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(claims domain.Claims) (string, error) {
	claims.SetExpiry()
	secretKey := claims.GetSecretKey()

	claimsToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := claimsToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string, claims domain.Claims) error {
	secretKey := claims.GetSecretKey()

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	_, ok := token.Claims.(domain.Claims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

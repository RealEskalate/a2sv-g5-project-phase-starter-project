package infrastructures

import (
	"errors"
	"loan-management/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateVerificationToken(email string, expirationTime time.Time) (string, error) {
	claims := &jwt.StandardClaims{Subject: email, ExpiresAt: expirationTime.Unix()}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	config, err := config.LoadConfig()
	if err != nil {
		return "", err
	}
	return token.SignedString([]byte(config.Jwt.Secret))
}

func ValidateVerificationToken(tokenString, expectedEmail string) error {
	config, err := config.LoadConfig()
	if err != nil {
		return err
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt.Secret), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		if claims.Subject != expectedEmail {
			return errors.New("token subject does not match expected email")
		}
		if claims.ExpiresAt < time.Now().Unix() {
			return errors.New("token has expired")
		}
		return nil
	}

	return errors.New("invalid token")
}

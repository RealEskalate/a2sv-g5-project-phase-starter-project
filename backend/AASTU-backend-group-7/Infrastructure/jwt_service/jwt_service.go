package jwtservice

import (
	"blogapp/Config"
	"blogapp/Domain"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(existingUser Domain.User) (string, error) {
	userclaims := &Domain.AccessClaims{
		ID:   existingUser.ID,
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}

	// Create a new JWT token with the user claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userclaims)

	// Ensure Config.JwtSecret is of type []byte
	jwtToken, err := token.SignedString([]byte(Config.JwtSecret))
	return jwtToken, err
}

func CreateRefreshToken(existingUser Domain.User) (refreshToken string, err error) {
	userclaims := &Domain.RefreshClaims{
		ID:   existingUser.ID,
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
		},
	}

	// Create a new JWT token with the user claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userclaims)

	// Ensure Config.JwtSecret is of type []byte
	jwtToken, err := token.SignedString([]byte(Config.JwtSecret))
	return jwtToken, err
}

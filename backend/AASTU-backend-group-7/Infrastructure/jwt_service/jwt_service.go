package jwtservice

import (
	config "blogapp/Config"
	"blogapp/Domain"
	"time"

	"github.com/golang-jwt/jwt"
)

func SignJwt(existingUser Domain.User) (string, error) {
	userclaims := &Domain.Claims{
		ID:   existingUser.ID,
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}

	// Create a new JWT token with the user claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userclaims)

	// Ensure config.JwtSecret is of type []byte
	jwtToken, err := token.SignedString([]byte(config.JwtSecret))
	return jwtToken, err
}

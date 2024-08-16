package jwtservice

import (
	Config "blogapp/Config"
	"blogapp/Domain"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func VerifyRefreshToken(tokenString string, userid primitive.ObjectID) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.JwtSecret), nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*Domain.RefreshClaims)
	if !ok || !token.Valid {
		fmt.Println(ok)
		fmt.Print(token.Valid)
		fmt.Println(claims.ID)
		fmt.Println("first one")
		return errors.New("Invalid refresh token")
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return errors.New("Refresh token has expired")
	}

	if claims.ID != userid {
		fmt.Println("second one")
		return errors.New("Invalid refresh token")
	}

	return nil
}

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
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
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
	fmt.Println("inside verify refresh token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.JwtSecret), nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("Invalid refresh token")
	}

	claimUserID, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return errors.New("Invalid token")
	}

	// Check if the token is expired
	expTime := int64(claims["exp"].(float64))
	if time.Unix(expTime, 0).Before(time.Now()) {
		return errors.New("Token expired")
	}

	if claimUserID != userid {
		return errors.New("Invalid refresh token")
	}

	return nil
}


package main

import (
	domain "blogs/Domain"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))

	// Create claims
	claims := &domain.JwtCustomClaims{
		Name: user.Full_Name,
		ID:   user.ID.Hex(),
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp), // Convert expiration time to *jwt.NumericDate
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}
func main(){
	id, _ :=  primitive.ObjectIDFromHex("66c0a0dcdb2272faca4591ae")
	fmt.Println(CreateAccessToken(&domain.User{
		Full_Name: "Full_Name",
		ID:       id,
		Role:      "user",
	}, "access_token_secret", 10))
}

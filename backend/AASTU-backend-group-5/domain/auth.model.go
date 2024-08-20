package domain

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	v"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	jwt.StandardClaims
	ID      primitive.ObjectID
	Name    string
	Email   string
	IsAdmin bool
}

type EmailUserClaims struct {
	v.Claims
	ID      primitive.ObjectID
	Email   string
}
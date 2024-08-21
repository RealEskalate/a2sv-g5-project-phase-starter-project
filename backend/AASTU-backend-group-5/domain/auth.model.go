package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserClaims struct {
	jwt.StandardClaims
	ID      primitive.ObjectID
	Name    string
	Email   string
	IsAdmin bool
}

type State struct{
	StateID string `json:"id" bson:"_id"`
	ExpiresAT time.Time `json:"expiresat" bson:"expiresat"`
}

type EmailUserClaims struct {
	jwt.StandardClaims
	ID    primitive.ObjectID `json:"_id"`
	Email string             `json:"email"`
}
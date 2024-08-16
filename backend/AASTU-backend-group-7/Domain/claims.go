package Domain

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccessClaims struct {
	ID             primitive.ObjectID `json:"id"`
	Role           string             `json:"role"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	ID             primitive.ObjectID `json:"id"`
	Role           string             `json:"role"`
	jwt.StandardClaims
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	userID       primitive.ObjectID `json:"user_id" bson:"_id,omitempty"`
	RefreshToken string `json:"refresh_token"`
}
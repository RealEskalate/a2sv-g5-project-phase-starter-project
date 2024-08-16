package Domain

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccessClaims struct {
	ID   primitive.ObjectID `json:"id"`
	Role string             `json:"role"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	ID   primitive.ObjectID `json:"id"`
	Role string             `json:"role"`
	jwt.StandardClaims
}

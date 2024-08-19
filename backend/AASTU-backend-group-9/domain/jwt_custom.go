package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtCustomClaims struct {
	UserID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string `json:"email"`
	Username   string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	Username string `json:"usernamae"`
	jwt.StandardClaims
}
package Domain

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	ID             primitive.ObjectID `json:"id"`
	Role           string             `json:"role"`
	StandardClaims jwt.StandardClaims `json:"standard_claims"`
}

func (c *Claims) Valid() error {
	return c.StandardClaims.Valid()
}

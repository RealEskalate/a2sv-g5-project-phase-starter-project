package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id" json:"userId"` // References the User ObjectID
	AccessToken  string             `bson:"access_token" json:"accessToken"`
	RefreshToken string             `bson:"refresh_token" json:"refreshToken"`
	ExpiresAt    time.Time          `bson:"expires_at" json:"expiresAt"`
	CreatedAt    time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updatedAt"`
}

type JwtCustomClaims struct {
	Authorized bool   `json:"authorized"`
	UserID     string `json:"user_id"`
	Role       string `json:"role"`
	Username   string `json:"username"`

	jwt.StandardClaims
}

// type JwtCustomClaimsInterface interface {
// 	Authorized() bool
// 	UserID() string
// 	Role() string
// 	Username() string
// }

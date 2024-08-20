package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshToken struct {
	// username     string    `bson:"token" json:"token"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Role      string             `bson:"role" json:"role"`
	ExpiresAt time.Time          `bson:"expires_at" json:"expires_at"`
}

type RefreshTokenRequest struct {
    RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenUsecaseInterface interface {
  RefreshToken(userID primitive.ObjectID, role string) (string, error)
	DeleteRefreshToken(userID primitive.ObjectID) error 
}

type TokenRepositoryInterface interface {
	SaveRefreshToken(refreshToken *RefreshToken) error
	FindRefreshToken(userID primitive.ObjectID) (*RefreshToken, error)
	DeleteRefreshTokenByUserID(userID primitive.ObjectID) error
}

// IsExpired checks if the refresh token is expired
// func (r *RefreshToken) IsExpired() bool {
// 	return time.Now().After(r.ExpiresAt)
// }

package domain

import (
    "time"
	"context"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
     ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    UserID       primitive.ObjectID `bson:"user_id"`
    AccessToken  string             `bson:"access_token" json:"access_token"`
    RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
    ExpiresAt    time.Time          `bson:"expires_at"`
    CreatedAt    time.Time          `bson:"created_at"`
}
type TokenRepository interface {
    SaveToken(ctx context.Context, token *Token) error
    FindTokenByAccessToken(ctx context.Context, accessToken string) (*Token, error)
    DeleteToken(ctx context.Context, tokenID primitive.ObjectID) error
    FindTokenByRefreshToken(ctx context.Context, refreshToken string) (*Token, error)
}
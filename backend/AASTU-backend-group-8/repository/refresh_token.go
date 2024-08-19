package repository

import (
	"context"
	"meleket/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenRepository struct {
	collection domain.Collection
}

func NewTokenRepository(col domain.Collection) *TokenRepository {
	return &TokenRepository{
		collection: col,
	}
}

// SaveRefreshToken saves the refresh token in the database
func (r *TokenRepository) SaveRefreshToken(refreshToken *domain.RefreshToken) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, refreshToken)
	return err
}

func (r *TokenRepository) FindRefreshToken(token string) (*domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken
	err := r.collection.FindOne(context.TODO(), bson.M{"token": token}).Decode(&refreshToken)
	return &refreshToken, err
}

func (r *TokenRepository) DeleteRefreshTokenByUserID(userID primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"user_id": userID})
	return err
}

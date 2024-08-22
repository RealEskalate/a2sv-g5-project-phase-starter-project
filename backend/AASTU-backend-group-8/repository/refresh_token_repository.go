package repository

import (
	"context"
	"meleket/domain"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenRepository struct {
	collection domain.Collection
	mutex		sync.RWMutex
}

func NewTokenRepository(col domain.Collection) *TokenRepository {
	return &TokenRepository{
		collection: col,
		mutex: 	sync.RWMutex{},
	}
}

// SaveRefreshToken saves the refresh token in the database
func (r *TokenRepository) SaveRefreshToken(refreshToken *domain.RefreshToken) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r.mutex.RLock()
	defer r.mutex.RUnlock()
	_, err := r.collection.InsertOne(ctx, refreshToken)
	return err
}

func (r *TokenRepository) FindRefreshToken(userID primitive.ObjectID) (*domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken
	err := r.collection.FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&refreshToken)
	return &refreshToken, err
}

func (r *TokenRepository) DeleteRefreshTokenByUserID(userID primitive.ObjectID) error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"user_id": userID})
	return err
}

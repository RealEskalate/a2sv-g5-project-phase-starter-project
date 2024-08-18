package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository struct {
	collection *mongo.Collection
}

// CreateRefreshToken implements interfaces.RefreshTokenRepository.
func (t *TokenRepository) CreateRefreshToken(refreshToken *entities.RefreshToken) (*entities.RefreshToken, error) {
	panic("unimplemented")
}

// DeleteRefreshTokenByUserId implements interfaces.RefreshTokenRepository.
func (t *TokenRepository) DeleteRefreshTokenByUserId(userId string) error {
	panic("unimplemented")
}

// FindRefreshTokenByUserId implements interfaces.RefreshTokenRepository.
func (t *TokenRepository) FindRefreshTokenByUserId(userId string) (*entities.RefreshToken, error) {
	panic("unimplemented")
}

func NewRefreshTokenRepository(collection *mongo.Collection) interfaces.RefreshTokenRepository {
	return &TokenRepository{
		collection: collection,
	}
}

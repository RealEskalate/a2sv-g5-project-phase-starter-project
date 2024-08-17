package repository

import (
	"blog/database"
	"blog/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

type MongoTokenRepository struct {
    db database.Collection
    Collection string
}

func NewMongoTokenRepository(db database.Collection) domain.TokenRepository {
    return &MongoTokenRepository{db: db}
}
func (repo *MongoTokenRepository) SaveToken(ctx context.Context, token *domain.Token) error {
    _, err := repo.db.InsertOne(ctx, token)
    return err
}

func (repo *MongoTokenRepository) FindTokenByAccessToken(ctx context.Context, accessToken string) (*domain.Token, error) {
    var token domain.Token
    err := repo.db.FindOne(ctx, bson.M{"access_token": accessToken}).Decode(&token)
    return &token, err
}

func (repo *MongoTokenRepository) FindTokenByRefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error) {
	var token domain.Token
	err := repo.db.FindOne(ctx, bson.M{"refresh_token": refreshToken}).Decode(&token)
	return &token, err
}

func (repo *MongoTokenRepository) DeleteToken(ctx context.Context, tokenID primitive.ObjectID) error {
	_, err := repo.db.DeleteOne(ctx, bson.M{"_id": tokenID})
	return err
}

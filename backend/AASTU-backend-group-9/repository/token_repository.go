package repository

import (
	"context"
	"errors"
	"blog/database"
	"blog/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoTokenRepository struct {
	db         database.Database
	collection string
}

func NewMongoTokenRepository(db database.Database) domain.TokenRepository {
	return &MongoTokenRepository{
		db:         db,
		collection: domain.TokenCollection,
	}
}

func (repo *MongoTokenRepository) SaveToken(ctx context.Context, token *domain.Token) error {
	collection := repo.db.Collection(repo.collection)
	_, err := collection.InsertOne(ctx, token)
	return err
}


func (repo *MongoTokenRepository) FindTokenByAccessToken(ctx context.Context, accessToken string) (*domain.Token, error) {
	var token domain.Token
	collection := repo.db.Collection(repo.collection)
	err := collection.FindOne(ctx, bson.M{"access_token": accessToken}).Decode(&token)
	return &token, err
}

func (repo *MongoTokenRepository) FindTokenByRefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error) {
	var token domain.Token
	collection := repo.db.Collection(repo.collection)
	err := collection.FindOne(ctx, bson.M{"refresh_token": refreshToken}).Decode(&token)
	return &token, err
}

func (repo *MongoTokenRepository) DeleteToken(ctx context.Context, tokenID primitive.ObjectID) error {
	collection := repo.db.Collection(repo.collection)
	deletedCount, err := collection.DeleteOne(ctx, bson.M{"_id": tokenID})
	if err != nil {
		return err
	}
	if deletedCount == 0 {
		return errors.New("no token found to delete")
	}
	return nil
}

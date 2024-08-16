package repository

import (
    "context"
	"errors"
    "blog/domain"
    "blog/database"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoTokenRepository struct {
    db database.Collection
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
    result := repo.db.FindOne(ctx, bson.M{"access_token": accessToken})
    if err := result.Decode(&token); err != nil {
        return nil, err
    }
    return &token, nil
}

func (repo *MongoTokenRepository) FindTokenByRefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error) {
    var token domain.Token
    result := repo.db.FindOne(ctx, bson.M{"refresh_token": refreshToken})
    if err := result.Decode(&token); err != nil {
        return nil, err
    }
    return &token, nil
}

func (repo *MongoTokenRepository) DeleteToken(ctx context.Context, tokenID primitive.ObjectID) error {
    deletedCount, err := repo.db.DeleteOne(ctx, bson.M{"_id": tokenID})
    if err != nil {
        return err
    }
    if deletedCount == 0 {
        return errors.New("no token found to delete")
    }
    return nil
}

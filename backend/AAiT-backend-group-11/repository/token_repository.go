package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type tokenRepository struct {
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewTokenRepository(db *mongo.Database) interfaces.RefreshTokenRepository {
	return &tokenRepository{Database: db, Collection: db.Collection("refresh_tokens")}
}

func (tr *tokenRepository) CreateRefreshToken(token *entities.RefreshToken) (*entities.RefreshToken, error) {
	userID := token.UserID

	filter := bson.D{{"userId", userID}}

	existed := tr.Collection.FindOne(context.TODO(), filter)
	if existed.Err() != nil {
		_, err := tr.Collection.InsertOne(context.TODO(), token)
		if err != nil {
			return nil, err
		}
		return token, nil
	}
	return nil, errors.New("refresh token already exists")
}

func (tr *tokenRepository) FindRefreshTokenByUserId(user_id string) (*entities.RefreshToken, error) {

	userID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"userId", userID}}

	result := tr.Collection.FindOne(context.TODO(), filter)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var token entities.RefreshToken
	err = result.Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (tr *tokenRepository) DeleteRefreshTokenByUserId(user_id string) error {
	userID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}

	filter := bson.D{{"userId", userID}}

	result := tr.Collection.FindOneAndDelete(context.TODO(), filter)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

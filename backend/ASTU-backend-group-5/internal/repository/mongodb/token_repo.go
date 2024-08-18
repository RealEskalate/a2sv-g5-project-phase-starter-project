package mongodb

import (
	"blogApp/internal/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTokenRepository struct {
	collection *mongo.Collection
}

func NewMongoTokenRepository(collection *mongo.Collection) *MongoTokenRepository {
	return &MongoTokenRepository{collection: collection}
}

func createIndexes(collection *mongo.Collection) error {
	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "token", Value: 1}, {Key: "token_type", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "expiry", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(0),
		},
	}
	_, err := collection.Indexes().CreateMany(context.Background(), indexModels)
	return err
}

func (r *MongoTokenRepository) BlacklistToken(ctx context.Context, token string, tokenType domain.TokenType, expiry time.Time) error {
	blacklistedToken := domain.BlacklistedToken{
		Token:     token,
		TokenType: tokenType,
		Expiry:    expiry,
		CreatedAt: time.Now(),
	}
	_, err := r.collection.InsertOne(ctx, blacklistedToken)
	return err
}

func (r *MongoTokenRepository) RemoveBlacklistedToken(ctx context.Context, token string, tokenType domain.TokenType) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"token": token, "token_type": tokenType})
	return err
}

func (r *MongoTokenRepository) IsTokenBlacklisted(ctx context.Context, token string, tokenType domain.TokenType) (bool, error) {
	var result domain.BlacklistedToken
	err := r.collection.FindOne(ctx, bson.M{"token": token, "token_type": tokenType}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	if result.Expiry.Before(time.Now()) {
		err = r.RemoveBlacklistedToken(ctx, token, tokenType)
		if err != nil {
			return true, err
		}
		return true, nil
	}
	return true, nil
}

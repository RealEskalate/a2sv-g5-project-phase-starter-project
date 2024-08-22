package repositories

import (
	"blog_project/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository struct {
	tokenCollection *mongo.Collection
}

func NewTokenRepository(tokenCollection *mongo.Collection) domain.ITokenRepository {
	return &TokenRepository{tokenCollection: tokenCollection}
}

func (tr *TokenRepository) BlacklistToken(ctx context.Context, token string) error {
	_, err := tr.tokenCollection.InsertOne(ctx, bson.M{"token": token})
	if err != nil {
		return err
	}
	return nil
}

func (tr *TokenRepository) IsBlacklisted(token string) (bool, error) {
	ctx := context.TODO()
	result := tr.tokenCollection.FindOne(ctx, bson.M{"token": token})
	println(result.Raw())
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, result.Err()
	}
	return true, nil
}

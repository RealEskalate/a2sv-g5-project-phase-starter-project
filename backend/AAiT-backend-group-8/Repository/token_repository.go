package repository

import (
	domain "AAiT-backend-group-8/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository struct {
	Collection *mongo.Collection
	Context    context.Context
}

func NewTokenRepository(collection *mongo.Collection, ctx context.Context) domain.ITokenRepository {
	// Assign the collection and context to the struct fields
	return &TokenRepository{
		Collection: collection,
		Context:    ctx,
	}
}

func (tr *TokenRepository) InsertRefresher(credential domain.Credential) error {
	_, err := tr.Collection.InsertOne(tr.Context, credential)
	return err
}

func (tr *TokenRepository) GetRefresher(email string) (string, error) {
	var existingCredential domain.Credential

	filter := bson.D{{Key: "email", Value: email}}

	err := tr.Collection.FindOne(tr.Context, filter).Decode(&existingCredential)

	if err != nil {
		return "", err
	}

	return existingCredential.Refresher, err
}

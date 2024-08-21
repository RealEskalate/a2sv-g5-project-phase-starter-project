package mongodb

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

func NewTokenRepository(collection *mongo.Collection, ctx context.Context) *TokenRepository {
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

func (tr *TokenRepository) StoreResetToken(email string, token string) error {
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"reset_token": token}}

	_, err := tr.Collection.UpdateOne(tr.Context, filter, update)
	return err
}

func (tr *TokenRepository) InvalidateResetToken(email string) error {
	filter := bson.M{"email": email}
	update := bson.M{"$unset": bson.M{"reset_token": ""}}

	_, err := tr.Collection.UpdateOne(tr.Context, filter, update)
	return err
}

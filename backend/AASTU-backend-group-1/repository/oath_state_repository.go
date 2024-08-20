package repository

import (
	"blogs/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OAuthRespository struct {
	collection *mongo.Collection
}

func NewOAuthRepository(database *mongo.Database) *OAuthRespository {
	return &OAuthRespository{collection: database.Collection("oauth_states")}
}

func (o *OAuthRespository) InsertState(state *domain.OAuthState) error {
	_, err := o.collection.InsertOne(context.Background(), state)
	if err != nil {
		return err
	}

	return nil
}

func (o *OAuthRespository) GetState(stateString string) (*domain.OAuthState, error) {
	result := &domain.OAuthState{}
	var filter = bson.M{"_id": stateString}

	err := o.collection.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *OAuthRespository) DeleteState(state *domain.OAuthState) error {
	_, err := o.collection.DeleteOne(context.Background(), state)
	if err != nil {
		return err
	}

	return nil
}

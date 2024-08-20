package repository

import (
	"context"
	"errors"
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepo struct {
	Collection database.CollectionInterface
}

func NewAuthRepo(coll database.CollectionInterface) *AuthRepo {
	return &AuthRepo{
		Collection: coll,
	}
}

func (repo *AuthRepo) SaveUser(user *domain.User) error {
	_, err := repo.Collection.InsertOne(context.TODO(), user)
	return err
}

func (repo *AuthRepo) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := repo.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (repo *AuthRepo) FindUserByID(id string) (*domain.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	var user domain.User
	err = repo.Collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (repo *AuthRepo) FindUserByOAuthID(provider, oauthID string) (*domain.User, error) {
	var user domain.User
	err := repo.Collection.FindOne(context.TODO(), bson.M{
		"oauth_provider": provider,
		"oauth_id":       oauthID,
	}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}



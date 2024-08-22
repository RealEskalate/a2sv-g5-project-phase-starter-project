package repository

import (
	"context"
	"errors"

	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthRepo struct {
	Collection database.CollectionInterface
}

func (repo *AuthRepo) EnsureIndexes() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, 
		Options: options.Index().SetUnique(true),
	}
	
	_, err := repo.Collection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}

func NewAuthRepo(coll database.CollectionInterface) (*AuthRepo, error) {
	AR := &AuthRepo{
		Collection : coll,
	}

	// Ensure indexes are created
	if err := AR.EnsureIndexes(); err != nil {
		return nil, err
	}

	return AR, nil
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




package repositories

import (
	domain "blogs/Domain"
	"blogs/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type activeUserRepository struct {
	database   mongo.Database
	collection string
}

func NewActiveUserRepository(db mongo.Database, collection string) domain.ActiveUserRepository {
	return &activeUserRepository{
		database:   db,
		collection: collection,
	}
}

// CreateActiveUser implements domain.ActiveUserRepository.
func (a *activeUserRepository) CreateActiveUser(au domain.ActiveUser, c context.Context) error {
	collection := a.database.Collection(a.collection)

	_, err := collection.InsertOne(c, au)

	return err
}

// DeleteActiveUser implements domain.ActiveUserRepository.
func (a *activeUserRepository) DeleteActiveUser(ids string, user_agent string, c context.Context) error {
	id, err := primitive.ObjectIDFromHex(ids)
	if err != nil {
		return err
	}
	collction := a.database.Collection(a.collection)
	_, err = collction.DeleteOne(c, bson.M{"id": id, "user_agent": user_agent})
	return err
}

// FindActiveUser implements domain.ActiveUserRepository.
func (a *activeUserRepository) FindActiveUser(ids string, user_agent string, c context.Context) (domain.ActiveUser, error) {
	collection := a.database.Collection(a.collection)
	id, err := primitive.ObjectIDFromHex(ids)
	if err != nil {
		return domain.ActiveUser{}, err
	}
	var au domain.ActiveUser
	err = collection.FindOne(c, bson.M{"id": id, "user_agent": user_agent}).Decode(&au)
	return au, err
}
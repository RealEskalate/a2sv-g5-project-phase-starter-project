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

// DeleteActiveUserById implements domain.ActiveUserRepository.
func (a *activeUserRepository) DeleteActiveUserById(id primitive.ObjectID, c context.Context) error {
	collction := a.database.Collection(a.collection)
	_, err := collction.DeleteOne(c, bson.M{"_id": id})
	return err
}

// FindActiveUserById implements domain.ActiveUserRepository.
func (a *activeUserRepository) FindActiveUserById(id primitive.ObjectID, c context.Context) (domain.ActiveUser, error) {
	collection := a.database.Collection(a.collection)
	var au domain.ActiveUser
	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&au)
	return au, err
}


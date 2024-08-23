package repositories

import (
	domain "blogs/Domain"
	"blogs/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type unverifiedUserRepo struct {
	database   mongo.Database
	collection string
}

func NewUnverifiedUserRepository(db mongo.Database, collection string) domain.UnverifiedUserRepository {
	return &unverifiedUserRepo{
		database:   db,
		collection: collection,
	}
}

// DeleteUnverifiedUser implements domain.UnverifiedUserRepository.
func (u *unverifiedUserRepo) DeleteUnverifiedUser(ctx context.Context, email string) error {
	collection := u.database.Collection(u.collection)
	_, err := collection.DeleteOne(ctx, bson.M{"email": email})
	return err
}

// FindUnverifiedUser implements domain.UnverifiedUserRepository.
func (u *unverifiedUserRepo) FindUnverifiedUser(ctx context.Context, email string) (domain.UnverifiedUser, error) {
	collection := u.database.Collection(u.collection)
	var uv domain.UnverifiedUser
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&uv)
	return uv, err
}

// StoreUnverifiedUser implements domain.UnverifiedUserRepository.
func (u *unverifiedUserRepo) StoreUnverifiedUser(ctx context.Context, uv domain.UnverifiedUser) error {
	collection := u.database.Collection(u.collection)
	_, err := collection.InsertOne(ctx, uv)
	return err
}

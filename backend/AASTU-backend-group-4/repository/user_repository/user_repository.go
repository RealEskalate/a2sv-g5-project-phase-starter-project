package user_repository

import (
	"blog-api/domain/user"
	"blog-api/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)


type UserRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (user.User, error) {
	var u user.User
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"email": email}
	err := collection.FindOne(ctx, filter).Decode(&u)
	return u, err
}

func (ur *UserRepository) GetByUsername(ctx context.Context, username string) (user.User, error) {
	var u user.User
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"username": username}
	err := collection.FindOne(ctx, filter).Decode(&u)
	return u, err
}

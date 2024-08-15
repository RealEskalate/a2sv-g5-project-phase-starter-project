package mongodb

import (
	"blogApp/internal/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMongo struct {
	Collection *mongo.Collection
}



func (r *UserRepositoryMongo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return user, err
}


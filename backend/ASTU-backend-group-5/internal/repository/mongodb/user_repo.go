package mongodb

import (
	"blogApp/internal/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMongo struct {
	Collection *mongo.Collection
}

func (r *UserRepositoryMongo) CreateUser(ctx context.Context, user *domain.User) error {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepositoryMongo) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, nil //expected if no user is found
	}
	if err != nil { //something went wrong
		return nil, err
	}
	return user, err //user exists
}

func (r *UserRepositoryMongo) UpdatePassword(ctx context.Context, email string, password string) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"password": password}})
	return err
}

func (r *UserRepositoryMongo) UpdateUser(ctx context.Context, user *domain.User) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	return err
}

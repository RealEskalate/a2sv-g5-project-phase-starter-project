package repositories

import (
	"context"
	"errors"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserRepositoryMongo(db *mongo.Database) interfaces.UserRepositoryInterface {
	return &UserRepositoryMongo{
		collection: db.Collection("Users"),
	}
}

func (r *UserRepositoryMongo) CreateUser(user *domain.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *UserRepositoryMongo) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *UserRepositoryMongo) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *UserRepositoryMongo) UpdateUser(user *domain.User) error {
	user.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": user.ID},
		bson.M{
			"$set": user,
		},
		options.Update().SetUpsert(false),
	)
	return err
}

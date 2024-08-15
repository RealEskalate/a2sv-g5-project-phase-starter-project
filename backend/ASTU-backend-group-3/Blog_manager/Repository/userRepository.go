package repository

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Save(user *Domain.User) error
	FindByEmail(email string) (*Domain.User, error)
	FindByUsername(username string) (*Domain.User, error)
	Update(username string, UpdatedUser *Domain.User) error
	Delete(userID string) error
	IsDbEmpty() (bool, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepository{collection: collection}
}

func (r *userRepository) Save(user *Domain.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *userRepository) FindByEmail(email string) (*Domain.User, error) {
	var user Domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (*Domain.User, error) {
	var user Domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(username string, updatedUser *Domain.User) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": updatedUser}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *userRepository) Delete(username string) error {
	filter := bson.M{"username": username}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	return err
}

func (r *userRepository) IsDbEmpty() (bool, error) {
	count, err := r.collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

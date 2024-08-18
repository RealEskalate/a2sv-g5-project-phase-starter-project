package repository

import (
	"blog/database"
	"blog/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepository struct {
	database   database.Database
	collection string
}

// CreateUser implements domain.UserRepository.
func (u *userRepository) CreateUser(c context.Context, user *domain.User) error {
	collection := u.database.Collection(u.collection)
	_, err := collection.InsertOne(c, user)
	return err
}

// DeleteUser implements domain.UserRepository.
func (u *userRepository) DeleteUser(c context.Context, id primitive.ObjectID) error {
	panic("unimplemented")
}

// GetUserByEmail implements domain.UserRepository.
func (u *userRepository) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	collection := u.database.Collection(u.collection)
	filter := bson.M{"email": email}
	user := &domain.User{}
	err := collection.FindOne(c, filter).Decode(user)
	return user, err
}

// GetUserByID implements domain.UserRepository.
func (u *userRepository) GetUserByID(c context.Context, id primitive.ObjectID) (*domain.User, error) {
	panic("unimplemented")
}

// GetUserByUsername implements domain.UserRepository.
func (u *userRepository) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	collection := u.database.Collection(u.collection)
	filter := bson.M{"username": username}
	user := &domain.User{}
	err := collection.FindOne(c, filter).Decode(user)
	return user, err
}

// UpdateUser implements domain.UserRepository.
func (u *userRepository) UpdateUser(c context.Context, user *domain.User) error {
	panic("unimplemented")
}

func NewUserRepository(db database.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database   *mongo.Database
	collection string
}

// CreateUser implements interfaces.UserRepository.
func (ur *userRepository) CreateUser(user *entities.User) (*entities.User, error) {
	collection := ur.database.Collection(ur.collection)
	ctx := context.Background()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser implements interfaces.UserRepository.
func (ur *userRepository) DeleteUser(userId string) error {
	collection := ur.database.Collection(ur.collection)
	ctx := context.Background()

	_, err := collection.DeleteOne(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}

// DemoteUserToRegular implements interfaces.UserRepository.
func (ur *userRepository) DemoteUserToRegular(userId string) error {
	panic("unimplemented")
}

// FindUserByEmail implements interfaces.UserRepository.
func (ur *userRepository) FindUserByEmail(email string) (*entities.User, error) {
	collection := ur.database.Collection(ur.collection)

	// Create a filter to find the user by email
	filter := bson.M{"email": email}

	var user entities.User

	// Find the user in the collection
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // No user found with this email
		}
		return nil, err // Some other error occurred
	}

	return &user, nil
}

// FindUserById implements interfaces.UserRepository.
func (ur *userRepository) FindUserById(userId string) (*entities.User, error) {
	panic("unimplemented")
}

// PromoteUserToAdmin implements interfaces.UserRepository.
func (ur *userRepository) PromoteUserToAdmin(userId string) error {
	panic("unimplemented")
}

// UpdateUser implements interfaces.UserRepository.
func (ur *userRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	collection := ur.database.Collection(ur.collection)

	// Create a filter to find the user by ID
	filter := bson.M{"_id": user.ID}

	// Create an update to replace the user with the new user
	update := bson.M{"$set": user}

	// Update the user in the collection
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserRepository(db *mongo.Database, collection string) interfaces.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

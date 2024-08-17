package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type userRepository struct{
	collection *mongo.Collection
	
}

func NewUserRepository(collection *mongo.Collection) interfaces.UserRepository {
	return &userRepository{
		collection: collection,
	}
}

func (repo *userRepository) CreateUser(user *entities.User) (*entities.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepository) FindUserByEmail(email string) (*entities.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entities.User
	err := repo.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) DeleteUser(userId string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.DeleteOne(ctx, bson.M{"_id": userId})
	if err != nil {
		return err
	}

	return nil
}




func (repo *userRepository) FindUserById(userId string) (*entities.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entities.User
	err := repo.collection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (repo *userRepository) UpdateUser(user *entities.User) (*entities.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}

	return user, nil

}


package repository

import (
	"context"
	"errors"
	"group3-blogApi/domain"
	"group3-blogApi/infrastracture"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepositoryImpl(coll *mongo.Collection) domain.UserRepository {
	return &UserRepositoryImpl{collection: coll}
}

func (ur *UserRepositoryImpl) Login(user domain.User) (domain.User, error) {
	var newUser domain.User
	err := ur.collection.FindOne(context.Background(), map[string]string{"username": user.Username}).Decode(&newUser)
	if err != nil {
		return domain.User{}, err
	}
	if infrastracture.CheckPasswordHash(user.Password, newUser.Password) {
		return newUser, nil
	}
	return domain.User{}, errors.New("invalid credentials")

}

func (ur *UserRepositoryImpl) Register(user domain.User) error {

	// isUserExist := ur.collection.FindOne(context.Background(), map[string]string{"username": user.Username}).Err()
	// if isUserExist == nil {
	// 	return errors.New("user already exists")
	// }
	_, err := ur.collection.InsertOne(context.Background(), user)
	return err
}

func (ur *UserRepositoryImpl) GetUserByUsernameOrEmail(username, email string) (domain.User, error) {
	var user domain.User
	err := ur.collection.FindOne(context.Background(),  bson.M{"username": username, "email": email}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}


func (ur *UserRepositoryImpl) AccountActivation(token string, email string) error {
	var user domain.User
	err := ur.collection.FindOne(context.Background(), map[string]string{"activation_token": token}).Decode(&user)
	if err != nil {
		return errors.New("invalid token or user not found")
	}

	if time.Since(user.TokenCreatedAt) > 24*time.Hour {
		return errors.New("token has expired")
	}

	

	// err = ur.collection.FindOneAndUpdate(context.Background(), bson.M{"_id": newID}, bson.M{"$set": user}).Decode(&updatedUser)
	_, err = ur.collection.UpdateOne(context.Background(), bson.M{"email": email}, bson.M{"$set": bson.M{"is_active": true}, "$unset": bson.M{"activation_token": ""}, "$currentDate": bson.M{"updated_at": true}})
	if err != nil {
		return errors.New("failed to activate account")
	}

	return nil
	
}






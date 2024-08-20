package repository

import (
	"context"
	// "fmt"
	"meleket/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection domain.Collection
}

func NewUserRepository(col domain.Collection) *UserRepository {
	return &UserRepository{
		collection: col,
	}
}

func (r *UserRepository) Create(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user domain.User
	// fmt.Printf("Searching for user with username: %s\n", username)
	err := r.collection.FindOne(ctx, bson.M{"name": username}).Decode(&user)
	// fmt.Printf("User found: %+v\n", user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email *string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"email": *email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]*domain.User, error) {
	var users []*domain.User
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(),&users)

	return users, err
}


func (r *UserRepository) UpdateProfile(id primitive.ObjectID, profile *domain.Profile) (*domain.Profile, error) {
	filter := bson.M{"_id": id}
    update := bson.M{"$set": profile}
    
    result := r.collection.FindOneAndUpdate(context.Background(), filter, update)
    if result.Err() != nil {
        return nil, result.Err()
    }

    var decoded domain.Profile
    if err := result.Decode(&decoded); err != nil {
        return nil, err
    }

	profile.ID = decoded.ID
	profile.UserID = decoded.UserID
	

    return profile, nil
}


func (r *UserRepository) DeleteUser(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}


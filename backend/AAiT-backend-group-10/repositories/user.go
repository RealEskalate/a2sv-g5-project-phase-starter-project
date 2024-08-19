package repositories

import (
	"context"
	"errors"
	"time"

	"aait.backend.g10/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collectionName string) *UserRepository {
	collection := db.Collection(collectionName)
	return &UserRepository{collection}
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var user domain.User
	filter := bson.D{{Key: "_id", Value: id}}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	return nil
}

func (r *UserRepository) DeleteUser(id uuid.UUID) error {
	return nil
}

func (r *UserRepository) PromoteUser(id uuid.UUID, makeAdmin bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isAdmin", Value: makeAdmin}}}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("username not found")
	} 

	return nil
}

func (r *UserRepository) GetAllUsersWithName(name string) ([]uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{{Key: "fullname", Value: bson.D{{Key: "$regex", Value: name}, {Key: "$options", Value: "i"}}}}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var userIDs []uuid.UUID
	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, user.ID)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return userIDs, nil
}

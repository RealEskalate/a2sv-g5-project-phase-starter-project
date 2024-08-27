package repositories

import (
	"context"
	"errors"
	"fmt"
	"loan-management/internal/domain"
	"time"

	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrFailedToInsert   = errors.New("failed to insert user")
	ErrFailedToUpdate   = errors.New("failed to update user")
	ErrFailedToDelete   = errors.New("failed to delete user")
	ErrFailedToRetrieve = errors.New("failed to retrieve users")
)

type userRepository struct {
	collection mongoifc.Collection
}

func NewUserRepository(db mongoifc.Database) domain.UserRepository {
	c := db.Collection(domain.UserCollection)
	c.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	})
	return &userRepository{collection: c}
}

func (r *userRepository) Create(user domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectIDFromTimestamp(time.Now()).Hex()
	_, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return domain.User{}, fmt.Errorf("user exists with the give email")
		}
		return domain.User{}, ErrFailedToInsert
	}
	return user, nil
}

func (r *userRepository) Update(id string, updateData domain.User) (domain.User, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{}}
	if updateData.Email != "" {
		return domain.User{}, fmt.Errorf("updating email is not allowed")
	}
	if updateData.Name != "" {
		update["$set"].(bson.M)["name"] = updateData.Name
	}
	if updateData.Password != "" {
		update["$set"].(bson.M)["password"] = updateData.Password
	}
	update["$set"].(bson.M)["is_active"] = updateData.IsActive
	update["$set"].(bson.M)["is_admin"] = updateData.IsAdmin
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.User{}, ErrFailedToUpdate
	}
	updatedUser, err := r.GetByID(id)
	if err != nil {
		return domain.User{}, ErrFailedToRetrieve
	}
	return updatedUser, nil
}

func (r *userRepository) Delete(id string) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return ErrFailedToDelete
	}
	return nil
}

func (r *userRepository) Get() ([]domain.User, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []domain.User
	for cursor.Next(context.TODO()) {
		var user domain.User
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetByID(id string) (domain.User, error) {
	var user domain.User
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	filter := bson.M{"email": email}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

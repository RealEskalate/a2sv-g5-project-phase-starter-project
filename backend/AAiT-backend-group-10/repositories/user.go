package repositories

import (
	"context"
	"errors"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
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

func (r *UserRepository) UpdateUser(user *dto.UserUpdate) error {
	update := bson.D{}
	if user.FullName != "" {
		update = append(update, bson.E{Key: "fullname", Value: user.FullName})
	}
	if user.Bio != "" {
		update = append(update, bson.E{Key: "bio", Value: user.Bio})
	}
	if user.ImageURL != "" {
		update = append(update, bson.E{Key: "imageUrl", Value: user.ImageURL})
	}
	if user.Password != "" {
		update = append(update, bson.E{Key: "password", Value: user.Password})
	}
	update = append(update, bson.E{Key: "updated_at", Value: user.UpdatedAt})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: user.ID}}
	_, err := r.collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) PromoteUser(id uuid.UUID, isPromote bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isAdmin", Value: isPromote}}}}
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


import (
	"context"
	"errors"
	"time"

	"aait.backend.g10/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

type UserRepositoryInterface interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
	UpdateUser(user *domain.User) error
}

func NewUserRepositoryMongo(db *mongo.Database) UserRepositoryInterface {
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

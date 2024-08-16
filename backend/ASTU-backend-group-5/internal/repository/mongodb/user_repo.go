package mongodb

import (
	"blogApp/internal/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewUserRepositoryMongo(collection *mongo.Collection) *UserRepositoryMongo {
	return &UserRepositoryMongo{
		Collection: collection,
	}
}

func (r *UserRepositoryMongo) CreateUser(ctx context.Context, user *domain.User) error {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepositoryMongo) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err //user exists
}

func (r *UserRepositoryMongo) FindUserById(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *UserRepositoryMongo) FindUserByUserName(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	err := r.Collection.FindOne(ctx, bson.M{"username": username}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err //user exists
}

func (r *UserRepositoryMongo) UpdateUser(ctx context.Context, user *domain.User) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	return err
}

func (r *UserRepositoryMongo) DeleteUser(ctx context.Context, id string) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *UserRepositoryMongo) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepositoryMongo) FilterUsers(ctx context.Context, filter map[string]interface{}) ([]*domain.User, error) {

	var users []*domain.User

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil

}

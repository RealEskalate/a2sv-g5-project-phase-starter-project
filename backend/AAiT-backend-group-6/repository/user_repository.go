package repository

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mongo"
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) CreateUser(c context.Context, user *domain.User) error {
	userCollection := ur.database.Collection(ur.collection)

	_, insertionErr := userCollection.InsertOne(c, user)
	return insertionErr
}

func (ur *userRepository) DeleteUser(c context.Context, id string) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.DeleteOne(c, bson.M{"_id": id})
	return err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByUsername(c context.Context, username string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"username": username}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetUsers(c context.Context) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []domain.User

	err = cursor.All(c, &users)
	if users == nil {
		return []domain.User{}, err
	}

	return users, err
}

func (ur *userRepository) UpdateUser(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)

	update := bson.M{}
    userValue := reflect.ValueOf(user).Elem()
    userType := reflect.TypeOf(user).Elem()

    for i := 0; i < userValue.NumField(); i++ {
        field := userValue.Field(i)
        fieldName := userType.Field(i).Tag.Get("bson")

        if !field.IsZero() && fieldName != "_id" { // Exclude zero values and ID field
            update[fieldName] = field.Interface()
        }
    }
	filter := bson.D{{Key: "_id", Value: user.ID}}

	_, result := collection.UpdateOne(c, filter, update)

	if result != nil {
		return result
	}
	
	return nil
}

func (ur *userRepository) LoginUser(c context.Context, user *domain.User) (domain.User, error) {
	panic("unimplemented")
}



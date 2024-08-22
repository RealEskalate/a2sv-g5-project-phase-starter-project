package repository

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mongo"
	"context"
	"reflect"
	"time"

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

func (ur *userRepository) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (ur *userRepository) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"username": username}).Decode(&user)
	return &user, err
}

func (ur *userRepository) GetUserByID(c context.Context, id string) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return &user, err
}

func (ur *userRepository) GetUsers(c context.Context) ([]*domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []*domain.User

	err = cursor.All(c, &users)
	if users == nil {
		return []*domain.User{}, err
	}

	return users, err
}

func (ur *userRepository) UpdateUser(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	setElements := bson.M{}
    userValue := reflect.ValueOf(user).Elem()
    userType := userValue.Type()

    for i := 0; i < userType.NumField(); i++ {
        field := userType.Field(i)
        jsonTag := field.Tag.Get("json")

        // Get the actual value of the field
        fieldValue := userValue.Field(i)

		if jsonTag == "token" || jsonTag == "refresh_token" {
			setElements[jsonTag] = fieldValue.Interface()
			continue
		}

		if jsonTag == "_id" || jsonTag == "" {
            continue
        }
        // Check for zero value and skip if so
        if !reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
            setElements[jsonTag] = fieldValue.Interface()
        }
    }

	setElements["updated_at"] = time.Now()

	update := bson.D{{Key: "$set", Value: setElements}}
	filter := bson.D{{Key: "_id", Value: user.ID}}

	_, result := collection.UpdateOne(c, filter, update)

	if result != nil {
		return result
	}
	
	return nil
}

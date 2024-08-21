package repository

import (
	"Blog_Starter/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	database   *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}

// UpdateSignup implements domain.UserRepository.
func (u *UserRepository) UpdateSignup(c context.Context, user *domain.User) error {
	collection := u.database.Collection(u.collection)

	// Define an update operation
	update := bson.M{
		"$set": user,
	}

	// Define options for the update operation (e.g., to perform an upsert)
	options := options.Update().SetUpsert(true)

	// Perform the update operation
	_, err := collection.UpdateOne(c, bson.M{"_id": user.UserID}, update, options)

	return err
}

// CreateUser implements domain.UserRepository.
func (u *UserRepository) CreateUser(c context.Context, user *domain.User) (*domain.User, error) {
	collection := u.database.Collection(u.collection)

	users, _ := collection.CountDocuments(context.Background(), bson.M{})
	if users == 0 {
		user.Role = "superAdmin" // Automatically make the first user a superAadmin
	} else {
		user.Role = "user"
	}

	// Perform the insert operation
	_, err := collection.InsertOne(c, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser implements domain.UserRepository.
func (u *UserRepository) DeleteUser(c context.Context, userID string) error {
	collection := u.database.Collection(u.collection)

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})
	return err
}

// GetAllUser implements domain.UserRepository.
func (u *UserRepository) GetAllUser(c context.Context) ([]*domain.User, error) {
	collection := u.database.Collection(u.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.M{}, opts)

	if err != nil {
		return nil, err
	}

	var users []*domain.User
	if err = cursor.All(c, &users); err != nil {
		return nil, err
	}

	return users, nil

}

// GetUserByEmail implements domain.UserRepository.
func (u *UserRepository) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	collection := u.database.Collection(u.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return &user, err

}

// GetUserByID implements domain.UserRepository.
func (u *UserRepository) GetUserByID(c context.Context, userID string) (*domain.User, error) {
	collection := u.database.Collection(u.collection)
	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return &user, err
}

// UpdatePassword implements domain.UserRepository.
func (u *UserRepository) UpdatePassword(c context.Context, password string, userID string) (*domain.User, error) {
	panic("unimplemented")
}

// UpdateProfile implements domain.UserRepository.
func (u *UserRepository) UpdateProfile(c context.Context, user *domain.UserUpdate, userID string) (*domain.User, error) {
	panic("unimplemented")
}

// UpdateToken implements domain.UserRepository.
func (u *UserRepository) UpdateToken(c context.Context, accessToken string, refreshToken string, userID string) (*domain.User, error) {
	panic("unimplemented")
}


// UpdateRole implements domain.UserRepository.
func (u *UserRepository) UpdateRole(c context.Context, role string, userID string) (*domain.User, error) {
	collection := u.database.Collection(u.collection)

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": idHex}
	update := bson.M{"$set": bson.M{"role": role}}

	_, err = collection.UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}

	return u.GetUserByID(c, userID)
}



package repository

import (
	"context"
	"errors"
	"time"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrUserNotFound = errors.New("user not found")

// mongodb implementation of UserRepository interface
type UserMongoRepository struct {
	Collection *mongo.Collection
}

// NewUserRepository creates a new UserMongoRepository
func NewUserRepository(db *mongo.Database) *UserMongoRepository {
	return &UserMongoRepository{
		Collection: db.Collection("user-Collection"),
	}
}

func (ur *UserMongoRepository) CreateUser(ctx context.Context, user *models.User) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := ur.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmailOrUsername fetches a user based on the username or email.
func (ur *UserMongoRepository) GetUserByEmailOrUsername(ctx context.Context, username, email string) (*User, error) {
	var user models.User
	filter := bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}

	err := ur.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByID fetches a user by their ID.
func (ur *UserMongoRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = ur.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates a user's information.
func (ur *UserMongoRepository) UpdateUser(ctx context.Context, user *models.User) error {
	filter := bson.M{"_id": user.ID}

	update := bson.M{}
	if user.Username != "" {
		update["username"] = user.Username
	}
	if user.Name != "" {
		update["name"] = user.Name
	}
	if user.Email != "" {
		update["email"] = user.Email
	}

	if len(update) == 0 {
		return nil
	}

	updateDocument := bson.M{
		"$set": update,
	}

	_, err := ur.Collection.UpdateOne(ctx, filter, updateDocument)
	return err
}

// DeleteUser deletes a user by their ID.
func (ur *UserMongoRepository) DeleteUser(ctx context.Context, userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = ur.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// PromoteUser promotes a user to a higher role.
func (ur *UserMongoRepository) PromoteUser(ctx context.Context, userID string) error {
	return ur.updateUserRole(ctx, userID, "Admin") // Example role
}

// DemoteUser demotes a user to a lower role.
func (ur *UserMongoRepository) DemoteUser(ctx context.Context, userID string) error {
	return ur.updateUserRole(ctx, userID, "User") // Example role
}

// updateUserRole is a helper method to update a user's role.
func (ur *UserMongoRepository) updateUserRole(ctx context.Context, userID, role string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{"role": role},
	}

	_, err = ur.Collection.UpdateOne(ctx, filter, update)
	return err
}

// StoreAccessToken stores an access token for the user.
func (ur *UserMongoRepository) StoreAccessToken(ctx context.Context, userID, token string) error {
	return ur.storeToken(ctx, userID, "access_token", token)
}

// StoreRefreshToken stores a refresh token for the user.
func (ur *UserMongoRepository) StoreRefreshToken(ctx context.Context, userID, token string) error {
	return ur.storeToken(ctx, userID, "refresh_token", token)
}

// storeToken is a helper method to store a token for a user.
func (ur *UserMongoRepository) storeToken(ctx context.Context, userID, tokenType, token string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{tokenType: token},
	}

	_, err = ur.Collection.UpdateOne(ctx, filter, update)
	return err
}

// DeleteTokensFromDB deletes both access and refresh tokens for a user.
func (ur *UserMongoRepository) DeleteTokensFromDB(ctx context.Context, userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$unset": bson.M{
			"access_token":  1,
			"refresh_token": 1,
		},
	}

	_, err = ur.Collection.UpdateOne(ctx, filter, update)
	return err
}

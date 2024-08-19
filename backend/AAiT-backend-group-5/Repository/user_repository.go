package repository

import (
	"context"
	"errors"
	"time"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
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
func NewUserRepository(db *mongo.Database) interfaces.UserRepository {
	return &UserMongoRepository{
		Collection: db.Collection("user-collection"),
	}
}

func (ur *UserMongoRepository) CreateUser(ctx context.Context, user *models.User) *models.ErrorResponse {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := ur.Collection.InsertOne(ctx, user)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	return models.Nil()
}

// GetUserByEmailOrUsername fetches a user based on the username or email.
func (ur *UserMongoRepository) GetUserByEmailOrUsername(ctx context.Context, username, email string) (*models.User, *models.ErrorResponse) {
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
			return nil, models.NotFound(err.Error())
		}
		return nil, models.NotFound(err.Error())
	}

	return &user, models.Nil()
}

// GetUserByID fetches a user by their ID.
func (ur *UserMongoRepository) GetUserByID(ctx context.Context, id string) (*models.User, *models.ErrorResponse) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	var user models.User
	err = ur.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, models.NotFound(err.Error())
	}

	return &user, models.Nil()
}

// UpdateUser updates a user's information.
func (ur *UserMongoRepository) UpdateUser(ctx context.Context, user *models.User, id string) *models.ErrorResponse {
	objID, Err := primitive.ObjectIDFromHex(id)
	if Err != nil {
		return models.InternalServerError(Err.Error())
	}
	filter := bson.M{"_id": objID}

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
		return models.Nil()
	}

	updateDocument := bson.M{
		"$set": update,
	}

	_, err := ur.Collection.UpdateOne(ctx, filter, updateDocument)
	if err != nil {
		return models.InternalServerError(err.Error())
	}
	return models.Nil()
}

// DeleteUser deletes a user by their ID.
func (ur *UserMongoRepository) DeleteUser(ctx context.Context, userID string) *models.ErrorResponse {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	_, err = ur.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	return models.Nil()
}

// updateUserRole is a helper method to update a user's role.
func (ur *UserMongoRepository) updateUserRole(ctx context.Context, userID, role string) *models.ErrorResponse {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{"role": role},
	}

	_, err = ur.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	return models.Nil()
}

// PromoteUser promotes a user to an admin role.
func (ur *UserMongoRepository) PromoteUser(ctx context.Context, userID string) *models.ErrorResponse {
	err := ur.updateUserRole(ctx, userID, "admin")
	return err
}

// DemoteUser demotes a user to a lower role.
func (ur *UserMongoRepository) DemoteUser(ctx context.Context, userID string) *models.ErrorResponse {
	err := ur.updateUserRole(ctx, userID, "admin")
	return err
}

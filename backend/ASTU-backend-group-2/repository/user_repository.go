package repository

import (
	"context"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) GetUser(c context.Context, userId string) (*domain.User, error) {
	return nil, nil
}

func (ur *userRepository) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (ur *userRepository) GetByID(c context.Context, userID string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}

func (ur *userRepository) RevokeRefreshToken(c context.Context, refreshToken string) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.UpdateOne(c, bson.M{"tokens": refreshToken}, bson.M{"$pull": bson.M{"tokens": refreshToken}})
	return err
}

func (ur *userRepository) UpdateUser(c context.Context, userID string, updatedUser *domain.User) error {
	return nil
}
func (ur *userRepository) DeleteUser(c context.Context, userID string) error {
	return nil
}
func (ur *userRepository) IsUserActive(c context.Context, userID string) (bool, error) {
	return false, nil
}
func (ur *userRepository) ResetUserPassword(c context.Context, userID string, resetPassword *domain.ResetPassword) error {
	return nil
}
func (ur *userRepository) UpdateUserPassword(c context.Context, userID string, updatePassword *domain.UpdatePassword) error {
	return nil
}
func (ur *userRepository) PromoteUserToAdmin(c context.Context, userID string) error {
	return nil
}
func (ur *userRepository) DemoteAdminToUser(c context.Context, userID string) error {
	return nil
}

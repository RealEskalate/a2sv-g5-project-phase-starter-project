package Repository

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Save(user *Domain.User) error
	FindByEmail(email string) (*Domain.User, error)
	FindByUsername(username string) (Domain.User, error)
	Update(username string, UpdatedUser bson.M) error
	Delete(userID string) error
	IsDbEmpty() (bool, error)
	InsertToken(username string, accessToke string, refreshToken string) error
	ExpireToken(token string) error
}

type userRepository struct {
	collection      *mongo.Collection
	tokenCollection *mongo.Collection
}

func NewUserRepository(collection, tokenCollection *mongo.Collection) UserRepository {
	return &userRepository{collection: collection, tokenCollection: tokenCollection}
}

func (r *userRepository) Save(user *Domain.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *userRepository) FindByEmail(email string) (*Domain.User, error) {
	var user Domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (Domain.User, error) {
	var user Domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return Domain.User{}, err
	}
	return user, nil
}

func (r *userRepository) Update(username string, updateFields bson.M) error {
	filter := bson.M{"username": username}

	// Only perform the update if there are fields to update
	if len(updateFields) == 0 {
		return nil // No update needed
	}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.M{"$set": updateFields})
	return err
}

func (r *userRepository) Delete(username string) error {
	filter := bson.M{"username": username}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	return err
}

func (r *userRepository) IsDbEmpty() (bool, error) {
	count, err := r.collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *userRepository) InsertToken(username string, accessToke string, refreshToken string) error {
	token := &Domain.Token{
		TokenID:      primitive.NewObjectID(),
		Username:     username,
		AccessToken:  accessToke,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(2 * time.Hour),
	}
	_, err := r.tokenCollection.InsertOne(context.TODO(), token)

	if err != nil {
		return err
	}

	return nil
}

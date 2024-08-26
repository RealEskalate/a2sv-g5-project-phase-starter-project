package repositories

import (
	domain "blogs/Domain"
	"blogs/mongo"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type unverifiedUserRepo struct {
	database   mongo.Database
	collection string
}

func NewUnverifiedUserRepository(db mongo.Database, collection string) domain.UnverifiedUserRepository {
	return &unverifiedUserRepo{
		database:   db,
		collection: collection,
	}
}

// DeleteUnverifiedUser implements domain.UnverifiedUserRepository.
func (u *unverifiedUserRepo) DeleteUnverifiedUser(ctx context.Context, email string) error {
	collection := u.database.Collection(u.collection)
	_, err := collection.DeleteOne(ctx, bson.M{"email": email})
	return err
}

// FindUnverifiedUser implements domain.UnverifiedUserRepository.
func (u *unverifiedUserRepo) FindUnverifiedUser(ctx context.Context, email string) (domain.UnverifiedUser, error) {
	collection := u.database.Collection(u.collection)
	var uv domain.UnverifiedUser
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&uv)
	return uv, err
}

// StoreUnverifiedUser implements domain.UnverifiedUserRepository.
func (u *unverifiedUserRepo) StoreUnverifiedUser(ctx context.Context, uv domain.UnverifiedUser) error {
	collection := u.database.Collection(u.collection)
	_, err := collection.InsertOne(ctx, uv)
	return err
}

//paste in repo

func (s *unverifiedUserRepo) UpdateOTP(ctx context.Context, email string, otp string, expiry time.Time) (domain.UnverifiedUserResponse, error) {
	collection := s.database.Collection(s.collection)
	var unverifiedUser domain.UnverifiedUserResponse

	// Filter to find the user by email
	filter := bson.M{"email": email}

	// Update to set the reset_token
	update := bson.M{"$set": bson.M{"otp": otp, "expiresat": expiry}}

	// Execute the update3
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return unverifiedUser, err
	}

	// Return the updated user
	return unverifiedUser, nil

}


func (r *unverifiedUserRepo) DeleteUnverifiedUsersBefore(ctx context.Context, cutoffDate time.Time) error {
    collection := r.database.Collection(r.collection)
    filter := bson.M{
        "created_at": bson.M{"$lt": cutoffDate},
    }
    result, err := collection.DeleteMany(ctx, filter)
	fmt.Println("WARNING: Deleting unverified users before", cutoffDate)
    if err != nil {
        return err
    }
    if result.DeletedCount == 0 {
        return errors.New("no ones Data expired to delete")
    }
    return nil
}
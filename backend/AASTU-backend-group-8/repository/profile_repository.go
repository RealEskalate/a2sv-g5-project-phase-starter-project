package repository

import (
	"context"
	"errors"
	"fmt"
	"meleket/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileRepository struct {
	collection domain.Collection
}

func NewProfileRepository(col domain.Collection) domain.ProfileRepository {
	return &ProfileRepository{
		collection: col,
	}
}

// SaveRefreshToken saves the refresh token in the database
func (r *ProfileRepository) SaveProfile(profile *domain.Profile) error {
	fmt.Println(profile)
	_, err := r.collection.InsertOne(context.TODO(), profile)
	return err
}

func (r *ProfileRepository) FindProfile(userID string) (*domain.Profile, error) {
	var profile domain.Profile
	err := r.collection.FindOne(context.TODO(), bson.M{"userid": userID}).Decode(&profile)
	return &profile, err
}

func (r *ProfileRepository) DeleteProfile(userID string) error {
	objid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"userid": objid})
	return err
}

func (r *ProfileRepository) UpdateProfile(profile *domain.Profile) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(profile)

	// Create a filter to find the profile by its unique identifier (e.g., ID)
	filter := bson.M{"userid": profile.UserID}

	// Create an update document specifying the fields to update
	update := bson.M{
		"$set": profile, // Update all fields in the profile
	}

	// Perform the update operation
	result, err := r.collection.UpdateOne(ctx, filter, update)
	fmt.Println(result)
	if err != nil {
		return err
	}

	// Check if any document was matched and updated
	if result.MatchedCount == 0 {
		return errors.New("No entry found with the given credentials")
	}

	return nil
}

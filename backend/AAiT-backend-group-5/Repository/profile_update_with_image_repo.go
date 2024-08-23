package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileRepository interface {
	SaveProfileImageKey(userID string, imageKey string) error
}

type profileRepository struct {
	Collection *mongo.Collection
}

func NewProfileRepository(db *mongo.Database) ProfileRepository {
	return &profileRepository{
		Collection: db.Collection("user-profile-collection"),
	}
}

func (pr *profileRepository) SaveProfileImageKey(userID string, imageKey string) error {
	// Update the user's profile with the new image key
	_, err := pr.Collection.UpdateOne(
		context.TODO(),
		bson.M{"user_id": userID}, // Filter by user ID
		bson.M{"$set": bson.M{"image_key": imageKey}},
	)
	return err
}

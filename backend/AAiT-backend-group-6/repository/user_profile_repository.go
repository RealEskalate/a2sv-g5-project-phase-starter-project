package repository

import (
	"AAiT-backend-group-6/domain" // Ensure this import is correct based on your project structure
	"AAiT-backend-group-6/mongo"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options" // Import options for upsert
)

type UserProfileRepository struct {
	database   mongo.Database
	collection string
}

func NewUserProfileRepository(db mongo.Database) *UserProfileRepository {
	return &UserProfileRepository{
		database:   db,
		collection: "Userprofile",
	}
}

func (ur *UserProfileRepository) Update(c context.Context, id string, updateData domain.UpdateData) error {
	// Convert the string ID to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}

	// Filter to match the document by its ID
	filter := bson.M{"user_id": objectID}

	// Update document fields based on the provided updateData
	update := bson.M{
		"$set": updateData,
	}

	// Define the options for the update operation, enabling upsert
	opts := options.Update().SetUpsert(true)

	// Perform the update-or-insert operation
	collection := ur.database.Collection(ur.collection)
	_, err = collection.UpdateOne(c, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

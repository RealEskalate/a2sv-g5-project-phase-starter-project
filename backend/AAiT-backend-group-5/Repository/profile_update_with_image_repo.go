package repository

import (
	"context"
	"fmt"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		Collection: db.Collection("user-collection"),
	}
}

func (pr *profileRepository) SaveProfileImageKey(userID string, imageKey string) error {
	objID, Err := primitive.ObjectIDFromHex(userID)

	fmt.Println(userID)
	fmt.Println(imageKey)
	if Err != nil {
		fmt.Println("You are here 1")
		return models.InternalServerError(Err.Error())
	}

	_, err := pr.Collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID}, // Filter by user ID
		bson.M{"$set": bson.M{"image_key": imageKey}},
	)
	return err
}

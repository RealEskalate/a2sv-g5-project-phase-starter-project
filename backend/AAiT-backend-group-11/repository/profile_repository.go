package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"backend-starter-project/mongo"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type profileRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
	context    context.Context
}

func NewProfileRepository(ctx context.Context, db *mongo.Database, collection *mongo.Collection) interfaces.ProfileRepository {
	return profileRepository{db: db, collection: collection, context: ctx}
}
func (repo profileRepository) GetUserProfile(user_id string) (*entities.Profile, error) {
	userID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"userId", userID}}
	user := (*repo.collection).FindOne(context.TODO(), filter)

	if err = mongo.ErrNoDocuments; err != nil {
		return &entities.Profile{}, errors.New("couldn't find the profile")
	}

	var profile entities.Profile

	err = user.Decode(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil

}

func (repo profileRepository) CreateUserProfile(profile *entities.Profile) (*entities.Profile, error) {
	if profile.UserID == primitive.NilObjectID {
		return nil, errors.New("user id is required")
	}
	if existed := (*repo.collection).FindOne(repo.context, bson.D{{"userId", profile.UserID}}); existed.Err() == nil {
		return nil, errors.New("profile already exists")
	}
	_, err := (*repo.collection).InsertOne(repo.context, profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (repo profileRepository) UpdateUserProfile(profile *entities.Profile) (*entities.Profile, error) {
	user_id := profile.UserID
	if user_id == primitive.NilObjectID {
		return nil, errors.New("user id is required")
	}
	filter := bson.D{{"userId", user_id}}

	data := bson.D{
		{"$set", bson.D{
			{"bio", profile.Bio},
			{"profilePicture", profile.ProfilePicture},
			{"contactInfo", profile.ContactInfo},
			{"updatedAt", profile.UpdatedAt},
		}},
	}
	_, err := (*repo.collection).UpdateOne(context.TODO(), filter, data)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (repo profileRepository) DeleteUserProfile(user_id string) error {
	userID, err := primitive.ObjectIDFromHex(user_id)
	filter := bson.D{{"userId", userID}}
	_, err = (*repo.collection).DeleteOne(repo.context, filter)
	if err != nil {
		return err
	}

	return nil
}

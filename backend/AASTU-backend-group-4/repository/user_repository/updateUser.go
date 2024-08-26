package user_repository

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser *domain.UpdateRequest) error {
	collection := ur.collection
	filter := bson.M{"_id": userID}

	update := bson.M{
		"$set": bson.M{
			"firstname":           updatedUser.Firstname,
			"lastname":            updatedUser.Lastname,
			"username":            updatedUser.Username,
			"bio":                 updatedUser.Bio,
			"profile_picture":     updatedUser.ProfilePicture,
			"contact_information": updatedUser.ContactInformation,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

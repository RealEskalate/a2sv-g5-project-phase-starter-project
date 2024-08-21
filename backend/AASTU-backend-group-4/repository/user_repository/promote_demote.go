package user_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *UserRepository) PromoteDemote(ctx context.Context, userID primitive.ObjectID, action string) error {
	var isAdmin bool
	if action == "promote" {
		isAdmin = true
	} else {
		isAdmin = false
	}

	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{
			"isAdmin": isAdmin,
		},
	}
	result, err := ur.collection.UpdateOne(ctx, filter, update)

	if result.MatchedCount == 0 {
		return errors.New("no user found with the specified id")
	}
	return err
}

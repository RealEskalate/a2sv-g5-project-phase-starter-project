package repository

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// storeToken is a helper method to store a token for a user.
func (ur *UserMongoRepository) storeToken(ctx context.Context, userID, tokenType, token string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{tokenType: token},
	}

	_, err = ur.Collection.UpdateOne(ctx, filter, update)
	if err != nil{
		return err
	}
	return nil	
}


// StoreAccessToken stores an access token for the user.
func (ur *UserMongoRepository) StoreAccessToken(ctx context.Context, userID, token string) *models.ErrorResponse {
	err := ur.storeToken(ctx, userID, "access_token", token)
	if err != nil{
		return models.InternalServerError(err.Error())
	}
	return models.Nil()
}

// StoreRefreshToken stores a refresh token for the user.
func (ur *UserMongoRepository) StoreRefreshToken(ctx context.Context, userID, token string) *models.ErrorResponse {
	err := ur.storeToken(ctx, userID, "refresh_token", token)
	if err != nil{
		return models.InternalServerError(err.Error())
	}
	return models.Nil()
}


// DeleteTokensFromDB deletes both access and refresh tokens for a user.
func (ur *UserMongoRepository) DeleteTokensFromDB(ctx context.Context, userID string) *models.ErrorResponse {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$unset": bson.M{
			"access_token":  1,
			"refresh_token": 1,
		},
	}

	_, err = ur.Collection.UpdateOne(ctx, filter, update)
	return models.InternalServerError(err.Error())
}

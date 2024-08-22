package repository

import (
	"context"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepo struct {
	Collection interfaces.Collection
}

func NewSessionRepository(db interfaces.Database) interfaces.SessionRepository {
	return &SessionRepo{
		Collection: db.Collection("session-collection"),
	}
}

// SaveToken saves a session token to the database.
func (sr *SessionRepo) SaveToken(ctx context.Context, session *models.Session) *models.ErrorResponse {
	var sessionExists models.Session
	err := sr.Collection.FindOne(ctx, bson.M{"user_id": session.UserID}).Decode(&sessionExists)

	if err == nil {
		filter := bson.M{"user_id": session.UserID}
		update := bson.M{
			"$set": bson.M{
				"refresh_token": session.RefreshToken,
				"access_token":  session.AccessToken,
			},
		}

		_, err = sr.Collection.UpdateOne(ctx, filter, update)

		if err != nil {
			return models.InternalServerError(err.Error())
		}

		return models.Nil()
	}

	_, nErr := sr.Collection.InsertOne(ctx, session)
	if nErr != nil {
		return models.InternalServerError(err.Error())
	}
	return models.Nil()

}

// UpdateToken updates an existing session token in the database.
func (sr *SessionRepo) UpdateToken(ctx context.Context, session *models.Session) *models.ErrorResponse {
	filter := bson.M{"user_id": session.UserID}
	update := bson.M{
		"$set": bson.M{
			"refresh_token": session.RefreshToken,
			"access_token":  session.AccessToken,
		},
	}

	_, err := sr.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return models.InternalServerError(err.Error())
	}
	return models.Nil()
}

// RemoveToken removes a session token from the database based on the user ID.
func (sr *SessionRepo) RemoveToken(ctx context.Context, userID string) *models.ErrorResponse {

	filter := bson.M{"user_id": userID}
	_, err := sr.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return models.InternalServerError(err.Error())
	}
	return models.Nil()
}

// GetToken retrieves a session token from the database based on the user ID.
func (sr *SessionRepo) GetToken(ctx context.Context, userID string) (*models.Session, *models.ErrorResponse) {
	filter := bson.M{"user_id": userID}
	var session models.Session
	err := sr.Collection.FindOne(ctx, filter).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, models.NotFound("Session not found")
		}
		return nil, models.InternalServerError(err.Error())
	}
	return &session, models.Nil()
}

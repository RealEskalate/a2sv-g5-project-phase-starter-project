package repository

import (
	"context"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepo struct {
	Collection *mongo.Collection
}

func NewSessionRepository(db *mongo.Database) interfaces.SessionRepository {
	return &SessionRepo{
		Collection: db.Collection("session-collection"),
	}
}

// SaveToken saves a session token to the database.
func (sr *SessionRepo) SaveToken(ctx context.Context, session *models.Session) *models.ErrorResponse {
	_, err := sr.Collection.InsertOne(ctx, session)
	if err != nil {
		return models.InternalServerError(err.Error())
	}
	return models.Nil()
}

// UpdateToken updates an existing session token in the database.
func (sr *SessionRepo) UpdateToken(ctx context.Context, session *models.Session) *models.ErrorResponse {
	filter := bson.M{"user_id": session.UserID}
	update := bson.M{
		"$set": bson.M{
			"access_token": session.AccessToken,
			"refresh_token": session.RefreshToken,
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
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	filter := bson.M{"user_id": objID}
	_, err = sr.Collection.DeleteOne(ctx, filter)
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

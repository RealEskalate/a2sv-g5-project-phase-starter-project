package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SessionRepository struct {
	collection *mongo.Collection
}

func NewSessionRespository(collection *mongo.Collection) domain.SessionRepository {
	return &SessionRepository{
		collection: collection,
	}
}
func (sessionRepo *SessionRepository) FindTokenById(cxt context.Context, id string) (*domain.Session, domain.Error) {
	sessionID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return &domain.Session{}, &domain.CustomError{Message: fmt.Sprintf("error parsing the session id. %v \n", errIDParse.Error()), Code: http.StatusInternalServerError}
	}
	filter := bson.D{{"_id", sessionID}}
	var fetchedSession domain.Session
	errFetchSession := sessionRepo.collection.FindOne(cxt, filter).Decode(&fetchedSession)
	if errFetchSession != nil {
		if errors.Is(errFetchSession, mongo.ErrNoDocuments) {
			return &domain.Session{}, &domain.CustomError{Message: fmt.Sprintf("user not found. %v \n", errFetchSession.Error()), Code: http.StatusNotFound}
		}
		return &domain.Session{}, &domain.CustomError{Message: fmt.Sprintf("error fetching the user. %v \n", errFetchSession.Error()), Code: http.StatusInternalServerError}
	}
	return &fetchedSession, nil
}

func (sessionRepo *SessionRepository) FindTokenByUserUsername(cxt context.Context, userID string) (*domain.Session, bool, domain.Error) {
	user_id, errIDParse := primitive.ObjectIDFromHex(userID)
	if errIDParse != nil {
		return &domain.Session{}, false, &domain.CustomError{Message: fmt.Sprintf("error parsing the user id. %v \n", errIDParse.Error()), Code: http.StatusInternalServerError}
	}
	filter := bson.D{{"user_id", user_id}}
	var fetchedSession domain.Session
	errFetchSession := sessionRepo.collection.FindOne(cxt, filter).Decode(&fetchedSession)
	if errFetchSession != nil {
		if errors.Is(errFetchSession, mongo.ErrNoDocuments) {
			return &domain.Session{}, false, nil
		}
		return &domain.Session{}, false, &domain.CustomError{Message: fmt.Sprintf("error fetching the session. %v \n", errFetchSession.Error()), Code: http.StatusInternalServerError}
	}
	return &fetchedSession, true, nil
}

func (sessionRepo *SessionRepository) CreateToken(cxt context.Context, session *domain.Session) (*domain.Session, domain.Error) {
	insertResult, errInsert := sessionRepo.collection.InsertOne(cxt, session)
	if errInsert != nil {
		if mongo.IsDuplicateKeyError(errInsert) {
			return &domain.Session{}, &domain.CustomError{Message: fmt.Sprintf("session already exists. %v \n", errInsert.Error()), Code: http.StatusConflict}
		}
		return &domain.Session{}, &domain.CustomError{Message: fmt.Sprintf("error inserting the session. %v \n", errInsert.Error()), Code: http.StatusInternalServerError}
	}
	returnedID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return &domain.Session{}, &domain.CustomError{Message: fmt.Sprintf("error getting the user id. %v \n", errInsert.Error()), Code: http.StatusInternalServerError}
	}
	session.ID = returnedID
	return session, nil
}

func (sessionRepo *SessionRepository) UpdateToken(cxt context.Context, id string, session *domain.Session) domain.Error {
	updateID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return &domain.CustomError{Message: fmt.Sprintf("error parsing the session id. %v \n", errIDParse.Error()), Code: http.StatusInternalServerError}
	}

	filter := bson.D{{"_id", updateID}}
	updateDoc := bson.D{{"$set", session}}
	opts := options.Update().SetUpsert(false)
	updateResult, errUpdate := sessionRepo.collection.UpdateOne(cxt, filter, updateDoc, opts)

	if errUpdate != nil {
		return &domain.CustomError{Message: fmt.Sprintf("error updating the session. %v \n", errUpdate.Error()), Code: http.StatusInternalServerError}
	}
	if updateResult.MatchedCount == 0 {
		return &domain.CustomError{Message: fmt.Sprintf("session not found. %v \n", errUpdate.Error()), Code: http.StatusNotFound}
	}

	return nil
}

func (sessionRepo *SessionRepository) DeleteToken(cxt context.Context, id string) domain.Error {
	deleteID, errIDParse := primitive.ObjectIDFromHex(id)
	if errIDParse != nil {
		return &domain.CustomError{Message: fmt.Sprintf("error parsing the session id. %v \n", errIDParse.Error()), Code: http.StatusInternalServerError}
	}

	filter := bson.D{{"_id", deleteID}}
	deleteResult, errDelete := sessionRepo.collection.DeleteOne(cxt, filter)

	if errDelete != nil {
		return &domain.CustomError{Message: fmt.Sprintf("error updating the session. %v \n", errDelete.Error()), Code: http.StatusInternalServerError}
	}
	if deleteResult.DeletedCount == 0 {
		return &domain.CustomError{Message: fmt.Sprintf("session not found. %v \n", errDelete.Error()), Code: http.StatusNotFound}
	}
	return nil
}

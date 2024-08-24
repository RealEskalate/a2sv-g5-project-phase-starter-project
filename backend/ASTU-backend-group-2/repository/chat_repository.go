package repository

import (
	"context"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type chatRepository struct {
	database   mongo.Database
	collection string
}

func NewChatRepository(db mongo.Database, collection string) entities.ChatRepository {
	return &chatRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *chatRepository) CreateChat(c context.Context, chat entities.Chat) (entities.Chat, error) {
	collection := cr.database.Collection(cr.collection)
	res, err := collection.InsertOne(c, chat)
	if err != nil {
		return entities.Chat{}, err
	}

	chat.ID = res.InsertedID.(primitive.ObjectID)

	return chat, nil
}

func (cr *chatRepository) GetChat(c context.Context, chatID string) (entities.Chat, error) {
	id, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return entities.Chat{}, err
	}

	collection := cr.database.Collection(cr.collection)
	filter := bson.M{"_id": id}
	var dbChat entities.Chat
	if err := collection.FindOne(c, filter).Decode(&dbChat); err == mongo.ErrNoDocuments {
		return entities.Chat{}, err
	} else if err != nil {
		return entities.Chat{}, err
	}

	return dbChat, nil
}

func (cr *chatRepository) GetChats(c context.Context, userID string, limit int64, page int64) ([]entities.Chat, mongopagination.PaginationData, error) {
	collection := cr.database.Collection(cr.collection)

	filter := bson.M{"user_id": userID}

	var chats []entities.Chat

	paginatedData, err := mongopagination.New(collection).Context(c).Limit(limit).Page(page).Filter(filter).Decode(&chats).Find()

	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return chats, paginatedData.Pagination, nil
}

func (cr *chatRepository) CreateMessage(c context.Context, chatID string, body entities.Message) error {

	id, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return err
	}

	collection := cr.database.Collection(cr.collection)
	query := bson.M{
		"$push": bson.M{"history": body},
	}

	_, err = collection.UpdateByID(c, id, query)
	if err == mongo.ErrNoDocuments {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (cr *chatRepository) DeleteChat(c context.Context, chatID string) error {
	id, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return err
	}

	collection := cr.database.Collection(cr.collection)
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(c, filter)
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return err
	}

	return nil
}

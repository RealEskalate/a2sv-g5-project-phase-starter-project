package repository

import (
	"context"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	mongopagination "github.com/gobeam/mongo-go-pagination"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type chatRepository struct {
	database       mongo.Database
	collectionName string
}

func NewChatRepository(db mongo.Database, collection string) entities.ChatRepository {
	return &chatRepository{
		database:       db,
		collectionName: collection,
	}
}

func (cr *chatRepository) CreateChat(c context.Context, chat entities.Chat) (entities.Chat, error) {
	collection := cr.database.Collection(cr.collectionName)
	res, err := collection.InsertOne(c, chat)
	if err != nil {
		return entities.Chat{}, custom_error.ErrCreatingChat
	}

	chat.ID = res.InsertedID.(primitive.ObjectID)

	return chat, nil
}

func (cr *chatRepository) GetChat(c context.Context, chatID primitive.ObjectID) (entities.Chat, error) {

	collection := cr.database.Collection(cr.collectionName)
	filter := bson.M{"_id": chatID}
	var dbChat entities.Chat
	if err := collection.FindOne(c, filter).Decode(&dbChat); err == mongo.ErrNoDocuments {
		return entities.Chat{}, custom_error.ErrChatNotFound
	} else if err != nil {
		return entities.Chat{}, custom_error.ErrChatNotFound
	}

	return dbChat, nil
}

func (cr *chatRepository) GetChats(c context.Context, userID primitive.ObjectID, limit int64, page int64) (*[]entities.Chat, mongopagination.PaginationData, error) {
	collection := cr.database.Collection(cr.collectionName)

	filter := bson.M{"user_id": userID}

	var chats []entities.Chat

	paginatedData, err := mongopagination.New(collection).Context(c).Limit(limit).Page(page).Filter(filter).Decode(&chats).Find()

	if err != nil {
		return nil, mongopagination.PaginationData{}, custom_error.ErrFilteringChats
	}

	return &chats, paginatedData.Pagination, nil
}

func (cr *chatRepository) CreateMessage(c context.Context, chatID primitive.ObjectID, body entities.Message) error {

	collection := cr.database.Collection(cr.collectionName)
	query := bson.M{
		"$push": bson.M{"messages": body},
	}

	_, err := collection.UpdateByID(c, chatID, query)
	if err == mongo.ErrNoDocuments {
		return custom_error.ErrChatNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

func (cr *chatRepository) DeleteChat(c context.Context, chatID primitive.ObjectID) error {

	collection := cr.database.Collection(cr.collectionName)
	filter := bson.M{"_id": chatID}

	deleteResult, err := collection.DeleteOne(c, filter)
	if err != nil || deleteResult.DeletedCount == 0 {
		return custom_error.ErrChatNotFound
	}

	return nil
}

func (cr *chatRepository) UpdateChat(ctx context.Context, chatID primitive.ObjectID, updatedChat entities.Chat) (entities.Chat, error) {
	collection := cr.database.Collection(cr.collectionName)
	filter := bson.M{"_id": chatID}
	update := bson.M{
		"$set": bson.M{
			"user_id":    updatedChat.UserID,
			"title":      updatedChat.Title,
			"messages":   updatedChat.Messages,
			"created_at": updatedChat.CreatedAt,
			"updated_at": updatedChat.UpdatedAt,
		},
	}

	_, err :=
		collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return entities.Chat{}, custom_error.ErrChatNotFound
	}

	return updatedChat, nil
}

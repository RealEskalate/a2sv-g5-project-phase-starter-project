package mongodb

import (
	"context"
	"math"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/chat"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "chats"

type ChatRepository struct {
	Database *mongo.Database
}

func NewChatRepository(database *mongo.Database, collectionName string) *ChatRepository {
	return &ChatRepository{
		Database: database,
	}
}

func (chatRepository *ChatRepository) AddMessage(ctx context.Context, chatID string, message chat.Message) error {
	id, err := primitive.ObjectIDFromHex(chatID)
	if err != nil{
		return chat.ErrInvalidID
	}
	
	collection := chatRepository.Database.Collection(collectionName)
	update := bson.M{
		"$push": bson.M{"history": message},
	}

	_, err = collection.UpdateByID(ctx, id, update)
	if err == mongo.ErrNoDocuments{
		return chat.ErrChatNotFound
	}

	if err != nil{
		return err
	}

	return nil
}

func (chatRepository *ChatRepository) CreateChat(ctx context.Context, newChat chat.Chat) (chat.Chat, error) {
	collection := chatRepository.Database.Collection(collectionName)
	_, err := collection.InsertOne(ctx, newChat)
	if err != nil {
		return chat.Chat{}, err
	}

	return newChat, nil
}

func (chatRepository *ChatRepository) DeleteChat(ctx context.Context, chatID string) error {
	id, err := primitive.ObjectIDFromHex(chatID)
	if err != nil{
		return chat.ErrInvalidID
	} 

	collection := chatRepository.Database.Collection(collectionName)
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return chat.ErrChatNotFound
	}

	return nil
}

func (chatRepository *ChatRepository) GetChat(ctx context.Context, chatID string) (chat.Chat, error) {
	id, err := primitive.ObjectIDFromHex(chatID)
	if err != nil{
		return chat.Chat{}, chat.ErrInvalidID
	} 

	collection := chatRepository.Database.Collection(collectionName)
	filter := bson.M{"_id": id}
	var retrievedChat chat.Chat
	if err := collection.FindOne(ctx, filter).Decode(&retrievedChat); err == mongo.ErrNoDocuments{
		return chat.Chat{}, chat.ErrChatNotFound
	}else if err != nil{
		return chat.Chat{}, err
	}

	return retrievedChat, nil
}

func (chatRepository *ChatRepository) GetChats(ctx context.Context, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[chat.Chat], error) {
	collection := chatRepository.Database.Collection(collectionName)
	findOptions := options.Find()
	findOptions.SetSkip(int64(pagination.Page-1) * int64(pagination.Limit))
	findOptions.SetLimit(int64(pagination.Limit))

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return infrastructure.PaginationResponse[chat.Chat]{}, err
	}

	totalPages := int(math.Ceil(float64(count) / float64(int64(pagination.Limit))))
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return infrastructure.PaginationResponse[chat.Chat]{}, err
	}

	var items []chat.Chat
	for cursor.Next(ctx) {
		var item chat.Chat
		err := cursor.Decode(&item)
		if err != nil {
			return infrastructure.PaginationResponse[chat.Chat]{}, err
		}

		items = append(items, item)
	}

	paginationReponse := infrastructure.PaginationResponse[chat.Chat]{
		Limit:      pagination.Limit,
		Page:       pagination.Page,
		Count:      count,
		TotalPages: totalPages,
		Items:      items,
	}

	return paginationReponse, nil
}

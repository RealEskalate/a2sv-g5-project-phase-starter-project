package repositories

import (
	domain "AAiT-backend-group-2/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type aiRepository struct {
	chatCollection *mongo.Collection	
}

func NewAIRepository(db *mongo.Database) domain.AIRepository{
	return &aiRepository{
		chatCollection: db.Collection("chats"),
	}
}

func (ar *aiRepository) CreateChat(c context.Context,  chat *domain.ChatContext) (error){
	_, err := ar.chatCollection.InsertOne(c, chat)
	if err != nil {
		return err
	}

	return nil
} 

func (ar *aiRepository) GetChat(c context.Context, id string) (*domain.ChatContext, error){
	var chat domain.ChatContext

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	err = ar.chatCollection.FindOne(c, filter).Decode(&chat)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func (ar *aiRepository) GetChatByUserId(c context.Context, userId string) (*domain.ChatContext, error){
	var chat domain.ChatContext

	filter := bson.M{"user_id": userId}

	err := ar.chatCollection.FindOne(c, filter).Decode(&chat)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func (ar *aiRepository) UpdateChat(c context.Context, messages []domain.ChatMessage, id string) (error){
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$push": bson.M{"chat_messages": bson.M{"$each": messages}}}

	_, err = ar.chatCollection.UpdateOne(c, filter, update)

	if err != nil {
		return err
	}

	return  nil
}

func (ar *aiRepository) DeleteChat(c context.Context, id string) error{
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	_, err = ar.chatCollection.DeleteOne(c, filter)
	if err != nil {
		return err
	}
	return nil
}
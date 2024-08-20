package repository

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type AIRepository struct{
	database mongo.Database
	collection string
}

func NewAIRepository(db mongo.Database) *AIRepository{
	return &AIRepository{
		database: db,
		collection: "chat_context",
	}
}

func (air *AIRepository) CreateChat(chat *domain.ChatContext) (*domain.ChatContext,error){
	ID := primitive.NewObjectID()
	chat.ID = ID
	coll := air.database.Collection(air.collection)
    _, err := coll.InsertOne(context.Background() ,chat)
	if err != nil {
		return nil, err
	}

	// Manually create the ChatContext struct based on the input data
	return chat, nil

	
}

func (air *AIRepository) GetChat(id string) (*domain.ChatContext, error){
	coll := air.database.Collection(air.collection)
	var chat domain.ChatContext
	objectId,err:= primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = coll.FindOne(context.Background(), bson.M{"_id":objectId}).Decode(&chat)
	return &chat, err
}

// for specific user
func (air *AIRepository) GetChats() ([]*domain.ChatContext, error){
	coll := air.database.Collection(air.collection)
	cursor, err := coll.Find(context.Background(), map[string]string{})
	if err != nil{
		return nil, err
	}
	var chats []*domain.ChatContext
	for cursor.Next(context.Background()){
		var chat domain.ChatContext
		err = cursor.Decode(&chat)
		if err != nil{
			return nil, err
		}
		chats = append(chats, &chat)
	}
	return chats, nil
}


func (air *AIRepository) UpdateChat(messages []domain.ChatMessage, id string) (*domain.ChatContext, error) {
	coll := air.database.Collection(air.collection)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	
	_, err = coll.UpdateOne(context.Background(), bson.M{"_id": objectId}, bson.M{"$set": bson.M{"chat_messages": messages}})
	if err != nil {
		return nil, err
	}

	// Manually create the ChatContext struct based on the input data
	chat := &domain.ChatContext{
		ID:           objectId,
		ChatMessages: messages,
		// Populate other fields of ChatContext if needed
	}

	return chat, nil
}


func (air *AIRepository) DeleteChat(id string) error{
	coll := air.database.Collection(air.collection)
	_, err := coll.DeleteOne(context.Background(), bson.M{"_id":id})
	return err
}
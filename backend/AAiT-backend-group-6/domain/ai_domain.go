package domain

import "go.mongodb.org/mongo-driver/bson/primitive"


type  ChatContext struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	ChatMessages []ChatMessage `json:"chat_messages" bson:"chat_messages"`
}

type ChatMessage struct {
	Content string `json:"content" bson:"content"`
	Role string `json:"role" bson:"role"`
}

type ChatUseCase interface {
	CreateChat(userquestion string) (*ChatContext,error)
	GetChat(id string) (*ChatContext, error)
	GetChats() ([]*ChatContext, error)
	UpdateChat(userquestion string, id string) (*ChatContext, error)
	DeleteChat(id string) error
}

type ChatRepository interface {
	CreateChat(chat *ChatContext) (*ChatContext,error)
	GetChat(id string) (*ChatContext, error)
	GetChats() ([]*ChatContext, error)
	UpdateChat(messages []ChatMessage, id string) (*ChatContext, error)
	DeleteChat(id string) error
}


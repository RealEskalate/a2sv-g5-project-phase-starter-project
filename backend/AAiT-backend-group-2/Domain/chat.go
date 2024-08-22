package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type  ChatContext struct {
	ID 				primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	UserID 			string				`json:"user_id" bson:"user_id"` 
	ChatMessages 	[]ChatMessage 		`json:"chat_messages" bson:"chat_messages"`
}

type ChatMessage struct {
	Content 	string 		`json:"content" bson:"content"`
	Role 		string 		`json:"role" bson:"role"`
}

type AIUseCase interface {
	CreateChat(c context.Context,userID, prompt string) ([]ChatMessage,error)
	GetChat(c context.Context, id string) (*ChatContext, error)
	GetChatByUserId(c context.Context, userId string) (*ChatContext, error)
	UpdateChat(c context.Context, userquestion string, id string) ([]ChatMessage, error)
	DeleteChat(c context.Context, id string) error
}


type AIRepository interface {
	CreateChat(c context.Context, chat *ChatContext) (error)
	GetChat(c context.Context, id string) (*ChatContext, error)
	GetChatByUserId(c context.Context, userId string) (*ChatContext, error)
	UpdateChat(c context.Context, messages []ChatMessage, id string) (error)
	DeleteChat(c context.Context, id string) error
}
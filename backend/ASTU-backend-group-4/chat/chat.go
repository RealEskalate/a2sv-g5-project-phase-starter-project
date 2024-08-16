package chat

import (
	"time"
	"context"
)

type Chat struct {
	ID        		string			`json:"id" bson:"_id" `
	Title     		string			`json:"title"`
	History   		[]Message		`json:"history"`
	UserID    		string			`json:"user_id"`
	CreatedAt 		time.Time		`json:"created_at"`
	UpdatedAt 		time.Time		`json:"updated_at"`
}


type Message struct{
	Text 			string			`json:"text"`
	Role			string			`json:"role"`	
	SentAt			time.Time		`json:"sent_at"`
}


type Usecase interface{
	CreateChat(userID, title string) (Chat, error)
	DeleteChat(chatID string) error
	GenerateChatTitle(message Message) (string, error)
	GetChat(chatID string) (Chat, error)
	GetChats() ([]Chat, error)
	SendMessage(chatID string, message Message) (error)
}


type Repository interface{
	CreateChat(userID, title string, ctx context.Context) (Chat, error)
	AddMessage(chatID string, message Message, ctx context.Context) (error)
	GetChat(chatID string, ctx context.Context) (Chat, error)
	GetChats(ctx context.Context) ([]Chat, error)
	DeleteChat(chatID string, ctx context.Context) (error)
}
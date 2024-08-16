package chat

import (
	"time"
	"context"
)

type Chat struct {
	ID        		string
	Title     		string
	History   		[]Message
	UserID    		string
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}


type Message struct{
	Text 			string
	Role			string
	SentAt			time.Time
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
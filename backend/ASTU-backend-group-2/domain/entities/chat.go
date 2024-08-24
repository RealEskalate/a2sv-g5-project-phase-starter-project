package entities

import (
	"context"
	"time"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" `
	UserID    primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Messages  []Message          `json:"messages" bson:"messages,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Message struct {
	Text      string    `json:"text,omitempty" bson:"text,omitempty"`
	Role      string    `json:"role,omitempty" bson:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type ChatUsecase interface {
	CreateChat(c context.Context, userId string) (Chat, error)
	GetChat(c context.Context, chatId string) (Chat, error)
	GetChats(c context.Context, userId string, limit int64, page int64) ([]Chat, mongopagination.PaginationData, error)
	CreateMessage(c context.Context, userId string, chatId string, body string) (Message, error)
	DeleteChat(c context.Context, chatId string) error
}

type ChatRepository interface {
	CreateChat(c context.Context, chat Chat) (Chat, error)
	GetChat(c context.Context, chatID string) (Chat, error)
	GetChats(c context.Context, userID string, limit int64, page int64) ([]Chat, mongopagination.PaginationData, error)
	CreateMessage(c context.Context, chatID string, body Message) error
	DeleteChat(c context.Context, chatID string) error
}

type AI interface {
	SendMessage(c context.Context, messages []Message, message Message) (Message, error)
}

package chat

import (
	"context"
	"time"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
)

type Chat struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty" `
	Title     string    `json:"title,omitempty"`
	History   []Message `json:"history"`
	UserID    string    `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Message struct {
	Text   string    `json:"text,omitempty"`
	Role   string    `json:"role,omitempty"`
	SentAt time.Time `json:"sent_at,omitempty"`
}

type Usecase interface {
	CreateChat(userID, title string) (Chat, error)
	DeleteChat(chatID string) error
	GenerateChatTitle(message Message) (string, error)
	GetChat(chatID string) (Chat, error)
	GetChats() ([]Chat, error)
	SendMessage(chatID string, message Message) error
}

type Repository interface {
	CreateChat(chat Chat, ctx context.Context) (Chat, error)
	AddMessage(chatID string, message Message, ctx context.Context) error
	GetChat(chatID string, ctx context.Context) (Chat, error)
	GetChats(ctx context.Context, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Chat], error)
	DeleteChat(chatID string, ctx context.Context) error
}

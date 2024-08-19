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
	CreateChat(ctx context.Context, userID, title string) (Chat, error)
	DeleteChat(ctx context.Context, chatID string) error
	GenerateChatTitle(ctx context.Context, message Message) (string, error)
	GetChat(ctx context.Context, chatID string) (Chat, error)
	GetChats(ctx context.Context, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Chat], error)
	SendMessage(ctx context.Context, chatID string, message Message) error
}

type Repository interface {
	AddMessage(ctx context.Context, chatID string, message Message) error
	CreateChat(ctx context.Context, chat Chat) (Chat, error)
	DeleteChat(ctx context.Context, chatID string) error
	GetChat(ctx context.Context, chatID string) (Chat, error)
	GetChats(ctx context.Context, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Chat], error)
}

package chat

import (
	"context"
	"time"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
)

type Chat struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty" `
	Title     string    `json:"title"`
	History   []Message `json:"history"`
	UserID    string    `json:"user_id" bson:"user_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Message struct {
	Text   string    `json:"text,omitempty"`
	Role   string    `json:"role,omitempty"`
	SentAt time.Time `json:"sent_at,omitempty"`
}

type Usecase interface {
	CreateChat(ctx context.Context, form UserIDForm) (Chat, error)
	DeleteChat(ctx context.Context, form DefaultChatForm) error
	GenerateChatTitle(ctx context.Context, form TextForm) (string, error)
	GetChat(ctx context.Context, form DefaultChatForm) (Chat, error)
	GetChats(ctx context.Context, form UserIDForm, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Chat], error)
	SendMessage(ctx context.Context, chatForm DefaultChatForm, textForm TextForm) (Message, error)
	UpdateChat(ctx context.Context, form DefaultChatForm, updatedChat Chat) (Chat, error)
}

type Repository interface {
	AddMessage(ctx context.Context, chatID string, message Message) error
	CreateChat(ctx context.Context, chat Chat) (Chat, error)
	DeleteChat(ctx context.Context, chatID string) error
	GetChat(ctx context.Context, chatID string) (Chat, error)
	GetChats(ctx context.Context, userID string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Chat], error)
	UpdateChat(ctx context.Context, chatID string, updatedChat Chat) (Chat, error)
}

type AIService interface {
	SendMessage(ctx context.Context, history []Message, message Message) (Message, error)
	GenerateChatTitle(ctx context.Context, text string) (string, error)
}

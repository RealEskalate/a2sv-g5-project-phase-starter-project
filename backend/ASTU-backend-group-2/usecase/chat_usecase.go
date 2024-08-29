package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	gemini "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/aiutil"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatUsecase struct {
	AIService      gemini.AI
	contextTimeout time.Duration
	ChatRepository entities.ChatRepository
}

func NewChatUsecase(repository entities.ChatRepository, aiService gemini.AI, timeout time.Duration) entities.ChatUsecase {
	return &ChatUsecase{
		ChatRepository: repository,
		AIService:      aiService,
		contextTimeout: timeout,
	}
}

func (cu *ChatUsecase) CreateChat(c context.Context, userId primitive.ObjectID) (entities.Chat, error) {

	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()

	command := fmt.Sprint(
		"Hi there! You can send me any of the following:",
		"",
		"1. **A brief overview of your blog topic:** I can help you create a polished and engaging blog post based on your ideas.",
		"2. **A draft or existing blog post:** I can offer comprehensive feedback to enhance its clarity, impact, and overall quality.",
		"",
		"Don’t hesitate to mention any specific guidelines or preferences you might have.",
	)

	firstMessage := entities.Message{
		Text:      command,
		Role:      "model",
		CreatedAt: time.Now(),
	}

	newChat := entities.Chat{
		Title:     "New Chat",
		UserID:    userId,
		Messages:  []entities.Message{firstMessage},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newChat, err := cu.ChatRepository.CreateChat(ctx, newChat)
	if err != nil {
		return entities.Chat{}, err
	}

	return newChat, nil
}

func (cu *ChatUsecase) DeleteChat(c context.Context, userId primitive.ObjectID, chatId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()

	chat, err := cu.ChatRepository.GetChat(ctx, chatId)
	if err != nil {
		return err
	}

	if chat.UserID != userId {
		return custom_error.ErrChatNotFound
	}

	return cu.ChatRepository.DeleteChat(ctx, chatId)
}

func (cu *ChatUsecase) GetChat(c context.Context, userId primitive.ObjectID, chatId primitive.ObjectID) (entities.Chat, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()

	chat, err := cu.ChatRepository.GetChat(ctx, chatId)
	if err != nil {
		return entities.Chat{}, err
	}

	if chat.UserID != userId {
		return entities.Chat{}, custom_error.ErrChatNotFound
	}

	return chat, nil
}

func (cu *ChatUsecase) GetChats(c context.Context, userId primitive.ObjectID, limit int64, skip int64) (*[]entities.Chat, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()

	return cu.ChatRepository.GetChats(ctx, userId, limit, skip)
}

func (cu *ChatUsecase) CreateMessage(c context.Context, userId primitive.ObjectID, chatId primitive.ObjectID, body string) (entities.Message, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()

	chat, err := cu.ChatRepository.GetChat(ctx, chatId)
	if err != nil {
		return entities.Message{}, err
	}

	if chat.UserID != userId {
		return entities.Message{}, custom_error.ErrChatNotFound
	}

	message := entities.Message{
		Text:      body,
		Role:      "user",
		CreatedAt: time.Now(),
	}

	if chat.Messages == nil {
		_, err = cu.ChatRepository.UpdateChat(ctx, chatId, chat)
		if err != nil {
			return entities.Message{}, err
		}
	}

	response, err := cu.AIService.SendMessage(ctx, chat.Messages, message)
	if err != nil {
		return entities.Message{}, err
	}

	if err := cu.ChatRepository.CreateMessage(ctx, chatId, message); err != nil {
		return entities.Message{}, err
	}

	return response, nil
}

func (cu *ChatUsecase) UpdateChat(c context.Context, userId primitive.ObjectID, chatId primitive.ObjectID, updatedChat entities.Chat) (entities.Chat, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()

	return cu.ChatRepository.UpdateChat(ctx, chatId, updatedChat)
}

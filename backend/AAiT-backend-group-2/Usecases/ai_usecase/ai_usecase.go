package ai_usecase

import (
	domain "AAiT-backend-group-2/Domain"
	"AAiT-backend-group-2/Infrastructure/services"
	"context"
	"time"

	"github.com/google/generative-ai-go/genai"
)



type aiUsecase struct {
	repo domain.AIRepository
	aiService services.AIService
	contextTimeout time.Duration
}

func NewAIUsecase(repo domain.AIRepository, aiService services.AIService, timeout time.Duration) domain.AIUseCase{
	return &aiUsecase{
		repo: repo,
		aiService: aiService,
		contextTimeout: timeout,
	}
}

func (au *aiUsecase) CreateChat(c context.Context, userId, prompt string) ([]domain.ChatMessage,error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	response, err := au.aiService.GenerateText(prompt, []genai.Part{})

	if err != nil {
		return nil, err
	}

	chat := &domain.ChatContext{
		UserID: userId,
		ChatMessages: []domain.ChatMessage{
			{Content: prompt, Role: "user"},
			{Content: response, Role: "system"},
		},
	}

	err = au.repo.CreateChat(ctx, chat)
	if err != nil {
		return nil, err
	}

	return []domain.ChatMessage{
		{Content: prompt, Role: "user"},
		{Content: response, Role: "system"},
	}, nil
}

func (au *aiUsecase) GetChat(c context.Context, id string) (*domain.ChatContext, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	return au.repo.GetChat(ctx, id) 
}

func (au *aiUsecase) GetChatByUserId(c context.Context, userId string) (*domain.ChatContext, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	return au.repo.GetChatByUserId(ctx, userId)
}


func (au *aiUsecase) UpdateChat(c context.Context, prompt, id string) ([]domain.ChatMessage, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	chat, err := au.repo.GetChat(ctx, id)
	
	if err != nil {
		return nil,err
	}
	chatMessages := []genai.Part{}
	for _, chatMessage := range chat.ChatMessages {
		chatMess := chatMessage.Role + " : " + chatMessage.Content
		chatMessages = append(chatMessages, genai.Text(chatMess))
	}
	response, err := au.aiService.GenerateText(prompt, chatMessages)
	
	if err != nil {
		return nil,err
	}
	chat.ChatMessages = append(chat.ChatMessages, domain.ChatMessage{Content: prompt, Role: "user"})
	chat.ChatMessages = append(chat.ChatMessages, domain.ChatMessage{Content: response, Role: "system"})

	err = au.repo.UpdateChat(ctx, chat.ChatMessages, id)
	if err != nil {
		return nil, err
	}

	return []domain.ChatMessage{
		{Content: prompt, Role: "user"},
		{Content: response, Role: "system"},
	}, nil
}

func (au *aiUsecase) DeleteChat(c context.Context, id string) error {
	return nil
}
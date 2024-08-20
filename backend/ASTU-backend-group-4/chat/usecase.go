package chat

import (
	"context"
	"time"
)

type ChatUsecase struct {
	Repository Repository
	AIService  AIService
}

func (usecase *ChatUsecase) NewUsecase(repository Repository, aiService AIService) *ChatUsecase {
	return &ChatUsecase{
		Repository: repository,
		AIService:  aiService,
	}
}

func (usecase *ChatUsecase) CreateChat(ctx context.Context, form CreateChatForm) (Chat, error) {
	
	newChat := Chat{
		Title: form.Title,
		UserID: form.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newChat, err := usecase.Repository.CreateChat(ctx, newChat)
	if err != nil{
		return Chat{}, err
	}

	return newChat, nil
}






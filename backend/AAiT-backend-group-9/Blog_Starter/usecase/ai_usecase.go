package usecase

import (
	"Blog_Starter/domain"
	"Blog_Starter/utils/infrastructure"
	"context"
	"errors"
)

type AIUseCase struct {
	aiService infrastructure.AIService
}

func NewAIUseCase(aiService infrastructure.AIService) domain.AiUsecase {
	return &AIUseCase{
		aiService: aiService,
	}
}

func (uc *AIUseCase) GenerateAIContent(ctx context.Context, content string, preText string) (string, error) {
	//check the length the content is not empity
	if len(content) == 0 {
		return "", errors.New("content is empty")
	}
	return uc.aiService.GenerateAIContent(ctx, content, preText)
}

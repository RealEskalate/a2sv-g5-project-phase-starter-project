package usecases

import (
	domain "blogs/Domain"
	"context"
	"fmt"
	"time"
)

type AIUsecase struct {
	aiService domain.AIConfig
	contextTimeout time.Duration

	
}

func NewAIUsecase(aiService domain.AIConfig , timeout time.Duration) domain.AIUsecase {
	return &AIUsecase{aiService: aiService,
					contextTimeout: timeout,}
}

func (u *AIUsecase) AskAI(c context.Context , request domain.AiRequest) interface{} {
	prompt := domain.Prompt_keyword
	formattedPrompt := fmt.Sprintf(prompt, request.Message)

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, err := u.aiService.Ask(ctx, formattedPrompt)

	if err != nil {
		fmt.Println(err)
		return &domain.ErrorResponse{Message: "Cannot answer Your Question Now", Status: 500}
	}

	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		response := result.Candidates[0].Content.Parts[0] // Convert the Part to a string
		
		return &domain.AiResponse{Response: response}
	}

	return &domain.ErrorResponse{Message: "No content generated", Status: 500}

}

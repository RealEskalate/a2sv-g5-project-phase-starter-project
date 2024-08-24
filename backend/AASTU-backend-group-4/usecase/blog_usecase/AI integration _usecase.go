package blog_usecase

import (
	"context"
	"errors"

	"github.com/google/generative-ai-go/genai"
)

func (uc *BlogUsecase) GenerateAIContent(c context.Context, prompt string) (*genai.GenerateContentResponse, error) {
	if prompt == "" {
		return nil, errors.New("prompt cannot be empty")
	}
	return uc.genAIService.GenerateContent(prompt)
}

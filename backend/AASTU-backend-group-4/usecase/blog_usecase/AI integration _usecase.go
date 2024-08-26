package blog_usecase

import (
	"context"
	"errors"

	"github.com/google/generative-ai-go/genai"
)

func (uc *BlogUsecase) GenerateAIContent(c context.Context, prompt string) (genai.Part, error) {
	if prompt == "" {
		return nil, errors.New("prompt cannot be empty")
	}
	response, err := uc.genAIService.GenerateContent(prompt, uc.Env.API_KEY)
	return response, err

}

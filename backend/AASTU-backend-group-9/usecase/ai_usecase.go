package usecase

import (
	"blog/domain"
	"context"
	"time"

	"github.com/google/generative-ai-go/genai"
)

type aiUsecase struct {
	contextTimeout time.Duration
	client         *genai.Client
}

func NewAIUsecase(timeout time.Duration, client *genai.Client) domain.AIUsecase {
	return &aiUsecase{
		contextTimeout: timeout,
		client:         client,
	}
}

func (au *aiUsecase) GenerateBlogContent(ctx context.Context, keywords string) (*genai.GenerateContentResponse, error) {
	model := au.client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(keywords))
    if err != nil {
        return nil, err
    }

    return resp, nil
}

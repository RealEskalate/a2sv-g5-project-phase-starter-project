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
	model := au.client.GenerativeModel("gemini-1.5-pro-latest")
	model.SetTemperature(0.9)
	model.SetTopP(0.5)
	model.SetTopK(20)
	model.SetMaxOutputTokens(100)
	model.SystemInstruction = genai.NewUserContent(genai.Text("You are a blog writer, you generate a content for a user for blog posting."))
	model.ResponseMIMEType = "application/json"
	resp, err := model.GenerateContent(ctx, genai.Text(keywords))
    if err != nil {
        return nil, err
    }

    return resp, nil
}

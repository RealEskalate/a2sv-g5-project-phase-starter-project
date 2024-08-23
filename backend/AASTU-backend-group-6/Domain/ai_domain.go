package domain

import (
	"context"

	"github.com/google/generative-ai-go/genai"
)

type AiRequest struct {
	Message string `json:"message"`
}
type BlogPost struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type AiResponse struct {
	Response genai.Part `json:"response"`
}

type AIUsecase interface {
	AskAI(c context.Context, request AiRequest) interface{}
}

type AIConfig interface {
	Ask(ctx context.Context, question string) (*genai.GenerateContentResponse, error)
}

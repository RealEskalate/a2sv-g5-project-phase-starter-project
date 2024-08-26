package domain

import "github.com/google/generative-ai-go/genai"

type AIContentGenerator interface {
	GenerateContent(prompt string, API_key string) (genai.Part, error)
}

type GenerateContentRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

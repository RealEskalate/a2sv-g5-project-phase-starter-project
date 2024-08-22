package domain

import "github.com/google/generative-ai-go/genai"

type AIContentGenerator interface {
	GenerateContent(prompt string) (*genai.GenerateContentResponse, error)
}

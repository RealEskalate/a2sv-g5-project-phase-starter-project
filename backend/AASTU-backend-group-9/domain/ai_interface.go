package domain

import (
	"context"

	"github.com/google/generative-ai-go/genai"
)

// AIUsecase defines the methods for AI content generation.
type AIUsecase interface {
	GenerateBlogContent(ctx context.Context, keywords string) ([]genai.Part, error)
}

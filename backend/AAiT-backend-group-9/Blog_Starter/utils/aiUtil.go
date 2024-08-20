package utils

import "context"

type AIService interface {
	GenerateAIContent(ctx context.Context, content string) (string, error)
}

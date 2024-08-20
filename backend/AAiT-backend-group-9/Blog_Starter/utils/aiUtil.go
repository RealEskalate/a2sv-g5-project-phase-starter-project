package utils

import "context"

type AIService interface {
	GenerateAIContent(ctx context.Context, content string, preText string) (string, error)
}

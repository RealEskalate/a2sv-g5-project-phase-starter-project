package domain

import "context"

type AiUsecase interface {
	GenerateAIContent(ctx context.Context, content string) (string, error)
}

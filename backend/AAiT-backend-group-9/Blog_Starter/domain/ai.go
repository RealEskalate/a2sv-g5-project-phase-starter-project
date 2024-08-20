package domain

import "context"

type AiUsecase interface {
	GenerateAIContent(ctx context.Context, content string, preText string) (string, error)
}

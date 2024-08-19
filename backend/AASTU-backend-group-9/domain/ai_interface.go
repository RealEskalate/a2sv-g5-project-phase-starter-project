package domain

import (
    "context"
)

// AIUsecase defines the methods for AI content generation.
type AIUsecase interface {
    GenerateBlogContent(ctx context.Context, keywords string) (string, error)
}
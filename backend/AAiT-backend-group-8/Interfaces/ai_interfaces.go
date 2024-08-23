package Interfaces

import (
	"AAiT-backend-group-8/Domain"
	"context"
)

type IAiService interface {
	GenerateContent(ctx context.Context, userInput string) (string, error)
	SuggestImprovements(ctx context.Context, title, body, tags string) (string, error)
}
type IAiUsecase interface {
	GenerateBlogContent(userInput string) (Domain.BlogResponse, error)
	SuggestImprovements(title, body string, tags []string) (Domain.SuggestionBlogResponse, error)
}

package domain

import "context"


type ContentSuggestion struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}


type AIUsecase interface {
	ContentSuggestions(ctx context.Context, userID string) ([]ContentSuggestion, error)
	ContentEnhancements(ctx context.Context, content string) (string, error)
}

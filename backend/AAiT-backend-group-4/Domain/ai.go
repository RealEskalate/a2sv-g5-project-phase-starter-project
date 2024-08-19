package domain

import "context"

type AiUsecase interface {
	GenerateTextWithTags(c context.Context, tags []Tag) (string, error)
	GenerateTextWithPromot(c context.Context, prompt string) (string, error)
	GenerateSuggestions(c context.Context, textContent string) (string, error)
	Chat(c context.Context, textContent string) (string, error)
}

type AiRepository interface {
	GenerateText(c context.Context, prompt string) (string, error)
	GenerateSuggestions(c context.Context, textContent string) (string, error)
	Chat(c context.Context, textContent string) (string, error)
}

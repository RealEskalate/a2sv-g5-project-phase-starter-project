package interfaces 

import "aait.backend.g10/domain"

type IAIService interface {
    GenerateContent(topic string, keywords []string) (*domain.BlogContentResponse, error)
	SuggestImprovements(content string) (*domain.SuggestionResponse, error)
}

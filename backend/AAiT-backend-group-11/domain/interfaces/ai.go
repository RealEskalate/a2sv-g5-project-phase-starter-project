package interfaces

import  "backend-starter-project/domain/entities"

type AIContentService interface {
	GenerateContentSuggestions(keywords []string) (*entities.ContentSuggestion, error)
	SuggestContentImprovements(blogPostId, content string) (*entities.ContentSuggestion, error)
}
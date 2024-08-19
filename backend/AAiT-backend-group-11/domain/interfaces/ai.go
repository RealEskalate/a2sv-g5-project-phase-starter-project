package interfaces

import (
	"backend-starter-project/domain/entities"
	"context"
)

type AIContentService interface {
	GenerateContentSuggestions(c context.Context,keywords []string) (*entities.ContentSuggestion, error)
	SuggestContentImprovements(c context.Context, blogPostId, content string) (*entities.ContentSuggestion, error)
}
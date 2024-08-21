package interfaces

import models "github.com/aait.backend.g5.main/backend/Domain/Models"

type AIContentSuggestionUsecase interface {
	SuggestContent(AI_query string) ([]string, *models.ErrorResponse)
	ImproveBlogContent(blogID string) ([]string, *models.ErrorResponse)
}

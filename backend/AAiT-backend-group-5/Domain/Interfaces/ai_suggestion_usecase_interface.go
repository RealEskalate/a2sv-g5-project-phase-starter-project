package interfaces

type AIContentSuggestionUsecase interface {
	SuggestContent(AI_query string) (string, error)
	ImproveBlogContent(blogID string) (string, error)
}

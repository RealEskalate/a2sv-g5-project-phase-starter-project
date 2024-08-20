package interfaces

type ContentSuggester interface {
	SuggestContent(input string) (string, error)
}

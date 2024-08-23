package Dtos

type PostDTO struct {
	Categories        []string `json:"categories"`
	MaxWordLimit      int      `json:"max_word_limit"`
	ParagraphLimit    int      `json:"paragraph_limit"`
	Title             string   `json:"title"`
	Keywords          []string `json:"keywords"`
	Tone              string   `json:"tone"`
	Format            string   `json:"format"`
	AdditionalContext string   `json:"additional_context"`
}

package Domain

type BlogRequest struct {
	UserInput string `json:"user_input"`
}

type BlogSuggestionRequest struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

type BlogResponse struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

type SuggestionBlogResponse struct {
	Comment string   `json:"comment"`
	Title   string   `json:"title"`
	Body    string   `json:"body"`
	Tags    []string `json:"tags"`
}

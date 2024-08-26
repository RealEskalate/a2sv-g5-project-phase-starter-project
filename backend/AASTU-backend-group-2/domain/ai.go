package domain

type PostResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AIService interface {
	GeneratePost(title, description string) (*PostResponse, error)
	Validate_Comment(comment string) error
	Validate_Blog(blog string) error
}

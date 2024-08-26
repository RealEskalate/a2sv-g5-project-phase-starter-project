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
	GeneratePost(title, description string) (*PostResponse, *AppError)
	Validate_Comment(comment string) *AppError
	Validate_Blog(blog string) *AppError
}

package domain


var AIBlog struct {
	Title string   `json:"title" binding:"required"`
	Tags  []string `json:"tags"`
}

type Content struct {
	Parts []string `json:"parts"`
	Role  string   `json:"role"`
}

type Candidate struct {
	Content *Content `json:"content"`
}

type Response struct {
	Candidates []*Candidate `json:"candidates"`
}


type AIUsecaseInterface interface {
	GenerateBlogContent(title string, tags []string) (string, error)
}
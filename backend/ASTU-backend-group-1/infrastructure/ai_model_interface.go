package infrastructure

type RecommendationResponse struct {
	Title   string
	Content string
	Tags    []string
}
type Data struct {
	Title   string
	Content string
	Tags    []string
}
type AIModel interface {
	Recommend(Data, opt string) (RecommendationResponse, error)
	Summarize(Data) (string, error)
	Validate(Data) error
	Chat(prompt string) (string, error)
}

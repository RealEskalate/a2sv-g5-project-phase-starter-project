package infrastructure

type Data struct {
	Title   string   `json:"title,omitempty"`
	Content string   `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}
type AIModel interface {
	Recommend(data Data, opt string) (interface{}, error)
	Summarize(Data) (string, error)
	Validate(Data) error
	Chat(prompt string) (string, error)
}

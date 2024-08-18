package domain

type AI_Output struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AI_Input struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// request payload to send to the gemini api
type GeminiRequest struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
	Model     string `json:"model"`
}

// response structure from the gemini api
type GeminiResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

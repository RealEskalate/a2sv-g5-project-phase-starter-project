package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIUtil interface {
	GenerateContentFromGemini(title string, description string, env bootstrap.Env) (string, error)
}

type AI struct{}

func NewAIUtil() AIUtil {
	return &AI{}
}

type Content struct {
	Parts []string `json:"parts"`
	Role  string   `json:"role"`
}

type Candidates struct {
	Content *Content `json:"content"`
}

type ContentResponse struct {
	Candidates *[]Candidates `json:"candidates"`
}

func (AI) GenerateContentFromGemini(title string, description string, env bootstrap.Env) (string, error) {

	client, err := genai.NewClient(context.TODO(), option.WithAPIKey(env.GeminiAPIKey))

	if err != nil {
		return "", fmt.Errorf("error initializing Gemini client: %v", err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	words, err := strconv.Atoi(env.GeminiWordCount)

	if err != nil {
		return "", fmt.Errorf("error converting word count to integer: %v", err)
	}

	prompt := []genai.Part{
		genai.Text(fmt.Sprintf("Create a blog post titled '%s' with approximately %d words of content. The post should focus on the following key points: '%s'. Please ensure the content is informative, engaging, and well-structured.",
			title,
			words,
			description)),
	}

	resp, err := model.GenerateContent(context.TODO(), prompt...)
	if err != nil {
		return "", fmt.Errorf("error generating content: %v", err)
	}

	marshalResponse, _ := json.MarshalIndent(resp, "", "  ")

	var generateResponse ContentResponse
	if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	var blogContent string
	for _, cad := range *generateResponse.Candidates {
		if cad.Content != nil {
			for _, part := range cad.Content.Parts {
				blogContent += part
			}
		}
	}

	return blogContent, nil
}

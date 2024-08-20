package infrastructure

import (
	"context"
	"fmt"
	"log"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiContentSuggester struct {
	Client *genai.Client
}

func NewAIContentSuggester(apiKey string) interfaces.ContentSuggester {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		log.Fatal(err)
	}
	return &GeminiContentSuggester{
		Client: client,
	}
}

func (g *GeminiContentSuggester) SuggestContent(input string) (string, error) {
	ctx := context.Background()
	model := g.Client.GenerativeModel("gemini-1.5-flash")

	model.SystemInstruction = genai.NewUserContent(genai.Text("Suggest the content of a blog post about"))
	resp, err := model.GenerateContent(ctx, genai.Text(input))
	if err != nil {
		log.Fatal(err)
	}

	// get the first from an array of responses
	return fmt.Sprintf("%v", resp.Candidates[0].Content), err
}

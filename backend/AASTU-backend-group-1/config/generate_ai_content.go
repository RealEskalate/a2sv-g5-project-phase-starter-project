package config

import (
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// GenerateAIContent generates AI content

func GenerateAIContent(prompt string) genai.Part {
	var generatedText genai.Part
	ctx := context.Background()
	client, err := genai.NewClient(context.Background(), option.WithAPIKey("AIzaSyCuwT3g4-x5xsIvl1VgCJKZ-VO48JUIvVY"))
	if err != nil {
		panic(err)
	}

	// Generate content
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if resp != nil {
		candidates := resp.Candidates

		for _, candidate := range candidates {
			content := candidate.Content

			if content != nil {
				text := content.Parts[0]
				generatedText = text
			}
		}

	}
	return generatedText
}

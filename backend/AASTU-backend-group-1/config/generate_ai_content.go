package config

import (
	"context"
	"encoding/json"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

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

func GenerateAIContent(prompt string) (string, error) {
	var generatedText string
	ctx := context.Background()
	client, err := genai.NewClient(context.Background(), option.WithAPIKey("AIzaSyCuwT3g4-x5xsIvl1VgCJKZ-VO48JUIvVY"))
	if err != nil {
		panic(err)
	}

	// Generate content
	model := client.GenerativeModel("gemini-1.5-flash")
	model.ResponseMIMEType = "text/plain"
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))

	if err != nil {
		panic(err)
	}

	marshalResp, err := json.MarshalIndent(resp, "", "  ")

	if err != nil {
		panic(err)
	}

	var generateResponse Response

	if err := json.Unmarshal(marshalResp, &generateResponse); err != nil {
		panic(err)
	}

	if resp != nil {

		for _, candcandidates := range generateResponse.Candidates {
			content := candcandidates.Content
			parts := content.Parts
			for _, part := range parts {
				generatedText += part

			}

		}
	}
	return generatedText, nil

}

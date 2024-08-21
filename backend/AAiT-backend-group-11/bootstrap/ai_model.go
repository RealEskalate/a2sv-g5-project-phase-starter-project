package bootstrap

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func NewAiModel(env *Env) *genai.GenerativeModel{
	client, err := genai.NewClient(context.TODO(), option.WithAPIKey(env.GeminiApiKey))
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	return model
}
package gemini

import (
	"context"
	"os"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type ReccomendationHandler struct {
	model genai.GenerativeModel
}

func NewReccomendationHandler() *ReccomendationHandler {
	return &ReccomendationHandler{}
}

func (h *ReccomendationHandler) Handle(cmd *RecommendationCommand) (*genai.GenerateContentResponse, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(
		ctx,
		genai.Text(*cmd.request),        
	)

	return resp, err
}
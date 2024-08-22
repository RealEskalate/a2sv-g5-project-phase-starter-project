package gemini

import "github.com/google/generative-ai-go/genai"

type RecommendationCommand struct {
	request *string 
	client *genai.Client

}

// new recommendation command 
func NewRecommendationCommand(request *string) *RecommendationCommand{
	return &RecommendationCommand{
		request: request,
	}
}
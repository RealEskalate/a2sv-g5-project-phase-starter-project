package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIContentService struct{}

func NewAIContentService() interfaces.AIContentService {
	return &AIContentService{}
}

func (acs *AIContentService) GenerateContentSuggestions(c context.Context,keywords []string) (*entities.ContentSuggestion, error) {
	
	client, err := genai.NewClient(c, option.WithAPIKey("AIzaSyDGyJ24ipTKSjoGbIsaMd_Np9dqvtd78XI"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(c, genai.Text("Write a story about a AI and magic"))
	if err != nil {
		log.Fatal(err)
	}

	candidates := resp.Candidates
	for _, candidate := range candidates {
		log.Printf("Generated text: %s", candidate.Content.Parts)
	}

	return nil,nil

}
func (acs *AIContentService) SuggestContentImprovements(c context.Context ,blogPostId, content string) (*entities.ContentSuggestion, error) {
	client, err := genai.NewClient(c, option.WithAPIKey("AIzaSyDGyJ24ipTKSjoGbIsaMd_Np9dqvtd78XI"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(c, genai.Text("Write a story about a AI and magic"))
	if err != nil {
		log.Fatal(err)
	}
	
	//create a string builder
	var generatedAnswer strings.Builder
	candidates := resp.Candidates
	for _, candidate := range candidates {
		//add the candidate to the string builder
		_ = candidate

	}
	_ = generatedAnswer

	return nil,nil


}

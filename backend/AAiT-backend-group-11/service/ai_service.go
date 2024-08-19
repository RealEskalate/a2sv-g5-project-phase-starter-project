package service

import (
	"context"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

type AIContentServiceInterface interface {
	GenerateContentSuggestions(keywords []string) (*genai.GenerateContentResponse, error)
	SuggestContentImprovements( blogPostId, content string) (*genai.GenerateContentResponse, error)
}

type AIContentService struct{
	ctx context.Context
	model genai.GenerativeModel
}

func NewAIContentService( ctx context.Context, model genai.GenerativeModel) AIContentServiceInterface {
	return &AIContentService{ ctx: ctx, model: model}
}

func (acs *AIContentService) GenerateContentSuggestions(keywords []string) (*genai.GenerateContentResponse, error) {

	// Construct the prompt from the keywords
	prompt := "Write blog content about: " + strings.Join(keywords, ", ")

	resp, err :=acs. model.GenerateContent(acs.ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}
	return resp,nil
}


func (acs *AIContentService) SuggestContentImprovements( blogPostId, content string) (*genai.GenerateContentResponse, error) {

	// Generate a prompt for content improvement
	prompt := "Improve the following blog: " + content

	resp, err := acs.model.GenerateContent(acs.ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	

	return resp, nil
}

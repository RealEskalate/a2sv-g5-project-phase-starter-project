package service

import (
	"backend-starter-project/domain/interfaces"
	"context"
	"errors"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

type AIContentServiceInterface interface {
	GenerateContentSuggestions(keywords []string) (*genai.GenerateContentResponse, error)
	SuggestContentImprovements( blogPostId string) (*genai.GenerateContentResponse, error)
}

type AIContentService struct{
	ctx context.Context
	model genai.GenerativeModel
	blogPostRepository interfaces.BlogRepository
}

func NewAIContentService( ctx context.Context, model genai.GenerativeModel, bpr interfaces.BlogRepository) AIContentServiceInterface {
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


func (acs *AIContentService) SuggestContentImprovements(blogPostId string) (*genai.GenerateContentResponse, error) {

	// Fetch the blog post content from the database
	 blogPost, err := acs.blogPostRepository.GetBlogPostById(blogPostId)
	if err != nil {
		return nil, errors.New("blog post not found")
	}

	content := blogPost.Content



	// Generate a prompt for content improvement
	prompt := "Improve the following blog: " + content

	resp, err := acs.model.GenerateContent(acs.ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	

	return resp, nil
}

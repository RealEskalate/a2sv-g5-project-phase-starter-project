package service

import (
	"backend-starter-project/domain/interfaces"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

type AIContentServiceInterface interface {
	GenerateContentSuggestions(keywords []string) (string, error)
	SuggestContentImprovements( blogPostId,instruction string) (string, error)
}

type AIContentService struct{
	ctx context.Context
	model *genai.GenerativeModel
	blogPostRepository interfaces.BlogRepository
}

func NewAIContentService( ctx context.Context, model *genai.GenerativeModel, bpr interfaces.BlogRepository) AIContentServiceInterface {
	return &AIContentService{ ctx: ctx, model: model, blogPostRepository: bpr}
}

func (acs *AIContentService) GenerateContentSuggestions(keywords []string) (string, error) {

	// Construct the prompt from the keywords

	prompt := "Write blog content about: " + strings.Join(keywords, ", ")
	resp, err :=acs. model.GenerateContent(acs.ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	ans := textGenerator(*resp)


	return ans,nil
}


func (acs *AIContentService) SuggestContentImprovements(blogPostId,instruction string) (string, error) {

	// Fetch the blog post content from the database
	 blogPost, err := acs.blogPostRepository.GetBlogPostById(blogPostId)
	if err != nil {
		return "", errors.New("blog post not found")
	}


	content := blogPost.Content

	

	// Generate a prompt for content improvement
	prompt := "Improve the following blog: " + content + "\n" + "by the following instructions: " + instruction + "\n"
	acs.model.ResponseMIMEType = "application/json"

	resp, err := acs.model.GenerateContent(acs.ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	ans := textGenerator(*resp)

	return ans, nil
}




func textGenerator(resp genai.GenerateContentResponse) string{
	var response strings.Builder
	for _, candidate := range resp.Candidates{
		if candidate != nil {
			content :=  candidate.Content
				if content != nil{
					text := fmt.Sprint(content.Parts)
					response.WriteString(text)
				}
			}
		}

	return response.String()
	}
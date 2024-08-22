package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"AAIT-backend-group-3/internal/domain/models"

	"github.com/google/generative-ai-go/genai"
)


type AiService struct {
	client *genai.Client

}

func NewAiService(client *genai.Client) *AiService {
	return &AiService{
		client: client,
	}

}

func (AI *AiService) GenerateBlog(blogDescription string) (models.GeneratedBlog, error) {

	ctx := context.Background()
	model := AI.client.GenerativeModel("gemini-1.5-flash")
	model.ResponseMIMEType = "application/json"
	model.SystemInstruction  = genai.NewUserContent(genai.Text(`
	You are a content creater and you want 
	to generate a blog post based on a description, don't make it HTML, just write a blog post.,
	 it should be at least 500 words and maximum of 1000 words. it should have fields like title, content and  tags. nothing more`))
	resp, err := model.GenerateContent(ctx, genai.Text(blogDescription))
	if err != nil {
		log.Fatal(err)
	}

	formatedResponse := formatResponse(resp)
	var formatedBlog models.GeneratedBlog
	err = json.Unmarshal([]byte(formatedResponse), &formatedBlog)
	if err != nil {
		return models.GeneratedBlog{}, err
	}
	return formatedBlog, nil
}

func (AI *AiService) EnhanceBlog(blogContent string) (models.GeneratedBlog, error) {

	ctx := context.Background()
	model := AI.client.GenerativeModel("gemini-1.5-flash")
	model.ResponseMIMEType = "application/json"
	model.SystemInstruction  = genai.NewUserContent(genai.Text(`
	You are a content creater and you want to enhance and rewrite the blog 
	from provided blog content, if the content seems to be unfinished 
	complete the blog with out change the core message. 
	if there are grammar or vocabulary errors, fix them. 
	it should be at least 100 words.
	it should have fields like title, content and  tags. nothing more
	choose the relevant title and tags for the blog post.
	`))
	resp, err := model.GenerateContent(ctx, genai.Text(blogContent))
	if err != nil {
		log.Fatal(err)
	}
	
	formatedResponse := formatResponse(resp)
	var formatedBlog models.GeneratedBlog
	err = json.Unmarshal([]byte(formatedResponse), &formatedBlog)
	if err != nil {
		return models.GeneratedBlog{}, err
	}
	return formatedBlog, nil
	
	}

func (AI *AiService) GenerateSummary(blogContent string) (string, error) {
	ctx := context.Background()
	model := AI.client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = genai.NewUserContent(genai.Text(`
	Your role is to generate a summary of a blog post into one short sentence. 
	Avoid starting with 'This blog...' and keep the summary under 10 words.`))

	model.ResponseMIMEType = "application/json"

	resp, err := model.GenerateContent(ctx, genai.Text(blogContent))
	if err != nil {
		log.Fatal(err)
	}

	formatedResponse := formatResponse(resp)
	if err != nil {
		return "", err
	}
	var summary map[string]string
	err = json.Unmarshal([]byte(formatedResponse), &summary)
	if err != nil {
		return "", err
	}

	return summary["summary"], nil
}

func (AI *AiService) GenerateTags(blogContent string) ([]string, error) {
	ctx := context.Background()
	model := AI.client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = genai.NewUserContent(genai.Text(`Your role is to generate tags for a blog post. 
	Tags should be single words or short phrases that describe 
	the content of the blog post. give at least 3 tags. maximum 10 words
	make the tags capitalized. your respone should have "tags" field with list of tags.
	`))
	model.ResponseMIMEType = "application/json"
	
	resp, err := model.GenerateContent(ctx, genai.Text( blogContent))
	if err != nil {
		log.Fatal(err)
	}

	type blogTags struct {
		Tags []string `json:"tags"`
	}
	
	formatedResponse := formatResponse(resp)
	var tags blogTags
	err = json.Unmarshal([]byte(formatedResponse), &tags)
	if err != nil {
		return []string{}, err
	}

	return tags.Tags, nil
	}

	
// taken from https://ai.google.dev/gemini-api/docs/text-generation?lang=go
func formatResponse(resp *genai.GenerateContentResponse) string {
	formated := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				formated += fmt.Sprintln(part)
			}
		}
	}
	return formated
}
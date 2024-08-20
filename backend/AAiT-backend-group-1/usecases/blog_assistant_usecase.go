package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	
	"strings"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type blogAssistantUsecase struct {
	client *genai.Client
}

func NewBlogAssistantUsecase(apiKey string) domain.BlogAssistantUseCase {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	return &blogAssistantUsecase{client: client}
}

func (b *blogAssistantUsecase) GenerateBlog(keywords []string, tone, audience string) (map[string]interface{}, domain.Error) {
	prompt := "Write a blog using these keywords: " + strings.Join(keywords, ", ") + ". Focus on " + tone + "for " + audience + "."

	model := b.client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = genai.NewUserContent(genai.Text("Your task is to create a well-structured and engaging blog post based on the provided keywords. The blog should be informative, original, and tailored to the target audience. Ensure the content flows logically, with an introduction, body, and conclusion. Incorporate the keywords naturally throughout the text while maintaining readability. The tone should align with the specified audience (e.g., professional, casual, or academic). Include relevant subheadings, bullet points, or lists where appropriate to improve readability."))
	model.ResponseMIMEType = "application/json"

	resp, err := model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		return nil, &domain.CustomError{Message: "error generating blog", Code: 500}
	}

	var blog map[string]interface{}

	err = json.Unmarshal([]byte(formatResponse(resp)), &blog)
	if err != nil {
		return nil, &domain.CustomError{Message: "error generating blog", Code: 500}
	}

	return blog, nil
}

func (b *blogAssistantUsecase) EnhanceBlog(content, command string) (map[string]interface{}, domain.Error) {
	prompt := "Enhance this blog: " + content + "." + " " + command
	model := b.client.GenerativeModel("gemini-1.5-flash")

	model.SystemInstruction = genai.NewUserContent(genai.Text("You are required to enhance an existing blog post by improving its structure, clarity, and engagement. Focus on refining the language, enhancing the flow of ideas, and ensuring that the content is more informative and compelling. Add relevant data, examples, or statistics to strengthen the points made. Ensure that the tone and style are consistent with the target audience. Incorporate any missing elements, such as a strong introduction or conclusion, and optimize the blog for SEO by naturally integrating keywords."))
	model.ResponseMIMEType = "application/json"

	resp, err := model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		return nil, &domain.CustomError{Message: "error enhancing blog", Code: 500}
	}

	var enhancedBlog map[string]interface{}

	err = json.Unmarshal([]byte(formatResponse(resp)), &enhancedBlog)
	if err != nil {
		return nil, &domain.CustomError{Message: "error enhancing blog", Code: 500}
	}

	return enhancedBlog, nil
}

func (b *blogAssistantUsecase) SuggestBlog(industry string) ([]map[string]interface{}, domain.Error) {
	if b.client == nil {
		log.Println("Client is not initialized.")
		return nil, &domain.CustomError{Message: "Client is not initialized.", Code: 500}
	}
	prompt := "Suggest blog topics for " + industry + ". Include brief descriptions and potential keywords."

	model := b.client.GenerativeModel("gemini-1.5-flash")

	model.SystemInstruction = genai.NewUserContent(genai.Text("Your role is to suggest blog topics that are relevant, engaging, and aligned with the specified niche or industry. Consider current trends, audience interests, and potential for high engagement when generating suggestions. Provide a brief description of each blog idea, explaining the angle or perspective that could be taken, the target audience, and why it would be of interest. Include potential keywords and subtopics that could be explored within each suggested blog."))
	model.ResponseMIMEType = "application/json"

	ctx := context.Background()
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Println(err)

		return nil, &domain.CustomError{Message: err.Error(), Code: 500}
	}

	var suggestions []map[string]interface{}

	err = json.Unmarshal([]byte(formatResponse(resp)), &suggestions)
	if err != nil {
		return nil, &domain.CustomError{Message: "error generating suggestions", Code: 500}
	}


	return suggestions, nil
}

func formatResponse(resp *genai.GenerateContentResponse) string {
	res := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				res += fmt.Sprintln(part)
			}
		}
	}
	return res
}

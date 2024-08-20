package usecase

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type blogAssistantUsecase struct {
	client *genai.Client
}

func NewBlogAssistantUsecase() domain.BlogAssistantUseCase {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	return &blogAssistantUsecase{client: client}
}

func (b *blogAssistantUsecase) GenerateBlog(keywords []string, tone, audience string) (string, domain.Error) {
	prompt := "Write a blog using these keywords: " + strings.Join(keywords, ", ") + ". Focus on " + tone + "for " + audience + "."

	model := b.client.GenerativeModel("gemini-1.5-flash")
	model.SetMaxOutputTokens(200)
	model.SystemInstruction = genai.NewUserContent(genai.Text("Your task is to create a well-structured and engaging blog post based on the provided keywords. The blog should be informative, original, and tailored to the target audience. Ensure the content flows logically, with an introduction, body, and conclusion. Incorporate the keywords naturally throughout the text while maintaining readability. The tone should align with the specified audience (e.g., professional, casual, or academic). Include relevant subheadings, bullet points, or lists where appropriate to improve readability."))

	resp, err := model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	return formatResponse(resp), nil
}

func (b *blogAssistantUsecase) EnhanceBlog(content, command string) (string, domain.Error) {
	prompt := "Enhance this blog: " + content + "." + " " + command
	model := b.client.GenerativeModel("gemini-1.5-flash")
	model.SetMaxOutputTokens(100)
	model.SystemInstruction = genai.NewUserContent(genai.Text("You are required to enhance an existing blog post by improving its structure, clarity, and engagement. Focus on refining the language, enhancing the flow of ideas, and ensuring that the content is more informative and compelling. Add relevant data, examples, or statistics to strengthen the points made. Ensure that the tone and style are consistent with the target audience. Incorporate any missing elements, such as a strong introduction or conclusion, and optimize the blog for SEO by naturally integrating keywords."))

	resp, err := model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	return formatResponse(resp), nil
}

func (b *blogAssistantUsecase) SuggestBlog(industry string) ([]string, domain.Error) {
	prompt := "Suggest blog topics for " + industry + ". Include brief descriptions and potential keywords."

	model := b.client.GenerativeModel("gemini-1.5-flash")
	model.SetMaxOutputTokens(100)
	model.SystemInstruction = genai.NewUserContent(genai.Text("Your role is to suggest blog topics that are relevant, engaging, and aligned with the specified niche or industry. Consider current trends, audience interests, and potential for high engagement when generating suggestions. Provide a brief description of each blog idea, explaining the angle or perspective that could be taken, the target audience, and why it would be of interest. Include potential keywords and subtopics that could be explored within each suggested blog."))

	resp, err := model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	suggestions := []string{}
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				suggestions = append(suggestions, fmt.Sprint(part))
			}
		}
	}

	return suggestions, nil
}

func formatResponse(resp *genai.GenerateContentResponse) string {
	res := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				res += fmt.Sprint(part) + " "
			}
		}
	}
	return res
}

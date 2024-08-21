package usecases

import (
	"context"
	"fmt"
	"log"
	"strings"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/google/generative-ai-go/genai"
)

type AISuggestionUsecase struct {
	Model          *genai.GenerativeModel
	BlogRepository interfaces.BlogRepository
}

func NewAISuggestionUsecase(client *genai.Client, blogRepository interfaces.BlogRepository) interfaces.AIContentSuggestionUsecase {
	// ctx := context.Background()
	// client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))

	// if err != nil {
	// 	log.Fatal(err)
	// }
	return &AISuggestionUsecase{
		Model:          client.GenerativeModel("gemini-1.5-flash"),
		BlogRepository: blogRepository,
	}
}

func (uc *AISuggestionUsecase) SuggestContent(AI_query string) (string, error) {
	ctx := context.Background()

	uc.Model.SystemInstruction = genai.NewUserContent(genai.Text("Suggest the content of a blog post about"))
	resp, err := uc.Model.GenerateContent(ctx, genai.Text(AI_query))
	if err != nil {
		log.Fatal(err)
	}

	// get the first from an array of responses
	return fmt.Sprintf("%v", resp.Candidates[0].Content), err
}

func (uc *AISuggestionUsecase) ImproveBlogContent(blogID string) (string, error) {
	ctx := context.Background()

	blog, e := uc.BlogRepository.GetBlog(ctx, blogID)
	if e != nil {
		return "", e
	}

	// generate content improvement queries
	uc.Model.SystemInstruction = genai.NewUserContent(genai.Text("Improve the content of the following blog post"))
	blogInfo := fmt.Sprintf("Blog title: %s\nBlog content: %s\nBlog tags: %s", blog.Title, blog.Content, strings.Join(blog.Tags, ", "))

	resp, err := uc.Model.GenerateContent(ctx, genai.Text(blogInfo))
	if err != nil {
		log.Fatal(err)
	}

	// get the first from an array of responses
	return fmt.Sprintf("%v", resp.Candidates[0].Content), err
}

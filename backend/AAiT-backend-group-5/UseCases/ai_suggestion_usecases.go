package usecases

import (
	"context"
	"fmt"
	"strings"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/google/generative-ai-go/genai"
)

type AISuggestionUsecase struct {
	Model          *genai.GenerativeModel
	BlogRepository interfaces.BlogRepository
}

func NewAISuggestionUsecase(client *genai.Client, blogRepository interfaces.BlogRepository) interfaces.AIContentSuggestionUsecase {
	return &AISuggestionUsecase{
		Model:          client.GenerativeModel("gemini-1.5-flash"),
		BlogRepository: blogRepository,
	}
}

func (uc *AISuggestionUsecase) SuggestContent(AI_query string) ([]string, *models.ErrorResponse) {
	ctx := context.Background()

	uc.Model.SystemInstruction = genai.NewUserContent(genai.Text("Generate a new blog post title, around 500 words long content and list of tags. Sepatate each of these with the '|' character. And separate each of the members of the tags with the character '&'. Generate everything continously without new line characters inside of the content section"))

	resp, err := uc.Model.GenerateContent(ctx, genai.Text(AI_query))
	if err != nil {
		return nil, models.InternalServerError("failed to connect to gemini")
	}

	// get the first response from an array of responses
	response := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts)
	response = strings.Trim(response, "&{}[]")
	response = strings.TrimSpace(response)

	response_array := strings.Split(response, "|") // [0]title, [1]content, [2]tags

	return response_array, nil
}

func (uc *AISuggestionUsecase) ImproveBlogContent(blogID string) ([]string, *models.ErrorResponse) {
	ctx := context.Background()

	blog, e := uc.BlogRepository.GetBlog(ctx, blogID)
	if e != nil {
		return nil, e
	}

	// generate content improvement queries
	uc.Model.SystemInstruction = genai.NewUserContent(genai.Text("Generate an improved blog post with the following format: a title, a single paragraph of around 500 words, and a list of tags, in that order. Separate the title, content, and tags with the '|' character. In the tags section, separate individual tags with the '&' character. Ensure the entire output is continuous, with no new lines."))
	blogInfo := fmt.Sprintf("Blog title: %s\nBlog content: %s\nBlog tags: %s", blog.Title, blog.Content, strings.Join(blog.Tags, ", "))

	resp, err := uc.Model.GenerateContent(ctx, genai.Text(blogInfo))
	if err != nil {
		return nil, models.InternalServerError("failed to connect to gemini")
	}

	// get the first response from an array of responses
	response := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts)
	response = strings.Trim(response, "&{}[]")
	response = strings.TrimSpace(response)

	response_array := strings.Split(response, "|") // [0]title, [1]content, [2]tags

	return response_array, nil
}

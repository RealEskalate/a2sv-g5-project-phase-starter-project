package usecases

import (
	// "meleket/domain"
	"meleket/infrastructure"
	"strings"
)

type AIUsecase struct {
	aiService infrastructure.AIService
}

func NewAIUsecase(aiService infrastructure.AIService) *AIUsecase {
	return &AIUsecase{
		aiService: aiService,
	}
}

func (u *AIUsecase) GenerateBlogContent(title string, tags []string) (string, error) {
	// Create the prompt based on title and tags
	prompt := "Title: " + title + "\nTags: " + tagsToString(tags) + "\nGenerate a blog post."
	return u.aiService.GenerateAIContent(prompt)
}

func tagsToString(tags []string) string {
	return "[" + strings.Join(tags, ", ") + "]"
}

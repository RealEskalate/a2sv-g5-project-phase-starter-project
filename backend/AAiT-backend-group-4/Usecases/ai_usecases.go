package usecases

import (
	domain "aait-backend-group4/Domain"
	"context"
	"time"
)

type aiUsecase struct {
	aiRepository domain.AiRepository
	timeout      time.Duration
}

func NewAiRepository(timeout time.Duration, aiRepository domain.AiRepository) domain.AiUsecase {
	return &aiUsecase{
		aiRepository: aiRepository,
		timeout:      timeout,
	}
}

func (au *aiUsecase) GenerateTextWithTags(c context.Context, tags []domain.Tag) (string, error) {
	prompt := "Generate a blog post using the following tags: "
	for i, tag := range tags {
		if i > 0 {
			prompt += ", "
		}
		prompt += string(tag)
	}
	prompt += ". The blog should be engaging, informative, and well-structured."

	blogContent, err := au.aiRepository.GenerateText(c, prompt)
	if err != nil {
		return "", err
	}

	return blogContent, nil
}

func (au *aiUsecase) GenerateTextWithPrompt(c context.Context, prompt string) (string, error) {
	blogContent, err := au.aiRepository.GenerateText(c, prompt)
	if err != nil {
		return "", err
	}

	return blogContent, nil
}

func (au *aiUsecase) GenerateSuggestions(c context.Context, textContent string) (string, error) {
	prompt := "Provide suggestions to improve the following blog content: "
	prompt += textContent
	prompt += ". The suggestions should be constructive, specific, and aimed at enhancing the quality and readability of the blog."

	suggestions, err := au.aiRepository.GenerateSuggestions(c, prompt)
	if err != nil {
		return "", err
	}

	return suggestions, nil
}

func (au *aiUsecase) Chat(c context.Context, textContent string) (string, error) {
	return "", nil
}

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

// NewAiUsecase creates a new instance of AiUsecase with the given timeout and aiRepository.
// It returns a pointer to the created AiUsecase.
func NewAiUsecase(timeout time.Duration, aiRepository domain.AiRepository) domain.AiUsecase {
	return &aiUsecase{
		aiRepository: aiRepository,
		timeout:      timeout,
	}
}

// GenerateTextWithTags generates a blog post using the provided tags.
// It takes a context.Context and a slice of domain.Tag as input parameters.
// The function returns a string containing the generated blog content and an error, if any.
// The function constructs a prompt by concatenating the tags and generates the blog content using the AI repository.
// The generated blog should be engaging, informative, and well-structured.
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

// GenerateTextWithPrompt generates text with the given prompt.
// It calls the GenerateText method of the aiRepository to generate the blog content.
// If an error occurs during the generation process, it returns an empty string and the error.
// Otherwise, it returns the generated blog content.
func (au *aiUsecase) GenerateTextWithPrompt(c context.Context, prompt string) (string, error) {
	blogContent, err := au.aiRepository.GenerateText(c, prompt)
	if err != nil {
		return "", err
	}

	return blogContent, nil
}

// GenerateSuggestions generates suggestions to improve the given blog content.
// The suggestions should be constructive, specific, and aimed at enhancing the quality and readability of the blog.
// It takes a context and the text content of the blog as input parameters.
// It returns the generated suggestions as a string and an error if any occurred.
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

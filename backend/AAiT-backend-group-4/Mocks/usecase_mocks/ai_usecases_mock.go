package mocks

import (
    "context"
    "errors"
    domain "aait-backend-group4/Domain"
)

// MockAiUsecase is a mock implementation of the AiUsecase interface.
type MockAiUsecase struct{}

func (m *MockAiUsecase) GenerateTextWithTags(c context.Context, tags []domain.Tag) (string, error) {
    if len(tags) == 0 {
        return "", errors.New("no tags provided")
    }
    return "Mocked blog content", nil
}

func (m *MockAiUsecase) GenerateTextWithPrompt(c context.Context, prompt string) (string, error) {
    if prompt == "" {
        return "", errors.New("empty prompt")
    }
    return "Mocked text content", nil
}

func (m *MockAiUsecase) GenerateSuggestions(c context.Context, textContent string) (string, error) {
    if textContent == "" {
        return "", errors.New("empty text content")
    }
    return "Mocked suggestions", nil
}

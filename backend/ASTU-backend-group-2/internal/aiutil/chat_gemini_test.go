package gemini

import (
	"testing"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
)

func TestGenerateContentFromGemini(t *testing.T) {
	title := "Test Title"
	description := "Test Description"
	env := bootstrap.Env{
		GeminiAPIKey:    "AIzaSyAPhVnC6ERQY7xILuWDEMfW0SKBIXaH_48",
		GeminiWordCount: "300",
	}
	ai := NewAIUtil(&env)

	content, err := ai.GenerateContentFromGemini(title, description, env)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if content == "" {
		t.Error("expected content to be non-empty")
	}

	t.Logf("Generated content: %s", content)

}

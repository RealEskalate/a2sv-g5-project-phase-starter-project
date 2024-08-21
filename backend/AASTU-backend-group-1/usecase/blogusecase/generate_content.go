package blogusecase

import "blogs/config"

func (b *BlogUsecase) GenerateAiContent(prompt string) (string, error) {
	result, err := config.GenerateAIContent(prompt)
	if err != nil {
		return "", err
	}

	return result, nil
}

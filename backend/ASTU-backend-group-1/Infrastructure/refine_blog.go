package infrastructure

import (
	"fmt"
)

func Refine(content string) (string, error) {
	prompt := fmt.Sprintf(`Refine this blog: `, content)
	refinedContent, err := SendPrompt(prompt)
	if err != nil {
		return "", err
	}
	return refinedContent, nil
}

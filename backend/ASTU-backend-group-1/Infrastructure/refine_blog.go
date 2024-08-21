package infrastructure

import (
	"astu-backend-g1/config"
	"context"
	"fmt"
)

func Refine(content string) (string, error) {
	prompt := fmt.Sprintf(`Please refine the following content to make it more engaging, clear, and concise. Focus on improving the flow, enhancing readability, and ensuring that the main points are emphasized effectively. Feel free to rephrase sentences, restructure paragraphs, and add any necessary transitions. The tone should remain professional yet approachable. And make sure that you don add any title please and no comments.: %v`, content)
	config, err := config.LoadConfig()
	if err != nil {
		return "", err
	}
	refinedContent, err := SendPrompt(prompt, config.Gemini.ApiKey, config.Gemini.Model, context.Background())
	if err != nil {
		return "", err
	}
	return refinedContent, nil
}

func Validate(content string) (string, error) {
	prompt := fmt.Sprintf(`
		You are an AI content moderator responsible for ensuring that blog content adheres to community guidelines. Please analyze the following text and determine if it complies with the following guidelines:

		Legal Compliance:

			No references to illegal activities, substances, or behaviors.
			No promotion or encouragement of illegal actions.
		Respect and Civility:

			No abusive, threatening, or harmful language.
			No hate speech, discrimination, or derogatory remarks targeting any individual or group based on race, ethnicity, nationality, religion, gender, sexual orientation, disability, or any other characteristic.
		Non-Violence:

			No content that promotes or glorifies violence, self-harm, or dangerous behaviors.
			No incitement to violence or harassment.
		Intellectual Property:

			No plagiarism or unauthorized use of copyrighted material.
		Proper attribution for quotes, images, or content that is not original.
		Privacy:

			No sharing of private or confidential information about individuals without their consent.
			No doxxing or revealing personal information that could lead to harm.
		Quality and Relevance:

		Content should be informative, relevant, and contribute positively to the community.
			am, misleading information, or clickbait.
		Age Appropriateness:

		Content should be suitable for a general audience, including minors.
			No explicit, sexual, or highly offensive material.
			
		Message to Validate\n: "%v" Based on these criteria, respond with:

		"true" if the content passes all the guidelines, followed by "valid content".

		"false" if the content fails any guideline, followed by a brief explanation of which guideline(s) the content violates and why.`, content)
	config, err := config.LoadConfig()
	if err != nil {
		return "", err
	}
	refinedContent, err := SendPrompt(prompt, config.Gemini.ApiKey, config.Gemini.Model, context.Background())
	if err != nil {
		return "", err
	}
	return refinedContent, nil
}

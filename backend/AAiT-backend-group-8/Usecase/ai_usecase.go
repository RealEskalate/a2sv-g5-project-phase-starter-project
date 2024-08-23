package usecase

import (
	"AAiT-backend-group-8/Domain"
	interfaces "AAiT-backend-group-8/Interfaces"
	"context"
	"fmt"
	"log"
	"strings"
)

type AiBlogUsecase struct {
	aiService interfaces.IAiService
}

func NewAiBlogUsecase(aiService interfaces.IAiService) interfaces.IAiUsecase{
	return &AiBlogUsecase{aiService: aiService}
}
func (u *AiBlogUsecase) GenerateBlogContent(userInput string) (Domain.BlogResponse, error) {
    resp, err := u.aiService.GenerateContent(context.Background(), userInput)
    if err != nil {
        return Domain.BlogResponse{}, err
    }

    // Log the raw AI response for debugging
    log.Printf("Raw AI response: %s", resp)

    // Initialize fields
    var title, body string
    var tags []string

    // Split the response into lines
    lines := strings.Split(resp, "\n")

    // Ensure we have at least a title
    if len(lines) > 0 {
        title = strings.TrimPrefix(strings.TrimSpace(lines[0]), "## ")
    }

    // Combine all lines after the title
    combinedLines := strings.Join(lines[1:], "\n")

    // Look for the "Tags:" label to separate body from tags
    tagsIndex := strings.Index(combinedLines, "Tags:")

    if tagsIndex != -1 {
        // Split the combined lines into body and tags based on the "Tags:" label
        body = strings.TrimSpace(combinedLines[:tagsIndex])
        tagsLine := strings.TrimSpace(combinedLines[tagsIndex+len("Tags:"):])
        tags = strings.Split(tagsLine, ", ")
        // Trim extra spaces from tags
        for i, tag := range tags {
            tags[i] = strings.TrimSpace(tag)
        }
    } else {
        // If no "Tags:" label, treat the entire combined lines as body
        body = combinedLines
    }

    // Return the processed blog response
    return Domain.BlogResponse{
        Title: title,
        Body:  body,
        Tags:  tags,
    }, nil
}








func (u *AiBlogUsecase) SuggestImprovements(title, body string, tags []string) (Domain.SuggestionBlogResponse, error) {
	tagsString := strings.Join(tags, ", ")
	resp, err := u.aiService.SuggestImprovements(context.Background(), title, body, tagsString)
	if err != nil {
		log.Printf("Error from AI service: %v", err)
		return Domain.SuggestionBlogResponse{}, fmt.Errorf("failed to suggest improvements")
	}

	// Log the raw AI response
	log.Printf("Raw AI response: %s", resp)

	// Process response
	suggestion := Domain.SuggestionBlogResponse{}
	lines := strings.Split(resp, "\n")

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])     // Remove any leading/trailing whitespace
		log.Printf("Processing line: %s", line) // Log each processed line

		if strings.HasPrefix(line, "## Comment:") {
			// Skip any blank lines until the actual comment is found
			for j := i + 1; j < len(lines); j++ {
				nextLine := strings.TrimSpace(lines[j])
				if nextLine != "" {
					suggestion.Comment = nextLine
					i = j // Update i to skip the processed line
					break
				}
			}
		} else if strings.HasPrefix(line, "**Title:**") {
			suggestion.Title = strings.TrimPrefix(line, "**Title:** ")
		} else if strings.HasPrefix(line, "**Body:**") {
			suggestion.Body = strings.TrimPrefix(line, "**Body:** ")
			// Continue collecting body content until the next tag
			for j := i + 1; j < len(lines); j++ {
				nextLine := strings.TrimSpace(lines[j])
				if strings.HasPrefix(nextLine, "**Tags:**") {
					break
				}
				suggestion.Body += " " + nextLine
				i = j // Update i to skip the processed lines
			}
		} else if strings.HasPrefix(line, "**Tags:**") {
			tagsLine := strings.TrimPrefix(line, "**Tags:** ")
			suggestion.Tags = strings.Split(tagsLine, ", ")
		}
	}

	// Log the parsed suggestion
	log.Printf("Parsed suggestion: %+v", suggestion)

	// Check if fields were populated; if not, handle it appropriately
	if suggestion.Comment == "" || suggestion.Title == "" || suggestion.Body == "" || len(suggestion.Tags) == 0 {
		log.Printf("Failed to parse AI response correctly: %+v", suggestion)
		return Domain.SuggestionBlogResponse{}, fmt.Errorf("failed to parse AI response correctly")
	}

	return suggestion, nil
}

package blogcmd

import (
	"github.com/google/uuid"
)

// AddCommand represents the data required to add a new blog.
// Fields:
// - Title: The title of the blog.
// - Content: The content of the blog.
// - Tags: A list of tags associated with the blog.
// - UserID: The UUID of the user who owns the blog.
type AddCommand struct {
	title   string
	content string
	tags    []string
	userID  uuid.UUID
}

// NewAddCommand creates a new AddCommand instance with the specified details.
func NewAddCommand(title, content string, tags []string, userID uuid.UUID) *AddCommand {
	return &AddCommand{
		title:   title,
		content: content,
		tags:    tags,
		userID:  userID,
	}
}

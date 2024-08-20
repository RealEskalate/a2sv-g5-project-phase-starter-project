package blogcmd

import (
	"github.com/google/uuid"
)

// Command represents the data needed to update an existing blog.
type UpdateCommand struct {
	id      uuid.UUID
	title   string
	content string
	tags    []string
}

// NewCommand creates a new Command instance with the provided blog details.
func NewUpdateCommand(id uuid.UUID, title, content string, tags []string) *UpdateCommand {
	return &UpdateCommand{
		id:      id,
		title:   title,
		content: content,
		tags:    tags,
	}
}

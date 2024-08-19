package updatecmd

import (
	"github.com/google/uuid"
)

// Command represents the data needed to update an existing blog.
type Command struct {
	id      uuid.UUID
	title   string
	content string
	tags    []string
}

// NewCommand creates a new Command instance with the provided blog details.
func NewCommand(id uuid.UUID, title, content string, tags []string) *Command {
	return &Command{
		id:      id,
		title:   title,
		content: content,
		tags:    tags,
	}
}

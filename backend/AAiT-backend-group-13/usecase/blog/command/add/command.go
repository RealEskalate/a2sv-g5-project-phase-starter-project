package addcmd

import (
	"github.com/google/uuid"
)

// Command represents the data required to add a new blog.
// Fields:
// - title: The title of the blog.
// - content: A Content that the blog has.
// - tags: different tags that the blog has.
// - userId: The userId of the user that this blog belongs to.
type Command struct {
	title   string
	content string
	tags    []string
	userId  uuid.UUID
}

// NewCommand creates a new Command instance with the specified details.
func NewCommand(title, content string, tags []string, userId uuid.UUID) *Command {
	return &Command{
		title:   title,
		content: content,
		tags:    tags,
		userId:  userId,
	}
}

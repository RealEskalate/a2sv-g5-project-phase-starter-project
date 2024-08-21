package addcom

import "github.com/google/uuid"

type Command struct {
	content string
	userId  uuid.UUID
	blogId  uuid.UUID
}

// NewCommand creates a new Command instance with the specified details.
func NewCommand(content string, blogId, userId uuid.UUID) *Command {
	return &Command{
		content: content,
		userId:  userId,
		blogId:  blogId,
	}
}

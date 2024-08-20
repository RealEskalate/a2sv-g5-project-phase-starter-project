package commentcmd

import "github.com/google/uuid"

// AddCommand represents a command to add a comment to a blog.
type AddCommand struct {
	content string
	userID  uuid.UUID
	blogID  uuid.UUID
}

// NewAddCommand creates a new AddCommand instance with the specified details.
func NewAddCommand(content string, blogID, userID uuid.UUID) *AddCommand {
	return &AddCommand{
		content: content,
		userID:  userID,
		blogID:  blogID,
	}
}

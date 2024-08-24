package reactioncmd

import "github.com/google/uuid"

type DeleteCommand struct {
	UserId uuid.UUID
	BlogId uuid.UUID
}

// NewDeleteCommand creates a new Command instance with the specified details.
func NewDeleteCommand(blogId, userId uuid.UUID) *DeleteCommand {
	return &DeleteCommand{
		UserId: userId,
		BlogId: blogId,
	}
}

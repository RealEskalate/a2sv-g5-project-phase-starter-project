package reactioncmd

import "github.com/google/uuid"

type UpdateCommand struct {
	IsLike bool
	UserId uuid.UUID
	BlogId uuid.UUID
}

// NewCommand creates a new Command instance with the specified details.
func NewUpdateCommand(isLike bool, blogId, userId uuid.UUID) *UpdateCommand {
	return &UpdateCommand{
		IsLike: isLike,
		UserId: userId,
		BlogId: blogId,
	}
}

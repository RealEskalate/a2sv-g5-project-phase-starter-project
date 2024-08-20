package updatereaction

import "github.com/google/uuid"

type Command struct {
	IsLike bool
	UserId uuid.UUID
	BlogId uuid.UUID
}

// NewCommand creates a new Command instance with the specified details.
func NewCommand(isLike bool, blogId, userId uuid.UUID) *Command {
	return &Command{
		IsLike: isLike,
		UserId: userId,
		BlogId: blogId,
	}
}

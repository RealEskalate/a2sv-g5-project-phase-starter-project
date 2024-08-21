package usercmd

import "github.com/google/uuid"

// PromoteCommand represents a command to promote a user with necessary data.
type PromoteCommand struct {
	username   string
	toAdmin    bool
	promoterID uuid.UUID
}

// NewPromoteCommand creates a new PromoteCommand instance with the given username, admin status, and promoter ID.
func NewPromoteCommand(username string, toAdmin bool, promoterId uuid.UUID) *PromoteCommand {
	return &PromoteCommand{
		username:   username,
		toAdmin:    toAdmin,
		promoterID: promoterId,
	}
}

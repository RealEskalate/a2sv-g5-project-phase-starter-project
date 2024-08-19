package promotcmd

import "github.com/google/uuid"

// Command represents a promote command with necessary data.
type Command struct {
	Username   string
	ToAdmin    bool
	PromoterID uuid.UUID
}

// NewCommand creates a new Command instance with the given username and promoter ID.
func NewCommand(username string, toAdmin bool, promoterId uuid.UUID) *Command {
	return &Command{
		Username:   username,
		ToAdmin:    toAdmin,
		PromoterID: promoterId,
	}
}

package resetpassword

import "github.com/google/uuid"

// Command represents the command to reset a user's password.
// It contains the user ID and the new password.
type Command struct {
	// Id is the unique identifier of the user whose password is being reset.
	Id uuid.UUID

	// NewPassword is the new password for the user.
	NewPassword string
}

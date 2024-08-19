package resetpassword

import "github.com/google/uuid"

// Command represents the command to reset a user's password.
// It contains the user ID, a reset token, and the new password.
type Command struct {
	// Id is the unique identifier of the user whose password is being reset.
	Id uuid.UUID

	// token is the secure token required for validating the password reset request.
	token string

	// NewPassword is the new password for the user.
	NewPassword string
}

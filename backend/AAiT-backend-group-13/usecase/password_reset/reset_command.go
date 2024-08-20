package passwordreset

// Command represents the command to reset a user's password.
// It contains the user ID, a reset token, and the new password.
type ResetCommand struct {
	token       string
	NewPassword string
}

func NewResetCommand(token string, newPassword string) *ResetCommand {
	return &ResetCommand{
		token:       token,
		NewPassword: newPassword,
	}
}

package passwordreset

import ()

// ValidateCodeCommand represents the command to validate a reset code.
// It includes the reset code and the user ID associated with it.
type ValidateCodeCommand struct {
	code  int
	email string
}

// NewValidateCodeCommand creates a new instance of ValidateCodeCommand.
func NewValidateCodeCommand(code int, email string) *ValidateCodeCommand {
	return &ValidateCodeCommand{
		code:  code,
		email: email,
	}
}

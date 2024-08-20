package passwordreset

import ()

// ValidateCodeCommand represents the command to validate a reset code.
// It includes the reset code and the user ID associated with it.
type ValidateCodeCommand struct {
	code  int64
	email string
}

// NewValidateCodeCommand creates a new instance of ValidateCodeCommand.
func NewValidateCodeCommand(code int64, email string) *ValidateCodeCommand {
	return &ValidateCodeCommand{
		code:  code,
		email: email,
	}
}

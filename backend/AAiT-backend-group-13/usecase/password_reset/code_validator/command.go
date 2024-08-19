package resetcodevalidate

import "github.com/google/uuid"

// Command represents the command to validate a reset code.
// It contains the reset code and the user ID associated with it.
type Command struct {
	// Code is the reset code that needs to be validated.
	Code int64

	// Id is the unique identifier of the user associated with the reset code.
	Id uuid.UUID
}

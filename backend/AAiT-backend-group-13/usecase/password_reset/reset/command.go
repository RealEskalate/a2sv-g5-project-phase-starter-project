package resetpassword

import "github.com/google/uuid"

type Command struct {
	Id          uuid.UUID
	NewPassword string
}

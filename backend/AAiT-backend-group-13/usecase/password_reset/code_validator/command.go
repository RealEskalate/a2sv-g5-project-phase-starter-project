package resetcodevalidate

import "github.com/google/uuid"

type Command struct {
	Code int64
	Id   uuid.UUID
}

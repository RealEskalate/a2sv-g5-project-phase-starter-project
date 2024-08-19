package chat

import "errors"

var ErrChatNotFound error = errors.New("chat not found")
var ErrInvalidID error = errors.New("invalid ID")

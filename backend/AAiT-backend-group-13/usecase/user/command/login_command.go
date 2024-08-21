package usercmd

// LoginCommand represents the command to log in a user.
type LoginCommand struct {
	username string
	password string
}

// NewLoginCommand creates a new instance of LoginCommand.
func NewLoginCommand(username, password string) *LoginCommand {
	return &LoginCommand{
		username: username,
		password: password,
	}
}


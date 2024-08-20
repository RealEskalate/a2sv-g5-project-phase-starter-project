package usercmd

// SignUpCommand represents the command to register a new user with necessary details.
type SignUpCommand struct {
	firstName string
	lastName  string
	username  string
	email     string
	password  string
}

// NewSignUpCommand creates a new SignUpCommand instance with the provided user details.
func NewSignUpCommand(username, firstName, lastName, email, password string) *SignUpCommand {
	return &SignUpCommand{
		firstName: firstName,
		lastName:  lastName,
		username:  username,
		email:     email,
		password:  password,
	}
}


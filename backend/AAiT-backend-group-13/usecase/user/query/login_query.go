package userqry

// LoginQuery represents the command to log in a user.
type LoginQuery struct {
	username string
	password string
}

// NewLoginQuery creates a new instance of LoginCommand.
func NewLoginQuery(username, password string) *LoginQuery {
	return &LoginQuery{
		username: username,
		password: password,
	}
}

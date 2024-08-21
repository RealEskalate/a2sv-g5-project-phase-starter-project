package logincommand


type LoginCommand struct {
	username string
	password string
}

func NewLoginCommand(username string, password string) LoginCommand {
	return LoginCommand{
		username: username,
		password: password,
	}
}
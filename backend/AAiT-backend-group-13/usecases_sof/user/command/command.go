package usercommand

type SignUpCommand struct {
	firstName string
	lastName  string
	username  string
	email     string
	password  string
}

func NewSignUpCommand(username string, firstName string, lastName string, email string, password string) SignUpCommand {
	return SignUpCommand{
		firstName: firstName,
		username:  username,
		lastName:  lastName,
		email:     email,
		password:  password,
	}
}

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

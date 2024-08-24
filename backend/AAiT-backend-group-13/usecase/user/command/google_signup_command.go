package usercmd

type GoogleSignupCommand struct {
	firstName  string
	lastName   string
	email      string
	isVerified bool
}

func NewGoogleSignupCommand(firstName string, lastname string, email string, isVerified bool) *GoogleSignupCommand {
	return &GoogleSignupCommand{
		firstName:  firstName,
		lastName:   lastname,
		email:      email,
		isVerified: isVerified,
	}
}

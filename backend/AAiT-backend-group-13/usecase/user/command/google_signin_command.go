package usercmd

type GoogleSigninCommand struct {
	email      string
	isVerified bool
}

func NewGoogleSigninCommand(email string, isVerified bool) *GoogleSigninCommand {
	return &GoogleSigninCommand{
		email:      email,
		isVerified: isVerified,
	}
}

package usercmd

// UpdateProfileCommand represents the command to update a user's profile with necessary details.
type UpdateProfileCommand struct {
	userid string 
	FirstName string
	LastName  string
	Email     string
	Password  string
	Username  string
}

// NewUpdateProfileCommand creates a new UpdateProfileCommand instance with the provided user details.
func NewUpdateProfileCommand(username, firstName, lastName, email, password string, userid string) *UpdateProfileCommand {
	return &UpdateProfileCommand{
		userid:     userid,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
	}
}


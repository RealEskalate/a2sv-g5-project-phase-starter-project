package tokenutil

func GetUserFromToken(token string) (string, error) {
	// get the user from the jwt token
	username, err := ParseToken(token)

	if err != nil {
		return "", err
	}

	// get the user from the database
	// TODO: implement GetUser
	user, err := "GetUser(username)"+username, nil

	return user, nil
}

func ParseToken(token string) (string, error) {
	// parse the token and get the user
	return "user", nil
}

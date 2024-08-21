package passwordservice

import "golang.org/x/crypto/bcrypt"

type PasswordS struct{}

func (*PasswordS) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
func (*PasswordS) ComparePassword(DBpassword string, InputPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(DBpassword), []byte(InputPassword))

	if err != nil {
		return false, err
	}

	return true, nil
}
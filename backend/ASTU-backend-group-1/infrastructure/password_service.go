package infrastructure

import "golang.org/x/crypto/bcrypt"

func PasswordHasher(Password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "Internal server error", err
	}
	hashedPassword := string(hashed)
	return hashedPassword, nil
}

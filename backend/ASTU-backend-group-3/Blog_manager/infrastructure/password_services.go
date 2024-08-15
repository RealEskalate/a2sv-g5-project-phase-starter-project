package utils

import "golang.org/x/crypto/bcrypt"

// ComparePasswords checks if the provided password matches the hashed password.
func ComparePasswords(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

// HashPassword takes a plaintext password and returns its hashed version.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

package infrastructure
import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt algorithm.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword compares the given password with the hashed password.
// Returns true if the passwords match, false otherwise.
func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
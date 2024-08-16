package infrastructure

import "golang.org/x/crypto/bcrypt"

//compares the inputted password from the existing hash
func PasswordComparator(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil
}

//hashes the password with a SHA-256 encryption
func PasswordHasher(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

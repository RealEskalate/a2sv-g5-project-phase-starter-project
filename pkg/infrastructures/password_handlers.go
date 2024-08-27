package infrastructures

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)) == nil
}

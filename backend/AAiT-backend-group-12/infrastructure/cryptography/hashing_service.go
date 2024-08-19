package cryptography

import (
	"blog_api/domain"

	"golang.org/x/crypto/bcrypt"
)

/*
Accepts a string and hashes it using the bcrypt
*/
func HashString(password string) (string, domain.CodedError) {
	hashedPwd, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		return "", domain.NewError("Internal server error: "+hashErr.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return string(hashedPwd), nil
}

/*
Accepts a hashed string and a plaintext string and validates the plaintext string
*/
func ValidateHashedString(hashedString string, plaintextString string) domain.CodedError {
	compErr := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(plaintextString))
	if compErr != nil {
		return domain.NewError("Invalid signature", domain.ERR_UNAUTHORIZED)
	}

	return nil
}

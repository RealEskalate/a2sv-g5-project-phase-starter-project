package result

import (
	"github.com/google/uuid"
)

type SignUpResult struct {
	ID        uuid.UUID
	Username  string
	FirstName string
	LastName  string
	IsAdmin   bool
}

func NewSignUpResult(ID uuid.UUID, Username string, firstName string, lastName string, IsAdmin bool) SignUpResult {
	return SignUpResult{
		ID:        ID,
		Username:  Username,
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   IsAdmin,
	}
}

type LoginInResult struct {
	Token        string
	RefreshToken string
}

func NewLoginInResult(token string, refreshToken string) LoginInResult {
	return LoginInResult{
		Token:        token,
		RefreshToken: refreshToken,
	}
}

type ValidateEmailResult struct {
	Token        string
	Refreshtoken string
}

func NewValidateEmailResult(token string, refreshtoken string) ValidateEmailResult {
	return ValidateEmailResult{
		Token:        token,
		Refreshtoken: refreshtoken,
	}
}

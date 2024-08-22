package result

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
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
	User         *models.User
	Token        string
	RefreshToken string
}

func NewLoginInResult(user *models.User, token string, refreshToken string) *LoginInResult {
	return &LoginInResult{
		User:         user,
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


type UpdateProfileResult struct {
	message string
}



func NewUpdateProfileResult(message string) UpdateProfileResult {
	return UpdateProfileResult{
		message: message,
	}
}
package usercontroller

import (
	"github.com/google/uuid"
	"github.com/group13/blog/usecase/user/result"
)

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID        uuid.UUID
	Username  string
	FirstName string
	LastName  string
	IsAdmin   bool
}

func NewLoginResponse(r *result.LoginInResult) *LoginResponse {
	return &LoginResponse{
		ID:        r.User.ID(),
		Username:  r.User.Username(),
		FirstName: r.User.FirstName(),
		LastName:  r.User.LastName(),
		IsAdmin:   r.User.IsAdmin(),
	}

}

type SignUpDto struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
}

type ValidateCodeDto struct {
	Code  int    `json:"code" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type ResetPasswordDto struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type ForgotPasswordDto struct {
	Email string `json:"email" binding:"required"`
}

type UpdateProfileDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}
type UserInfo struct {
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
}

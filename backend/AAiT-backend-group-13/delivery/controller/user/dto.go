package usercontroller

import "github.com/google/uuid"

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpDto struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
}

type ValidateCodeDto struct {
	Code  int64  `json:"code" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type ResetPasswordDto struct {
	Id          uuid.UUID `json:"id" binding:"required"`
	Token       string    `json:"token" binding:"required"`
	NewPassword string    `json:"newPassword" binding:"required"`
}

type ForgotPasswordDto struct {
	Email string `json:"email" binding:"required"`
}

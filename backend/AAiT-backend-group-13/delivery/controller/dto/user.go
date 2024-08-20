package dto

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


type ResetPasswordDto struct {
	Code    int64  `json:"code" binding:"required"`
	Id	  uuid.UUID `json:"id" binding:"required"`
}


type ForgotPasswordDto struct {
	Id uuid.UUID `json:"id" binding:"required"`
	Token string `json:"token" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}



package domain

import (
	"context"
)
type ForgetPWRequest struct{
	Email  string	`json:"email" validate:"email,required"`
}

type ResetPWRequest struct{
	Email    string	`form:"email" binding:"required,email"`
	Password string	`json:"password" validate:"required,min=8,max=32"`
}

type ForgetPWUsecase interface {
	ForgetPW(c context.Context, email string, server_address string) error
	ResetPW(c context.Context, request ResetPWRequest) error
	VerifyForgetPWRequest(c context.Context, username string, recover_token string) error
	GenerateRecoveryLink(server_address string, username string, recoveryToken string) (recoveryLink string)
}
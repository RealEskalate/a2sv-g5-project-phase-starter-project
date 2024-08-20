package service

import "backend-starter-project/domain/interfaces"

type passwordResetService struct {
}

func NewPasswordResetService() interfaces.PasswordResetService{
	return &passwordResetService{}
}

func (pass_service *passwordResetService) RequestPasswordReset(email string) error{
	return nil
}

func (pass_service *passwordResetService) ResetPassword(token,newPass string)error{
	return nil
}
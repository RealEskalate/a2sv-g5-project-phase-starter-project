package interfaces

import "aait.backend.g10/domain"

type IEmailService interface {
	SendResetEmail(email string, resetToken string) *domain.CustomError
	SendActivationEmail(email string, activationToken string) *domain.CustomError
}

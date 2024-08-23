package interfaces

import "aait.backend.g10/domain"

type IEmailService interface {
	SendResetEmail(email string, resetToken string) *domain.CustomError
}
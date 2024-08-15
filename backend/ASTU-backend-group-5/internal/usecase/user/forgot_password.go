package user

import (
	"blogApp/internal/config"
	"blogApp/pkg/email"
	"log"
)

// Generate a random token
func generateToken() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return chars
}

func (u *UserUsecase) RequestPasswordResetUsecase(userEmail string) error {
	var emailSender email.EmailSender

	Config, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	emailProvider := Config.EMAIL_PROVIDER

	switch emailProvider {
	case "simple":
		emailSender = email.NewSimpleEmailSender(
			Config.SMTP_HOST,
			Config.SMTP_PORT,
			Config.EMAIL_SENDER_EMAIL,
			Config.EMAIL_SENDER_PASSWORD,
		)
	default:
		emailSender = email.NewSimpleEmailSender(
			Config.SMTP_HOST,
			Config.SMTP_PORT,
			Config.EMAIL_SENDER_EMAIL,
			Config.EMAIL_SENDER_PASSWORD,
		)
	}

	resetToken := generateToken()

	err = emailSender.SendPasswordResetEmail(userEmail, resetToken)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
	return err
}

func (u *UserUsecase) ResetPassword(resetToken string, newPassword string, email string) error {
	return nil
}

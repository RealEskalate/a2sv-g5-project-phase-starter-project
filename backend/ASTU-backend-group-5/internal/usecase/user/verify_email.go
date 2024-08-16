package user

import (
	"blogApp/internal/config"
	"blogApp/internal/domain"
	"blogApp/pkg/email"
	"blogApp/pkg/jwt"
	"context"
	"errors"
	"log"
)

func TokenGenerator() string {
	return "token"
}

func (u *UserUsecase) RequestEmailVerification(user domain.User) error {
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

	go func() {
		err := emailSender.SendVerificationEmail(user.Email, TokenGenerator())
		if err != nil {
			log.Printf("Failed to send verification email: %v", err)
		}
	}()

	return nil
}

func (u *UserUsecase) VerifyEmail(token string, email string) error {
	claims, err := jwt.ValidateToken(token)
	if err != nil {
		return err
	}
	issuerEmail := claims.Email
	if issuerEmail != email {
		return errors.New("invalid token")
	}
	user, err := u.repo.FindUserByEmail(context.Background(), email)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	// user. = true
	err = u.repo.UpdateUser(context.Background(), user)
	if err != nil {
		return err
	}
	return nil

}

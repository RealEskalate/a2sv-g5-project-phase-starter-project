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

func (u *UserUsecase) RequestEmailVerification(user domain.User) error {
	var emailSender email.EmailSender

	Config, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	dbUSer, err := u.repo.FindUserByEmail(context.Background(), user.Email)

	if err != nil {
		return err
	}

	if dbUSer == nil {
		return errors.New("user with this email does not exist")
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
		token, err := jwt.GenerateJWT("email-verification", user.Email, "email-verification", "email-verification")
		if err != nil {
			log.Printf("Failed to generate token: %v", err)
		}
		err = emailSender.SendVerificationEmail(user.Email, token)
		if err != nil {
			log.Printf("Failed to send verification email: %v", err)
		}
	}()

	return nil
}

func (u *UserUsecase) VerifyEmail(token string) error {
	claims, err := jwt.ValidateToken(token)
	if err != nil {
		return err
	}
	issuerEmail := claims.Email
	user, err := u.repo.FindUserByEmail(context.Background(), issuerEmail)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user with this email does not exist")
	}
	// user. = true
	user.Verified = true
	err = u.repo.UpdateUser(context.Background(), user)
	if err != nil {
		return err
	}
	return nil

}

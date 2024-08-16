package user

import (
	"blogApp/internal/config"
	"blogApp/pkg/email"
	"blogApp/pkg/hash"
	"blogApp/pkg/jwt"
	"context"
	"errors"
	"fmt"
	"log"
)

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
	accessToken, err := jwt.GenerateJWT("password-reset", userEmail, "password-reset", "password-reset")
	if err != nil {
		return err
	}
	go func() {
		err := emailSender.SendPasswordResetEmail(userEmail, accessToken)
		if err != nil {
			log.Printf("Failed to send password reset email: %v", err)
		}
	}()

	return nil
}

func (u *UserUsecase) ResetPassword(resetToken string, newPassword string, email string) error {
	claims, err := jwt.ValidateToken(resetToken)
	if err != nil {
		return err
	}
	issuerEmail := claims.Email
	fmt.Println(issuerEmail, email)
	if issuerEmail != email {
		return errors.New("invalid token")
	}
	hashedPassword, err := hash.HashPassword(newPassword)
	if err != nil {
		return err
	}
	user, err := u.repo.FindUserByEmail(context.Background(), email)

	if user != nil && err == nil {
		user.Password = hashedPassword
		err = u.repo.UpdateUser(context.Background(), user)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}

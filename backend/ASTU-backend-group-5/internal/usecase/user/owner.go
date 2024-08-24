package user

import (
	"blogApp/internal/config"
	"blogApp/pkg/email"
	"context"
	"errors"
	"fmt"
	"log"
)

func (u *UserUsecase) PromoteToAdmin(UserId string) error {
	var emailSender email.EmailSender

	user, err := u.repo.FindUserById(context.Background(), UserId)
	fmt.Println(user)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	user.Role = "admin"
	err = u.repo.UpdateUser(context.Background(), user)
	if err != nil {
		return err
	}

	Config, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		return nil
	}
	emailProvider := Config.EMAIL_PROVIDER

	switch emailProvider {
	case "simple":
		emailSender = email.NewSimpleEmailSender(
			Config.SMTP_HOST,
			Config.EMAIL_PORT,
			Config.EMAIL_SENDER_EMAIL,
			Config.EMAIL_SENDER_PASSWORD,
		)
	default:
		emailSender = email.NewSimpleEmailSender(
			Config.SMTP_HOST,
			Config.EMAIL_PORT,
			Config.EMAIL_SENDER_EMAIL,
			Config.EMAIL_SENDER_PASSWORD,
		)
	}

	go func() {
		err := emailSender.SendPromotionToAdminEmail(user.Email)
		if err != nil {
			log.Printf("Failed to send password reset email: %v", err)
		}
	}()

	return nil
}

func (u *UserUsecase) DemoteFromAdmin(UserId string) error {
	var emailSender email.EmailSender
	user, err := u.repo.FindUserById(context.Background(), UserId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	user.Role = "user"
	err = u.repo.UpdateUser(context.Background(), user)
	if err != nil {
		return err
	}

	Config, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		return nil
	}

	emailProvider := Config.EMAIL_PROVIDER

	switch emailProvider {
	case "simple":
		emailSender = email.NewSimpleEmailSender(
			Config.SMTP_HOST,
			Config.EMAIL_PORT,
			Config.EMAIL_SENDER_EMAIL,
			Config.EMAIL_SENDER_PASSWORD,
		)
	default:
		emailSender = email.NewSimpleEmailSender(
			Config.SMTP_HOST,
			Config.EMAIL_PORT,
			Config.EMAIL_SENDER_EMAIL,
			Config.EMAIL_SENDER_PASSWORD,
		)
	}

	go func() {
		err := emailSender.SendDemotionFromAdminEmail(user.Email)
		if err != nil {
			log.Printf("Failed to send password reset email: %v", err)
		}
	}()

	return nil
}

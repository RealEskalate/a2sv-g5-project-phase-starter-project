package user

import (
	"blogApp/internal/config"
	"blogApp/pkg/email"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TokenGenerator() string {
	return "token"
}

type User struct {
	Email    string
	ID       primitive.ObjectID
	UserName string
}

type UserUsecase struct {
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}
func (u *UserUsecase) RequestEmailVerification(user User) error {
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
	return nil
}

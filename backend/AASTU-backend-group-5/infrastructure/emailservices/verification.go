package emailservices

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	gmail "gopkg.in/gomail.v2"
)


func SendVerificationEmail(to, subject, body string) error {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
	}

	username := os.Getenv("USER_EMAIL")
	password := os.Getenv("EMAIL_PASS")

	m := gmail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gmail.NewDialer("smtp.gmail.com", 587, username, password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}


func IsValidVerificationToken(strtoken string) (string , error) {
	var err error = godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
		return "",nil
	}
	var SecretKey = []byte(os.Getenv("SECRETKEY"))
	token , err := jwt.ParseWithClaims(strtoken, &domain.EmailUserClaims{} ,func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return "" , err
	}

	if !token.Valid {
		return "",errors.New("token not valid")
	}

	payload,ok := token.Claims.(*domain.EmailUserClaims); 
	if !ok {
		return "",errors.New("token payload not valid")
	}

	return payload.ID.Hex(),nil
}
package infrastructure

import (
	domain "aait-backend-group4/Domain"
	"bytes"
	"crypto/rand"
	"fmt"
	"text/template"

	"gopkg.in/gomail.v2"
)

type otpService struct{}

func NewOTPService() domain.OtpInfrastructure {
	return &otpService{}
}

func (os *otpService) CreateOTP(otp *domain.UserOTPRequest) (otpCode string, err error) {
	otpC, err := generateRandomOTP(6)
	if err != nil {
		return "", err
	}
	return otpC, nil
}

func (os *otpService) SendEmail(email string, subject string, key string, otp string) error {
	var b bytes.Buffer
	t, err := template.ParseFiles("../templates/email_confirmation.html")
	if err != nil {
		return fmt.Errorf("template file not found")
	}
	t.Execute(&b, struct {
		Subject string
		Otp     string
	}{Subject: subject,
		Otp: otp,
	})

	// Send with Go mail
	m := gomail.NewMessage()
	m.SetHeader("From", "AAiT Backend group 4")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", b.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "Example@gmail.com", key)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// generateRandomOTP generates a random numeric OTP of the specified length
func generateRandomOTP(length int) (string, error) {
	otp := make([]byte, length)
	_, err := rand.Read(otp)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		otp[i] = otp[i]%10 + '0' // Convert to a digit between '0' and '9'
	}

	return string(otp), nil
}

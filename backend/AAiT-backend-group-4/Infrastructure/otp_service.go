package infrastructure

import (
	domain "aait-backend-group4/Domain"
	"bytes"
	"crypto/rand"
	"text/template"

	"gopkg.in/gomail.v2"
)

type otpService struct{}

func NewOTPService() domain.OtpInfrastructure {
	return &otpService{}
}

// CreateOTP generates a new OTP (One-Time Password) for the given UserOTPRequest.
// It returns the generated OTP code and any error encountered during the process.
func (os *otpService) CreateOTP(otp *domain.UserOTPRequest) (otpCode string, err error) {
	otpC, err := generateRandomOTP(6)
	if err != nil {
		return "", err
	}
	return otpC, nil
}

// SendEmail sends an email to the specified recipient with the given subject, key, and OTP.
// It uses a template file to generate the email body and sends it using Go mail.
// The email is sent using the SMTP server at smtp.gmail.com on port 587.
// The sender's email address is set as "AAiT Backend group 4".
// The recipient's email address is set as the value of the 'email' parameter.
// The subject of the email is set as the value of the 'subject' parameter.
// The email body is generated using the 'email_confirmation.html' template file,
// with the 'Subject' and 'Otp' fields set to the values of the 'subject' and 'otp' parameters, respectively.
// The email body is sent as HTML content.
// The SMTP server is authenticated using the provided 'key' parameter, which is the password for the sender's email account.
// If an error occurs while sending the email, it is returned as an error.
// If the template file is not found, an error is returned with the message "template file not found".
func (os *otpService) SendEmail(email string, subject string, key string, otp string) error {
	var b bytes.Buffer
	t, err := template.ParseFiles("templates/email_confirmation.html")
	if err != nil {
		return err
	}
	t.Execute(&b, struct {
		Subject string
		Otp     string
	}{
		Subject: subject,
		Otp:     otp,
	})

	// Send with Go mail
	m := gomail.NewMessage()
	m.SetHeader("From", "solomonjohna21@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", b.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "solomonjohna21@gmail.com", key)

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

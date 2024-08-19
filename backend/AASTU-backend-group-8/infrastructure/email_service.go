package infrastructure

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func SendOTPEmail(email, otp string) error {
	// Set up the email
	m := gomail.NewMessage()
	m.SetHeader("From", "kalkidanamare11a@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is %s", otp))

	// Set up the SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "kalkidanamare11a@gmail.com", "jcwf vfzi njtd rayo")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

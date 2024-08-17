package infrastructure

import (
	"crypto/tls"
	"fmt"
	"meleket/domain"
	"meleket/utils"

	gomail "gopkg.in/mail.v2"
)

type EmailService struct{
	dialer *gomail.Dialer
}

//
func (s *EmailService) SendOTPEmail(user *domain.User) error {
	otp := utils.GenerateOTP(6)

	// Store OTP in the database
	if err := s.userRepo.StoreOTP(user.ID, otp); err != nil {
		return err
	}

	// Set up the email
	m := gomail.NewMessage()
	m.SetHeader("From", "kalkidanamare11a@gmail.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is %s", otp))

	// Set up the SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "kalidanamare11a@gmail.com", "axbs xtrk xuqm vvsa")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	return d.DialAndSend(m)
}

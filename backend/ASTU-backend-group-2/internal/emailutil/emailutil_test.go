package emailutil

import (
	"testing"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
)

func TestSendVerificationEmail(t *testing.T) {
	recipientEmail := "nahomderese@gmail.com"
	verificationToken := "abc123"
	env := &bootstrap.Env{
		SenderEmail:    "wendmnewsaleamlak@gmail.com",
		SenderPassword: "nhba zihx apjr zqpx",
		SmtpHost:       "smtp.gmail.com",
		SmtpPort:       "587",
	}

	err := SendVerificationEmail(recipientEmail, verificationToken, env)
	if err != nil {
		t.Errorf("Failed to send verification email: %v", err)
	}

}

func TestSendOtpVerificationEmail(t *testing.T) {
	recipientEmail := "saleamlakwendmnew55@gmail.com"
	otp := "abc123"
	env := &bootstrap.Env{
		SenderEmail:    "wendmnewsaleamlak@gmail.com",
		SenderPassword: "nhba zihx apjr zqpx",
		SmtpHost:       "smtp.gmail.com",
		SmtpPort:       "587",
	}

	err := SendOtpVerificationEmail(recipientEmail, otp, env)
	if err != nil {
		t.Errorf("Failed to send otp verification email: %v", err)
	}

}


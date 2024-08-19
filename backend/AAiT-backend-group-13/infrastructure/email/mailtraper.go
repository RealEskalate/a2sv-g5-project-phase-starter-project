package email

import (
	iemail "github.com/group13/blog/usecase/common/i_email"
	"gopkg.in/gomail.v2"
)

type MailTrapService struct {
}

func (es *MailTrapService) Send(mail iemail.Mail) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "no-reply@onlinemarketplace.com")
	for _, reciever := range mail.To {
		m.SetHeader("To", reciever)
	}
	m.SetHeader("Subject", "Activate your account")
	m.SetBody("text/html", mail.HtmlBody)

	d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, "91f20fe8a3a824", "d468719b417627")

	return d.DialAndSend(m)
}

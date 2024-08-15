package email

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"
)

type SimpleEmailSender struct {
	smtpHost     string
	smtpPort     string
	senderEmail  string
	senderPasswd string
}

func NewSimpleEmailSender(smtpHost, smtpPort, senderEmail, senderPasswd string) *SimpleEmailSender {
	return &SimpleEmailSender{
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		senderEmail:  senderEmail,
		senderPasswd: senderPasswd,
	}
}

func (s *SimpleEmailSender) SendVerificationEmail(userEmail string, token string) error {
	subject := "Verify Your Email"
	html, err := s.loadHTML("pkg/email/templates/verify_email.html")
	if err != nil {
		return err
	}
	html = strings.Replace(html, "{{token}}", token, -1)
	return s.sendMail(userEmail, subject, html)
}

func (s *SimpleEmailSender) SendPasswordResetEmail(userEmail string, token string) error {
	subject := "Reset Your Password"
	html, err := s.loadHTML("pkg/email/templates/password_reset.html")
	if err != nil {
		return err
	}
	html = strings.Replace(html, "{{token}}", token, -1)
	return s.sendMail(userEmail, subject, html)
}

func (s *SimpleEmailSender) loadHTML(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *SimpleEmailSender) sendMail(to string, subject string, html string) error {
	auth := smtp.PlainAuth("", s.senderEmail, s.senderPasswd, s.smtpHost)
	from := s.senderEmail
	toList := []string{to}

	header := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: multipart/alternative; boundary=\"boundary\"\r\n\r\n", from, toList[0], subject)
	body := fmt.Sprintf("--boundary\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s\r\n--boundary--", html)

	msg := header + body

	err := smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, from, toList, []byte(msg))
	return err
}

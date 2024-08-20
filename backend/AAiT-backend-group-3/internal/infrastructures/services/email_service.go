package services

import (
    "fmt"
    "net/smtp"
)

type IEmailService interface {
    SendEmail(to string, subject string, body string) error
}

type SMTPEmailService struct {
    smtpHost     string
    smtpPort     string
    senderEmail  string
    senderPass   string
}

func NewSMTPEmailService(host, port, email, password string) IEmailService {
    return &SMTPEmailService{
        smtpHost:    host,
        smtpPort:    port,
        senderEmail: email,
        senderPass:  password,
    }
}

func (s *SMTPEmailService) SendEmail(to string, subject string, body string) error {
    auth := smtp.PlainAuth("", s.senderEmail, s.senderPass, s.smtpHost)
    
    msg := "From: " + s.senderEmail + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)
    
    if err := smtp.SendMail(addr, auth, s.senderEmail, []string{to}, []byte(msg)); err != nil {
        return err
    }
    return nil
}


package email

import (
	"fmt"

	"github.com/mailjet/mailjet-apiv3-go"
)

type MailjetSender struct {
	Client *mailjet.Client
}

func NewMailjetSender(apiKey, apiSecret string) *MailjetSender {
	client := mailjet.NewMailjetClient(apiKey, apiSecret)
	return &MailjetSender{Client: client}
}

func (m *MailjetSender) SendVerificationEmail(userEmail string, token string) error {
	messages := mailjet.MessagesV31{
		Info: []mailjet.InfoMessagesV31{
			{
				From: &mailjet.RecipientV31{
					Email: "noreply@yourapp.com",
					Name:  "YourApp",
				},
				To: &mailjet.RecipientsV31{
					{
						Email: userEmail,
						Name:  "New User",
					},
				},
				Subject:  "Verify Your Email",
				TextPart: "Please verify your email using the link below.",
				HTMLPart: fmt.Sprintf(`
                    <p>Please verify your email by clicking the link below:</p>
                    <a href="https://yourapp.com/verify?token=%s">Verify Email</a>
                `, token),
			},
		},
	}

	_, err := m.Client.SendMailV31(&messages)
	return err
}

func (m *MailjetSender) SendPasswordResetEmail(userEmail string, token string) error {
	messages := mailjet.MessagesV31{
		Info: []mailjet.InfoMessagesV31{
			{
				From: &mailjet.RecipientV31{
					Email: "noreply@yourapp.com",
					Name:  "YourApp",
				},
				To: &mailjet.RecipientsV31{
					{
						Email: userEmail,
						Name:  "User",
					},
				},
				Subject:  "Reset Your Password",
				TextPart: "You requested to reset your password.",
				HTMLPart: fmt.Sprintf(`
                    <p>You requested to reset your password. Click the link below to reset it:</p>
                    <a href="https://yourapp.com/reset-password?token=%s">Reset Password</a>
                `, token),
			},
		},
	}

	_, err := m.Client.SendMailV31(&messages)
	return err
}

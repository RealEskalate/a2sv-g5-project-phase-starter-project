package iemail

import "fmt"

type Mail struct {
  From    string
  To      []string
  Subject string
  Body    string
  Data    string
}

func NewMail(from string, to []string, subject, body, data string) *Mail {
  return &Mail{
    From:    from,
    To:      to,
    Subject: subject,
    Body:    body,
    Data:    data,
  }
}






func NewSignUpEmail(data string, to []string, link string) *Mail {
    return &Mail{
      From:    "no-reply@example.com",
      To:      to,
      Subject: "Welcome to EBS Blog!",
      Body: fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <head>
          <title>Welcome to EBS Blog</title>
        </head>
        <body>
          <p>Hi there,</p>
          <p>Thank you for signing up for <strong>EBS Blog</strong>! We're excited to have you on board.</p>
          <p>To get started, please confirm your email address by clicking the button below:</p>
          <p style="text-align: center;">
            <a href="%s" style="display: inline-block; padding: 10px 20px; font-size: 16px; color: #ffffff; background-color: #007BFF; text-decoration: none; border-radius: 5px;">Confirm Email</a>
          </p>				
          <p>Best regards,<br>EBS Blog Team</p>
        </body>
        </html>
      `, link),
      Data: data,
    }
  }
  

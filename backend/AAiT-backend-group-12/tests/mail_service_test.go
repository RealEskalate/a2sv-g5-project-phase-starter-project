package tests

import (
	mail_service "blog_api/infrastructure/mail"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type MailServiceTestSuite struct {
	suite.Suite
	mailService *mail_service.MailService
}

func (suite *MailServiceTestSuite) SetupSuite() {
	suite.mailService = mail_service.NewMailService("smtp.gmail.com", "password")
}

func (suite *MailServiceTestSuite) TestSendMail() {
	err := suite.mailService.SendMail("from", "to", "mailContent")
	suite.Error(err, "error after SendMail has been called with incorrect SMTP credentials")
}

func (suite *MailServiceTestSuite) TestEmailVerificationTemplate() {
	hostUrl := "http://localhost:8080"
	username := "username"
	token := "token"
	str := suite.mailService.EmailVerificationTemplate(hostUrl, username, token)
	suite.Contains(str, hostUrl)
	suite.Contains(str, username)
	suite.Contains(str, token)
	suite.Contains(str, fmt.Sprint(time.Now().Year()))
}

func (suite *MailServiceTestSuite) TestPasswordResetTemplate() {
	token := "token"
	str := suite.mailService.PasswordResetTemplate(token)
	suite.Contains(str, token)
	suite.Contains(str, fmt.Sprint(time.Now().Year()))
}

func TestMailService(t *testing.T) {
	suite.Run(t, new(MailServiceTestSuite))
}

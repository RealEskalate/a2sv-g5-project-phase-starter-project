package tests

import (
	"testing"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure/mail"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type MailTestSuite struct {
	suite.Suite
	mailService domain.EmailService
}

func (suite *MailTestSuite) SetupTest() {
	mailSV := mail.NewEmailService("../infrastructure/mail/templates/")
	suite.mailService = mailSV
}

func (suite *MailTestSuite) TestSendEmail_Success() {
	data := map[string]string{
		"Name":             "testUser",
		"VerificationLink": "testLink",
	}
	errSend := suite.mailService.SendMail("testmail@blog.com", "testing the mail service", "verification.html", data)
	suite.NoError(errSend)
}

func (suite *MailTestSuite) TestSendVerificationEmail() {
	errSend := suite.mailService.SendVerificationEmail("test@example.com", "testUser", "testLink")
	suite.NoError(errSend)
}

func (suite *MailTestSuite) TestSendPasswordResetEmail() {
	errSend := suite.mailService.SendPasswordResetEmail("test@example.com", "testUser", "resetLink", "resetCode")
	suite.NoError(errSend)
}

func TestMailServiceSuite(t *testing.T) {
	godotenv.Load("../.env")
	suite.Run(t, new(MailTestSuite))
}

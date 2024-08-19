package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	toEmail := "abel.wendmu@a2sv.org"
	title := "Test Email"
	body := "This is a test email"
	hashedpwd := "hashedpassword"

	err := SendEmail(toEmail, title, body, hashedpwd)
	assert.NoError(t, err)
}

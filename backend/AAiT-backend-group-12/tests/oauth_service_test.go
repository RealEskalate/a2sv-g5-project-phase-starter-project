package tests

import (
	google_auth "blog_api/infrastructure/oauth"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type OAuthTestSuite struct {
	suite.Suite
}

func (suite *OAuthTestSuite) TestGoogleAuth_Positive() {
	clientId := "test_client_id"
	tokenEmail := "test_email@gmail.com"
	expiresAt := time.Now().Round(0).Add(time.Hour * 24).Unix()
	googleUser := google_auth.GoogleUser{
		ISS:   "accounts.google.com",
		AUD:   clientId,
		Email: tokenEmail,
		Exp:   fmt.Sprint(expiresAt),
	}

	err := google_auth.VerifyResponseContents(googleUser, clientId, tokenEmail)
	suite.NoError(err)
}

func (suite *OAuthTestSuite) TestGoogleAuth_Negative_InvalidIssuer() {
	clientId := "test_client_id"
	tokenEmail := "test_email@gmail.com"
	expiresAt := time.Now().Round(0).Add(time.Hour * 24).Unix()
	googleUser := google_auth.GoogleUser{
		ISS:   "not.google.com",
		AUD:   clientId,
		Email: tokenEmail,
		Exp:   fmt.Sprint(expiresAt),
	}

	err := google_auth.VerifyResponseContents(googleUser, clientId, tokenEmail)
	suite.Error(err)
}

func (suite *OAuthTestSuite) TestGoogleAuth_Negative_InvalidClientID() {
	clientId := "test_client_id"
	tokenEmail := "test_email@gmail.com"
	expiresAt := time.Now().Round(0).Add(time.Hour * 24).Unix()
	googleUser := google_auth.GoogleUser{
		ISS:   "accounts.google.com",
		AUD:   clientId,
		Email: tokenEmail,
		Exp:   fmt.Sprint(expiresAt),
	}

	err := google_auth.VerifyResponseContents(googleUser, clientId+":SDKLF", tokenEmail)
	suite.Error(err)
}

func (suite *OAuthTestSuite) TestGoogleAuth_Negative_InvalidEmail() {
	clientId := "test_client_id"
	tokenEmail := "test_email@gmail.com"
	expiresAt := time.Now().Round(0).Add(time.Hour * 24).Unix()
	googleUser := google_auth.GoogleUser{
		ISS:   "accounts.google.com",
		AUD:   clientId,
		Email: tokenEmail,
		Exp:   fmt.Sprint(expiresAt),
	}

	err := google_auth.VerifyResponseContents(googleUser, clientId, "real"+tokenEmail)
	suite.Error(err)
}

func TestOAuthService(t *testing.T) {
	suite.Run(t, new(OAuthTestSuite))
}

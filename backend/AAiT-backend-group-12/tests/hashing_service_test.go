package tests

import (
	"blog_api/infrastructure/cryptography"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HashServiceTestSuite struct {
	suite.Suite
	hashingService *cryptography.HashingService
}

func (suite *HashServiceTestSuite) SetupSuite() {
	suite.hashingService = cryptography.NewHashingService()
}

func (suite *HashServiceTestSuite) TestHashString() {
	password := "plain_text password"

	hashedPwd, err := suite.hashingService.HashString(password)

	suite.NoError(err, "no errors when given password")
	suite.NotEqual(hashedPwd, password)
}

func (suite *HashServiceTestSuite) TestValidatePassword_Positive() {
	password := "plain_text password"

	hashedPwd, _ := suite.hashingService.HashString(password)
	err := suite.hashingService.ValidateHashedString(hashedPwd, password)

	suite.NoError(err, "no errors when given correct password")
}

func (suite *HashServiceTestSuite) TestValidatePassword_Negative() {
	password := "plain_text password"

	hashedPwd, _ := suite.hashingService.HashString(password)
	err := suite.hashingService.ValidateHashedString(hashedPwd, "wrong plain_text password")

	suite.Error(err, "error when given incorrect password")
}

func TestHashService(t *testing.T) {
	suite.Run(t, new(HashServiceTestSuite))
}

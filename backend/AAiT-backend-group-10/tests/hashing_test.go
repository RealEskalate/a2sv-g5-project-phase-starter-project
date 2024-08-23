package tests

import (
	"testing"

	"aait.backend.g10/usecases/interfaces"
	"aait.backend.g10/infrastructures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PasswordServiceSuite struct {
	suite.Suite
	service interfaces.IHashingService
}

func (suite *PasswordServiceSuite) SetupTest() {
	suite.service = &infrastructures.HashingService{}
}

// Test HashPassword

func (suite *PasswordServiceSuite) TestHashPassword_Success() {
	password := "mySecureP@ssw0rd"
	hashedPassword, err := suite.service.HashPassword(password)
	assert.Nil(suite.T(), err)
	assert.NotEmpty(suite.T(), hashedPassword)

	// Verify that the hashed password is different from the plain password
	assert.NotEqual(suite.T(), password, hashedPassword)
}

// Test ComparePassword

func (suite *PasswordServiceSuite) TestComparePassword_Success() {
	password := "mySecureP@ssw0rd"
	hashedPassword, _ := suite.service.HashPassword(password)

	isMatch := suite.service.CheckPasswordHash(password, hashedPassword)
	assert.True(suite.T(), isMatch)
}

func (suite *PasswordServiceSuite) TestComparePassword_Failure() {
	password := "mySecureP@ssw0rd"
	hashedPassword, _ := suite.service.HashPassword(password)

	// Incorrect password should not match
	isMatch := suite.service.CheckPasswordHash("wrongPassword", hashedPassword)
	assert.False(suite.T(), isMatch)
}

func TestPasswordServiceSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceSuite))
}

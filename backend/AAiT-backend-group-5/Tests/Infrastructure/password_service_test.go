package infrastructure_test

import (
	"testing"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	"github.com/stretchr/testify/suite"
)

type PasswordServiceTestSuite struct {
	suite.Suite
	passwordService interfaces.PasswordService
}

func (suite *PasswordServiceTestSuite) SetupSuite() {
	suite.passwordService = infrastructure.NewPasswordService()
}

func (suite *PasswordServiceTestSuite) TestEncryptPassword_Success() {
	password := "Password123!"
	encryptedPassword, err := suite.passwordService.EncryptPassword(password)
	suite.NoError(err)
	suite.NotEmpty(encryptedPassword)
}

func (suite *PasswordServiceTestSuite) TestValidatePassword_Success() {
	password := "Password123!"
	encryptedPassword, _ := suite.passwordService.EncryptPassword(password)
	isValid := suite.passwordService.ValidatePassword(password, encryptedPassword)
	suite.True(isValid)
}

func (suite *PasswordServiceTestSuite) TestValidatePasswordStrength_TooShort() {
	password := "short"
	errResponse := suite.passwordService.ValidatePasswordStrength(password)
	suite.NotNil(errResponse)
	suite.Equal("Password must be at least 8 characters long.", errResponse.Message)
}

func (suite *PasswordServiceTestSuite) TestValidatePasswordStrength_NoLowercase() {
	password := "PASSWORD123!"
	errResponse := suite.passwordService.ValidatePasswordStrength(password)
	suite.NotNil(errResponse)
	suite.Equal("Password must contain at least one lowercase letter, one uppercase letter, one digit, and one special character.", errResponse.Message)
}

func (suite *PasswordServiceTestSuite) TestValidatePasswordStrength_NoUppercase() {
	password := "password123!"
	errResponse := suite.passwordService.ValidatePasswordStrength(password)
	suite.NotNil(errResponse)
	suite.Equal("Password must contain at least one lowercase letter, one uppercase letter, one digit, and one special character.", errResponse.Message)
}

func (suite *PasswordServiceTestSuite) TestValidatePasswordStrength_NoDigit() {
	password := "Password!"
	errResponse := suite.passwordService.ValidatePasswordStrength(password)
	suite.NotNil(errResponse)
	suite.Equal("Password must contain at least one lowercase letter, one uppercase letter, one digit, and one special character.", errResponse.Message)
}

func (suite *PasswordServiceTestSuite) TestValidatePasswordStrength_NoSymbol() {
	password := "Password123"
	errResponse := suite.passwordService.ValidatePasswordStrength(password)
	suite.NotNil(errResponse)
	suite.Equal("Password must contain at least one lowercase letter, one uppercase letter, one digit, and one special character.", errResponse.Message)
}

func (suite *PasswordServiceTestSuite) TestValidatePasswordStrength_Success() {
	password := "Password123!"
	errResponse := suite.passwordService.ValidatePasswordStrength(password)
	suite.Nil(errResponse)
}

func TestPasswordServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceTestSuite))
}

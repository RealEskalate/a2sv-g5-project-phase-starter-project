package tests

import (
	"testing"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure"
	"github.com/stretchr/testify/suite"
)

type PasswordServiceTestSuite struct {
	suite.Suite
	passwordService domain.PasswordService
}

func (suite *PasswordServiceTestSuite) SetupTest() {
	suite.passwordService = infrastructure.NewPasswordService()

}

func (suite *PasswordServiceTestSuite) TestHashPassword() {
	password := "securepassword"
	_, err := suite.passwordService.HashPassword(password)

	suite.NoError(err, "Error hashing password")
}

func (suite *PasswordServiceTestSuite) TestHashPassword_Fail() {
	password := "re"
	_, err := suite.passwordService.HashPassword(password)

	suite.Error(err, "Error hashing password")
}

func (suite *PasswordServiceTestSuite) TestComparePassword() {
	password := "securepassword"
	hashedPassword, _ := suite.passwordService.HashPassword(password)

	match, err := suite.passwordService.ComparePassword(hashedPassword, password)

	suite.True(match, "Passwords do not match")
	suite.NoError(err, "Error comparing passwords")
}

func (suite *PasswordServiceTestSuite) TestComparePassword_Fail() {
	password := "securepassword"
	hashedPassword, _ := suite.passwordService.HashPassword(password)

	match, err := suite.passwordService.ComparePassword(hashedPassword, "wrongpassword")

	suite.False(match, "Passwords match")
	suite.Error(err, "Error comparing passwords")
}

func TestPasswordServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceTestSuite))
}

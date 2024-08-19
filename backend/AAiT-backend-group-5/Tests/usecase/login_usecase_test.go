package usecase

import (
	"context"
	"errors"
	"testing"

	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type PasswordServiceTestSuite struct {
	suite.Suite
	ctrl                *gomock.Controller
	mockPasswordService *mocks.MockPasswordService
	mockPasswordUsecase *mocks.MockPasswordUsecase
	mockUserRepository  *mocks.MockUserRepository
}

func (suite *PasswordServiceTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockPasswordService = mocks.NewMockPasswordService(suite.ctrl)
	suite.mockPasswordUsecase = mocks.NewMockPasswordUsecase(suite.ctrl)
	suite.mockUserRepository = mocks.NewMockUserRepository(suite.ctrl)
}

func (suite *PasswordServiceTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

// Test cases for PasswordService
func (suite *PasswordServiceTestSuite) TestEncryptPasswordSuccess() {
	password := "myPassword"
	hashedPassword := "hashedPassword"

	suite.mockPasswordService.EXPECT().
		EncryptPassword(password).
		Return(hashedPassword, nil)

	result, err := suite.mockPasswordService.EncryptPassword(password)
	
	if err != nil {
		suite.T().Errorf("Error: %v", err)
	}
	suite.Equal(hashedPassword, result)
}

func (suite *PasswordServiceTestSuite) TestEncryptPasswordFailure() {
	password := "myPassword"
	expectedErr := errors.New("encryption error")

	suite.mockPasswordService.EXPECT().
		EncryptPassword(password).
		Return("", expectedErr)

	result, err := suite.mockPasswordService.EncryptPassword(password)
	suite.Equal(expectedErr, err)
	suite.Empty(result)
}

func (suite *PasswordServiceTestSuite) TestValidatePasswordSuccess() {
	password := "myPassword"
	hashedPassword := "hashedPassword"

	suite.mockPasswordService.EXPECT().
		ValidatePassword(password, hashedPassword).
		Return(true)

	result := suite.mockPasswordService.ValidatePassword(password, hashedPassword)
	suite.True(result)
}

func (suite *PasswordServiceTestSuite) TestValidatePasswordFailure() {
	password := "myPassword"
	hashedPassword := "hashedPassword"

	suite.mockPasswordService.EXPECT().
		ValidatePassword(password, hashedPassword).
		Return(false)

	result := suite.mockPasswordService.ValidatePassword(password, hashedPassword)
	suite.False(result)
}

// Test cases for PasswordUsecase with UserRepository
func (suite *PasswordServiceTestSuite) TestGenerateResetURLSuccess() {
	email := "user@example.com"
	resetURL := "http://example.com/reset"

	suite.mockPasswordUsecase.EXPECT().
		GenerateResetURL(context.Background(), email).
		Return(resetURL, nil)

	result, err := suite.mockPasswordUsecase.GenerateResetURL(context.Background(), email)
	
	if err != nil {
		suite.T().Errorf("Error: %v", err)
	}
	suite.Equal(resetURL, result)
}

func (suite *PasswordServiceTestSuite) TestSendResetEmailSuccess() {
	email := "user@example.com"
	resetURL := "http://example.com/reset"

	suite.mockPasswordUsecase.EXPECT().
		SendResetEmail(context.Background(), email, resetURL).
		Return(nil)

	err := suite.mockPasswordUsecase.SendResetEmail(context.Background(), email, resetURL)
	
	if err != nil {
		suite.T().Errorf("Error: %v", err)
	}
}

func (suite *PasswordServiceTestSuite) TestSetPasswordSuccess() {
	shortURLCode := "shortURLCode"
	password := "newPassword"

	suite.mockPasswordUsecase.EXPECT().
		SetPassword(context.Background(), shortURLCode, password).
		Return(nil)

	err := suite.mockPasswordUsecase.SetPassword(context.Background(), shortURLCode, password)
	
	if err != nil {
		suite.T().Errorf("Error: %v", err)
	}
}

func TestPasswordServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceTestSuite))
}

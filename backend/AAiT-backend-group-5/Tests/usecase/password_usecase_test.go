package usecase

import (
	"context"
	"testing"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type SetupPasswordTestSuite struct {
	suite.Suite
	ctrl                 *gomock.Controller
	mockURLService       *mocks.MockURLService
	mockJwtService       *mocks.MockJwtService
	mockEmailService     *mocks.MockEmailService
	mockPasswordService  *mocks.MockPasswordService
	mockUserRepository   *mocks.MockUserRepository
	setupPasswordUsecase interfaces.PasswordUsecase
}

func (suite *SetupPasswordTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockURLService = mocks.NewMockURLService(suite.ctrl)
	suite.mockJwtService = mocks.NewMockJwtService(suite.ctrl)
	suite.mockEmailService = mocks.NewMockEmailService(suite.ctrl)
	suite.mockPasswordService = mocks.NewMockPasswordService(suite.ctrl)
	suite.mockUserRepository = mocks.NewMockUserRepository(suite.ctrl)
	suite.setupPasswordUsecase = usecases.NewSetupPassword(
		suite.mockURLService,
		suite.mockJwtService,
		suite.mockUserRepository,
		suite.mockEmailService,
		suite.mockPasswordService,
	)
}

func (suite *SetupPasswordTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *SetupPasswordTestSuite) TestGenerateResetURLSuccess() {
	email := "user@example.com"
	user := &models.User{ID: "user123", Email: email}
	token := "token123"
	resetURL := "http://example.com/reset?token=" + token

	suite.mockUserRepository.EXPECT().
		GetUserByEmailOrUsername(context.Background(), email, email).
		Return(user, nil)

	suite.mockJwtService.EXPECT().
		CreateAccessToken(*user, 3600).
		Return(token, nil)

	suite.mockURLService.EXPECT().
		GenerateURL(token).
		Return(resetURL, nil)

	resetURLResult, errResp := suite.setupPasswordUsecase.GenerateResetURL(context.Background(), email)

	if errResp != nil {
		suite.T().Errorf("Expected no error but got, %v", errResp.Message)
	}
	suite.Equal(resetURL, resetURLResult)
}

func (suite *SetupPasswordTestSuite) TestGenerateResetURLUserNotFound() {
	email := "user@example.com"

	suite.mockUserRepository.EXPECT().
		GetUserByEmailOrUsername(context.Background(), email, email).
		Return(nil, models.BadRequest("User not found"))

	resetURL, errResp := suite.setupPasswordUsecase.GenerateResetURL(context.Background(), email)

	suite.Error(errResp)
	suite.Empty(resetURL)
}

func (suite *SetupPasswordTestSuite) TestSendResetEmailSuccess() {
	email := "user@example.com"
	resetURL := "http://example.com/reset?token=token123"
	subject := "Password Reset"
	body := "Click the link below to reset your password\n" + resetURL + "\n\nThis link will expire in 1 hour"

	suite.mockEmailService.EXPECT().
		IsValidEmail(email).
		Return(true)

	suite.mockEmailService.EXPECT().
		SendEmail(email, subject, body).
		Return(nil)

	errResp := suite.setupPasswordUsecase.SendResetEmail(context.Background(), email, resetURL)

	if errResp != nil {
		suite.T().Errorf("Expected no error but got, %v", errResp.Message)
	}
}

func (suite *SetupPasswordTestSuite) TestSendResetEmailInvalidEmail() {
	email := "invalid-email"
	resetURL := "http://example.com/reset?token=token123"

	suite.mockEmailService.EXPECT().
		IsValidEmail(email).
		Return(false)

	errResp := suite.setupPasswordUsecase.SendResetEmail(context.Background(), email, resetURL)

	suite.Error(errResp)
	suite.Equal("Invalid email address", errResp.Message)
}

func (suite *SetupPasswordTestSuite) TestSetPasswordInvalidToken() {
	shortURLCode := "shortURLCode"
	password := "newpassword"
	invalidTokenError := models.BadRequest("Invalid token")

	suite.mockURLService.EXPECT().
		GetURL(shortURLCode).
		Return(nil, invalidTokenError)

	errResp := suite.setupPasswordUsecase.SetPassword(context.Background(), shortURLCode, password)

	suite.Error(errResp)
	suite.Equal("Invalid token", errResp.Message)
}

func TestSetupPasswordTestSuite(t *testing.T) {
	suite.Run(t, new(SetupPasswordTestSuite))
}

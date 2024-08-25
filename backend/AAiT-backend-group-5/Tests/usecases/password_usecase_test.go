package usecases_test

import (
	"context"
	"testing"

	config "github.com/aait.backend.g5.main/backend/Config"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type SetupPasswordTestSuite struct {
	suite.Suite
	urlServiceMock      *mocks.MockURLService
	jwtServiceMock      *mocks.MockJwtService
	emailServiceMock    *mocks.MockEmailService
	passwordServiceMock *mocks.MockPasswordService
	repoMock            *mocks.MockUserRepository
	otpServiceMock      *mocks.MockOTPService
	env                 config.Env
	setupPassword       interfaces.PasswordUsecase
	ctr                 *gomock.Controller
}

func (suite *SetupPasswordTestSuite) SetupSuite() {
	suite.ctr = gomock.NewController(suite.T())
	suite.urlServiceMock = mocks.NewMockURLService(suite.ctr)
	suite.jwtServiceMock = mocks.NewMockJwtService(suite.ctr)
	suite.emailServiceMock = mocks.NewMockEmailService(suite.ctr)
	suite.passwordServiceMock = mocks.NewMockPasswordService(suite.ctr)
	suite.repoMock = mocks.NewMockUserRepository(suite.ctr)
	suite.otpServiceMock = mocks.NewMockOTPService(suite.ctr)

	suite.env = config.Env{}

	suite.setupPassword = usecases.NewSetupPassword(
		suite.urlServiceMock,
		suite.jwtServiceMock,
		suite.repoMock,
		suite.emailServiceMock,
		suite.passwordServiceMock,
		suite.otpServiceMock,
	)
}

func (suite *SetupPasswordTestSuite) TearDownSuite() {
	suite.ctr.Finish()
}

func (suite *SetupPasswordTestSuite) TestGenerateResetURL_Success_Web() {
	ctx := context.Background()
	email := "user@example.com"
	agent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
	user := &models.User{
		Email: email,
	}

	token := "reset_token"
	resetURL := "http://example.com/resetPassword?token=" + token

	suite.repoMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, email, email).
		Return(user, nil)

	suite.jwtServiceMock.
		EXPECT().
		CreateAccessToken(*user, 3600).
		Return(token, nil)

	suite.urlServiceMock.
		EXPECT().
		GenerateURL(token, "resetPassword").
		Return(resetURL, nil)

	resultURL, err := suite.setupPassword.GenerateResetURL(ctx, email, agent)
	suite.Nil(err)
	suite.Equal(resetURL, resultURL)
}

func (suite *SetupPasswordTestSuite) TestGenerateResetURL_Success_Mobile() {
	ctx := context.Background()
	email := "user@example.com"
	agent := "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E5239e Safari/602.1"
	user := &models.User{
		Email: email,
	}

	token := "reset_token"
	otp := "123456"

	suite.repoMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, email, email).
		Return(user, nil)

	suite.jwtServiceMock.
		EXPECT().
		CreateAccessToken(*user, 3600).
		Return(token, nil)

	suite.otpServiceMock.
		EXPECT().
		GenerateOTP(token).
		Return(otp, nil)

	resultURL, err := suite.setupPassword.GenerateResetURL(ctx, email, agent)
	suite.Nil(err)
	suite.Equal(otp, resultURL)
}

func (suite *SetupPasswordTestSuite) TestGenerateResetURL_UserNotFound() {
	ctx := context.Background()
	email := "user@example.com"
	agent := "Mozilla/5.0"

	suite.repoMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, email, email).
		Return(nil, models.NotFound("User not found"))

	resultURL, err := suite.setupPassword.GenerateResetURL(ctx, email, agent)
	suite.Equal(models.NotFound("User not found"), err)
	suite.Empty(resultURL)
}

func (suite *SetupPasswordTestSuite) TestGenerateResetURL_TokenError() {
	ctx := context.Background()
	email := "user@example.com"
	agent := "Mozilla/5.0"
	user := &models.User{
		Email: email,
	}

	suite.repoMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, email, email).
		Return(user, nil)

	suite.jwtServiceMock.
		EXPECT().
		CreateAccessToken(*user, 3600).
		Return("", models.InternalServerError("Error generating token"))

	resultURL, err := suite.setupPassword.GenerateResetURL(ctx, email, agent)
	suite.Equal(models.InternalServerError("An error occurred while generating the reset URL"), err)
	suite.Empty(resultURL)
}

func (suite *SetupPasswordTestSuite) TestSendResetEmail_Success() {
	ctx := context.Background()
	email := "user@example.com"
	resetURL := "http://example.com/resetPassword?token=reset_token"

	suite.emailServiceMock.
		EXPECT().
		IsValidEmail(email).
		Return(true)

	suite.emailServiceMock.
		EXPECT().
		SendEmail(email, "Password Reset", "you can use the below to reset your password \n"+resetURL+"\nThis link will expire in 1 hour").
		Return(nil)

	err := suite.setupPassword.SendResetEmail(ctx, email, resetURL)
	suite.Nil(err)
}

func (suite *SetupPasswordTestSuite) TestSendResetEmail_InvalidEmail() {
	ctx := context.Background()
	email := "invalid-email"
	resetURL := "http://example.com/resetPassword?token=reset_token"

	suite.emailServiceMock.
		EXPECT().
		IsValidEmail(email).
		Return(false)

	err := suite.setupPassword.SendResetEmail(ctx, email, resetURL)
	suite.Equal(models.BadRequest("Invalid email address"), err)
}

func (suite *SetupPasswordTestSuite) TestSendResetEmail_ErrorSendingEmail() {
	ctx := context.Background()
	email := "user@example.com"
	resetURL := "http://example.com/resetPassword?token=reset_token"

	suite.emailServiceMock.
		EXPECT().
		IsValidEmail(email).
		Return(true)

	suite.emailServiceMock.
		EXPECT().
		SendEmail(email, "Password Reset", "you can use the below to reset your password \n"+resetURL+"\nThis link will expire in 1 hour").
		Return(models.InternalServerError("Error sending email"))

	err := suite.setupPassword.SendResetEmail(ctx, email, resetURL)
	suite.Equal(models.InternalServerError("An error occurred while sending the reset email"), err)
}

func (suite *SetupPasswordTestSuite) TestSetNewUserPassword_Success() {
	ctx := context.Background()
	shortURlCode := "short_url_code"
	password := "new_password"

	user := &models.URLTokenCustom{
		Name:     "John Doe",
		Username: "johndoe",
		Email:    "user@example.com",
		Role:     "user",
	}

	hashedPassword := "hashed_password"

	suite.urlServiceMock.
		EXPECT().
		GetURL(shortURlCode).
		Return(&models.URL{Token: "valid_token"}, nil)

	suite.jwtServiceMock.
		EXPECT().
		ValidateURLToken("valid_token").
		Return(user, nil)

	suite.passwordServiceMock.
		EXPECT().
		ValidatePasswordStrength(password).
		Return(nil)

	suite.passwordServiceMock.
		EXPECT().
		EncryptPassword(password).
		Return(hashedPassword, nil)

	suite.repoMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, user.Username, user.Email).
		Return(nil, nil)

	suite.repoMock.
		EXPECT().
		CreateUser(ctx, &models.User{
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			Password: hashedPassword,
			Role:     models.RoleUser,
		}).
		Return(nil)

	suite.urlServiceMock.
		EXPECT().
		RemoveURL(shortURlCode).
		Return(nil)

	err := suite.setupPassword.SetNewUserPassword(ctx, shortURlCode, password)
	suite.Nil(err)
}


func TestPasswordUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(SetupPasswordTestSuite))
}

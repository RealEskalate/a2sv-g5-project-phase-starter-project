package tests

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"testing"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/mocks"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/usecases"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecaseTest struct {
	suite.Suite
	usecase         domain.UserUseCase
	repository      *mocks.UserRepository
	passwordService *mocks.PasswordService
	jwtService      *mocks.JwtService
	mailService     *mocks.EmailService
	redisService    *mocks.CacheService
	sessionRP       *mocks.SessionRepository
	utilsMock       *mocks.Utils
}

func (suite *UserUsecaseTest) SetupTest() {
	userRP := new(mocks.UserRepository)
	sessionRP := new(mocks.SessionRepository)
	passwordService := new(mocks.PasswordService)
	jwtService := new(mocks.JwtService)
	mailService := new(mocks.EmailService)
	redisService := new(mocks.CacheService)
	utilsMock := new(mocks.Utils)
	suite.usecase = usecases.NewUserUseCase(userRP, sessionRP, passwordService, jwtService, mailService, redisService, utilsMock)
	suite.repository = userRP
	suite.passwordService = passwordService
	suite.jwtService = jwtService
	suite.mailService = mailService
	suite.redisService = redisService
	suite.sessionRP = sessionRP
	suite.utilsMock = utilsMock
}

func (suite *UserUsecaseTest) TestRegisterStart_Positive() {
	user := domain.User{
		Username: "username",
		Password: "passworD.2132",
		Email:    "username@user.com",
		Role:     "user",
	}

	suite.repository.On("FindByEmail", mock.Anything, user.Email).Return(&domain.User{}, domain.CustomError{Message: "", Code: http.StatusNotFound})
	suite.repository.On("FindByUsername", mock.Anything, user.Username).Return(&domain.User{}, domain.CustomError{Message: "", Code: http.StatusNotFound})
	suite.passwordService.On("HashPassword", user.Password).Return("passworD.2132", nil)
	suite.jwtService.On("GenerateVerificationToken", user).Return("token", nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, user.Username).Return(&domain.Session{
		Username:          user.Username,
		RefreshToken:      "refreshToken",
		VerificationToken: "token",
	}, false, domain.CustomError{Message: "", Code: http.StatusNotFound})
	suite.sessionRP.On("CreateToken", mock.Anything, &domain.Session{
		Username:          user.Username,
		VerificationToken: "token",
	}).Return(&domain.Session{
		Username:          user.Username,
		RefreshToken:      "refreshToken",
		VerificationToken: "token",
	}, nil)
	suite.mailService.On("SendVerificationEmail", user.Email, user.Username, "http://localhost:8080/user/verify/token").Return(nil)

	errReg := suite.usecase.RegisterStart(context.TODO(), &user)
	suite.NoError(errReg, "Error should be nil")
}

func (suite *UserUsecaseTest) TestRegisterStart_InvalidCredentials() {
	user := domain.User{
		Username: "username",
		Password: "password.2132",
		Email:    "username@user.com",
		Role:     "user",
	}

	errReg := suite.usecase.RegisterStart(context.TODO(), &user)
	suite.Error(errReg, "Error should be nil")
	suite.Equal(http.StatusBadRequest, errReg.StatusCode(), "Error code should be 400")

	user = domain.User{
		Username: "username",
		Password: "passworD.2132",
		Email:    "usernameuser.com",
		Role:     "user",
	}

	errReg = suite.usecase.RegisterStart(context.TODO(), &user)
	suite.Error(errReg, "Error should be nil")
	suite.Equal(http.StatusBadRequest, errReg.StatusCode(), "Error code should be 400")
}

func (suite *UserUsecaseTest) TestRegisterStart_ExistingUser() {
	user := domain.User{
		Username: "user",
		Password: "Password.2132",
		Email:    "user@user.com",
		Role:     "user",
	}

	suite.repository.On("FindByEmail", mock.Anything, user.Email).Return(&domain.User{}, nil)
	suite.repository.On("FindByUsername", mock.Anything, user.Username).Return(&domain.User{
		Username: user.Username,
		Password: "Password.2132",
		Email:    "user@user.com",
		Role:     "user",
	}, nil)

	errReg := suite.usecase.RegisterStart(context.TODO(), &user)
	suite.Error(errReg, "Error should not be nil")
	suite.Equal(http.StatusBadRequest, errReg.StatusCode(), "Error code should be 400")
}

func (suite *UserUsecaseTest) TestRegisterEnd() {
	mockUsername := "testUser"
	mockToken := "mockValidToken"
	mockClaims := jwt.MapClaims{
		"username": "testUser",
		"email":    "test@example.com",
		"role":     "user",
		"password": "hashedPassword",
	}
	mockParsedToken := &jwt.Token{Claims: mockClaims}
	mockSession := &domain.Session{
		ID:                 primitive.NilObjectID,
		VerificationToken:  mockToken,
		Username:           mockUsername,
		PasswordResetToken: "",
	}

	suite.jwtService.On("ValidateVerificationToken", mockToken).Return(mockParsedToken, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)
	suite.repository.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil, nil)

	errReg := suite.usecase.RegisterEnd(context.TODO(), mockToken)
	suite.NoError(errReg, "RegisterEnd should succeed")

}

func (suite *UserUsecaseTest) TestRegisterEnd_InvalidToken() {
	mockToken := "invalid Token"
	suite.jwtService.On("ValidateVerificationToken", mockToken).Return(nil, domain.CustomError{
		Message: "invalid token",
		Code:    http.StatusBadRequest,
	})

	errReg := suite.usecase.RegisterEnd(context.TODO(), mockToken)
	suite.Error(errReg, "error should not be nil")
}

func (suite *UserUsecaseTest) TestRegisterEnd_InvaildClaim() {
	mockToken := "invalid Token"
	mockParsedToken := &jwt.Token{}
	suite.jwtService.On("ValidateVerificationToken", mockToken).Return(mockParsedToken, nil)

	errReg := suite.usecase.RegisterEnd(context.TODO(), mockToken)
	suite.Error(errReg, "error should not be nil, error should rise for invalid claim")
}

func (suite *UserUsecaseTest) TestRegisterEnd_TokenMismatch() {
	mockToken := "invalid Token"
	mockClaim := jwt.MapClaims{
		"username": "testUser",
		"email":    "test@example.com",
		"role":     "user",
		"password": "hashedPassword",
	}

	mockParsedToken := jwt.Token{Claims: mockClaim}
	suite.jwtService.On("ValidateVerificationToken", mockToken).Return(&mockParsedToken, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, "testUser").Return(&domain.Session{
		Username:          "testUser",
		VerificationToken: "another Token",
	}, true, nil)

	errReg := suite.usecase.RegisterEnd(context.TODO(), mockToken)
	suite.Error(errReg, "RegisterEnd should fail if token does not match")
}

func (suite *UserUsecaseTest) TestRegisterEnd_UserRegisterFail() {
	mockToken := "invalid Token"
	mockClaim := jwt.MapClaims{
		"username": "testUser",
		"email":    "test@example.com",
		"role":     "user",
		"password": "hashedPassword",
	}

	mockParsedToken := jwt.Token{Claims: mockClaim}
	suite.jwtService.On("ValidateVerificationToken", mockToken).Return(&mockParsedToken, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, "testUser").Return(&domain.Session{
		Username:          "testUser",
		VerificationToken: "another Token",
	}, true, nil)

	suite.repository.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(&domain.User{}, domain.CustomError{
		Message: "error creating a user",
		Code:    http.StatusInternalServerError,
	})

	errReg := suite.usecase.RegisterEnd(context.TODO(), mockToken)
	suite.Error(errReg, "RegisterEnd should fail if user creation fails")
}

func (suite *UserUsecaseTest) TestLogin_Success() {
	mockUsername := "testUser"
	mockPassword := "testPassword"
	mockHashedPassword := "hashedPassword"
	mockAccessToken := "mockAccessToken"
	mockRefreshToken := "mockRefreshToken"

	mockUser := domain.User{
		Username: mockUsername,
		Password: mockHashedPassword,
		Email:    "user@user.com",
		Role:     "user",
	}

	mockSession := &domain.Session{
		ID:                 primitive.NilObjectID,
		VerificationToken:  "mockVerificationToken",
		PasswordResetToken: "mockPasswordResetToken",
		Username:           mockUsername,
	}

	suite.repository.On("FindByUsername", mock.Anything, mockUsername).Return(&mockUser, nil)
	suite.passwordService.On("ComparePassword", mockHashedPassword, mockPassword).Return(true, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)
	suite.jwtService.On("GenerateAccessTokenWithPayload", mockUser).Return(mockAccessToken, nil)
	suite.jwtService.On("GenerateRefreshTokenWithPayload", mockUser).Return(mockRefreshToken, nil)

	updatedSession := *mockSession
	updatedSession.RefreshToken = mockRefreshToken

	suite.sessionRP.On("UpdateToken", mock.Anything, mockSession.ID.Hex(), &updatedSession).Return(nil)

	returnValue, errLogin := suite.usecase.Login(context.TODO(), mockUsername, mockPassword)

	suite.NoError(errLogin, "Login should succeed with valid credentials")
	suite.Equal(mockAccessToken, returnValue["access_token"], "Access token should be equal")
	suite.Equal(mockRefreshToken, returnValue["refresh_token"], "Refresh token should be equal")
}

func (suite *UserUsecaseTest) TestLogin_UserNotFound() {
	mockUsername := "testUser"
	mockPassword := "testPassword"

	suite.repository.On("FindByUsername", mock.Anything, mockUsername).Return(nil, &domain.CustomError{Message: "User not found", Code: http.StatusNotFound})

	result, err := suite.usecase.Login(context.TODO(), mockUsername, mockPassword)

	suite.Error(err)
	suite.Equal(http.StatusNotFound, err.StatusCode())
	suite.Empty(result)
}

func (suite *UserUsecaseTest) TestLogin_InvalidPassword() {
	mockUsername := "testUser"
	mockPassword := "testPassword"
	mockHashedPassword := "hashedPassword"

	mockUser := &domain.User{
		Username: mockUsername,
		Password: mockHashedPassword,
	}

	suite.repository.On("FindByUsername", mock.Anything, mockUsername).Return(mockUser, nil)
	suite.passwordService.On("ComparePassword", mockHashedPassword, mockPassword).Return(false, nil)

	result, err := suite.usecase.Login(context.TODO(), mockUsername, mockPassword)

	suite.Error(err)
	suite.Equal(http.StatusUnauthorized, err.StatusCode())
	suite.Empty(result)
}

func (suite *UserUsecaseTest) TestLogin_ErrorGeneratingAccessToken() {
	mockUsername := "testUser"
	mockPassword := "testPassword"
	mockHashedPassword := "hashedPassword"

	mockUser := &domain.User{
		Username: mockUsername,
		Password: mockHashedPassword,
	}

	mockSession := &domain.Session{
		ID:                 primitive.NewObjectID(),
		VerificationToken:  "mockVerificationToken",
		PasswordResetToken: "mockPasswordResetToken",
		Username:           mockUsername,
	}

	suite.repository.On("FindByUsername", mock.Anything, mockUsername).Return(mockUser, nil)
	suite.passwordService.On("ComparePassword", mockHashedPassword, mockPassword).Return(true, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)
	suite.jwtService.On("GenerateAccessTokenWithPayload", *mockUser).Return("", domain.CustomError{Message: "error generating access token", Code: http.StatusInternalServerError})

	result, err := suite.usecase.Login(context.TODO(), mockUsername, mockPassword)

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Empty(result)
}

func (suite *UserUsecaseTest) TestLogin_ErrorUpdatingToken() {
	mockUsername := "testUser"
	mockPassword := "testPassword"
	mockHashedPassword := "hashedPassword"
	mockAccessToken := "mockAccessToken"
	mockRefreshToken := "mockRefreshToken"

	mockUser := &domain.User{
		Username: mockUsername,
		Password: mockHashedPassword,
	}

	mockSession := &domain.Session{
		ID:                 primitive.NewObjectID(),
		VerificationToken:  "mockVerificationToken",
		PasswordResetToken: "mockPasswordResetToken",
		Username:           mockUsername,
	}

	suite.repository.On("FindByUsername", mock.Anything, mockUsername).Return(mockUser, nil)
	suite.passwordService.On("ComparePassword", mockHashedPassword, mockPassword).Return(true, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)
	suite.jwtService.On("GenerateAccessTokenWithPayload", *mockUser).Return(mockAccessToken, nil)
	suite.jwtService.On("GenerateRefreshTokenWithPayload", *mockUser).Return(mockRefreshToken, nil)
	suite.sessionRP.On("UpdateToken", mock.Anything, mockSession.ID.Hex(), mock.AnythingOfType("*domain.Session")).Return(domain.CustomError{Message: "error updating token", Code: http.StatusInternalServerError})

	result, err := suite.usecase.Login(context.TODO(), mockUsername, mockPassword)

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Empty(result)
}

func (suite *UserUsecaseTest) TestRefreshToken_Success() {
	mockUserID := "mockUserID"
	mockRefreshToken := "mockRefreshToken"
	mockAccessToken := "mockAccessToken"
	mockUsername := "mockUsername"
	mockUser := &domain.User{ID: primitive.NilObjectID, Username: mockUsername}
	mockSession := &domain.Session{Username: mockUsername, RefreshToken: mockRefreshToken}

	suite.jwtService.On("ValidateRefreshToken", mockRefreshToken).Return(&jwt.Token{
		Claims: jwt.MapClaims{"user_id": mockUserID},
	}, nil)
	suite.repository.On("FindById", mock.Anything, mockUserID).Return(mockUser, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)
	suite.jwtService.On("GenerateAccessTokenWithPayload", *mockUser).Return(mockAccessToken, nil)
	tokenMap, err := suite.usecase.RefreshToken(context.TODO(), mockRefreshToken)

	suite.NoError(err)
	suite.Equal(mockAccessToken, tokenMap["refresh_token"])

	suite.jwtService.AssertExpectations(suite.T())
	suite.repository.AssertExpectations(suite.T())
	suite.sessionRP.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestRefreshToken_InvalidToken() {
	mockRefreshToken := "invalidToken"

	mockError := &domain.CustomError{Message: "invalid token", Code: http.StatusUnauthorized}
	suite.jwtService.On("ValidateRefreshToken", mockRefreshToken).Return(nil, mockError)

	tokenMap, err := suite.usecase.RefreshToken(context.TODO(), mockRefreshToken)

	suite.Error(err)
	suite.EqualError(err, mockError.Error())
	suite.Empty(tokenMap)

	suite.jwtService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestRefreshToken_SessionNotFound() {
	mockUserID := "mockUserID"
	mockRefreshToken := "mockRefreshToken"
	mockUsername := "mockUsername"
	mockUser := &domain.User{ID: primitive.NilObjectID, Username: mockUsername}

	suite.jwtService.On("ValidateRefreshToken", mockRefreshToken).Return(&jwt.Token{
		Claims: jwt.MapClaims{"user_id": mockUserID},
	}, nil)

	suite.repository.On("FindById", mock.Anything, mockUserID).Return(mockUser, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(nil, false, nil)

	tokenMap, err := suite.usecase.RefreshToken(context.TODO(), mockRefreshToken)

	suite.Error(err)
	suite.Contains(err.Error(), "session not found")
	suite.Empty(tokenMap)

	suite.jwtService.AssertExpectations(suite.T())
	suite.repository.AssertExpectations(suite.T())
	suite.sessionRP.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestForgotPassword_Success() {
	mockEmail := "mock@example.com"
	mockUsername := "mockUsername"
	mockUser := &domain.User{ID: primitive.NewObjectID(), Email: mockEmail, Username: mockUsername}
	mockResetToken := "mockResetToken"
	mockSession := &domain.Session{Username: mockUsername, PasswordResetToken: mockResetToken}

	expectedResetURL := "http://localhost:8080/user/reset/%s"
	suite.repository.On("FindByEmail", mock.Anything, mockEmail).Return(mockUser, nil)
	suite.jwtService.On("GenerateResetToken", mockEmail).Return(mockResetToken, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)
	suite.sessionRP.On("UpdateToken", mock.Anything, mockSession.ID.Hex(), mock.Anything).Return(nil)
	suite.mailService.On("SendPasswordResetEmail", mockEmail, mockUsername, fmt.Sprintf(expectedResetURL, mockResetToken), mock.Anything).Return(nil)

	err := suite.usecase.ForgotPassword(context.TODO(), mockEmail)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())
	suite.utilsMock.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.sessionRP.AssertExpectations(suite.T())
	suite.mailService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestForgotPassword_UserNotFound() {
	mockEmail := "mock@example.com"

	mockError := &domain.CustomError{Message: "user not found", Code: http.StatusNotFound}
	suite.repository.On("FindByEmail", mock.Anything, mockEmail).Return(nil, mockError)

	err := suite.usecase.ForgotPassword(context.TODO(), mockEmail)

	suite.Error(err)
	suite.EqualError(err, mockError.Error())

	suite.repository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestForgotPassword_ErrorSendingEmail() {
	mockEmail := "mock@example.com"
	mockUsername := "mockUsername"
	mockUser := &domain.User{ID: primitive.NewObjectID(), Email: mockEmail, Username: mockUsername}
	mockResetToken := "mockResetToken"
	mockSession := &domain.Session{Username: mockUsername, PasswordResetToken: mockResetToken}
	expectedResetURL := "http://localhost:8080/user/reset/%s"

	suite.repository.On("FindByEmail", mock.Anything, mockEmail).Return(mockUser, nil)
	suite.jwtService.On("GenerateResetToken", mockEmail).Return(mockResetToken, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)
	suite.sessionRP.On("UpdateToken", mock.Anything, mockSession.ID.Hex(), mock.Anything).Return(nil)
	mockError := &domain.CustomError{Message: "failed to send email", Code: http.StatusInternalServerError}
	suite.mailService.On("SendPasswordResetEmail", mockEmail, mockUsername, fmt.Sprintf(expectedResetURL, mockResetToken), mock.Anything).Return(mockError)

	err := suite.usecase.ForgotPassword(context.TODO(), mockEmail)

	suite.Error(err)
	suite.EqualError(err, mockError.Error())

	suite.repository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.sessionRP.AssertExpectations(suite.T())
	suite.mailService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestResetPassword_Success() {
	mockEmail := "mock@example.com"
	mockUsername := "mockUsername"
	mockToken := "validToken"
	mockCode := 123456
	mockNewPassword := "newPassword"
	mockConfirmPassword := "newPassword"
	mockHashedPassword := "hashedPassword"

	mockUser := &domain.User{ID: primitive.NewObjectID(), Email: mockEmail, Username: mockUsername}
	mockSession := &domain.Session{
		Username:           mockUsername,
		PasswordResetToken: mockToken,
		ResetPasswordToken: mockCode,
	}

	suite.jwtService.On("ValidateResetToken", mockToken).Return(&jwt.Token{Claims: jwt.MapClaims{"email": mockEmail}}, nil)
	suite.repository.On("FindByEmail", mock.Anything, mockEmail).Return(mockUser, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)
	suite.passwordService.On("HashPassword", mockNewPassword).Return(mockHashedPassword, nil)
	suite.repository.On("UpdatePassword", mock.Anything, mockUser.ID.Hex(), mockHashedPassword).Return(nil)

	err := suite.usecase.ResetPassword(context.TODO(), mockNewPassword, mockConfirmPassword, mockToken, mockCode)

	suite.NoError(err)

	suite.jwtService.AssertExpectations(suite.T())
	suite.repository.AssertExpectations(suite.T())
	suite.sessionRP.AssertExpectations(suite.T())
	suite.passwordService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestResetPassword_InvalidToken() {
	mockToken := "invalidToken"

	suite.jwtService.On("ValidateResetToken", mockToken).Return(nil, domain.CustomError{Message: "invalid token", Code: http.StatusBadRequest})

	err := suite.usecase.ResetPassword(context.TODO(), "newPassword", "newPassword", mockToken, 123456)

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, err.StatusCode())

	suite.jwtService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestResetPassword_PasswordMismatch() {
	mockEmail := "mock@example.com"
	mockUsername := "mockUsername"
	mockToken := "validToken"
	mockCode := 123456
	mockNewPassword := "newPassword"
	mockConfirmPassword := "differentPassword"

	mockUser := &domain.User{ID: primitive.NewObjectID(), Email: mockEmail, Username: mockUsername}
	mockSession := &domain.Session{
		Username:           mockUsername,
		PasswordResetToken: mockToken,
		ResetPasswordToken: mockCode,
	}

	suite.jwtService.On("ValidateResetToken", mockToken).Return(&jwt.Token{Claims: jwt.MapClaims{"email": mockEmail}}, nil)
	suite.repository.On("FindByEmail", mock.Anything, mockEmail).Return(mockUser, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)

	err := suite.usecase.ResetPassword(context.TODO(), mockNewPassword, mockConfirmPassword, mockToken, mockCode)

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, err.(*domain.CustomError).Code)

	suite.jwtService.AssertExpectations(suite.T())
	suite.repository.AssertExpectations(suite.T())
	suite.sessionRP.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestResetPassword_InvalidTokenOrCode() {
	mockEmail := "mock@example.com"
	mockUsername := "mockUsername"
	mockToken := "validToken"
	mockCode := 123456
	mockNewPassword := "newPassword"
	mockConfirmPassword := "newPassword"

	mockUser := &domain.User{ID: primitive.NewObjectID(), Email: mockEmail, Username: mockUsername}
	mockSession := &domain.Session{
		Username:           mockUsername,
		PasswordResetToken: "differentToken",
		ResetPasswordToken: 654321,
	}

	suite.jwtService.On("ValidateResetToken", mockToken).Return(&jwt.Token{Claims: jwt.MapClaims{"email": mockEmail}}, nil)
	suite.repository.On("FindByEmail", mock.Anything, mockEmail).Return(mockUser, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)

	err := suite.usecase.ResetPassword(context.TODO(), mockNewPassword, mockConfirmPassword, mockToken, mockCode)

	suite.Error(err)
	suite.Equal(http.StatusUnauthorized, err.(*domain.CustomError).Code)

	suite.jwtService.AssertExpectations(suite.T())
	suite.repository.AssertExpectations(suite.T())
	suite.sessionRP.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestResetPassword_SessionNotFound() {
	mockEmail := "mock@example.com"
	mockUsername := "mockUsername"
	mockToken := "validToken"
	mockCode := 123456
	mockNewPassword := "newPassword"
	mockConfirmPassword := "newPassword"

	mockUser := &domain.User{ID: primitive.NewObjectID(), Email: mockEmail, Username: mockUsername}

	suite.jwtService.On("ValidateResetToken", mockToken).Return(&jwt.Token{Claims: jwt.MapClaims{"email": mockEmail}}, nil)
	suite.repository.On("FindByEmail", mock.Anything, mockEmail).Return(mockUser, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(nil, false, nil)

	err := suite.usecase.ResetPassword(context.TODO(), mockNewPassword, mockConfirmPassword, mockToken, mockCode)

	suite.Error(err)
	suite.Equal(http.StatusNotFound, err.(*domain.CustomError).Code)

	suite.jwtService.AssertExpectations(suite.T())
	suite.repository.AssertExpectations(suite.T())
	suite.sessionRP.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestResetPassword_HashingError() {
	mockEmail := "mock@example.com"
	mockUsername := "mockUsername"
	mockToken := "validToken"
	mockCode := 123456
	mockNewPassword := "newPassword"
	mockConfirmPassword := "newPassword"
	mockError := errors.New("hashing error")

	mockUser := &domain.User{ID: primitive.NewObjectID(), Email: mockEmail, Username: mockUsername}
	mockSession := &domain.Session{
		Username:           mockUsername,
		PasswordResetToken: mockToken,
		ResetPasswordToken: mockCode,
	}

	suite.jwtService.On("ValidateResetToken", mockToken).Return(&jwt.Token{Claims: jwt.MapClaims{"email": mockEmail}}, nil)
	suite.repository.On("FindByEmail", mock.Anything, mockEmail).Return(mockUser, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)
	suite.passwordService.On("HashPassword", mockNewPassword).Return("", mockError)

	err := suite.usecase.ResetPassword(context.TODO(), mockNewPassword, mockConfirmPassword, mockToken, mockCode)

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.(*domain.CustomError).Code)

	suite.jwtService.AssertExpectations(suite.T())
	suite.repository.AssertExpectations(suite.T())
	suite.sessionRP.AssertExpectations(suite.T())
	suite.passwordService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestLogout_Success() {
	mockAccessToken := "validAccessToken"
	mockRefreshToken := "validRefreshToken"
	mockUsername := "mockUsername"
	mockSession := &domain.Session{
		ID:                 primitive.NewObjectID(),
		Username:           mockUsername,
		RefreshToken:       mockRefreshToken,
		VerificationToken:  "",
		PasswordResetToken: "",
	}

	suite.redisService.On("Set", mockAccessToken, mockAccessToken, time.Minute*15).Return(nil)
	suite.jwtService.On("ValidateAccessToken", mockAccessToken).Return(&jwt.Token{Claims: jwt.MapClaims{"username": mockUsername}}, nil)
	suite.jwtService.On("ValidateRefreshToken", mockRefreshToken).Return(&jwt.Token{}, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)
	suite.sessionRP.On("UpdateToken", mock.Anything, mockSession.ID.Hex(), mockSession).Return(nil)

	err := suite.usecase.Logout(context.TODO(), map[string]string{
		"access_token":  mockAccessToken,
		"refresh_token": mockRefreshToken,
	})

	suite.NoError(err)

	suite.redisService.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.sessionRP.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestLogout_MissingAccessToken() {
	err := suite.usecase.Logout(context.TODO(), map[string]string{
		"refresh_token": "someRefreshToken",
	})

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, err.StatusCode())
	suite.Equal("access token not found", err.(*domain.CustomError).Message)
}

func (suite *UserUsecaseTest) TestLogout_MissingRefreshToken() {
	err := suite.usecase.Logout(context.TODO(), map[string]string{
		"access_token": "someAccessToken",
	})

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, err.StatusCode())
	suite.Equal("refresh token not found", err.Error())
}

func (suite *UserUsecaseTest) TestLogout_ErrorInvalidatingAccessToken() {
	mockAccessToken := "validAccessToken"
	mockRefreshToken := "validRefreshToken"

	suite.redisService.On("Set", mockAccessToken, mockAccessToken, time.Minute*15).Return(domain.CustomError{Message: "redis error", Code: http.StatusInternalServerError})

	err := suite.usecase.Logout(context.TODO(), map[string]string{
		"access_token":  mockAccessToken,
		"refresh_token": mockRefreshToken,
	})

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Equal("redis error", err.Error())
}

func (suite *UserUsecaseTest) TestLogout_InvalidAccessToken() {
	mockAccessToken := "invalidAccessToken"
	mockRefreshToken := "validRefreshToken"

	suite.redisService.On("Set", mockAccessToken, mockAccessToken, time.Minute*15).Return(nil)
	suite.jwtService.On("ValidateAccessToken", mockAccessToken).Return(nil, domain.CustomError{Message: "invalid token", Code: http.StatusUnauthorized})

	err := suite.usecase.Logout(context.TODO(), map[string]string{
		"access_token":  mockAccessToken,
		"refresh_token": mockRefreshToken,
	})

	suite.Error(err)
	suite.Equal(http.StatusUnauthorized, err.StatusCode())
	suite.Equal("invalid token", err.Error())
}

func (suite *UserUsecaseTest) TestLogout_InvalidRefreshToken() {
	mockAccessToken := "validAccessToken"
	mockRefreshToken := "invalidRefreshToken"

	suite.redisService.On("Set", mockAccessToken, mockAccessToken, time.Minute*15).Return(nil)
	suite.jwtService.On("ValidateAccessToken", mockAccessToken).Return(&jwt.Token{Claims: jwt.MapClaims{"username": "mockUsername"}}, nil)
	suite.jwtService.On("ValidateRefreshToken", mockRefreshToken).Return(nil, domain.CustomError{Message: "invalid token", Code: http.StatusUnauthorized})

	err := suite.usecase.Logout(context.TODO(), map[string]string{
		"access_token":  mockAccessToken,
		"refresh_token": mockRefreshToken,
	})

	suite.Error(err)
	suite.Equal(http.StatusUnauthorized, err.StatusCode())
	suite.Equal("invalid token", err.Error())
}

func (suite *UserUsecaseTest) TestLogout_SessionNotFound() {
	mockAccessToken := "validAccessToken"
	mockRefreshToken := "validRefreshToken"
	mockUsername := "mockUsername"

	suite.redisService.On("Set", mockAccessToken, mockAccessToken, time.Minute*15).Return(nil)
	suite.jwtService.On("ValidateAccessToken", mockAccessToken).Return(&jwt.Token{Claims: jwt.MapClaims{"username": mockUsername}}, nil)
	suite.jwtService.On("ValidateRefreshToken", mockRefreshToken).Return(&jwt.Token{}, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(nil, false, nil)

	err := suite.usecase.Logout(context.TODO(), map[string]string{
		"access_token":  mockAccessToken,
		"refresh_token": mockRefreshToken,
	})

	suite.Error(err)
	suite.Equal(http.StatusNotFound, err.StatusCode())
	suite.Equal("session not found", err.Error())
}

func (suite *UserUsecaseTest) TestLogout_ErrorRetrievingSession() {
	mockAccessToken := "validAccessToken"
	mockRefreshToken := "validRefreshToken"
	mockUsername := "mockUsername"

	suite.redisService.On("Set", mockAccessToken, mockAccessToken, time.Minute*15).Return(nil)
	suite.jwtService.On("ValidateAccessToken", mockAccessToken).Return(&jwt.Token{Claims: jwt.MapClaims{"username": mockUsername}}, nil)
	suite.jwtService.On("ValidateRefreshToken", mockRefreshToken).Return(&jwt.Token{}, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(nil, false, domain.CustomError{Message: "retrieval error", Code: http.StatusInternalServerError})

	err := suite.usecase.Logout(context.TODO(), map[string]string{
		"access_token":  mockAccessToken,
		"refresh_token": mockRefreshToken,
	})

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Equal("retrieval error", err.Error())
}

func (suite *UserUsecaseTest) TestLogout_ErrorUpdatingSession() {
	mockAccessToken := "validAccessToken"
	mockRefreshToken := "validRefreshToken"
	mockUsername := "mockUsername"
	mockSession := &domain.Session{
		ID:                 primitive.NewObjectID(),
		Username:           mockUsername,
		RefreshToken:       mockRefreshToken,
		VerificationToken:  "",
		PasswordResetToken: "",
	}

	suite.redisService.On("Set", mockAccessToken, mockAccessToken, time.Minute*15).Return(nil)
	suite.jwtService.On("ValidateAccessToken", mockAccessToken).Return(&jwt.Token{Claims: jwt.MapClaims{"username": mockUsername}}, nil)
	suite.jwtService.On("ValidateRefreshToken", mockRefreshToken).Return(&jwt.Token{}, nil)
	suite.sessionRP.On("FindTokenByUserUsername", mock.Anything, mockUsername).Return(mockSession, true, nil)
	suite.sessionRP.On("UpdateToken", mock.Anything, mockSession.ID.Hex(), mockSession).Return(domain.CustomError{Message: "error deleting", Code: http.StatusInternalServerError})

	err := suite.usecase.Logout(context.TODO(), map[string]string{
		"access_token":  mockAccessToken,
		"refresh_token": mockRefreshToken,
	})

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Equal("error deleting", err.Error())
}

func (suite *UserUsecaseTest) TestPromoteUser_Success() {
	userID := "someUserID"
	suite.repository.On("UpdateRole", mock.Anything, userID, "admin").Return(nil)

	err := suite.usecase.PromoteUser(context.TODO(), userID)

	suite.NoError(err)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestPromoteUser_ErrorUpdatingRole() {
	userID := "someUserID"
	suite.repository.On("UpdateRole", mock.Anything, userID, "admin").Return(domain.CustomError{Message: "error updating role", Code: http.StatusInternalServerError})

	err := suite.usecase.PromoteUser(context.TODO(), userID)

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Equal("error updating role", err.Error())
}

func (suite *UserUsecaseTest) TestDemoteUser_Success() {
	userID := "someUserID"
	suite.repository.On("UpdateRole", mock.Anything, userID, "user").Return(nil)

	err := suite.usecase.DemoteUser(context.TODO(), userID)

	suite.NoError(err)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestDemoteUser_ErrorUpdatingRole() {
	userID := "someUserID"
	suite.repository.On("UpdateRole", mock.Anything, userID, "user").Return(domain.CustomError{Message: "error updating role", Code: http.StatusInternalServerError})

	err := suite.usecase.DemoteUser(context.TODO(), userID)

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Equal("error updating role", err.Error())
}

func (suite *UserUsecaseTest) TestUpdateProfile_Success() {
	userID := "someUserID"
	user := map[string]interface{}{
		"username": "newUsername",
		"email":    "newEmail@example.com",
		"bio":      "Some bio",
	}

	suite.repository.On("CheckExistence", mock.Anything, userID).Return(1, nil)
	suite.repository.On("CountByUsername", mock.Anything, "newUsername").Return(0, nil)
	suite.repository.On("CountByEmail", mock.Anything, "newEmail@example.com").Return(0, nil)
	suite.repository.On("UpdateProfile", mock.Anything, userID, user).Return(nil)

	err := suite.usecase.UpdateProfile(context.TODO(), userID, user)

	suite.NoError(err)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestUpdateProfile_UserNotFound() {
	userID := "someUserID"
	user := map[string]interface{}{
		"username": "newUsername",
		"email":    "newEmail@example.com",
		"bio":      "Some bio",
	}

	suite.repository.On("CheckExistence", mock.Anything, userID).Return(0, nil)

	err := suite.usecase.UpdateProfile(context.TODO(), userID, user)

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, err.StatusCode())
	suite.Equal("user by the given user id doesn't exist", err.Error())
}

func (suite *UserUsecaseTest) TestUpdateProfile_UsernameAlreadyTaken() {
	userID := "someUserID"
	user := map[string]interface{}{
		"username": "takenUsername",
		"email":    "newEmail@example.com",
		"bio":      "Some bio",
	}

	suite.repository.On("CheckExistence", mock.Anything, userID).Return(1, nil)
	suite.repository.On("CountByUsername", mock.Anything, "takenUsername").Return(1, nil)

	err := suite.usecase.UpdateProfile(context.TODO(), userID, user)

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, err.StatusCode())
	suite.Equal("Username already taken", err.Error())
}

func (suite *UserUsecaseTest) TestUpdateProfile_EmailAlreadyTaken() {
	userID := "someUserID"
	user := map[string]interface{}{
		"username": "newUsername",
		"email":    "takenEmail@example.com",
		"bio":      "Some bio",
	}

	suite.repository.On("CheckExistence", mock.Anything, userID).Return(1, nil)
	suite.repository.On("CountByUsername", mock.Anything, "newUsername").Return(0, nil)
	suite.repository.On("CountByEmail", mock.Anything, "takenEmail@example.com").Return(1, nil)

	err := suite.usecase.UpdateProfile(context.TODO(), userID, user)

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, err.StatusCode())
	suite.Equal("Email already taken", err.Error())
}

func (suite *UserUsecaseTest) TestUpdateProfile_ErrorUpdatingProfile() {
	userID := "someUserID"
	user := map[string]interface{}{
		"username": "newUsername",
		"email":    "newEmail@example.com",
		"bio":      "Some bio",
	}

	suite.repository.On("CheckExistence", mock.Anything, userID).Return(1, nil)
	suite.repository.On("CountByUsername", mock.Anything, "newUsername").Return(0, nil)
	suite.repository.On("CountByEmail", mock.Anything, "newEmail@example.com").Return(0, nil)
	suite.repository.On("UpdateProfile", mock.Anything, userID, user).Return(domain.CustomError{Message: "update error", Code: http.StatusInternalServerError})

	err := suite.usecase.UpdateProfile(context.TODO(), userID, user)

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Equal("update error", err.Error())
}

func (suite *UserUsecaseTest) TestImageUpload_Success() {
	file := new(multipart.File)
	header := &multipart.FileHeader{Filename: "test_image.png"}
	id := "someUserID"

	suite.utilsMock.On("IsValidFileFormat", header, "image/png", "image/jpeg").Return(true)
	suite.repository.On("FindById", mock.Anything, id).Return(&domain.User{}, nil)
	suite.utilsMock.On("SaveImage", *file, header.Filename, mock.Anything).Return(&uploader.UploadResult{
		SecureURL: "https://example.com/image.jpg",
		PublicID:  "somePublicID",
	}, nil)
	suite.repository.On("UploadProfilePicture", mock.Anything, mock.Anything, id).Return(nil)

	err := suite.usecase.ImageUpload(context.TODO(), file, header, id)

	suite.NoError(err)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTest) TestImageUpload_InvalidFileFormat() {
	file := new(multipart.File)
	header := &multipart.FileHeader{Filename: "image.txt"}
	id := "someUserID"

	suite.utilsMock.On("IsValidFileFormat", header, "image/png", "image/jpeg").Return(false)

	err := suite.usecase.ImageUpload(context.TODO(), file, header, id)

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, err.StatusCode())
	suite.Equal("invalid file format", err.Error())
}

func (suite *UserUsecaseTest) TestImageUpload_ErrorFindingUser() {
	file := new(multipart.File)
	header := &multipart.FileHeader{Filename: "image.jpg"}
	id := "someUserID"

	suite.utilsMock.On("IsValidFileFormat", header, "image/png", "image/jpeg").Return(true)
	suite.repository.On("FindById", mock.Anything, id).Return(nil, domain.CustomError{Message: "user not found", Code: http.StatusNotFound})

	err := suite.usecase.ImageUpload(context.TODO(), file, header, id)

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Equal("user not found", err.Error())
}

func (suite *UserUsecaseTest) TestImageUpload_ErrorDeletingExistingImage() {
	file := new(multipart.File)
	header := &multipart.FileHeader{Filename: "image.jpg"}
	id := "someUserID"

	existingUser := &domain.User{
		ProfilePicture: domain.Photo{
			FilePath:  "https://example.com/old_image.jpg",
			Public_id: "oldPublicID",
		},
	}

	suite.utilsMock.On("IsValidFileFormat", header, "image/png", "image/jpeg").Return(true)
	suite.repository.On("FindById", mock.Anything, id).Return(existingUser, nil)
	suite.utilsMock.On("DeleteImage", existingUser.ProfilePicture.Public_id, mock.Anything).Return(domain.CustomError{Message: "delete error", Code: http.StatusInternalServerError})

	err := suite.usecase.ImageUpload(context.TODO(), file, header, id)

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Equal("delete error", err.Error())
}

func (suite *UserUsecaseTest) TestImageUpload_ErrorSavingImage() {
	file := new(multipart.File)
	header := &multipart.FileHeader{Filename: "image.jpg"}
	id := "someUserID"

	suite.utilsMock.On("IsValidFileFormat", header, "image/png", "image/jpeg").Return(true)
	suite.repository.On("FindById", mock.Anything, id).Return(&domain.User{}, nil)
	suite.utilsMock.On("SaveImage", *file, header.Filename, mock.Anything).Return(nil, domain.CustomError{Message: "save error", Code: http.StatusInternalServerError})

	err := suite.usecase.ImageUpload(context.TODO(), file, header, id)

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Equal("save error", err.Error())
}

func (suite *UserUsecaseTest) TestImageUpload_ErrorUpdatingProfilePicture() {
	file := new(multipart.File)
	header := &multipart.FileHeader{Filename: "image.jpg"}
	id := "someUserID"

	suite.utilsMock.On("IsValidFileFormat", header, "image/png", "image/jpeg").Return(true)
	suite.repository.On("FindById", mock.Anything, id).Return(&domain.User{}, nil)
	suite.utilsMock.On("SaveImage", *file, header.Filename, mock.Anything).Return(&uploader.UploadResult{
		SecureURL: "https://example.com/image.jpg",
		PublicID:  "somePublicID",
	}, nil)
	suite.repository.On("UploadProfilePicture", mock.Anything, mock.Anything, id).Return(domain.CustomError{Message: "update error", Code: http.StatusInternalServerError})

	err := suite.usecase.ImageUpload(context.TODO(), file, header, id)

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.StatusCode())
	suite.Equal("update error", err.Error())
}

func TestUserUsecaseTest(t *testing.T) {
	errLoad := godotenv.Load("../.env")
	if errLoad != nil {
		fmt.Println("error loading the env file")
	}
	suite.Run(t, new(UserUsecaseTest))
}

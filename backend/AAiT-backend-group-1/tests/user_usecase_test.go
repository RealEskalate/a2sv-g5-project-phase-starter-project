package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/mocks"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/usecases"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTest struct {
	suite.Suite
	usecase         domain.UserUseCase
	repository      *mocks.UserRepository
	passwordService *mocks.PasswordService
	jwtService      *mocks.JwtService
	mailService     *mocks.EmailService
	redisService    *mocks.MockCacheService
	sessionRP       *mocks.SessionRepository
}

func (suite *UserUsecaseTest) SetupSuite() {
	userRP := mocks.UserRepository{}
	sessionRP := mocks.SessionRepository{}
	passwordService := mocks.PasswordService{}
	jwtService := mocks.JwtService{}
	mailService := mocks.EmailService{}
	redisService := mocks.MockCacheService{}
	suite.usecase = usecases.NewUserUseCase(&userRP, &sessionRP, &passwordService, &jwtService, &mailService, &redisService)
	suite.repository = &userRP
	suite.passwordService = &passwordService
	suite.jwtService = &jwtService
	suite.mailService = &mailService
	suite.redisService = &redisService
	suite.sessionRP = &sessionRP
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

func TestUserUsecaseTest(t *testing.T) {
	errLoad := godotenv.Load("../.env")
	if errLoad != nil {
		fmt.Println("error loading the env file")
	}
	suite.Run(t, new(UserUsecaseTest))
}

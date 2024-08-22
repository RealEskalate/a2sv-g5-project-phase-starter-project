package tests

import (
	"blog_api/domain"
	"blog_api/mocks"
	"blog_api/usecase"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	VALID_GOOGLE_TOKEN = "valid_google_token"
	VALID_DELETED_FILE = "valid_deleted_file.png"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	Usecase             usecase.UserUsecase
	mockUserRepository  domain.UserRepositoryInterface
	mockCacheRepository domain.CacheRepositoryInterface
	mockMailService     domain.MailServiceInterface
	mockHashService     domain.HashingServiceInterface
	mockJWTService      domain.JWTServiceInterface
	GenerateToken       func(int) (string, error)
	VerifyIdToken       func(string, string, string) error
	DeleteFile          func(string) error
	ENV                 domain.EnvironmentVariables
}

func (suite *UserUsecaseTestSuite) SetupSuite() {
	suite.mockUserRepository = new(mocks.UserRepositoryInterface)
	suite.mockCacheRepository = new(mocks.CacheRepositoryInterface)
	suite.mockMailService = new(mocks.MailServiceInterface)
	suite.mockHashService = new(mocks.HashingServiceInterface)
	suite.mockJWTService = new(mocks.JWTServiceInterface)

	suite.GenerateToken = func(num int) (string, error) {
		return "__generated_token__" + fmt.Sprint(num), nil
	}

	suite.VerifyIdToken = func(token string, apiKey string, projectID string) error {
		if token != VALID_GOOGLE_TOKEN {
			return fmt.Errorf("invalid token")
		}

		return nil
	}

	suite.DeleteFile = func(filename string) error {
		if filename == VALID_DELETED_FILE {
			return nil
		}

		return fmt.Errorf("file not found")
	}

	suite.Usecase = *usecase.NewUserUsecase(
		suite.mockUserRepository,
		suite.mockCacheRepository,
		suite.GenerateToken,
		suite.mockMailService,
		suite.mockJWTService,
		suite.mockHashService,
		suite.VerifyIdToken,
		suite.DeleteFile,
		suite.ENV,
	)
}

// Reset the environment variables before each test
func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.ENV = domain.EnvironmentVariables{}
}

func (suite *UserUsecaseTestSuite) TestValidatePassword_Positive() {
	testPw := "cR@zyP@ssw0rd"
	err := suite.Usecase.ValidatePassword(testPw)
	suite.Nil(err)

}

func (suite *UserUsecaseTestSuite) TestValidatePassword_Negative_TooShort() {
	testPw := "cR@0"
	err := suite.Usecase.ValidatePassword(testPw)
	suite.NotNil(err)
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "short")
}

// checked because there is cap on how long the password can be for the bcrypt library to hash it
func (suite *UserUsecaseTestSuite) TestValidatePassword_Negative_TooLong() {
	testPw := strings.Repeat("a", 100)
	err := suite.Usecase.ValidatePassword(testPw)
	suite.NotNil(err)
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "long")
}

func (suite *UserUsecaseTestSuite) TestValidatePassword_Negative_NoLowerCase() {
	testPw := "CR@ZYP@SSW0RD"
	err := suite.Usecase.ValidatePassword(testPw)
	suite.NotNil(err)
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "lower")
}

func (suite *UserUsecaseTestSuite) TestValidatePassword_Negative_NoUpperCase() {
	testPw := "cr@zyp@ssw0rd"
	err := suite.Usecase.ValidatePassword(testPw)
	suite.NotNil(err)
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "upper")
}

func (suite *UserUsecaseTestSuite) TestValidatePassword_Negative_NoNumbers() {
	testPw := "cr@zyp@ssWOrd"
	err := suite.Usecase.ValidatePassword(testPw)
	suite.NotNil(err)
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "number")
}

func (suite *UserUsecaseTestSuite) TestValidatePassword_Negative_NoSpecialCharacters() {
	testPw := "crazypassW0rd"
	err := suite.Usecase.ValidatePassword(testPw)
	suite.NotNil(err)
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "special")
}

func (suite *UserUsecaseTestSuite) TestValidateUsername_Positive() {
	testUsername := "12timid_"
	err := suite.Usecase.ValidateUsername(testUsername)
	suite.Nil(err)
}

func (suite *UserUsecaseTestSuite) TestValidateUsername_Negative_TooShort() {
	testUsername := "ti"
	err := suite.Usecase.ValidateUsername(testUsername)
	suite.NotNil(err)
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "short")
}

func (suite *UserUsecaseTestSuite) TestValidateUsername_Negative_TooLong() {
	testUsername := strings.Repeat("a", 100)
	err := suite.Usecase.ValidateUsername(testUsername)
	suite.NotNil(err)
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "long")
}

func (suite *UserUsecaseTestSuite) TestValidateUsername_Negative_NoSpecialCharacters() {
	testUsername := "timid-"
	err := suite.Usecase.ValidateUsername(testUsername)
	suite.NotNil(err)
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "must contain only")
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

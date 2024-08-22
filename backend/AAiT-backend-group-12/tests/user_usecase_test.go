package tests

import (
	"blog_api/domain"
	"blog_api/domain/dtos"
	"blog_api/mocks"
	"blog_api/usecase"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_GOOGLE_TOKEN = "valid_google_token"
	VALID_DELETED_FILE = "valid_deleted_file.png"
)

var TEST_USER = domain.User{
	Username:    "  timid_  ",
	Email:       " timid_  @  gmail.com         ",
	Bio:         "   cartifan20 ",
	Password:    "  cR@zyP@ssw0rd  ",
	PhoneNumber: " +256 6 45 2 10  21         ",
}

var TEST_GOOGLE_RES = dtos.GoogleResponse{
	RawData: struct {
		Email         string `json:"email" binding:"required"`
		ID            string `json:"id" binding:"required"`
		Picture       string `json:"picture"`
		VerifiedEmail bool   `json:"verified_email" binding:"required"`
	}{
		Email:         TEST_USER.Email,
		ID:            "google_id",
		Picture:       "google_picture",
		VerifiedEmail: true,
	},
	Provider:     "google",
	Email:        TEST_USER.Email,
	UserID:       "google_id",
	AccessToken:  "google_access_token",
	RefreshToken: "google_refresh_token",
	ExpiresAt:    "google_expires_at",
	IDToken:      VALID_GOOGLE_TOKEN,
}

type UserUsecaseTestSuite struct {
	suite.Suite
	Usecase             usecase.UserUsecase
	mockUserRepository  *mocks.UserRepositoryInterface
	mockCacheRepository *mocks.CacheRepositoryInterface
	mockMailService     *mocks.MailServiceInterface
	mockHashService     *mocks.HashingServiceInterface
	mockJWTService      *mocks.JWTServiceInterface
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
	suite.mockCacheRepository.ExpectedCalls = []*mock.Call{}
	suite.mockHashService.ExpectedCalls = []*mock.Call{}
	suite.mockMailService.ExpectedCalls = []*mock.Call{}
	suite.mockUserRepository.ExpectedCalls = []*mock.Call{}
	suite.mockJWTService.ExpectedCalls = []*mock.Call{}
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

func (suite *UserUsecaseTestSuite) TestValidateEmail_Positive() {
	testEmail := "timid_@gmail.com"
	err := suite.Usecase.ValidateEmail(testEmail)
	suite.Nil(err)
}

func (suite *UserUsecaseTestSuite) TestValidateEmail_Negative_InvalidEmail() {
	testEmail := "timid_gmail.com"
	err := suite.Usecase.ValidateEmail(testEmail)
	suite.NotNil(err)
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "Email")
}

func (suite *UserUsecaseTestSuite) TestSanitizeUserFields() {
	user := TEST_USER
	oldPwd := user.Password

	suite.Usecase.SantizeUserFields(&user)
	suite.Equal(user.Username, "timid_")
	suite.Equal(user.Email, "timid_@gmail.com")
	suite.Equal(user.Bio, "cartifan20")
	suite.Equal(user.PhoneNumber, "+25664521021")
	suite.Equal(oldPwd, user.Password)
}

func (suite *UserUsecaseTestSuite) TestSanitizeAndValidateNewUser_Positive() {
	user := TEST_USER
	oldPwd := TEST_USER.Password

	err := suite.Usecase.SanitizeAndValidateNewUser(&user)
	suite.Nil(err, "error should be nil")
	suite.Equal(user.Username, "timid_")
	suite.Equal(user.Email, "timid_@gmail.com")
	suite.Equal(user.Bio, "cartifan20")
	suite.Equal(user.PhoneNumber, "+25664521021")
	suite.Equal(oldPwd, user.Password)
}

func (suite *UserUsecaseTestSuite) TestSanitizeAndValidateNewUser_Negative_InvalidUsername() {
	user := TEST_USER
	user.Username = "ti"

	err := suite.Usecase.SanitizeAndValidateNewUser(&user)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "Username")
}

func (suite *UserUsecaseTestSuite) TestSanitizeAndValidateNewUser_Negative_InvalidEmail() {
	user := TEST_USER
	user.Email = "ti"

	err := suite.Usecase.SanitizeAndValidateNewUser(&user)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "Email")
}

func (suite *UserUsecaseTestSuite) TestSanitizeAndValidateNewUser_Negative_InvalidPassword() {
	user := TEST_USER
	user.Password = "ti"

	err := suite.Usecase.SanitizeAndValidateNewUser(&user)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "Password")
}

func (suite *UserUsecaseTestSuite) TestSanitizeAndValidateNewUser_Negative_InvalidBio() {
	user := TEST_USER
	user.Bio = ""

	err := suite.Usecase.SanitizeAndValidateNewUser(&user)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "Bio")
}

func (suite *UserUsecaseTestSuite) TestSanitizeAndValidateNewUser_Negative_InvalidPhoneNumber() {
	user := TEST_USER
	user.PhoneNumber = ""

	err := suite.Usecase.SanitizeAndValidateNewUser(&user)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Contains(err.Error(), "PhoneNumber")
}

func (suite *UserUsecaseTestSuite) TestGetVerificationData_Positive() {
	verificationType := "verification_type"
	expiresAt := time.Now().Round(0).Add(time.Hour)
	tokenLength := 10

	verificationData, err := suite.Usecase.GetVerificationData(context.Background(), verificationType, expiresAt, tokenLength)
	suite.Nil(err)
	suite.Equal(verificationData.Type, verificationType)
	suite.Equal(verificationData.ExpiresAt, expiresAt)
	suite.LessOrEqual(0, len(verificationData.Token))
}

func (suite *UserUsecaseTestSuite) TestSignup_Positive() {
	user := TEST_USER

	suite.Usecase.SanitizeAndValidateNewUser(&user)
	mail := "mail_template"
	hostUrl := "host_url"
	suite.mockHashService.On("HashString", user.Password).Return("hashed_str", nil).Once()
	suite.mockUserRepository.On("CreateUser", context.Background(), mock.AnythingOfType("*domain.User")).Return(nil).Once()
	genToken, _ := suite.GenerateToken(32)
	suite.mockMailService.On("EmailVerificationTemplate", mock.AnythingOfType("string"), user.Username, genToken).Return(mail).Once()
	suite.mockMailService.On("SendMail", mock.AnythingOfType("string"), user.Email, mail).Return(nil).Once()

	err := suite.Usecase.Signup(context.Background(), &user, hostUrl)
	suite.Nil(err)

	// check the updated values
	suite.False(user.IsVerified)
	suite.NotEqual("", user.VerificationData.Token)
	suite.GreaterOrEqual(user.CreatedAt, time.Now().Add(-1*time.Minute))
	suite.GreaterOrEqual(user.VerificationData.ExpiresAt, time.Now())
	suite.NotEqual(user.Password, TEST_USER.Password)
	suite.Equal(user.Role, domain.RoleUser)

	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockMailService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestSignup_Negative_HashError() {
	user := TEST_USER

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	hostUrl := "host_url"
	suite.mockHashService.On("HashString", user.Password).Return("hashed_str", sampleErr).Once()

	uErr := suite.Usecase.Signup(context.Background(), &user, hostUrl)

	suite.NotNil(uErr, "error during hash")
	suite.Equal(uErr.GetCode(), domain.ERR_INTERNAL_SERVER)
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestSignup_Negative_RepositoryErr() {
	user := TEST_USER

	hostUrl := "host_url"
	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockHashService.On("HashString", user.Password).Return("hashed_str", nil).Once()
	suite.mockUserRepository.On("CreateUser", context.Background(), mock.AnythingOfType("*domain.User")).Return(sampleErr).Once()

	err := suite.Usecase.Signup(context.Background(), &user, hostUrl)
	suite.NotNil(err, "error during repository call")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())

	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestSignup_Negative_MailError() {
	user := TEST_USER

	suite.Usecase.SanitizeAndValidateNewUser(&user)
	mail := "mail_template"
	hostUrl := "host_url"
	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockHashService.On("HashString", user.Password).Return("hashed_str", nil).Once()
	suite.mockUserRepository.On("CreateUser", context.Background(), mock.AnythingOfType("*domain.User")).Return(nil).Once()
	suite.mockUserRepository.On("DeleteUser", context.Background(), user.Username).Return(nil).Once()

	genToken, _ := suite.GenerateToken(32)
	suite.mockMailService.On("EmailVerificationTemplate", mock.AnythingOfType("string"), user.Username, genToken).Return(mail).Once()
	suite.mockMailService.On("SendMail", mock.AnythingOfType("string"), user.Email, mail).Return(sampleErr).Once()

	err := suite.Usecase.Signup(context.Background(), &user, hostUrl)
	suite.NotNil(err, "error during mail send")
	suite.Equal(err.GetCode(), domain.ERR_INTERNAL_SERVER)

	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockMailService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestOAuthSignup_Positive() {
	user := TEST_USER
	google_res := TEST_GOOGLE_RES
	suite.Usecase.SantizeUserFields(&user)
	oauthDto := dtos.OAuthSignup{
		Username: user.Username,
		Password: user.Password,
	}

	suite.mockHashService.On("HashString", user.Password).Return("hashed_str", nil).Once()
	suite.mockUserRepository.On("CreateUser", context.Background(), mock.AnythingOfType("*domain.User")).Return(nil).Once()
	suite.mockUserRepository.On("VerifyUser", context.Background(), user.Username).Return(nil).Once()

	err := suite.Usecase.OAuthSignup(context.Background(), &google_res, &oauthDto)
	suite.Nil(err, "error should be nil")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestOAuthSignup_Negative_VerificationError() {
	user := TEST_USER
	google_res := TEST_GOOGLE_RES
	suite.Usecase.SantizeUserFields(&user)
	oauthDto := dtos.OAuthSignup{
		Username: user.Username,
		Password: user.Password,
	}

	google_res.IDToken = "random"

	err := suite.Usecase.OAuthSignup(context.Background(), &google_res, &oauthDto)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestOAuthSignup_Negative_HashError() {
	user := TEST_USER
	google_res := TEST_GOOGLE_RES
	suite.Usecase.SantizeUserFields(&user)
	oauthDto := dtos.OAuthSignup{
		Username: user.Username,
		Password: user.Password,
	}

	suite.mockHashService.On("HashString", user.Password).Return("hashed_str", domain.NewError("", domain.ERR_INTERNAL_SERVER)).Once()

	err := suite.Usecase.OAuthSignup(context.Background(), &google_res, &oauthDto)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_INTERNAL_SERVER)
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}
func (suite *UserUsecaseTestSuite) TestOAuthSignup_RepositoryError() {
	user := TEST_USER
	google_res := TEST_GOOGLE_RES
	suite.Usecase.SantizeUserFields(&user)
	oauthDto := dtos.OAuthSignup{
		Username: user.Username,
		Password: user.Password,
	}

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockHashService.On("HashString", user.Password).Return("hashed_str", nil).Once()
	suite.mockUserRepository.On("CreateUser", context.Background(), mock.AnythingOfType("*domain.User")).Return(sampleErr).Once()

	err := suite.Usecase.OAuthSignup(context.Background(), &google_res, &oauthDto)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}

// func (suite *UserUsecaseTestSuite) TestLogin_Positive() {
// 	user := TEST_USER
// 	user.IsVerified = true
// 	suite.Usecase.SantizeUserFields(&user)

// 	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{
// 		Username: user.Username,
// 		Email:    user.Email,
// 	}).Return(&user, nil).Once()

// 	suite.mockUserRepository.AssertExpectations(suite.T())
// }

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

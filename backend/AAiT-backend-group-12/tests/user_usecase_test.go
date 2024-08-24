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

	"github.com/golang-jwt/jwt"
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
	Role:        domain.RoleUser,
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
	ENV                 *domain.EnvironmentVariables
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

	suite.ENV = &domain.EnvironmentVariables{
		ACCESS_TOKEN_LIFESPAN:  12,
		REFRESH_TOKEN_LIFESPAN: 13,
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
		*suite.ENV,
	)
}

// Reset the environment variables before each test
func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.mockCacheRepository.ExpectedCalls = []*mock.Call{}
	suite.mockHashService.ExpectedCalls = []*mock.Call{}
	suite.mockMailService.ExpectedCalls = []*mock.Call{}
	suite.mockUserRepository.ExpectedCalls = []*mock.Call{}
	suite.mockJWTService.ExpectedCalls = []*mock.Call{}
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

func (suite *UserUsecaseTestSuite) TestLogin_Positive() {
	user := TEST_USER
	user.IsVerified = true
	suite.Usecase.SantizeUserFields(&user)

	suite.mockUserRepository.On("FindUser", context.Background(), &user).Return(user, nil).Once()
	suite.mockHashService.On("ValidateHashedString", user.Password, user.Password).Return(nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "accessToken", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return("acc.ess.Token", nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "refreshToken", time.Hour*time.Duration(suite.ENV.REFRESH_TOKEN_LIFESPAN)).Return("ref.reshT.oken", nil).Once()
	suite.mockHashService.On("HashString", "oken").Return("hashed_str", nil).Once()
	suite.mockUserRepository.On("SetRefreshToken", context.Background(), &user, "hashed_str").Return(nil).Once()

	ack, rfk, err := suite.Usecase.Login(context.Background(), &user)
	suite.Nil(err, "error should be nil")
	suite.Equal(ack, "acc.ess.Token")
	suite.Equal(rfk, "ref.reshT.oken")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLogin_Negative_RepositoryError_Find() {
	user := TEST_USER
	user.IsVerified = true
	suite.Usecase.SantizeUserFields(&user)

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &user).Return(user, sampleErr).Once()

	ack, rfk, err := suite.Usecase.Login(context.Background(), &user)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.Equal(ack, "")
	suite.Equal(rfk, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLogin_Negative_HashError() {
	user := TEST_USER
	user.IsVerified = true
	suite.Usecase.SantizeUserFields(&user)

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &user).Return(user, nil).Once()
	suite.mockHashService.On("ValidateHashedString", user.Password, user.Password).Return(sampleErr).Once()

	ack, rfk, err := suite.Usecase.Login(context.Background(), &user)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.Equal(ack, "")
	suite.Equal(rfk, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLogin_Negative_JWTError() {
	user := TEST_USER
	user.IsVerified = true
	suite.Usecase.SantizeUserFields(&user)

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &user).Return(user, nil).Once()
	suite.mockHashService.On("ValidateHashedString", user.Password, user.Password).Return(nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "accessToken", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return("acc.ess.Token", sampleErr).Once()

	ack, rfk, err := suite.Usecase.Login(context.Background(), &user)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.Equal(ack, "")
	suite.Equal(rfk, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLogin_Negative_RepositoryError_SetRefreshToken() {
	user := TEST_USER
	user.IsVerified = true
	suite.Usecase.SantizeUserFields(&user)

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &user).Return(user, nil).Once()
	suite.mockHashService.On("ValidateHashedString", user.Password, user.Password).Return(nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "accessToken", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return("acc.ess.Token", nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "refreshToken", time.Hour*time.Duration(suite.ENV.REFRESH_TOKEN_LIFESPAN)).Return("ref.reshT.oken", nil).Once()
	suite.mockHashService.On("HashString", "oken").Return("hashed_str", nil).Once()
	suite.mockUserRepository.On("SetRefreshToken", context.Background(), &user, "hashed_str").Return(sampleErr).Once()

	ack, rfk, err := suite.Usecase.Login(context.Background(), &user)
	suite.NotNil(err, "error should be not nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.Equal(ack, "")
	suite.Equal(rfk, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLogin_Negative_NoEmailAndUsername() {
	ack, rfk, err := suite.Usecase.Login(context.Background(), &domain.User{})

	suite.NotNil(err, "error should be not nil")
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Equal(ack, "")
	suite.Equal(rfk, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestOAuthLogin_Positive() {
	google_data := TEST_GOOGLE_RES
	user := TEST_USER
	user.IsVerified = true

	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Email: google_data.Email}).Return(user, nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "accessToken", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return("acc.ess.Token", nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "refreshToken", time.Hour*time.Duration(suite.ENV.REFRESH_TOKEN_LIFESPAN)).Return("ref.reshT.oken", nil).Once()
	suite.mockHashService.On("HashString", "oken").Return("hashed_str", nil).Once()
	suite.mockUserRepository.On("SetRefreshToken", context.Background(), &user, "hashed_str").Return(nil).Once()

	ack, rfk, err := suite.Usecase.OAuthLogin(context.Background(), &google_data)
	suite.Nil(err, "error should be nil")
	suite.Equal(ack, "acc.ess.Token")
	suite.Equal(rfk, "ref.reshT.oken")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestOAuthLogin_Negative_RepositoryError_Find() {
	google_data := TEST_GOOGLE_RES
	user := TEST_USER
	user.IsVerified = true

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Email: google_data.Email}).Return(user, sampleErr).Once()

	ack, rfk, err := suite.Usecase.OAuthLogin(context.Background(), &google_data)
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.Equal(rfk, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestOAuthLogin_Negative_VerificationError() {
	google_data := TEST_GOOGLE_RES
	google_data.IDToken = "random"
	user := TEST_USER
	// user.IsVerified = true

	// sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Email: google_data.Email}).Return(user, nil).Once()

	ack, rfk, err := suite.Usecase.OAuthLogin(context.Background(), &google_data)
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.Equal(rfk, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestOAuthLogin_Negative_JWTError() {
	google_data := TEST_GOOGLE_RES
	user := TEST_USER
	user.IsVerified = true

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Email: google_data.Email}).Return(user, nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "accessToken", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return("acc.ess.Token", sampleErr).Once()

	ack, rfk, err := suite.Usecase.OAuthLogin(context.Background(), &google_data)
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.Equal(rfk, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestOAuthLogin_Negative_HashError() {
	google_data := TEST_GOOGLE_RES
	user := TEST_USER
	user.IsVerified = true

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Email: google_data.Email}).Return(user, nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "accessToken", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return("acc.ess.Token", nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "refreshToken", time.Hour*time.Duration(suite.ENV.REFRESH_TOKEN_LIFESPAN)).Return("ref.reshT.oken", nil).Once()
	suite.mockHashService.On("HashString", "oken").Return("hashed_str", sampleErr).Once()

	ack, rfk, err := suite.Usecase.OAuthLogin(context.Background(), &google_data)
	suite.Equal(err.GetCode(), domain.ERR_INTERNAL_SERVER)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.Equal(rfk, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestOAuthLogin_Negative_RepositoryError_SetRefreshToken() {
	google_data := TEST_GOOGLE_RES
	user := TEST_USER
	user.IsVerified = true

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Email: google_data.Email}).Return(user, nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "accessToken", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return("acc.ess.Token", nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "refreshToken", time.Hour*time.Duration(suite.ENV.REFRESH_TOKEN_LIFESPAN)).Return("ref.reshT.oken", nil).Once()
	suite.mockHashService.On("HashString", "oken").Return("hashed_str", nil).Once()
	suite.mockUserRepository.On("SetRefreshToken", context.Background(), &user, "hashed_str").Return(sampleErr).Once()

	ack, rfk, err := suite.Usecase.OAuthLogin(context.Background(), &google_data)
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.Equal(rfk, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Positive() {
	refreshToken := "refr.esht.oken"
	user := TEST_USER
	user.IsVerified = true
	user.RefreshToken = "hashed_str"
	tk := &jwt.Token{}

	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, nil).Once()
	suite.mockJWTService.On("GetTokenType", tk).Return("refreshToken", nil).Once()
	suite.mockJWTService.On("GetUsername", tk).Return(user.Username, nil).Once()
	suite.mockJWTService.On("GetExpiryDate", tk).Return(time.Now().Round(0).Add(time.Minute), nil).Once()
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: user.Username}).Return(user, nil).Once()
	suite.mockHashService.On("ValidateHashedString", "hashed_str", "oken").Return(nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "accessToken", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return("acc.ess.Token", nil).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.Nil(err, "error should be nil")
	suite.Equal(ack, "acc.ess.Token")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Negative_JWTError() {
	refreshToken := "refr.esht.oken"
	tk := &jwt.Token{}

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, sampleErr).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.Equal(ack, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Negative_GetTokenTypeError() {
	refreshToken := "refr.esht.oken"
	tk := &jwt.Token{}

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, nil).Once()
	suite.mockJWTService.On("GetTokenType", tk).Return("refreshToken", sampleErr).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Negative_InvalidTokenType() {
	refreshToken := "refr.esht.oken"
	tk := &jwt.Token{}

	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, nil).Once()
	suite.mockJWTService.On("GetTokenType", tk).Return("INVALID", nil).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.Equal(err.GetCode(), domain.ERR_FORBIDDEN)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Negative_GetUsernameError() {
	refreshToken := "refr.esht.oken"
	user := TEST_USER
	tk := &jwt.Token{}

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, nil).Once()
	suite.mockJWTService.On("GetTokenType", tk).Return("refreshToken", nil).Once()
	suite.mockJWTService.On("GetUsername", tk).Return(user.Username, sampleErr).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Negative_GetExpiryDateError() {
	refreshToken := "refr.esht.oken"
	user := TEST_USER
	tk := &jwt.Token{}

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, nil).Once()
	suite.mockJWTService.On("GetTokenType", tk).Return("refreshToken", nil).Once()
	suite.mockJWTService.On("GetUsername", tk).Return(user.Username, nil).Once()
	suite.mockJWTService.On("GetExpiryDate", tk).Return(time.Now().Round(0).Add(time.Minute), sampleErr).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Negative_ExpiredToken() {
	refreshToken := "refr.esht.oken"
	user := TEST_USER
	tk := &jwt.Token{}

	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, nil).Once()
	suite.mockJWTService.On("GetTokenType", tk).Return("refreshToken", nil).Once()
	suite.mockJWTService.On("GetUsername", tk).Return(user.Username, nil).Once()
	suite.mockJWTService.On("GetExpiryDate", tk).Return(time.Now().Round(0).Add(time.Minute*-2), nil).Once()
	suite.mockUserRepository.On("SetRefreshToken", context.Background(), &domain.User{Username: user.Username}, "").Return(nil).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.Contains(err.Error(), "expired")
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Negative_RepositoryError_FindUser() {
	refreshToken := "refr.esht.oken"
	user := TEST_USER
	tk := &jwt.Token{}

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, nil).Once()
	suite.mockJWTService.On("GetTokenType", tk).Return("refreshToken", nil).Once()
	suite.mockJWTService.On("GetUsername", tk).Return(user.Username, nil).Once()
	suite.mockJWTService.On("GetExpiryDate", tk).Return(time.Now().Round(0).Add(time.Minute), nil).Once()
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: user.Username}).Return(user, sampleErr).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Negative_NotVerified() {
	refreshToken := "refr.esht.oken"
	user := TEST_USER
	// user.IsVerified = true
	user.RefreshToken = "hashed_str"
	tk := &jwt.Token{}

	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, nil).Once()
	suite.mockJWTService.On("GetTokenType", tk).Return("refreshToken", nil).Once()
	suite.mockJWTService.On("GetUsername", tk).Return(user.Username, nil).Once()
	suite.mockJWTService.On("GetExpiryDate", tk).Return(time.Now().Round(0).Add(time.Minute), nil).Once()
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: user.Username}).Return(user, nil).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Negative_NoRefreshToken() {
	refreshToken := "refr.esht.oken"
	user := TEST_USER
	user.IsVerified = true
	// user.RefreshToken = "hashed_str"
	tk := &jwt.Token{}

	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, nil).Once()
	suite.mockJWTService.On("GetTokenType", tk).Return("refreshToken", nil).Once()
	suite.mockJWTService.On("GetUsername", tk).Return(user.Username, nil).Once()
	suite.mockJWTService.On("GetExpiryDate", tk).Return(time.Now().Round(0).Add(time.Minute), nil).Once()
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: user.Username}).Return(user, nil).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.Equal(err.GetCode(), domain.ERR_NOT_FOUND)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Negative_HashError() {
	refreshToken := "refr.esht.oken"
	user := TEST_USER
	user.IsVerified = true
	user.RefreshToken = "hashed_str"
	tk := &jwt.Token{}

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, nil).Once()
	suite.mockJWTService.On("GetTokenType", tk).Return("refreshToken", nil).Once()
	suite.mockJWTService.On("GetUsername", tk).Return(user.Username, nil).Once()
	suite.mockJWTService.On("GetExpiryDate", tk).Return(time.Now().Round(0).Add(time.Minute), nil).Once()
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: user.Username}).Return(user, nil).Once()
	suite.mockHashService.On("ValidateHashedString", "hashed_str", "oken").Return(sampleErr).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRenewAccessToken_Negative_() {
	refreshToken := "refr.esht.oken"
	user := TEST_USER
	user.IsVerified = true
	user.RefreshToken = "hashed_str"
	tk := &jwt.Token{}

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockJWTService.On("ValidateAndParseToken", refreshToken).Return(tk, nil).Once()
	suite.mockJWTService.On("GetTokenType", tk).Return("refreshToken", nil).Once()
	suite.mockJWTService.On("GetUsername", tk).Return(user.Username, nil).Once()
	suite.mockJWTService.On("GetExpiryDate", tk).Return(time.Now().Round(0).Add(time.Minute), nil).Once()
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: user.Username}).Return(user, nil).Once()
	suite.mockHashService.On("ValidateHashedString", "hashed_str", "oken").Return(nil).Once()
	suite.mockJWTService.On("SignJWTWithPayload", user.Username, user.Role, "accessToken", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return("acc.ess.Token", sampleErr).Once()

	ack, err := suite.Usecase.RenewAccessToken(context.Background(), refreshToken)
	suite.Equal(err.GetCode(), domain.ERR_INTERNAL_SERVER)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(ack, "")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestUpdateUser_Positive() {
	user := TEST_USER
	updates := dtos.UpdateUser{
		PhoneNumber: "+2347030000000",
		Bio:         "I",
		ProfilePicture: dtos.ProfilePicture{
			FileName: "profile_picture",
		},
	}

	resUpdated := map[string]string{
		"PhoneNumber":    updates.PhoneNumber,
		"Bio":            updates.Bio,
		"ProfilePicture": updates.ProfilePicture.FileName,
	}

	suite.mockUserRepository.On("UpdateUser", context.Background(), user.Username, &updates).Return(resUpdated, "", nil).Once()

	updatedData, err := suite.Usecase.UpdateUser(context.Background(), user.Username, user.Username, &updates)
	suite.Nil(err, "error should be nil")
	suite.Equal(updatedData["PhoneNumber"], updates.PhoneNumber)
	suite.Equal(updatedData["Bio"], updates.Bio)
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestUpdateUser_Positive_WithOldPFP() {
	user := TEST_USER
	updates := dtos.UpdateUser{
		PhoneNumber: "+2347030000000",
		Bio:         "I",
		ProfilePicture: dtos.ProfilePicture{
			FileName: "profile_picture",
		},
	}

	resUpdated := map[string]string{
		"PhoneNumber":    updates.PhoneNumber,
		"Bio":            updates.Bio,
		"ProfilePicture": updates.ProfilePicture.FileName,
	}

	suite.mockUserRepository.On("UpdateUser", context.Background(), user.Username, &updates).Return(resUpdated, "oldpfp", nil).Once()

	updatedData, err := suite.Usecase.UpdateUser(context.Background(), user.Username, user.Username, &updates)
	suite.Nil(err, "error should be nil")
	suite.Equal(updatedData["PhoneNumber"], updates.PhoneNumber)
	suite.Equal(updatedData["Bio"], updates.Bio)
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestUpdateUser_Negative_NotTheOwner() {
	user := TEST_USER
	updates := dtos.UpdateUser{
		PhoneNumber: "+2347030000000",
		Bio:         "I",
		ProfilePicture: dtos.ProfilePicture{
			FileName: "profile_picture",
		},
	}

	updatedData, err := suite.Usecase.UpdateUser(context.Background(), user.Username, user.Username+"random", &updates)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_FORBIDDEN)
	suite.Equal(updatedData["PhoneNumber"], "")
	suite.Equal(updatedData["Bio"], "")
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestUpdateUser_Negative_InvalidPhoneNumber() {
	user := TEST_USER
	updates := dtos.UpdateUser{
		PhoneNumber: "+2dasfasdf0",
		Bio:         "I",
		ProfilePicture: dtos.ProfilePicture{
			FileName: "profile_picture",
		},
	}

	updatedData, err := suite.Usecase.UpdateUser(context.Background(), user.Username, user.Username, &updates)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Equal(updatedData["PhoneNumber"], "")
	suite.Equal(updatedData["Bio"], "")
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestUpdateUser_Negative_EmptyUpdate() {
	user := TEST_USER
	updates := dtos.UpdateUser{}

	updatedData, err := suite.Usecase.UpdateUser(context.Background(), user.Username, user.Username, &updates)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.Equal(updatedData["PhoneNumber"], "")
	suite.Equal(updatedData["Bio"], "")
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestPromoteUser_Positive() {
	user := TEST_USER

	suite.mockUserRepository.On("ChangeRole", context.Background(), user.Username, domain.RoleAdmin).Return(nil).Once()

	err := suite.Usecase.PromoteUser(context.Background(), user.Username)
	suite.Nil(err, "error should be nil")
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestPromoteUser_Negative() {
	user := TEST_USER
	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("ChangeRole", context.Background(), user.Username, domain.RoleAdmin).Return(sampleErr).Once()

	err := suite.Usecase.PromoteUser(context.Background(), user.Username)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestDemoteUser_Positive() {
	user := TEST_USER

	suite.mockUserRepository.On("ChangeRole", context.Background(), user.Username, domain.RoleUser).Return(nil).Once()

	err := suite.Usecase.DemoteUser(context.Background(), user.Username)
	suite.Nil(err, "error should be nil")
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestDemoteUser_Negative() {
	user := TEST_USER
	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("ChangeRole", context.Background(), user.Username, domain.RoleUser).Return(sampleErr).Once()

	err := suite.Usecase.DemoteUser(context.Background(), user.Username)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestVerifyEmail_Positive() {
	user := TEST_USER
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Add(time.Minute)
	token := "token1"
	hostUrl := "http://localhost:8080"

	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: strings.TrimSpace(user.Username)}).Return(user, nil).Once()
	suite.mockUserRepository.On("VerifyUser", context.Background(), strings.TrimSpace(user.Username)).Return(nil).Once()

	err := suite.Usecase.VerifyEmail(context.Background(), user.Username, token, hostUrl)
	suite.Nil(err, "error should be nil")
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestVerifyEmail_Negative_RespositoryError_FindUser() {
	user := TEST_USER
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Add(time.Minute)
	token := "token1"
	hostUrl := "http://localhost:8080"

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: strings.TrimSpace(user.Username)}).Return(user, sampleErr).Once()

	err := suite.Usecase.VerifyEmail(context.Background(), user.Username, token, hostUrl)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestVerifyEmail_Negative_RespositoryError_InvalidToken() {
	user := TEST_USER
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Add(time.Minute)
	token := "token2"
	hostUrl := "http://localhost:8080"

	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: strings.TrimSpace(user.Username)}).Return(user, nil).Once()

	err := suite.Usecase.VerifyEmail(context.Background(), user.Username, token, hostUrl)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestVerifyEmail_Negative_RespositoryError_AlreadyVerified() {
	user := TEST_USER
	user.IsVerified = true
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Add(time.Minute)
	token := "token1"
	hostUrl := "http://localhost:8080"

	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: strings.TrimSpace(user.Username)}).Return(user, nil).Once()

	err := suite.Usecase.VerifyEmail(context.Background(), user.Username, token, hostUrl)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestVerifyEmail_Negative_ExpiredMailError() {
	user := TEST_USER
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Add(time.Minute * -1)
	mail := "mail"
	token := "token1"
	hostUrl := "http://localhost:8080"

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: strings.TrimSpace(user.Username)}).Return(user, nil).Once()
	suite.mockUserRepository.On("UpdateVerificationDetails", context.Background(), strings.TrimSpace(user.Username), mock.AnythingOfType("domain.VerificationData")).Return(nil).Once()
	suite.mockMailService.On("EmailVerificationTemplate", hostUrl, strings.TrimSpace(user.Username), mock.AnythingOfType("string")).Return(mail).Once()
	suite.mockMailService.On("SendMail", mock.AnythingOfType("string"), user.Email, mail).Return(sampleErr).Once()

	err := suite.Usecase.VerifyEmail(context.Background(), user.Username, token, hostUrl)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_INTERNAL_SERVER)
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestVerifyEmail_Negative_RepositoryError_UpdateVerification() {
	user := TEST_USER
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Add(time.Minute * -1)
	token := "token1"
	hostUrl := "http://localhost:8080"

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: strings.TrimSpace(user.Username)}).Return(user, nil).Once()
	suite.mockUserRepository.On("UpdateVerificationDetails", context.Background(), strings.TrimSpace(user.Username), mock.AnythingOfType("domain.VerificationData")).Return(sampleErr).Once()

	err := suite.Usecase.VerifyEmail(context.Background(), user.Username, token, hostUrl)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestVerifyEmail_Negative_VerificationDataExpired() {
	user := TEST_USER
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Add(time.Minute * -1)
	mail := "mail"
	token := "token1"
	hostUrl := "http://localhost:8080"

	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: strings.TrimSpace(user.Username)}).Return(user, nil).Once()
	suite.mockUserRepository.On("UpdateVerificationDetails", context.Background(), strings.TrimSpace(user.Username), mock.AnythingOfType("domain.VerificationData")).Return(nil).Once()
	suite.mockMailService.On("EmailVerificationTemplate", hostUrl, strings.TrimSpace(user.Username), mock.AnythingOfType("string")).Return(mail).Once()
	suite.mockMailService.On("SendMail", mock.AnythingOfType("string"), user.Email, mail).Return(nil).Once()

	err := suite.Usecase.VerifyEmail(context.Background(), user.Username, token, hostUrl)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestVerifyEmail_Negative_RespositoryError_VerifyUser() {
	user := TEST_USER
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Add(time.Minute)
	token := "token1"
	hostUrl := "http://localhost:8080"

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: strings.TrimSpace(user.Username)}).Return(user, nil).Once()
	suite.mockUserRepository.On("VerifyUser", context.Background(), strings.TrimSpace(user.Username)).Return(sampleErr).Once()

	err := suite.Usecase.VerifyEmail(context.Background(), user.Username, token, hostUrl)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.mockUserRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestInitResetPassword_Positive() {
	user := TEST_USER
	user.IsVerified = true
	hostUrl := "host_url"
	mail := "password_reset_mail"
	token, _ := suite.GenerateToken(12)

	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: user.Username, Email: user.Email}).Return(user, nil).Once()
	suite.mockUserRepository.On("UpdateVerificationDetails", context.Background(), user.Username, mock.AnythingOfType("domain.VerificationData")).Return(nil).Once()
	suite.mockMailService.On("PasswordResetTemplate", token).Return(mail).Once()
	suite.mockMailService.On("SendMail", mock.AnythingOfType("string"), user.Email, mail).Return(nil).Once()

	err := suite.Usecase.InitResetPassword(context.Background(), user.Username, user.Email, hostUrl)
	suite.Nil(err, "error should be nil")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockMailService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestInitResetPassword_Negative_UserNotFound() {
	user := TEST_USER
	user.IsVerified = true
	hostUrl := "host_url"

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: user.Username, Email: user.Email}).Return(user, sampleErr).Once()

	err := suite.Usecase.InitResetPassword(context.Background(), user.Username, user.Email, hostUrl)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockMailService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestInitResetPassword_Negative_NotVerified() {
	user := TEST_USER
	// user.IsVerified = true
	hostUrl := "host_url"

	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: user.Username, Email: user.Email}).Return(user, nil).Once()

	err := suite.Usecase.InitResetPassword(context.Background(), user.Username, user.Email, hostUrl)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockMailService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestInitResetPassword_Negative_RepositoryError_UpdateVerificationDetails() {
	user := TEST_USER
	user.IsVerified = true
	hostUrl := "host_url"

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: user.Username, Email: user.Email}).Return(user, nil).Once()
	suite.mockUserRepository.On("UpdateVerificationDetails", context.Background(), user.Username, mock.AnythingOfType("domain.VerificationData")).Return(sampleErr).Once()

	err := suite.Usecase.InitResetPassword(context.Background(), user.Username, user.Email, hostUrl)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockMailService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestInitResetPassword_Negative_MailError() {
	user := TEST_USER
	user.IsVerified = true
	hostUrl := "host_url"
	mail := "password_reset_mail"
	token, _ := suite.GenerateToken(12)

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: user.Username, Email: user.Email}).Return(user, nil).Once()
	suite.mockUserRepository.On("UpdateVerificationDetails", context.Background(), user.Username, mock.AnythingOfType("domain.VerificationData")).Return(nil).Once()
	suite.mockMailService.On("PasswordResetTemplate", token).Return(mail).Once()
	suite.mockMailService.On("SendMail", mock.AnythingOfType("string"), user.Email, mail).Return(sampleErr).Once()

	err := suite.Usecase.InitResetPassword(context.Background(), user.Username, user.Email, hostUrl)
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_INTERNAL_SERVER)
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockMailService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestResetPassword_Positive() {
	user := TEST_USER
	user.VerificationData.Type = domain.ResetPasswordType
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Round(0).Add(time.Minute)
	newPwd := "new_paSsword@123"
	hashedPwd := "hashed_pwd"
	resetData := dtos.ResetPassword{
		Username:    user.Username,
		NewPassword: newPwd,
	}

	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: resetData.Username}).Return(user, nil).Once()
	suite.mockHashService.On("HashString", resetData.NewPassword).Return(hashedPwd, nil).Once()
	suite.mockUserRepository.On("UpdatePassword", context.Background(), resetData.Username, hashedPwd).Return(nil).Once()

	err := suite.Usecase.ResetPassword(context.Background(), resetData, "token1")
	suite.Nil(err, "error should be nil")
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestResetPassword_Negative_RepositoryError_FindUser() {
	user := TEST_USER
	user.VerificationData.Type = domain.ResetPasswordType
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Round(0).Add(time.Minute)
	newPwd := "new_paSsword@123"
	resetData := dtos.ResetPassword{
		Username:    user.Username,
		NewPassword: newPwd,
	}

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: resetData.Username}).Return(user, sampleErr).Once()

	err := suite.Usecase.ResetPassword(context.Background(), resetData, "token1")
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}
func (suite *UserUsecaseTestSuite) TestResetPassword_Negative_TokenContentMismatch() {
	user := TEST_USER
	user.VerificationData.Type = domain.ResetPasswordType
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Round(0).Add(time.Minute)
	newPwd := "new_paSsword@123"
	resetData := dtos.ResetPassword{
		Username:    user.Username,
		NewPassword: newPwd,
	}

	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: resetData.Username}).Return(user, nil).Once()

	err := suite.Usecase.ResetPassword(context.Background(), resetData, "token2")
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestResetPassword_Negative_TokenTypeMismatch() {
	user := TEST_USER
	user.VerificationData.Type = "invalid_token_type"
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Round(0).Add(time.Minute)
	newPwd := "new_paSsword@123"
	resetData := dtos.ResetPassword{
		Username:    user.Username,
		NewPassword: newPwd,
	}

	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: resetData.Username}).Return(user, nil).Once()

	err := suite.Usecase.ResetPassword(context.Background(), resetData, "token1")
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestResetPassword_Negative_TokenExpired() {
	user := TEST_USER
	user.VerificationData.Type = domain.ResetPasswordType
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Round(0).Add(time.Minute * -1)
	newPwd := "new_paSsword@123"
	resetData := dtos.ResetPassword{
		Username:    user.Username,
		NewPassword: newPwd,
	}

	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: resetData.Username}).Return(user, nil).Once()

	err := suite.Usecase.ResetPassword(context.Background(), resetData, "token1")
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_UNAUTHORIZED)
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestResetPassword_Negative_HashError() {
	user := TEST_USER
	user.VerificationData.Type = domain.ResetPasswordType
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Round(0).Add(time.Minute)
	newPwd := "new_paSsword@123"
	hashedPwd := "hashed_pwd"
	resetData := dtos.ResetPassword{
		Username:    user.Username,
		NewPassword: newPwd,
	}

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: resetData.Username}).Return(user, nil).Once()
	suite.mockHashService.On("HashString", resetData.NewPassword).Return(hashedPwd, sampleErr).Once()

	err := suite.Usecase.ResetPassword(context.Background(), resetData, "token1")
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), domain.ERR_INTERNAL_SERVER)
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestResetPassword_Negative_RepositoryError_UpdatePassword() {
	user := TEST_USER
	user.VerificationData.Type = domain.ResetPasswordType
	user.VerificationData.Token = "token1"
	user.VerificationData.ExpiresAt = time.Now().Round(0).Add(time.Minute)
	newPwd := "new_paSsword@123"
	hashedPwd := "hashed_pwd"
	resetData := dtos.ResetPassword{
		Username:    user.Username,
		NewPassword: newPwd,
	}

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockUserRepository.On("FindUser", context.Background(), &domain.User{Username: resetData.Username}).Return(user, nil).Once()
	suite.mockHashService.On("HashString", resetData.NewPassword).Return(hashedPwd, nil).Once()
	suite.mockUserRepository.On("UpdatePassword", context.Background(), resetData.Username, hashedPwd).Return(sampleErr).Once()

	err := suite.Usecase.ResetPassword(context.Background(), resetData, "token1")
	suite.NotNil(err, "error should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.mockUserRepository.AssertExpectations(suite.T())
	suite.mockHashService.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLogout_Positive() {
	user := TEST_USER
	accessToken := "acc.esst.oken"

	suite.mockCacheRepository.On("CacheData", accessToken, "", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return(nil).Once()
	suite.mockUserRepository.On("SetRefreshToken", context.Background(), &domain.User{Username: user.Username}, "").Return(nil).Once()

	err := suite.Usecase.Logout(context.Background(), user.Username, accessToken)
	suite.Nil(err, "err should be nil")
	suite.mockCacheRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLogout_Negative_CacheError() {
	user := TEST_USER
	accessToken := "acc.esst.oken"

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockCacheRepository.On("CacheData", accessToken, "", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return(sampleErr).Once()

	err := suite.Usecase.Logout(context.Background(), user.Username, accessToken)
	suite.NotNil(err, "err should not be nil")
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.mockCacheRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLogout_Negative_ReposioryError_SetRefreshToken() {
	user := TEST_USER
	accessToken := "acc.esst.oken"

	sampleErr := domain.NewError("this a sample error", domain.ERR_BAD_REQUEST)
	suite.mockCacheRepository.On("CacheData", accessToken, "", time.Minute*time.Duration(suite.ENV.ACCESS_TOKEN_LIFESPAN)).Return(nil).Once()
	suite.mockUserRepository.On("SetRefreshToken", context.Background(), &domain.User{Username: user.Username}, "").Return(sampleErr).Once()

	err := suite.Usecase.Logout(context.Background(), user.Username, accessToken)
	suite.Equal(err.GetCode(), sampleErr.GetCode())
	suite.Equal(err.Error(), sampleErr.Error())
	suite.NotNil(err, "err should not be nil")

	suite.mockCacheRepository.AssertExpectations(suite.T())
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

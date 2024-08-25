package controller

import (
    "Blog_Starter/domain"
    "Blog_Starter/domain/mocks"
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
	"fmt"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpControllerTestSuite struct {
    suite.Suite
    signUpController *SignUpController
    mockSignupUsecase *mocks.SignupUsecase
    mockOtpUsecase    *mocks.OtpUsecase
}

func (suite *SignUpControllerTestSuite) SetupTest() {
    suite.mockSignupUsecase = new(mocks.SignupUsecase)
    suite.mockOtpUsecase = new(mocks.OtpUsecase)
    suite.signUpController = NewSignUpController(suite.mockSignupUsecase, suite.mockOtpUsecase)
}

func (suite *SignUpControllerTestSuite) TestSignUp_Success() {
    userSignUp := domain.UserSignUp{
        Username: "test_user",
        Email:    "test@example.com",
        Password: "Test@1234",
    }

    user := &domain.User{
        UserID: primitive.NewObjectID(),
        Username:  userSignUp.Username,
        Email:     userSignUp.Email,
        Password:  userSignUp.Password,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    suite.mockSignupUsecase.On("CreateUser", mock.Anything, &userSignUp).Return(user, nil)
    suite.mockOtpUsecase.On("GetOtpByEmail", mock.Anything, userSignUp.Email).Return(domain.Otp{}, errors.New("not found"))
    suite.mockOtpUsecase.On("SaveOtp", mock.Anything, mock.Anything).Return(nil)

    body, _ := json.Marshal(userSignUp)
    req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/signup", suite.signUpController.SignUp)
    router.ServeHTTP(rr, req)

	fmt.Println(rr.Body)

    assert.Equal(suite.T(), http.StatusCreated, rr.Code)
    assert.Contains(suite.T(), rr.Body.String(), "succesfuly sent otp")
}

func (suite *SignUpControllerTestSuite) TestSignUp_ValidationError() {
    userSignUp := domain.UserSignUp{
        Username: "test_user",
        Email:    "invalid-email",
        Password: "short",
    }

    body, _ := json.Marshal(userSignUp)
    req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/signup", suite.signUpController.SignUp)
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
    assert.Contains(suite.T(), rr.Body.String(), "invalid email format")
}

func (suite *SignUpControllerTestSuite) TestVerifyEmail_Success() {
    verifyEmailRequest := domain.VerifyEmailRequest{
        Email: "test@example.com",
        OTP:   "1234",
    }

    otp := domain.Otp{
        Email:      verifyEmailRequest.Email,
        Otp:        verifyEmailRequest.OTP,
        Expiration: time.Now().Add(5 * time.Minute),
    }

    user := &domain.UserResponse{
        UserID:   primitive.NewObjectID(),
        Username: "test_user",
        Email:    verifyEmailRequest.Email,
    }

    suite.mockOtpUsecase.On("GetOtpByEmail", mock.Anything, verifyEmailRequest.Email).Return(otp, nil)
    suite.mockSignupUsecase.On("VerifyEmail", mock.Anything, &verifyEmailRequest).Return(user, nil)

    body, _ := json.Marshal(verifyEmailRequest)
    req, err := http.NewRequest(http.MethodPost, "/verify-email", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/verify-email", suite.signUpController.VerifyEmail)
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusOK, rr.Code)
    assert.Contains(suite.T(), rr.Body.String(), "Email verified")
}

func (suite *SignUpControllerTestSuite) TestVerifyEmail_InvalidOTP() {
    verifyEmailRequest := domain.VerifyEmailRequest{
        Email: "test@example.com",
        OTP:   "1234",
    }

    otp := domain.Otp{
        Email:      verifyEmailRequest.Email,
        Otp:        "5678",
        Expiration: time.Now().Add(5 * time.Minute),
    }

    suite.mockOtpUsecase.On("GetOtpByEmail", mock.Anything, verifyEmailRequest.Email).Return(otp, nil)

    body, _ := json.Marshal(verifyEmailRequest)
    req, err := http.NewRequest(http.MethodPost, "/verify-email", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/verify-email", suite.signUpController.VerifyEmail)
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
    assert.Contains(suite.T(), rr.Body.String(), "Invalid OTP")
}

func (suite *SignUpControllerTestSuite) TestResendOTP_Success() {
    resendOTPRequest := domain.ResendOTPRequest{
        Email: "test@example.com",
    }

    otp := domain.Otp{
        Email:      resendOTPRequest.Email,
        Otp:        "1234",
        Expiration: time.Now().Add(5 * time.Minute),
    }

    suite.mockSignupUsecase.On("ResendOTP", mock.Anything, &resendOTPRequest).Return(nil)
    suite.mockOtpUsecase.On("GetOtpByEmail", mock.Anything, resendOTPRequest.Email).Return(otp, nil)
    suite.mockOtpUsecase.On("SaveOtp", mock.Anything, mock.Anything).Return(nil)

    body, _ := json.Marshal(resendOTPRequest)
    req, err := http.NewRequest(http.MethodPost, "/resend-otp", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/resend-otp", suite.signUpController.ResendOTP)
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusOK, rr.Code)
    assert.Contains(suite.T(), rr.Body.String(), "succesfuly sent otp")
}

func (suite *SignUpControllerTestSuite) TestResendOTP_UserNotFound() {
    resendOTPRequest := domain.ResendOTPRequest{
        Email: "test@example.com",
    }

    suite.mockSignupUsecase.On("ResendOTP", mock.Anything, &resendOTPRequest).Return(errors.New("mongo: no documents in result"))

    body, _ := json.Marshal(resendOTPRequest)
    req, err := http.NewRequest(http.MethodPost, "/resend-otp", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/resend-otp", suite.signUpController.ResendOTP)
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusNotFound, rr.Code)
    assert.Contains(suite.T(), rr.Body.String(), "user not found")
}

func TestSignUpControllerTestSuite(t *testing.T) {
    suite.Run(t, new(SignUpControllerTestSuite))
}
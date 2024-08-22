package controllers_test

import (
	domain "blogs/Domain"
	controllers "blogs/delivery/controllers"
	"blogs/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type OauthControllerTestSuite struct {
	suite.Suite
	mockOauthUsecase *mocks.OauthUsecase
	mockLoginService *mocks.LoginUsecase
	oauthController  *controllers.OauthController
}

func (suite *OauthControllerTestSuite) SetupTest() {
	suite.mockOauthUsecase = new(mocks.OauthUsecase)
	suite.mockLoginService = new(mocks.LoginUsecase)
	suite.oauthController = &controllers.OauthController{
		OauthUsecase: suite.mockOauthUsecase,
		Login:        suite.mockLoginService,
		Config:       nil,
	}
}


func (suite *OauthControllerTestSuite) TestGoogleAuth_Success() {
	// Mock the OauthService to return a URL
	expectedURL := "https://example.com/oauth?state=oauthStateString"
	suite.mockOauthUsecase.On("OauthService").Return(&domain.URL{URL: expectedURL}).Once()

	// Prepare the request
	req, err := http.NewRequest(http.MethodGet, "/auth/signin/google", nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the GoogleAuth method
	suite.oauthController.GoogleAuth(c)

	// Verify the response
	suite.Equal(http.StatusTemporaryRedirect, w.Code)
	suite.Equal(expectedURL, w.Header().Get("Location"))

	// Assert that the mock was called
	suite.mockOauthUsecase.AssertCalled(suite.T(), "OauthService")
}

// func (suite *OauthControllerTestSuite) TestGoogleCallback_Success() {
// 	// Mocking the OauthUsecase and LoginService responses
// 	verifiedUser := domain.User{
// 		Email:    "test@example.com",
// 		Verified: true,
// 	}

// 	// Hard-coding the OauthSecret check
// 	expectedOauthSecret := "expectedState"

// 	suite.mockOauthUsecase.On("OauthCallback", mock.Anything, "authCode").Return(&domain.UserResponse{User: verifiedUser}).Once()
// 	suite.mockLoginService.On("CreateAccessToken", &verifiedUser, mock.Anything, mock.Anything).Return("accessToken", nil).Once()
// 	suite.mockLoginService.On("CreateRefreshToken", &verifiedUser, mock.Anything, mock.Anything).Return("refreshToken", nil).Once()
// 	suite.mockLoginService.On("SaveAsActiveUser", mock.AnythingOfType("domain.ActiveUser"), "refreshToken", mock.Anything).Return(nil).Once()

// 	// Prepare the request
// 	req, err := http.NewRequest(http.MethodGet, "/auth/callback?state=expectedState&code=authCode", nil)
// 	if err != nil {
// 		suite.T().Fatal(err)
// 	}

// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = req

// 	// Directly checking the state against the expected hard-coded value
// 	if state := c.Query("state"); state != expectedOauthSecret {
// 		suite.T().Fatal("state does not match expected OauthSecret")
// 	}

// 	// Call the GoogleCallback method
// 	suite.oauthController.GoogleCallback(c)

// 	// Convert response.Data to map[string]interface{}
// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	if err != nil {
// 		suite.T().Fatal(err)
// 	}

// 	// Ensure that "Email" and "Verified" are not nil before asserting
// 	dataMap, ok := response["data"].(map[string]interface{})
// 	if !ok {
// 		suite.T().Fatal("expected Data to be of type map[string]interface{}")
// 	}

// 	email, emailOk := dataMap["email"].(string)
// 	if !emailOk || email == "" {
// 		suite.T().Fatal("expected Email to be a non-empty string")
// 	}
// 	verified, verifiedOk := dataMap["verified"].(bool)
// 	if !verifiedOk {
// 		suite.T().Fatal("expected Verified to be a bool")
// 	}

// 	// Convert map to domain.User
// 	decodedUser := domain.User{
// 		Email:    email,
// 		Verified: verified,
// 	}

// 	// Assert that the user was verified successfully
// 	suite.Equal(verifiedUser.Email, decodedUser.Email)
// 	suite.True(decodedUser.Verified)

// 	// Assert that the mock was called with the correct parameters
// 	suite.mockOauthUsecase.AssertCalled(suite.T(), "OauthCallback", mock.Anything, "authCode")
// 	suite.mockLoginService.AssertCalled(suite.T(), "CreateAccessToken", &verifiedUser, mock.Anything, mock.Anything)
// 	suite.mockLoginService.AssertCalled(suite.T(), "CreateRefreshToken", &verifiedUser, mock.Anything, mock.Anything)
// 	suite.mockLoginService.AssertCalled(suite.T(), "SaveAsActiveUser", mock.AnythingOfType("domain.ActiveUser"), "refreshToken", mock.Anything)
// }



func TestOauthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(OauthControllerTestSuite))
}



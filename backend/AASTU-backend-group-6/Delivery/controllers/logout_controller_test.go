package controllers

import (
	infrastructure "blogs/Infrastructure"
	"blogs/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LogoutControllerTestSuite struct {
	suite.Suite
	controller        *LogoutController
	mockLogoutUsecase *mocks.LogoutUsecase
	env               *infrastructure.Config
}

func (suite *LogoutControllerTestSuite) SetupTest() {
	suite.mockLogoutUsecase = new(mocks.LogoutUsecase)
	suite.env = &infrastructure.Config{}
	suite.controller = &LogoutController{
		LogoutUsecase: suite.mockLogoutUsecase,
		Env:           suite.env,
	}
}

// func (suite *LogoutControllerTestSuite) TestLogout_Success() {
// 	// Setup
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	id := "testID"
// 	userAgent := "testUserAgent"
// 	c.Request, _ = http.NewRequest(http.MethodGet, "/logout", nil)
// 	c.Request.Header.Set("user_id", id)
// 	c.Request.Header.Set("User-Agent", userAgent)

// 	// Mocking
// 	suite.mockLogoutUsecase.On("CheckActiveUser", mock.Anything, id, userAgent).Return(&domain.ActiveUser{}, nil)
// 	suite.mockLogoutUsecase.On("Logout", mock.Anything, id, userAgent).Return(nil)

// 	// Execute
// 	suite.controller.Logout(c)

// 	// Assertions
// 	suite.Equal(http.StatusOK, w.Code)
// 	expectedResponse := gin.H{"message": "Logout success"}
// 	var actualResponse gin.H
// 	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
// 	if err != nil {
// 		suite.T().Fatal(err)
// 	}
// 	suite.Equal(expectedResponse, actualResponse)
// 	suite.mockLogoutUsecase.AssertExpectations(suite.T())
// }

// func (suite *LogoutControllerTestSuite) TestLogout_UserNotFound() {
// 	// Setup
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	id := "testID"
// 	userAgent := "testUserAgent"
// 	c.Request, _ = http.NewRequest(http.MethodGet, "/logout", nil)
// 	c.Request.Header.Set("user_id", id)
// 	c.Request.Header.Set("User-Agent", userAgent)

// 	// Mocking
// 	suite.mockLogoutUsecase.On("CheckActiveUser", mock.Anything, id, userAgent).Return(nil, errors.New("user not found"))

// 	// Execute
// 	suite.controller.Logout(c)

// 	// Assertions
// 	suite.Equal(http.StatusNotFound, w.Code)
// 	expectedResponse := gin.H{"error": "User not found, page not found, login before logout"}
// 	var actualResponse gin.H
// 	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
// 	if err != nil {
// 		suite.T().Fatal(err)
// 	}
// 	suite.Equal(expectedResponse, actualResponse)
// 	suite.mockLogoutUsecase.AssertExpectations(suite.T())
// }

// func (suite *LogoutControllerTestSuite) TestLogout_InternalServerError() {
// 	// Setup
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	id := "testID"
// 	userAgent := "testUserAgent"
// 	c.Request, _ = http.NewRequest(http.MethodGet, "/logout", nil)
// 	c.Request.Header.Set("user_id", id)
// 	c.Request.Header.Set("User-Agent", userAgent)

// 	// Mocking
// 	suite.mockLogoutUsecase.On("CheckActiveUser", mock.Anything, id, userAgent).Return(&domain.ActiveUser{}, nil)
// 	suite.mockLogoutUsecase.On("Logout", mock.Anything, id, userAgent).Return(errors.New("internal error"))

// 	// Execute
// 	suite.controller.Logout(c)

// 	// Assertions
// 	suite.Equal(http.StatusInternalServerError, w.Code)
// 	expectedResponse := gin.H{"error": "internal error"}
// 	var actualResponse gin.H
// 	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
// 	if err != nil {
// 		suite.T().Fatal(err)
// 	}
// 	suite.Equal(expectedResponse, actualResponse)
// 	suite.mockLogoutUsecase.AssertExpectations(suite.T())
// }

func TestLogoutControllerTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutControllerTestSuite))
}

package handlers_test

import (
	"blogApp/internal/domain"
	"blogApp/internal/http/handlers"
	"blogApp/mocks/usecase"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userHandlerSuite struct {
	suite.Suite
	handler       *handlers.UserHandler
	usecase       *mocks.UserUseCaseInterface
	testingServer *httptest.Server
}

func (suite *userHandlerSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.usecase = new(mocks.UserUseCaseInterface)

	suite.handler = &handlers.UserHandler{
		UserUsecase: suite.usecase,
	}

	router := gin.Default()
	router.POST("/register", suite.handler.Register)
	router.POST("/login", suite.handler.Login)
	router.PUT("/update", suite.handler.UpdateUser)
	suite.testingServer = httptest.NewServer(router)
}

func (suite *userHandlerSuite) TearDownSuite() {
	suite.testingServer.Close()
}

func (suite *userHandlerSuite) TestRegister_Positive() {
	userID := primitive.NewObjectID()
	user := &domain.User{
		ID:       userID,
		Email:    "johndoe@gmail.com",
		Password: "Password123@@",
	}

	suite.usecase.On("RegisterUser", user).Return(user, nil)

	requestBody, err := json.Marshal(user)
	suite.NoError(err, "cannot marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()
	
	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	
	suite.NoError(err, "cannot unmarshal response body")

	
	//fmt.Println(err, ".............................++++++++++++++++++++++++++++++==========================..........................................................................")
	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal(userID.Hex(), responseBody["id"])
	suite.Equal(user.Email, responseBody["email"])
	suite.Equal(user.Role, responseBody["role"])
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *userHandlerSuite) TestLogin_Positive() {
	userID := primitive.NewObjectID()
	user := &domain.User{
		ID:       userID,
		Email:    "johndoe@gmail.com",
		Password: "Password123@@@",
	}

	token := &domain.Token{
		AccessToken:  "mockAccessToken",
		RefreshToken: "",
	}

	suite.usecase.On("Login", user.Email, user.Password).Return(user, token, nil)

	requestBody, err := json.Marshal(map[string]string{
		"email":    user.Email,
		"password": user.Password,
	})
	suite.NoError(err, "cannot marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "cannot unmarshal response body")

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal(userID.Hex(), responseBody["id"])
	suite.Equal(user.Email, responseBody["email"])
	suite.Equal(user.Role, responseBody["role"])

	tokenResponse, ok := responseBody["token"].(map[string]interface{})
	suite.True(ok, "token should be a map with access_token and refresh_token")

	suite.Equal(token.AccessToken, tokenResponse["access_token"])
	suite.Equal(token.RefreshToken, tokenResponse["refresh_token"])
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *userHandlerSuite) TestUpdateUser_Success() {
	userID := primitive.NewObjectID()
	user := &domain.User{
		ID:       userID,
		Email:    "johndoe@example.com",
		Password: "newpassword123",
		Role:     "user",
	}

	suite.usecase.On("UpdateUser", user).Return(nil)

	claims := &domain.Claims{
		UserID: userID.Hex(),
	}

	requestBody, err := json.Marshal(user)
	suite.NoError(err, "cannot marshal struct to json")

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/update", suite.testingServer.URL), bytes.NewBuffer(requestBody))
	suite.NoError(err, "cannot create request")
	req.Header.Set("Authorization", "Bearer mocktoken") 

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("claims", claims)
	suite.handler.UpdateUser(c)

	suite.Equal(http.StatusOK, w.Code)
	var responseBody map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&responseBody)
	suite.NoError(err, "cannot unmarshal response body")
	suite.Equal("User updated successfully", responseBody["message"])

	suite.usecase.AssertExpectations(suite.T())
}

func (suite *userHandlerSuite) TestUpdateUser_InvalidInput() {
	claims := &domain.Claims{
		UserID: primitive.NewObjectID().Hex(),
	}

	
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/update", suite.testingServer.URL), bytes.NewBuffer([]byte(`{invalid}`)))
	suite.NoError(err, "cannot create request")
	req.Header.Set("Authorization", "Bearer mocktoken") 

	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("claims", claims)

	
	suite.handler.UpdateUser(c)

	
	suite.Equal(http.StatusBadRequest, w.Code)
	var responseBody map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&responseBody)
	suite.NoError(err, "cannot unmarshal response body")
	suite.Equal("Invalid input data", responseBody["error"])
}

func (suite *userHandlerSuite) TestUpdateUser_InternalServerError() {
	userID := primitive.NewObjectID()
	user := &domain.User{
		ID:       userID,
		Email:    "johndoe@example.com",
		Password: "newpassword123",
		Role:     "user",
	}

	suite.usecase.On("UpdateUser", user).Return(fmt.Errorf("internal error"))

	
	claims := &domain.Claims{
		UserID: userID.Hex(),
	}

	
	requestBody, err := json.Marshal(user)
	suite.NoError(err, "cannot marshal struct to json")

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/update", suite.testingServer.URL), bytes.NewBuffer(requestBody))
	suite.NoError(err, "cannot create request")
	req.Header.Set("Authorization", "Bearer mocktoken") 

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("claims", claims)

	suite.handler.UpdateUser(c)

	
	suite.Equal(http.StatusInternalServerError, w.Code)
	var responseBody map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&responseBody)
	suite.NoError(err, "cannot unmarshal response body")
	suite.Equal("internal error", responseBody["error"])

	suite.usecase.AssertExpectations(suite.T())
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(userHandlerSuite))
}

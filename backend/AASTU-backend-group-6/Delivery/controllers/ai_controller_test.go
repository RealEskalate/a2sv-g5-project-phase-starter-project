package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	domain "blogs/Domain"
// 	"blogs/Infrastructure"
// 	"blogs/controllers"
// 	"blogs/mocks"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// )

// type AiControllerTestSuite struct {
// 	suite.Suite
// 	mockAiUsecase *mocks.AIUsecase
// 	aiController *controllers.AiController
// }

// func (suite *AiControllerTestSuite) SetupTest() {
// 	suite.mockAiUsecase = new(mocks.AIUsecase)
// 	suite.aiController = &controllers.AiController{
// 		Config:    &infrastructure.Config{}, // Adjust as needed
// 		AiUsecase: suite.mockAiUsecase,
// 	}
// }

// func (suite *AiControllerTestSuite) TestAsk_Success() {
// 	// Prepare a valid request
// 	requestPayload := domain.AiRequest{
// 		Query: "What is the capital of France?",
// 	}
// 	body, _ := json.Marshal(requestPayload)

// 	req, err := http.NewRequest(http.MethodPost, "/ai/ask", bytes.NewReader(body))
// 	if err != nil {
// 		suite.T().Fatal(err)
// 	}

// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = req

// 	// Mock the AskAI method
// 	expectedResponse := domain.AiResponse{
// 		Answer: "The capital of France is Paris.",
// 	}
// 	suite.mockAiUsecase.On("AskAI", mock.Anything, requestPayload).Return(&domain.SuccessResponse{
// 		Message: "Request successful",
// 		Data:    expectedResponse,
// 		Status:  http.StatusOK,
// 	}).Once()

// 	// Call the controller method
// 	suite.aiController.Ask(c)

// 	// Validate response
// 	suite.Equal(http.StatusOK, w.Code)
// 	var response domain.SuccessResponse
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	suite.NoError(err)
// 	suite.Equal("Request successful", response.Message)
// 	suite.Equal(expectedResponse, response.Data)
// }

// func (suite *AiControllerTestSuite) TestAsk_BadRequest() {
// 	// Prepare an invalid request
// 	body := []byte(`{"invalid_field": "value"}`)

// 	req, err := http.NewRequest(http.MethodPost, "/ask", bytes.NewReader(body))
// 	if err != nil {
// 		suite.T().Fatal(err)
// 	}

// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = req

// 	// Call the controller method
// 	suite.aiController.Ask(c)

// 	// Validate response
// 	suite.Equal(http.StatusBadRequest, w.Code)
// 	var response map[string]string
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	suite.NoError(err)
// 	suite.Contains(response, "error")
// }

// func TestAiControllerTestSuite(t *testing.T) {
// 	suite.Run(t, new(AiControllerTestSuite))
// }

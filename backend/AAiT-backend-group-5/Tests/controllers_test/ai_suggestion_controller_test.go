package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks/usecase_mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// Test Suite
type ContentSuggestionControllerTestSuite struct {
	suite.Suite
	mockUsecase *mocks.AIContentSuggestionUsecase
	controller  *controllers.ContentSuggestionController
	router      *gin.Engine
}

func (suite *ContentSuggestionControllerTestSuite) SetupSuite() {
	suite.mockUsecase = new(mocks.AIContentSuggestionUsecase)
	suite.controller = controllers.NewContentSuggestionController(suite.mockUsecase)
	suite.router = gin.Default()

	// Define the routes
	suite.router.GET("/suggest", suite.controller.HandleSuggestion)
	suite.router.GET("/improve/:id", suite.controller.HandleContentImprovement)
}

func (suite *ContentSuggestionControllerTestSuite) TearDownSuite() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *ContentSuggestionControllerTestSuite) TestHandleSuggestion_Success() {
	query := "test query"
	expectedSuggestion := []string{"Title", "Content", "Tag1&Tag2"}

	suite.mockUsecase.On("SuggestContent", query).Return(expectedSuggestion, nil).Once()

	request, _ := http.NewRequest(http.MethodGet, "/suggest?query="+query, nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Title")
	suite.Contains(responseWriter.Body.String(), "Content")
	suite.Contains(responseWriter.Body.String(), "Tag1")
	suite.Contains(responseWriter.Body.String(), "Tag2")
}

func (suite *ContentSuggestionControllerTestSuite) TestHandleSuggestion_BadRequest() {
	request, _ := http.NewRequest(http.MethodGet, "/suggest", nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusBadRequest, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Query parameter is required")
}

func (suite *ContentSuggestionControllerTestSuite) TestHandleSuggestion_InternalServerError() {
	query := "test query"
	expectedError := &models.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}

	suite.mockUsecase.On("SuggestContent", query).Return(nil, expectedError).Once()

	request, _ := http.NewRequest(http.MethodGet, "/suggest?query="+query, nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(expectedError.Code, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "internal server error")
}

func (suite *ContentSuggestionControllerTestSuite) TestHandleContentImprovement_Success() {
	blogID := "123"
	expectedSuggestion := []string{"Improved Title", "Improved Content", "Tag3&Tag4"}

	suite.mockUsecase.On("ImproveBlogContent", blogID).Return(expectedSuggestion, nil).Once()

	request, _ := http.NewRequest(http.MethodGet, "/improve/"+blogID, nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Improved Title")
	suite.Contains(responseWriter.Body.String(), "Improved Content")
	suite.Contains(responseWriter.Body.String(), "Tag3")
	suite.Contains(responseWriter.Body.String(), "Tag4")
}

func (suite *ContentSuggestionControllerTestSuite) TestHandleContentImprovement_InternalSeverError() {
	blogID := "123"
	expectedError := &models.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}

	suite.mockUsecase.On("ImproveBlogContent", blogID).Return(nil, expectedError).Once()

	request, _ := http.NewRequest(http.MethodGet, "/improve/"+blogID, nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(expectedError.Code, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "internal server error")
}

func TestContentSuggestionControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ContentSuggestionControllerTestSuite))
}

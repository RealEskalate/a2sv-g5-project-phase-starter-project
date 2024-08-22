package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks/usecase_mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type PromoteDemoteControllerTestSuite struct {
	suite.Suite
	mockUsecase *mocks.PromoteDemoteUserUsecase
	controller  *controllers.PromoteDemoteController
	router      *gin.Engine
}

func (suite *PromoteDemoteControllerTestSuite) SetupSuite() {
	suite.mockUsecase = new(mocks.PromoteDemoteUserUsecase)
	suite.controller = controllers.NewPromoteDemoteController(suite.mockUsecase)
	suite.router = gin.Default()

	// Define the routes
	suite.router.PUT("/promote/:id", suite.controller.PromoteUser)
	suite.router.PUT("/demote/:id", suite.controller.DemoteUser)
}

func (suite *PromoteDemoteControllerTestSuite) TearDownSuite() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *PromoteDemoteControllerTestSuite) TestPromoteUser_Success() {
	userID := "123"
	suite.mockUsecase.On("PromoteUser", mock.Anything, userID).Return(nil).Once()

	request, _ := http.NewRequest(http.MethodPut, "/promote/"+userID, nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "user successfully promoted")
}

func (suite *PromoteDemoteControllerTestSuite) TestPromoteUser_Error() {
	userID := "123"
	expectedError := &models.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "promotion error",
	}

	suite.mockUsecase.On("PromoteUser", mock.Anything, userID).Return(expectedError).Once()

	request, _ := http.NewRequest(http.MethodPut, "/promote/"+userID, nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(expectedError.Code, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "promotion error")
}

func (suite *PromoteDemoteControllerTestSuite) TestDemoteUser_Success() {
	userID := "123"
	suite.mockUsecase.On("DemoteUser", mock.Anything, userID).Return(nil).Once()

	request, _ := http.NewRequest(http.MethodPut, "/demote/"+userID, nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "user successfully demoted")
}

func (suite *PromoteDemoteControllerTestSuite) TestDemoteUser_Error() {
	userID := "123"
	expectedError := &models.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "demotion error",
	}

	suite.mockUsecase.On("DemoteUser", mock.Anything, userID).Return(expectedError).Once()

	request, _ := http.NewRequest(http.MethodPut, "/demote/"+userID, nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(expectedError.Code, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "demotion error")
}

func TestPromoteDemoteControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PromoteDemoteControllerTestSuite))
}

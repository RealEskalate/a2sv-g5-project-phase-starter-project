package usecase

import (
	"context"
	"testing"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type LogoutUsecaseTestSuite struct {
	suite.Suite
	ctrl                  *gomock.Controller
	mockJwtService        *mocks.MockJwtService
	mockSessionRepository *mocks.MockSessionRepository
	logoutUsecase         interfaces.LogoutUsecase
}

func (suite *LogoutUsecaseTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockJwtService = mocks.NewMockJwtService(suite.ctrl)
	suite.mockSessionRepository = mocks.NewMockSessionRepository(suite.ctrl)
	suite.logoutUsecase = usecases.NewLogoutUsecase(suite.mockJwtService, suite.mockSessionRepository)
}

func (suite *LogoutUsecaseTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

// Test cases for LogoutUser
func (suite *LogoutUsecaseTestSuite) TestLogoutUserSuccess() {
	userID := "user123"

	suite.mockSessionRepository.EXPECT().
		RemoveToken(context.Background(), userID).
		Return(nil)

	errResp := suite.logoutUsecase.LogoutUser(context.Background(), userID)
	if errResp != nil {
		suite.T().Errorf("expected no error but got" + errResp.Error())
	}
}

func (suite *LogoutUsecaseTestSuite) TestLogoutUserFailure() {
	userID := "user123"
	expectedError := &models.ErrorResponse{
		Message: "Error removing user token",
	}

	// Setup expectation for token removal failure
	suite.mockSessionRepository.EXPECT().
		RemoveToken(context.Background(), userID).
		Return(expectedError)

	// Call the method under test
	errResp := suite.logoutUsecase.LogoutUser(context.Background(), userID)

	suite.Error(errResp)
	suite.Equal(expectedError.Message, errResp.Message)
}

func TestLogoutUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutUsecaseTestSuite))
}

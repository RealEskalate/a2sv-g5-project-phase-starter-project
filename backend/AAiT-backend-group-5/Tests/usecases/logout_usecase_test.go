package usecases_test

import (
	"context"
	"testing"

	config "github.com/aait.backend.g5.main/backend/Config"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type LogoutUsecaseTestSuite struct {
	suite.Suite
	jwtServiceMock *mocks.MockJwtService
	repositoryMock *mocks.MockSessionRepository
	env            config.Env
	logoutUsecase  interfaces.LogoutUsecase
	ctr            *gomock.Controller
}

func (suite *LogoutUsecaseTestSuite) SetupSuite() {
	suite.ctr = gomock.NewController(suite.T())
	suite.jwtServiceMock = mocks.NewMockJwtService(suite.ctr)
	suite.repositoryMock = mocks.NewMockSessionRepository(suite.ctr)

	suite.env = config.Env{}

	suite.logoutUsecase = usecases.NewLogoutUsecase(
		suite.jwtServiceMock,
		suite.repositoryMock,
	)
}

func (suite *LogoutUsecaseTestSuite) TearDownSuite() {
	suite.ctr.Finish()
}

func (suite *LogoutUsecaseTestSuite) TestLogoutUser_Success() {
	ctx := context.Background()
	userID := "user1"

	suite.repositoryMock.
		EXPECT().
		RemoveToken(ctx, userID).
		Return(nil)

	result := suite.logoutUsecase.LogoutUser(ctx, userID)
	if result != nil {
		suite.T().Errorf("Error removing user token: %v", result)
	}
}

func (suite *LogoutUsecaseTestSuite) TestLogoutUser_ErrorRemovingToken() {
	ctx := context.Background()
	userID := "user1"

	suite.repositoryMock.
		EXPECT().
		RemoveToken(ctx, userID).
		Return(models.InternalServerError("Error removing user token"))

	result := suite.logoutUsecase.LogoutUser(ctx, userID)
	suite.Equal(models.InternalServerError("Error removing user token"), result)
}

func TestLogoutUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutUsecaseTestSuite))
}

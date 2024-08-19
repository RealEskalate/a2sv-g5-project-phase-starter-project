package usecase

import (
	"context"
	"testing"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	ctrl            *gomock.Controller
	mockUserRepo    *mocks.MockUserRepository
	userUsecase     interfaces.PromoteUserUsecase
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockUserRepo = mocks.NewMockUserRepository(suite.ctrl)
	suite.userUsecase = usecases.NewUserUsecase(suite.mockUserRepo)
}

func (suite *UserUsecaseTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *UserUsecaseTestSuite) TestPromoteUserSuccess() {
	ctx := context.Background()
	userID := "user123"
	user := &models.User{ID: userID, Role: "user"}

	suite.mockUserRepo.EXPECT().
		GetUserByID(ctx, userID).
		Return(user, nil)

	suite.mockUserRepo.EXPECT().
		PromoteUser(ctx, userID).
		Return(nil)

	errResp := suite.userUsecase.PromoteUser(ctx, userID)

	suite.Nil(errResp)
}

func (suite *UserUsecaseTestSuite) TestPromoteUserAlreadyAdmin() {
	ctx := context.Background()
	userID := "admin123"
	user := &models.User{ID: userID, Role: "admin"}

	suite.mockUserRepo.EXPECT().
		GetUserByID(ctx, userID).
		Return(user, nil)

	errResp := suite.userUsecase.PromoteUser(ctx, userID)

	suite.NotNil(errResp)
	suite.Equal("User is already an admin", errResp.Message)
}

func (suite *UserUsecaseTestSuite) TestPromoteUserNotFound() {
	ctx := context.Background()
	userID := "unknownUser"

	suite.mockUserRepo.EXPECT().
		GetUserByID(ctx, userID).
		Return(nil, models.BadRequest("User not found"))

	errResp := suite.userUsecase.PromoteUser(ctx, userID)

	suite.NotNil(errResp)
	suite.Equal("User not found", errResp.Message)
}

func (suite *UserUsecaseTestSuite) TestDemoteUserSuccess() {
	ctx := context.Background()
	userID := "admin123"
	user := &models.User{ID: userID, Role: "admin"}

	suite.mockUserRepo.EXPECT().
		GetUserByID(ctx, userID).
		Return(user, nil)

	suite.mockUserRepo.EXPECT().
		DemoteUser(ctx, userID).
		Return(nil)

	errResp := suite.userUsecase.DemoteUser(ctx, userID)

	suite.Nil(errResp)
}

func (suite *UserUsecaseTestSuite) TestDemoteUserNotAdmin() {
	ctx := context.Background()
	userID := "user123"
	user := &models.User{ID: userID, Role: "user"}

	suite.mockUserRepo.EXPECT().
		GetUserByID(ctx, userID).
		Return(user, nil)

	errResp := suite.userUsecase.DemoteUser(ctx, userID)

	suite.NotNil(errResp)
	suite.Equal("User is not an admin", errResp.Message)
}

func (suite *UserUsecaseTestSuite) TestDemoteUserNotFound() {
	ctx := context.Background()
	userID := "unknownUser"

	suite.mockUserRepo.EXPECT().
		GetUserByID(ctx, userID).
		Return(nil, models.BadRequest("User not found"))

	errResp := suite.userUsecase.DemoteUser(ctx, userID)

	suite.NotNil(errResp)
	suite.Equal("User not found", errResp.Message)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

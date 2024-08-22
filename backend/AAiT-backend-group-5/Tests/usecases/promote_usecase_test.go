package usecases_test

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

type UserUsecaseTestSuite struct {
	suite.Suite
	repoMock *mocks.MockUserRepository
	setup    interfaces.PromoteDemoteUserUsecase
	ctrl     *gomock.Controller
}

func (suite *UserUsecaseTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repoMock = mocks.NewMockUserRepository(suite.ctrl)
	suite.setup = usecases.NewUserUsecase(suite.repoMock)
}

func (suite *UserUsecaseTestSuite) TearDownSuite() {
	suite.ctrl.Finish()
}

func (suite *UserUsecaseTestSuite) TestPromoteUser_Success() {
	ctx := context.Background()
	userID := "user1"
	user := &models.User{
		Role: "user",
	}

	suite.repoMock.
		EXPECT().
		GetUserByID(ctx, userID).
		Return(user, nil)

	suite.repoMock.
		EXPECT().
		PromoteUser(ctx, userID).
		Return(nil)

	err := suite.setup.PromoteUser(ctx, userID)
	suite.Nil(err)
}

func (suite *UserUsecaseTestSuite) TestPromoteUser_AlreadyAdmin() {
	ctx := context.Background()
	userID := "user1"
	user := &models.User{
		Role: "admin",
	}

	suite.repoMock.
		EXPECT().
		GetUserByID(ctx, userID).
		Return(user, nil)

	err := suite.setup.PromoteUser(ctx, userID)
	suite.Equal(models.BadRequest("User is already an admin"), err)
}

func (suite *UserUsecaseTestSuite) TestPromoteUser_GetUserError() {
	ctx := context.Background()
	userID := "user1"

	suite.repoMock.
		EXPECT().
		GetUserByID(ctx, userID).
		Return(nil, models.InternalServerError("Error fetching user"))

	err := suite.setup.PromoteUser(ctx, userID)
	suite.Equal(models.InternalServerError("Error fetching user"), err)
}

func (suite *UserUsecaseTestSuite) TestDemoteUser_Success() {
	ctx := context.Background()
	userID := "user1"
	user := &models.User{
		Role: "admin",
	}

	suite.repoMock.
		EXPECT().
		GetUserByID(ctx, userID).
		Return(user, nil)

	suite.repoMock.
		EXPECT().
		DemoteUser(ctx, userID).
		Return(nil)

	err := suite.setup.DemoteUser(ctx, userID)
	suite.Nil(err)
}

func (suite *UserUsecaseTestSuite) TestDemoteUser_NotAdmin() {
	ctx := context.Background()
	userID := "user1"
	user := &models.User{
		Role: "user",
	}

	suite.repoMock.
		EXPECT().
		GetUserByID(ctx, userID).
		Return(user, nil)

	err := suite.setup.DemoteUser(ctx, userID)
	suite.Equal(models.BadRequest("User is not an admin"), err)
}

func (suite *UserUsecaseTestSuite) TestDemoteUser_GetUserError() {
	ctx := context.Background()
	userID := "user1"

	suite.repoMock.
		EXPECT().
		GetUserByID(ctx, userID).
		Return(nil, models.InternalServerError("Error fetching user"))

	err := suite.setup.DemoteUser(ctx, userID)
	suite.Equal(models.InternalServerError("Error fetching user"), err)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

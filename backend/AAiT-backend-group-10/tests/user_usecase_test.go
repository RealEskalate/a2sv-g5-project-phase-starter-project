package tests

import (
	"testing"

	"aait.backend.g10/domain"
	"aait.backend.g10/tests/mocks"
	"aait.backend.g10/usecases"
	"aait.backend.g10/usecases/dto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseSuite struct {
	suite.Suite
	userRepo    *mocks.IUserRepository
	userUseCase usecases.IUserUseCase
}

func (suite *UserUseCaseSuite) SetupTest() {
	suite.userRepo = new(mocks.IUserRepository)
	suite.userUseCase = usecases.NewUserUseCase(suite.userRepo)
}
func (suite *UserUseCaseSuite) TearDownTest() {
	suite.userRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseSuite) TestCreateUser_Positive() {
	user := &domain.User{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "123",
	}

	suite.userRepo.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil)
	_, err := suite.userUseCase.CreateUser(user)
	suite.Nil(err)
}

func (suite *UserUseCaseSuite) TestCreateUser_Negative() {
	user := &domain.User{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "123",
	}

	suite.userRepo.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(domain.ErrUserCreationFailed)
	_, err := suite.userUseCase.CreateUser(user)
	suite.Equal(err, domain.ErrUserCreationFailed)
}
func (suite *UserUseCaseSuite) TestGetUserByID_Positive() {
	userID := uuid.New()
	user := &domain.User{
		ID:       userID,
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Bio:      "A test bio",
		ImageURL: "http://example.com/image.jpg",
		IsAdmin:  false,
	}

	suite.userRepo.On("GetUserByID", userID).Return(user, nil)

	result, err := suite.userUseCase.GetUserByID(userID)
	suite.Nil(err)
	suite.Equal(userID, result.ID)
	suite.Equal("John Doe", result.FullName)
	suite.Equal("johndoe@example.com", result.Email)
	suite.Equal("A test bio", result.Bio)
	suite.Equal("http://example.com/image.jpg", result.ImageUrl)
	suite.Equal(false, result.IsAdmin)
}

func (suite *UserUseCaseSuite) TestGetUserByID_Negative() {
	userID := uuid.New()
	suite.userRepo.On("GetUserByID", userID).Return(nil, domain.ErrUserNotFound)
	result, err := suite.userUseCase.GetUserByID(userID)
	suite.Equal(err, domain.ErrUserNotFound)
	suite.Nil(result)
}

func (suite *UserUseCaseSuite) TestUpdateUser_Positive() {
	requesterID := uuid.New()
	userUpdate := &dto.UserUpdate{
		ID:       requesterID,
		FullName: "Jane Doe",
	}

	suite.userRepo.On("UpdateUser", userUpdate).Return(nil)

	err := suite.userUseCase.UpdateUser(requesterID, userUpdate)
	suite.Nil(err)
}
func (suite *UserUseCaseSuite) TestUpdateUser_UnAuthorized_Negative() {
	requesterID := uuid.New()
	userUpdate := &dto.UserUpdate{
		ID:       uuid.New(),
		FullName: "Jane Doe",
	}
	err := suite.userUseCase.UpdateUser(requesterID, userUpdate)
	suite.Equal(err, domain.ErrUnAuthorized)
}

func (suite *UserUseCaseSuite) TestUpdateUser_ServerError_Negative() {
	requesterID := uuid.New()
	userUpdate := &dto.UserUpdate{
		ID:       requesterID,
		FullName: "Jane Doe",
	}
	suite.userRepo.On("UpdateUser", userUpdate).Return(domain.ErrUserUpdateFailed)
	err := suite.userUseCase.UpdateUser(requesterID, userUpdate)
	suite.Equal(err, domain.ErrUserUpdateFailed)
}

func (suite *UserUseCaseSuite) TestPromoteUser_Positive() {
	userID := uuid.New()
	suite.userRepo.On("PromoteUser", userID, true).Return(nil)

	err := suite.userUseCase.PromoteUser(userID, true)
	suite.Nil(err)
}

func (suite *UserUseCaseSuite) TestPromoteUser_Negative() {
	userID := uuid.New()
	suite.userRepo.On("PromoteUser", userID, true).Return(domain.ErrUserPromotionFailed)

	err := suite.userUseCase.PromoteUser(userID, true)
	suite.Equal(err, domain.ErrUserPromotionFailed)
}

// Test Suite Entry Point
func TestUserUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseSuite))
}

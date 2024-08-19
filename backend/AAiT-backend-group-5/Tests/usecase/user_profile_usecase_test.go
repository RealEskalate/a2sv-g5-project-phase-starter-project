package usecase

import (
	"context"
	"testing"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type UserProfileUpdateTestSuite struct {
	suite.Suite
	ctrl                 *gomock.Controller
	mockUserRepository   *mocks.MockUserRepository
	mockPasswordService  *mocks.MockPasswordService
	userProfileUsecase   interfaces.UserProfileUpdateUsecase
}

func (suite *UserProfileUpdateTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockUserRepository = mocks.NewMockUserRepository(suite.ctrl)
	suite.mockPasswordService = mocks.NewMockPasswordService(suite.ctrl)
	suite.userProfileUsecase = usecases.NewUserProfileUpdateUsecase(
		suite.mockUserRepository,
		suite.mockPasswordService,
	)
}

func (suite *UserProfileUpdateTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *UserProfileUpdateTestSuite) TestUpdateUserProfile_Success() {
	userID := "user123"
	updateRequest := &dtos.ProfileUpdateRequest{
		Username: "newusername",
		Name:     "New Name",
		Password: "newpassword",
	}

	existingUser := &models.User{
		ID:       userID,
		Username: "oldusername",
		Name:     "Old Name",
		Password: "oldpassword",
	}

	hashedPassword := "hashednewpassword"

	suite.mockUserRepository.EXPECT().
		GetUserByID(context.Background(), userID).
		Return(existingUser, nil)

	suite.mockPasswordService.EXPECT().
		EncryptPassword(updateRequest.Password).
		Return(hashedPassword, nil)

	suite.mockUserRepository.EXPECT().
		UpdateUser(context.Background(), &models.User{
			ID:       userID,
			Username: updateRequest.Username,
			Name:     updateRequest.Name,
			Password: hashedPassword,
		}, userID).
		Return(nil)

	err := suite.userProfileUsecase.UpdateUserProfile(context.Background(), userID, updateRequest)

	if err != nil {
		suite.T().Errorf("Expected no error but got %v", err.Message)
	}
}

func (suite *UserProfileUpdateTestSuite) TestUpdateUserProfile_UserNotFound() {
	userID := "user123"
	updateRequest := &dtos.ProfileUpdateRequest{
		Username: "newusername",
	}

	suite.mockUserRepository.EXPECT().
		GetUserByID(context.Background(), userID).
		Return(nil, models.NotFound("User not found"))

	err := suite.userProfileUsecase.UpdateUserProfile(context.Background(), userID, updateRequest)

	suite.Error(err)
	suite.Equal("User not found", err.Message)
}

func (suite *UserProfileUpdateTestSuite) TestUpdateUserProfile_PasswordEncryptionError() {
	userID := "user123"
	updateRequest := &dtos.ProfileUpdateRequest{
		Password: "newpassword",
	}

	existingUser := &models.User{
		ID: userID,
	}

	suite.mockUserRepository.EXPECT().
		GetUserByID(context.Background(), userID).
		Return(existingUser, nil)

	suite.mockPasswordService.EXPECT().
		EncryptPassword(updateRequest.Password).
		Return("", models.InternalServerError("Encryption failed"))

	err := suite.userProfileUsecase.UpdateUserProfile(context.Background(), userID, updateRequest)

	suite.Error(err)
	suite.Equal("Something went wrong", err.Message)
}

func (suite *UserProfileUpdateTestSuite) TestUpdateUserProfile_UpdateError() {
	userID := "user123"
	updateRequest := &dtos.ProfileUpdateRequest{
		Username: "newusername",
	}

	existingUser := &models.User{
		ID: userID,
	}

	suite.mockUserRepository.EXPECT().
		GetUserByID(context.Background(), userID).
		Return(existingUser, nil)

	suite.mockUserRepository.EXPECT().
		UpdateUser(context.Background(), &models.User{
			ID:       userID,
			Username: updateRequest.Username,
		}, userID).
		Return(models.InternalServerError("Update failed"))

	err := suite.userProfileUsecase.UpdateUserProfile(context.Background(), userID, updateRequest)

	suite.Error(err)
	suite.Equal("Update failed", err.Message)
}

func TestUserProfileUpdateTestSuite(t *testing.T) {
	suite.Run(t, new(UserProfileUpdateTestSuite))
}

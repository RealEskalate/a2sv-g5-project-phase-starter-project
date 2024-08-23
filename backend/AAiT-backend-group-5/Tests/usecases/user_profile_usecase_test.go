package usecases_test

// import (
// 	"context"
// 	"testing"

// 	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
// 	models "github.com/aait.backend.g5.main/backend/Domain/Models"
// 	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
// 	mocks "github.com/aait.backend.g5.main/backend/Mocks"
// 	usecases "github.com/aait.backend.g5.main/backend/UseCases"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/suite"
// )

// type UserProfileUpdateUsecaseTestSuite struct {
// 	suite.Suite
// 	repositoryMock    *mocks.MockUserRepository
// 	passwordServiceMock *mocks.MockPasswordService
// 	userProfileUpdateUsecase interfaces.UserProfileUpdateUsecase
// 	ctrl               *gomock.Controller
// }

// func (suite *UserProfileUpdateUsecaseTestSuite) SetupTest() {
// 	suite.ctrl = gomock.NewController(suite.T())
// 	suite.repositoryMock = mocks.NewMockUserRepository(suite.ctrl)
// 	suite.passwordServiceMock = mocks.NewMockPasswordService(suite.ctrl)
// 	suite.userProfileUpdateUsecase = usecases.NewUserProfileUpdateUsecase(
// 		suite.repositoryMock,
// 		suite.passwordServiceMock,
// 	)
// }

// func (suite *UserProfileUpdateUsecaseTestSuite) TearDownTest() {
// 	suite.ctrl.Finish()
// }

// func (suite *UserProfileUpdateUsecaseTestSuite) TestUpdateUserProfile_Success() {
// 	ctx := context.Background()
// 	userID := "user1"
// 	password := "new_password"
// 	hashedPassword := "hashed_password"

// 	updateRequest := &dtos.ProfileUpdateRequest{
// 		Password: password,
// 	}

// 	existingUser := &models.User{
// 		ID: userID,
// 	}

// 	suite.repositoryMock.
// 		EXPECT().
// 		GetUserByID(ctx, userID).
// 		Return(existingUser, nil)

// 	suite.passwordServiceMock.
// 		EXPECT().
// 		EncryptPassword(password).
// 		Return(hashedPassword, nil)

// 	suite.repositoryMock.
// 		EXPECT().
// 		UpdateUser(ctx, existingUser, userID).
// 		Return(nil)

// 	err := suite.userProfileUpdateUsecase.UpdateUserProfile(ctx, userID, updateRequest)
// 	suite.Nil(err)
// 	suite.Equal(hashedPassword, existingUser.Password)
// }

// func (suite *UserProfileUpdateUsecaseTestSuite) TestUpdateUserProfile_NoPasswordChange() {
// 	ctx := context.Background()
// 	userID := "user1"

// 	updateRequest := &dtos.ProfileUpdateRequest{
// 		Password: "",
// 	}

// 	existingUser := &models.User{
// 		ID: userID,
// 	}

// 	suite.repositoryMock.
// 		EXPECT().
// 		GetUserByID(ctx, userID).
// 		Return(existingUser, nil)

// 	suite.repositoryMock.
// 		EXPECT().
// 		UpdateUser(ctx, existingUser, userID).
// 		Return(nil)

// 	err := suite.userProfileUpdateUsecase.UpdateUserProfile(ctx, userID, updateRequest)
// 	suite.Nil(err)
// 	suite.Empty(existingUser.Password)
// }

// func (suite *UserProfileUpdateUsecaseTestSuite) TestUpdateUserProfile_GetUserError() {
// 	ctx := context.Background()
// 	userID := "user1"

// 	updateRequest := &dtos.ProfileUpdateRequest{
// 		Password: "new_password",
// 	}

// 	suite.repositoryMock.
// 		EXPECT().
// 		GetUserByID(ctx, userID).
// 		Return(nil, models.InternalServerError("Error fetching user"))

// 	err := suite.userProfileUpdateUsecase.UpdateUserProfile(ctx, userID, updateRequest)
// 	suite.Equal(models.InternalServerError("Error fetching user"), err)
// }

// func (suite *UserProfileUpdateUsecaseTestSuite) TestUpdateUserProfile_EncryptPasswordError() {
// 	ctx := context.Background()
// 	userID := "user1"
// 	password := "new_password"

// 	updateRequest := &dtos.ProfileUpdateRequest{
// 		Password: password,
// 	}

// 	existingUser := &models.User{
// 		ID: userID,
// 	}

// 	suite.repositoryMock.
// 		EXPECT().
// 		GetUserByID(ctx, userID).
// 		Return(existingUser, nil)

// 	suite.passwordServiceMock.
// 		EXPECT().
// 		EncryptPassword(password).
// 		Return("", models.InternalServerError("Error encrypting password"))

// 	err := suite.userProfileUpdateUsecase.UpdateUserProfile(ctx, userID, updateRequest)
// 	suite.Equal(models.InternalServerError("Something went wrong"), err)
// }

// func (suite *UserProfileUpdateUsecaseTestSuite) TestUpdateUserProfile_UpdateUserError() {
// 	ctx := context.Background()
// 	userID := "user1"
// 	password := "new_password"
// 	hashedPassword := "hashed_password"

// 	updateRequest := &dtos.ProfileUpdateRequest{
// 		Password: password,
// 	}

// 	existingUser := &models.User{
// 		ID: userID,
// 	}

// 	suite.repositoryMock.
// 		EXPECT().
// 		GetUserByID(ctx, userID).
// 		Return(existingUser, nil)

// 	suite.passwordServiceMock.
// 		EXPECT().
// 		EncryptPassword(password).
// 		Return(hashedPassword, nil)

// 	suite.repositoryMock.
// 		EXPECT().
// 		UpdateUser(ctx, existingUser, userID).
// 		Return(models.InternalServerError("Error updating user"))

// 	err := suite.userProfileUpdateUsecase.UpdateUserProfile(ctx, userID, updateRequest)
// 	suite.Equal(models.InternalServerError("Error updating user"), err)
// }

// func TestUserProfileUpdateUsecaseTestSuite(t *testing.T) {
// 	suite.Run(t, new(UserProfileUpdateUsecaseTestSuite))
// }

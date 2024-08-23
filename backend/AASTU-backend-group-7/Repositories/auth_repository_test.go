package Repositories_test

import (
	"blogapp/Domain"
	"blogapp/Dtos"
	"blogapp/Infrastructure/password_services"
	"blogapp/mocks"
	"context"
	"errors"
	"net/http"
	"testing"

	repo "blogapp/Repositories"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepositoryTestSuite struct {
	suite.Suite
	repo            Domain.AuthRepository
	usercollection  *mocks.Collection
	tokencollection *mocks.Collection
	userRepo        *mocks.UserRepository
}

func (suite *AuthRepositoryTestSuite) SetupTest() {

	// Initialize the repository and mock dependencies
	suite.usercollection = new(mocks.Collection)
	suite.tokencollection = new(mocks.Collection)
	suite.userRepo = new(mocks.UserRepository)
	suite.repo = repo.NewAuthRepository(suite.usercollection, suite.tokencollection, suite.userRepo)

	suite.userRepo.On("SendActivationEmail", mock.Anything, mock.Anything).Return(nil, http.StatusOK)
	suite.userRepo.On("CheckPasswordStrength", mock.Anything).Return(nil)
	suite.userRepo.On("GenerateTokenFromUser", mock.Anything, mock.Anything).Return(Domain.AccessClaims{}, nil)
}

func (suite *AuthRepositoryTestSuite) TearDownTest() {
	suite.usercollection = nil
	suite.tokencollection = nil
	suite.userRepo = nil
	// Cleanup resources if needed
}

func (suite *AuthRepositoryTestSuite) TestLogin() {
	ctx := context.Background()

	suite.Run("InvalidCredentials", func() {
		// Mock the UserCollection to return an error when finding the user
		expectedUser := Domain.User{ID: primitive.NewObjectID(), Role: "admin"}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(errors.New("user not found")).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.User)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult)

		user := &Dtos.LoginUserDto{
			Email:    "test@example.com",
			Password: "password123",
		}
		// Call the method under test
		tokens, err, status := suite.repo.Login(ctx, user)

		// Assertions
		suite.Error(err)
		suite.Equal(http.StatusBadRequest, status)
		suite.Equal(Domain.Tokens{}, tokens)
	})

	suite.Run("ValidCredentials", func() {
		// Mock the UserCollection to return an existing user
		expectedUser := &Domain.User{
			ID:            primitive.NewObjectID(),
			Email:         "test@example.com",
			Password:      "password123",
			Role:          "admin",
			EmailVerified: true,
		}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.User)
			*userPtr = *expectedUser
		})

		suite.usercollection.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult)

		// Mock the CompareHashAndPasswordCustom method
		password_services.CompareHashAndPasswordCustom = func(hashedPassword string, plainPassword string) bool {
			return true
		}

		// Call the method under test
		tokens, err, status := suite.repo.Login(ctx, &Dtos.LoginUserDto{
			Email:    "test@example.com",
			Password: "password123",
		})

		// Assertions
		suite.Error(err)
		suite.NotEqual(http.StatusOK, status)
		suite.Equal(Domain.Tokens{}, tokens)
	})

}
func (suite *AuthRepositoryTestSuite) TestRegister_InvalidUserInput() {
	ctx := context.TODO()
	user := &Dtos.RegisterUserDto{
		Email:    "invalid-email",
		Password: "weakpassword",
		UserName: "testuser",
	}

	result, err, status := suite.repo.Register(ctx, user)

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, status)
	suite.Equal(&Domain.OmitedUser{}, result)
}

func (suite *AuthRepositoryTestSuite) TestRegister_EmailAlreadyTaken() {
	ctx := context.TODO()
	user := &Dtos.RegisterUserDto{
		Email:    "test@example.com",
		Password: "StrongPassword123!",
		UserName: "testuser",
	}

	suite.usercollection.On("CountDocuments", ctx, mock.Anything).Return(int64(1), nil)

	result, err, status := suite.repo.Register(ctx, user)

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, status)
	suite.Equal(&Domain.OmitedUser{}, result)
}

func (suite *AuthRepositoryTestSuite) TestRegister_Success() {
	ctx := context.TODO()
	user := &Dtos.RegisterUserDto{
		Email:    "test@example.com",
		Password: "StrongPassword123!",
		UserName: "testuser",
	}

	suite.usercollection.On("CountDocuments", ctx, mock.Anything).Return(int64(0), nil)
	suite.userRepo.On("CheckPasswordStrength", user.Password).Return(nil)
	suite.userRepo.On("GenerateFromPasswordCustom", user.Password).Return([]byte("hashedpassword"), nil)
	suite.usercollection.On("InsertOne", ctx, user).Return(&mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}, nil)

	mockSingleResult := new(mocks.SingleResult)
	mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		userPtr := args.Get(0).(*Domain.OmitedUser)
		*userPtr = Domain.OmitedUser{Email: user.Email}
	})
	suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)
	suite.userRepo.On("SendActivationEmail", user.Email).Return(nil, http.StatusOK)

	result, err, status := suite.repo.Register(ctx, user)

	suite.NoError(err)
	suite.Equal(http.StatusOK, status)
	suite.Equal(user.Email, result.Email)
}

func (suite *AuthRepositoryTestSuite) TestLogout() {
	ctx := context.Background()
	userID := primitive.NewObjectID()

	suite.Run("DeleteTokenError", func() {
		// Mock the TokenRepository to return an error when deleting the token
		err := errors.New("delete token error")
		statusCode := http.StatusInternalServerError
		count := int64(0)
		// Mock the DeleteOne method
		delResult := &mongo.DeleteResult{DeletedCount: count}
		suite.usercollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()
		suite.tokencollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()
		// Call the method under test
		err, status := suite.repo.Logout(ctx, userID)

		// Assertions
		suite.Equal(err, err)
		suite.NotEqual(statusCode, status)
	})

	suite.Run("DeleteTokenSuccess", func() {
		// Mock the TokenRepository to return no error when deleting the token
		statusCode := http.StatusOK

		count := int64(1)
		delResult := &mongo.DeleteResult{DeletedCount: count}
		suite.usercollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()
		suite.tokencollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()
		// Call the method under test
		err, status := suite.repo.Logout(ctx, userID)

		// Assertions
		suite.Equal(err, err)
		suite.Equal(statusCode, status)
	})
}

func TestAuthRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AuthRepositoryTestSuite))
}

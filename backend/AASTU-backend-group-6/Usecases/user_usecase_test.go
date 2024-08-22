package usecases

import (
	domain "blogs/Domain"
	"blogs/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	UserUsecase          domain.UserUseCase
	mockUserRepo         *mocks.UserRepository
	contextTimeout       time.Duration
	mockPasswordService  *mocks.PasswordService
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.mockUserRepo = new(mocks.UserRepository)
	suite.contextTimeout = time.Second * 5
	suite.mockPasswordService = new(mocks.PasswordService)
	suite.UserUsecase = NewUserUseCase(suite.mockUserRepo, suite.contextTimeout , suite.mockPasswordService)
}

func (suite *UserUsecaseTestSuite) TestPromoteandDemoteUser() {
	// passed
	suite.Run("Unauthorized", func() {
		// Setup
		userId := primitive.NewObjectID().Hex()
		promotion := domain.UserPromotionRequest{Action: "promote"}
		ctx := context.Background()

		// Execute
		result := suite.UserUsecase.PromoteandDemoteUser(ctx, userId, promotion, "user")

		// Assert
		suite.IsType(&domain.ErrorResponse{}, result)
		suite.Equal(&domain.ErrorResponse{Message: "Unauthorized to promote/demote user", Status: 403}, result)
	})

	// passed
	suite.Run("UserNotFound", func() {
		// Setup
		userId := primitive.NewObjectID().Hex()
		promotion := domain.UserPromotionRequest{Action: "promote"}
		ctx := context.Background()

		// Mock
		suite.mockUserRepo.On("FindUserByID", mock.Anything, userId).Return(domain.User{}, errors.New("User not found")).Once()

		// Execute
		result := suite.UserUsecase.PromoteandDemoteUser(ctx, userId, promotion, "admin")

		// Assert
		suite.IsType(&domain.ErrorResponse{}, result)
		suite.Equal(&domain.ErrorResponse{Message: "User not found", Status: 404}, result)
	})
	
	// passed
	suite.Run("DemoteUser_Success", func() {
		// Setup
		userId := primitive.NewObjectID().Hex()
		promotion := domain.UserPromotionRequest{Action: "demote"}
		ctx := context.Background()

		// Mock
		suite.mockUserRepo.On("FindUserByID", mock.Anything, userId).Return(domain.User{}, nil).Once()
		suite.mockUserRepo.On("PromoteandDemoteUser", mock.Anything, userId, "user").Return(nil).Once()

		// Execute
		result := suite.UserUsecase.PromoteandDemoteUser(ctx, userId, promotion, "admin")

		// Assert
		suite.IsType(&domain.SuccessResponse{}, result)
		suite.Equal(&domain.SuccessResponse{Message: "User demoted successfully", Status: 200}, result)
	})
	// passed
	suite.Run("PromoteUser_Success", func() {
		// Setup
		userId := primitive.NewObjectID().Hex()
		promotion := domain.UserPromotionRequest{Action: "promote"}
		ctx := context.Background()

		// Mock
		suite.mockUserRepo.On("FindUserByID", mock.Anything, userId).Return(domain.User{}, nil).Once()
		suite.mockUserRepo.On("PromoteandDemoteUser", mock.Anything, userId, "admin").Return(nil).Once()

		// Execute
		result := suite.UserUsecase.PromoteandDemoteUser(ctx, userId, promotion, "admin")

		// Assert
		suite.IsType(&domain.SuccessResponse{}, result)
		suite.Equal(&domain.SuccessResponse{Message: "User promoted successfully", Status: 200}, result)
	})

	suite.Run("InvalidAction", func() {
		// Setup
		userId := primitive.NewObjectID().Hex()
		promotion := domain.UserPromotionRequest{Action: "invalid_action"}
		ctx := context.Background()

		suite.mockUserRepo.On("FindUserByID", mock.Anything, userId).Return(domain.User{}, errors.New("Invalid action")).Once()
		// Execute
		result := suite.UserUsecase.PromoteandDemoteUser(ctx, userId, promotion, "admin")

		// Assert
		suite.IsType(&domain.ErrorResponse{}, result)
		suite.Equal(&domain.ErrorResponse{Message: "Invalid action", Status: 404}, result)
	})

	suite.Run("PromoteUser_Error", func() {
		// Setup
		userId := primitive.NewObjectID().Hex()
		promotion := domain.UserPromotionRequest{Action: "promote"}
		ctx := context.Background()

		// Mock
		suite.mockUserRepo.On("FindUserByID", mock.Anything, userId).Return(domain.User{}, nil).Once()
		suite.mockUserRepo.On("PromoteandDemoteUser", mock.Anything, userId, "admin").Return(errors.New("some error")).Once()

		// Execute
		result := suite.UserUsecase.PromoteandDemoteUser(ctx, userId, promotion, "admin")

		// Assert
		suite.IsType(&domain.ErrorResponse{}, result)
		suite.Equal(&domain.ErrorResponse{Message: "some error", Status: 500}, result)
	})
}

func (suite *UserUsecaseTestSuite) TestUpdateUser() {
	
	suite.Run("UserNotFound", func() {
		// Setup
		req := domain.UserUpdateRequest{
			ID:       primitive.NewObjectID().Hex(),
			Username: "john",
		}
		ctx := context.Background()

		// Mock
		suite.mockUserRepo.On("FindUserByID", mock.Anything, req.ID).Return(domain.User{}, errors.New("User not found")).Once()

		// Execute
		result := suite.UserUsecase.UpdateUser(ctx, req)

		// Assert
		suite.IsType(&domain.ErrorResponse{}, result)
		suite.Equal(&domain.ErrorResponse{Message: "User not found", Status: 404}, result)
	})

	// passed
	suite.Run("UnauthorizedToUpdateUser", func() {
		// Setup
		req := domain.UserUpdateRequest{
			ID:       primitive.NewObjectID().Hex(),
			Username: "john",
		}
		ctx := context.Background()
		user := domain.User{
			ID:       primitive.NewObjectID(),
			Username: "jane",
		}

		// Mock
		suite.mockUserRepo.On("FindUserByID", mock.Anything, req.ID).Return(user, nil).Once()

		// Execute
		result := suite.UserUsecase.UpdateUser(ctx, req)

		// Assert
		suite.IsType(&domain.ErrorResponse{}, result)
		suite.Equal(&domain.ErrorResponse{Message: "Unauthorized to update this user", Status: 403}, result)
	})
	
	// passed
		suite.Run("UsernameAlreadyTaken", func() {
		// Setup
		id := primitive.NewObjectID().Hex()
		objectid, _ := primitive.ObjectIDFromHex(id)
		req := domain.UserUpdateRequest{
			ID:       id,
			Username: "john",
		}
		ctx := context.Background()
		user := domain.User{
			ID:       objectid, // Ensure this ID matches `req.ID`
			Username: "john",
		}

		previousUser := domain.User{
			ID: 	  primitive.NewObjectID(),
			Username: "john",
		}

		// Mock
		suite.mockUserRepo.On("FindUserByID", mock.Anything, req.ID).Return(user, nil).Once()
		// Simulate finding a user with the same username
		suite.mockUserRepo.On("FindUserByUsername", mock.Anything, req.Username).Return(previousUser, nil).Once()
		suite.mockUserRepo.On("UpdateUser" , mock.Anything , user).Return(domain.User{} , nil).Once()
		// Execute
		result := suite.UserUsecase.UpdateUser(ctx, req)

		// Assert
		suite.IsType(&domain.ErrorResponse{}, result)
		suite.Equal(&domain.ErrorResponse{Message: "Username is already taken", Status: 409}, result)
	})

	suite.Run("UpdateUser_Success", func() {
    // Setup
    id := primitive.NewObjectID().Hex()
    objid, _ := primitive.ObjectIDFromHex(id)
    req := domain.UserUpdateRequest{
        ID:               id,
        Username:         "john",
        Full_Name:        "John Doe",
        Password:         "password",
        Profile_image_url: "https://example.com/profile.jpg",
        Contact:          "john@example.com",
        Bio:              "Hello, I'm John",
    }
    ctx := context.Background()
    user := domain.User{
        ID:       objid,
        Username: "john",
    }

    // Mock: Find the existing user by ID
    suite.mockUserRepo.On("FindUserByID", mock.Anything, req.ID).Return(user, nil).Once()

    // Mock: Check if the username is already taken
    suite.mockUserRepo.On("FindUserByUsername", mock.Anything, req.Username).Return(user, nil).Once()

    // Mock: Validate the password
    suite.mockPasswordService.On("ValidatePassword", req.Password).Return(nil).Once()

    // Mock: Hash the new password
    suite.mockPasswordService.On("HashPassword", req.Password).Return("hashedpassword", nil).Once()

    // Prepare the updated user object with all fields modified as per the request
    updatedUser := domain.User{
        ID:               objid,
        Full_Name:        "John Doe",
        Username:         "john",
        Password:         "", // Expect the password to be empty in the returned user object
        Profile_image_url: "https://example.com/profile.jpg",
        Contact:          "john@example.com",
        Bio:              "Hello, I'm John",
    }

    // Mock: Update the user in the repository
    suite.mockUserRepo.On("UpdateUser", mock.Anything, mock.MatchedBy(func(u domain.User) bool {
        // Ensure that the password is hashed before updating in the database
        return u.Password == "hashedpassword"
    })).Return(updatedUser, nil).Once()

    // Execute
    result := suite.UserUsecase.UpdateUser(ctx, req)

    // Assert
    suite.IsType(&domain.SuccessResponse{}, result)
    suite.Equal(&domain.SuccessResponse{Message: "User updated successfully", Status: 200, Data: updatedUser}, result)
})


}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
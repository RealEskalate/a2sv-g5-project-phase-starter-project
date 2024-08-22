package controllers_test

import (
	"blogapp/Domain"
	"blogapp/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "blogapp/Delivery/controllers"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileControllerTestSuite struct {
	suite.Suite
	mockProfileUseCase *mocks.ProfileUseCases
	profileController  *controllers.Profile_controller
	patch              *monkey.PatchGuard
	userID             primitive.ObjectID
	Role               string
}

func (suite *ProfileControllerTestSuite) SetupTest() {
	suite.mockProfileUseCase = new(mocks.ProfileUseCases)
	var err error
	suite.profileController = controllers.NewProfileController(suite.mockProfileUseCase)
	assert.NoError(suite.T(), err)
	suite.userID = user_id
	suite.Role = "admin"
	suite.patch = monkey.Patch(controllers.Getclaim, mockExtractUser)
}

func (suite *ProfileControllerTestSuite) TestGetProfile() {
	// Arrange
	loged_user := Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	}
	// Create a mock user with type Domain.OmitedUser
	mockUser := Domain.OmitedUser{}

	suite.mockProfileUseCase.On("GetProfile", mock.Anything, loged_user.ID, loged_user).Return(mockUser, nil, 200).Once()

	req, _ := http.NewRequest(http.MethodGet, "/me", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: loged_user.ID.Hex()},
	}
	c.Set("claim", &Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	})

	// Act
	suite.profileController.GetProfile(c)

	// Assert
	suite.Equal(http.StatusOK, w.Code)
	suite.mockProfileUseCase.AssertExpectations(suite.T())
}

func (suite *ProfileControllerTestSuite) TestUpdateProfile() {
	// Arrange
	userID := suite.userID
	user := Domain.User{
		ID:    userID,
		Email: "updated@example.com",
	}

	updatedUser := Domain.OmitedUser{
		ID:    userID,
		Email: "updated@example.com",
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPut, "/me", bytes.NewBuffer(jsonUser))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: userID.Hex()},
	}

	c.Set("claim", &Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	})

	// Ensure the mock expectation matches the actual call
	suite.mockProfileUseCase.On("UpdateProfile", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(updatedUser, nil, 200).Once()

	// Act
	suite.profileController.UpdateProfile(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockProfileUseCase.AssertExpectations(suite.T())
}
func (suite *ProfileControllerTestSuite) TestDeleteProfile() {
	// Arrange
	currentUser := Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	}

	req, _ := http.NewRequest(http.MethodDelete, "/me", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: currentUser.ID.Hex()},
	}
	c.Set("claim", &Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	})

	suite.mockProfileUseCase.On("DeleteProfile", c, mock.Anything, mock.Anything).Return(nil, 200).Once()

	// Act
	suite.profileController.DeleteProfile(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	// assert.Equal(suite.T(), "{\"message\":\"User deleted successfully\"}", w.Body.String())
	suite.mockProfileUseCase.AssertExpectations(suite.T())
}

func TestProfileController(t *testing.T) {
	suite.Run(t, new(ProfileControllerTestSuite))
}

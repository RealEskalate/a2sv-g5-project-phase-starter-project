package controller_test

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/entities"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type MockProfileService struct {
	mock.Mock
}

func (m *MockProfileService) GetUserProfile(userId string) (*entities.Profile, error) {
	args := m.Called(userId)
	profile, ok := args.Get(0).(*entities.Profile)
	if !ok {
		return nil, args.Error(1)
	}
	return profile, args.Error(1)
}

func (m *MockProfileService) UpdateUserProfile(profile *dto.UpdateProfileDto) (*entities.Profile, error) {
	args := m.Called(profile)
	profile_, ok := args.Get(0).(*entities.Profile)
	if !ok {
		return nil, args.Error(1)
	}
	return profile_, args.Error(1)
}

func (m *MockProfileService) CreateUserProfile(profile *dto.CreateProfileDto) (*entities.Profile, error) {
	args := m.Called(profile)
	profile_, ok := args.Get(0).(*entities.Profile)
	if !ok {
		return nil, args.Error(1)
	}
	return profile_, args.Error(1)
}

func (m *MockProfileService) DeleteUserProfile(user_id string) error {
	args := m.Called(user_id)
	return args.Error(0)
}

type ProfileControllerTestSuite struct {
	suite.Suite
	controller    controller.ProfileController
	mockService   *MockProfileService
}

func (suite *ProfileControllerTestSuite) SetupTest() {
	suite.mockService = new(MockProfileService)
	suite.controller = controller.NewProfileController(suite.mockService)
}

func (suite *ProfileControllerTestSuite) TestCreateUserProfile_Success() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userId", "user123")

	jsonBody := `{"profilePicture": "pic.com", "bio": "new bio", "phoneNumber": "0912131415", "email": "example@example.com", "address": "Some Address"}`
	c.Request = httptest.NewRequest("POST", "/profile", strings.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	profileDto := &dto.CreateProfileDto{
		UserID:         "user123",
		ProfilePicture: "pic.com",
		Bio:            "new bio",
		PhoneNumber:    "0912131415",
		Email:          "example@example.com",
		Address:        "Some Address",
	}

	objId := primitive.NewObjectID()
	expectedProfile := &entities.Profile{
		UserID:         objId,
		ProfilePicture: "pic.com",
		Bio:            "new bio",
		ContactInfo:    entities.ContactInfo{PhoneNumber: "0912131415", Email: "example@example.com", Address: "Some Address"},
		UpdatedAt:      time.Now(),
	}
	suite.mockService.On("CreateUserProfile", profileDto).Return(expectedProfile, nil)

	suite.controller.CreateUserProfile(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *ProfileControllerTestSuite) TestCreateUserProfile_Fail() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userId", "user123")

	jsonBody := `{"profilePicture": "pic.com", "bio": "new bio", "phoneNumber": "0912131415", "email": "example@example.com", "address": "Some Address"}`
	c.Request = httptest.NewRequest("POST", "/profile", strings.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	profileDto := &dto.CreateProfileDto{
		UserID:         "user123",
		ProfilePicture: "pic.com",
		Bio:            "new bio",
		PhoneNumber:    "0912131415",
		Email:          "example@example.com",
		Address:        "Some Address",
	}

	suite.mockService.On("CreateUserProfile", profileDto).Return(nil, errors.New("some error"))

	suite.controller.CreateUserProfile(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *ProfileControllerTestSuite) TestGetUserProfile_Success() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "userId", Value: "user123"}}
	objId := primitive.NewObjectID()
	expectedProfile := &entities.Profile{
		UserID:         objId,
		ProfilePicture: "pic.com",
		Bio:            "new bio",
		ContactInfo:    entities.ContactInfo{PhoneNumber: "0912131415", Email: "example@example.com", Address: "Some Address"},
		UpdatedAt:      time.Now(),
	}
	suite.mockService.On("GetUserProfile", "user123").Return(expectedProfile, nil)

	suite.controller.GetUserProfile(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *ProfileControllerTestSuite) TestGetUserProfile_Fail() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "userId", Value: "user123"}}

	suite.mockService.On("GetUserProfile", "user123").Return(nil, errors.New("some error"))

	suite.controller.GetUserProfile(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *ProfileControllerTestSuite) TestUpdateUserProfile_Success() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "userId", Value: "user123"}}
	c.Set("userId", "user123")

	jsonBody := `{"profilePicture": "pic.com", "bio": "new bio", "phoneNumber": "0912131415", "address": "Some Address"}`
	c.Request = httptest.NewRequest("PUT", "/profile/user123", strings.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	profileDto := &dto.UpdateProfileDto{
		UserID:         "user123",
		ProfilePicture: "pic.com",
		Bio:            "new bio",
		Address:        "Some Address",
	}

	objId := primitive.NewObjectID()
	updatedProfile := &entities.Profile{
		UserID:         objId,
		ProfilePicture: "pic.com",
		Bio:            "new bio",
		ContactInfo:    entities.ContactInfo{PhoneNumber: "0912131415", Address: "Some Address"},
		UpdatedAt:      time.Now(),
	}

	suite.mockService.On("UpdateUserProfile", profileDto).Return(updatedProfile, nil)

	suite.controller.UpdateUserProfile(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *ProfileControllerTestSuite) TestUpdateUserProfile_Fail() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "userId", Value: "user123"}}
	c.Set("userId", "user123")

	jsonBody := `{"profilePicture": "pic.com", "bio": "new bio", "phoneNumber": "0912131415", "address": "Some Address"}`
	c.Request = httptest.NewRequest("PUT", "/profile/user123", strings.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	profileDto := &dto.UpdateProfileDto{
		UserID:         "user123",
		ProfilePicture: "pic.com",
		Bio:            "new bio",
		Address:        "Some Address",
	}

	suite.mockService.On("UpdateUserProfile", profileDto).Return(nil, errors.New("some error"))

	suite.controller.UpdateUserProfile(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *ProfileControllerTestSuite) TestDeleteUserProfile_Success() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "userId", Value: "user123"}}
	c.Set("userId", "user123")

	suite.mockService.On("DeleteUserProfile", "user123").Return(nil)

	suite.controller.DeleteUserProfile(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *ProfileControllerTestSuite) TestDeleteUserProfile_Fail() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "userId", Value: "user123"}}
	c.Set("userId", "user123")

	suite.mockService.On("DeleteUserProfile", "user123").Return(errors.New("some error"))

	suite.controller.DeleteUserProfile(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func TestProfileControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ProfileControllerTestSuite))
	}
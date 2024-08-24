package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/domain"
	"aait.backend.g10/tests/mocks"
	"aait.backend.g10/usecases/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserControllerSuite struct {
	suite.Suite
	controller  *controllers.UserController
	mockUseCase *mocks.IUserUseCase
}

func (s *UserControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.mockUseCase = new(mocks.IUserUseCase)
	s.controller = controllers.NewUserController(s.mockUseCase)
}

func (s *UserControllerSuite) TearDownTest() {
	s.mockUseCase.AssertExpectations(s.T())
}

func (s *UserControllerSuite) TestPromoteUser_Success() {
	userID := uuid.New()
	isPromote := true

	s.mockUseCase.On("PromoteUser", userID, isPromote).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/promote", bytes.NewBufferString(`{
		"id": "`+userID.String()+`",
		"is_promote": true
	}`))
	c.Request.Header.Set("Content-Type", "application/json")

	s.controller.PromoteUser(c)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Equal(s.T(), `{"message":"User promoted successfully"}`, w.Body.String())
}

func (s *UserControllerSuite) TestPromoteUser_InvalidJSON() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/promote", bytes.NewBufferString(`{invalid-json}`))
	c.Request.Header.Set("Content-Type", "application/json")

	s.controller.PromoteUser(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
	assert.Contains(s.T(), w.Body.String(), "invalid character")
}

func (s *UserControllerSuite) TestPromoteUser_Failure() {
	userID := uuid.New()
	isPromote := true

	s.mockUseCase.On("PromoteUser", userID, isPromote).Return(&domain.CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to promote user",
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/promote", bytes.NewBufferString(`{
		"id": "`+userID.String()+`",
		"is_promote": true
	}`))
	c.Request.Header.Set("Content-Type", "application/json")

	s.controller.PromoteUser(c)

	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
	assert.Equal(s.T(), `{"error":"failed to promote user"}`, w.Body.String())
}

func (s *UserControllerSuite) TestUpdateProfile_Success() {
	requesterID := uuid.New()
	userUpdate := dto.UserUpdate{
		ID:       requesterID,
		FullName: "Updated Name",
		Bio:      "Human",
	}

	s.mockUseCase.On("UpdateUser", requesterID, &userUpdate).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/user/"+requesterID.String(), bytes.NewBufferString(`{
		"fullname": "Updated Name",
		"bio": "Human"
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: requesterID.String()}}
	c.Set("id", requesterID.String())

	s.controller.UpdateProfile(c)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Equal(s.T(), `{"message":"User profile updated successfully"}`, w.Body.String())
}

func (s *UserControllerSuite) TestUpdateProfile_InvalidJSON() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/user/invalid-uuid", bytes.NewBufferString(`{invalid-json}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: uuid.New().String()}}

	s.controller.UpdateProfile(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
	assert.Contains(s.T(), w.Body.String(), "invalid character")
}

func (s *UserControllerSuite) TestUpdateProfile_InvalidUserID() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/user/invalid-uuid", nil)
	c.Params = gin.Params{{Key: "id", Value: "invalid-uuid"}}

	s.controller.UpdateProfile(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
	assert.Equal(s.T(), `{"error":"Invalid ID"}`, w.Body.String())
}

func (s *UserControllerSuite) TestUpdateProfile_Failure() {
	requesterID := uuid.New()
	userUpdate := dto.UserUpdate{
		ID:       requesterID,
		FullName: "Updated Name",
		Bio:      "Human",
	}

	s.mockUseCase.On("UpdateUser", requesterID, &userUpdate).Return(&domain.CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to update profile",
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/user/"+requesterID.String(), bytes.NewBufferString(`{
		"fullname": "Updated Name",
		"bio": "Human"
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: requesterID.String()}}
	c.Set("id", requesterID.String())

	s.controller.UpdateProfile(c)

	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
	assert.Equal(s.T(), `{"error":"failed to update profile"}`, w.Body.String())
}

func (s *UserControllerSuite) TestGetUserByID_Success() {
	userID := uuid.New()
	mockUser := dto.GetUserResponseDto{
		ID:       userID,
		FullName: "Test User",
		Email:    "test.user@example.com",
	}

	s.mockUseCase.On("GetUserByID", userID).Return(&mockUser, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/user/"+userID.String(), nil)
	c.Params = gin.Params{{Key: "id", Value: userID.String()}}

	s.controller.GetUserByID(c)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	var response dto.GetUserResponseDto
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockUser, response)
}

func (s *UserControllerSuite) TestGetUserByID_InvalidUserID() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/user/invalid-uuid", nil)
	c.Params = gin.Params{{Key: "id", Value: "invalid-uuid"}}

	s.controller.GetUserByID(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
	assert.Equal(s.T(), `{"error":"Invalid ID"}`, w.Body.String())
}

func (s *UserControllerSuite) TestGetUserByID_Failure() {
	userID := uuid.New()

	s.mockUseCase.On("GetUserByID", userID).Return(nil, &domain.CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to get user",
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/user/"+userID.String(), nil)
	c.Params = gin.Params{{Key: "id", Value: userID.String()}}

	s.controller.GetUserByID(c)

	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
	assert.Equal(s.T(), `{"error":"failed to get user"}`, w.Body.String())
}

func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}

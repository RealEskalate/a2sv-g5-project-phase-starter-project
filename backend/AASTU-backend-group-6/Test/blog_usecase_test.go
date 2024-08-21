package test

// package controllers_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	controllers "task_manager/Delivery/Controllers"
// 	domain "task_manager/Domain"
// 	"task_manager/mocks"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/suite"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type ControllerTestSuite struct {
// 	suite.Suite
// 	mockTaskUseCase *mocks.TaskUseCase
// 	mockUserUseCase *mocks.UserUseCase
// 	router          *gin.Engine
// 	controller *controllers.Controller
// }

// func MockAuthMiddleware(role string, userID string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Set("role", role)
// 		c.Set("user_id", userID)
// 		c.Next()
// 	}
// }

// func (suite *ControllerTestSuite) SetupTest() {
// 	suite.mockTaskUseCase = new(mocks.TaskUseCase)
// 	suite.mockUserUseCase = new(mocks.UserUseCase)
// 	suite.router = gin.Default()
// 	controller := controllers.NewController(suite.mockUserUseCase, suite.mockTaskUseCase)
// 	suite.controller = controller
// }
// func (suite *ControllerTestSuite) TestGetAllTaskHandler_UserSuccess() {
// 	mockTasks := []domain.Task{
// 		{
// 			ID: primitive.NewObjectID(),
// 			Title: "Nothing",
// 			Description: "Something",
// 			DueDate: time.Now(),
// 			Status: "Pending",
// 			UserId: primitive.NewObjectID(),
// 		},
// 		{
// 			ID: primitive.NewObjectID(),
// 			Title: "Nothing",
// 			Description: "Something",
// 			DueDate: time.Now(),
// 			Status: "In Progress",
// 			UserId: primitive.NewObjectID(),
// 		},
// 	}

// 	suite.router.GET("/tasks", MockAuthMiddleware("USER", "user123"), suite.controller.GetAllTaskHandler)

// 	suite.mockTaskUseCase.On("GetAllTasks", "USER", "user123").Return(mockTasks, nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/tasks", nil)
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.Contains(w.Body.String(), "Nothing")
// 	suite.Contains(w.Body.String(), "Something")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestGetAllTaskHandler_UserNoTasks() {
// 	suite.router.GET("/tasks", MockAuthMiddleware("USER", "user123"), suite.controller.GetAllTaskHandler)
// 	suite.mockTaskUseCase.On("GetAllTasks", "USER", "user123").Return([]domain.Task{}, nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/tasks", nil)
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.Contains(w.Body.String(), "Task not found")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }
// func (suite *ControllerTestSuite) TestGetAllTaskHandler_AdminSuccess() {
// 	mockTasks := []domain.Task{
// 		{
// 			ID: primitive.NewObjectID(),
// 			Title: "Admin Task 1",
// 			Description: "Admin Description 1",
// 			DueDate: time.Now(),
// 			Status: "Pending",
// 			UserId: primitive.NewObjectID(),
// 		},
// 	}
// 	suite.router.GET("/tasks", MockAuthMiddleware("ADMIN", "admin123"), suite.controller.GetAllTaskHandler)
// 	suite.mockTaskUseCase.On("GetAllTasks", "ADMIN", "").Return(mockTasks, nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/tasks", nil)
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.Contains(w.Body.String(), "Admin Task 1")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestGetAllTaskHandler_AdminError() {
// 	suite.mockTaskUseCase.On("GetAllTasks", "ADMIN", "").Return(nil, errors.New("error fetching tasks"))
// 	suite.router.GET("/tasks", MockAuthMiddleware("ADMIN", "admin123"), suite.controller.GetAllTaskHandler)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/tasks", nil)
// 	suite.router.ServeHTTP(w, req)
// 	suite.Equal(http.StatusBadRequest, w.Code)
// 	suite.Contains(w.Body.String(), "error fetching tasks")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestTaskByIdHandler_UserSuccess() {
// 	task_id := primitive.NewObjectID()
// 	mockTask := domain.Task{
// 		ID: task_id,
// 		Title: "User Task",
// 		Description: "User Task Description",
// 		DueDate: time.Now(),
// 		Status: "Pending",
// 		UserId: primitive.NewObjectID(),
// 	}

// 	suite.mockTaskUseCase.On("GetTaskById", task_id.Hex(), "user123", "USER").Return(mockTask, nil)
// 	suite.router.GET("/tasks/:id",  MockAuthMiddleware("USER", "user123"), suite.controller.TaskByIdHandler)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/tasks/"+mockTask.ID.Hex(), nil)
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.Contains(w.Body.String(), "User Task")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestTaskByIdHandler_UserNotFound() {
// 	suite.mockTaskUseCase.On("GetTaskById", "id", "user123", "USER").Return(domain.Task{}, errors.New("task not found"))
// 	suite.router.GET("/tasks/:id", MockAuthMiddleware("USER", "user123"), suite.controller.TaskByIdHandler)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/tasks/id", nil)
// 	suite.router.ServeHTTP(w, req)
// 	suite.Equal(http.StatusNotFound, w.Code)
// 	suite.Contains(w.Body.String(), "task not found")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestTaskByIdHandler_AdminSuccess() {
// 	taskID := primitive.NewObjectID()
// 	mockTask := domain.Task{
// 		ID:          taskID,
// 		Title:       "Admin Task",
// 		Description: "Admin Task Description",
// 		DueDate:     time.Now(),
// 		Status:      "Pending",
// 		UserId:      primitive.NewObjectID(),
// 	}
// 	suite.router.GET("/tasks/:id", MockAuthMiddleware("ADMIN", "admin123"), suite.controller.TaskByIdHandler)
// 	suite.mockTaskUseCase.On("GetTaskById", taskID.Hex(), "", "ADMIN").Return(mockTask, nil)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/tasks/"+mockTask.ID.Hex(), nil)
// 	suite.router.ServeHTTP(w, req)
// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.Contains(w.Body.String(), "Admin Task")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestTaskByIdHandler_AdminNotFound() {
// 	suite.mockTaskUseCase.On("GetTaskById", "id", "", "ADMIN").Return(domain.Task{}, errors.New("task not found"))
// 	suite.router.GET("/tasks/:id", MockAuthMiddleware("ADMIN", "admin123"), suite.controller.TaskByIdHandler)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/tasks/id", nil)
// 	req.Header.Set("role", "ADMIN")
// 	req.Header.Set("user_id", "admin123")
// 	suite.router.ServeHTTP(w, req)

// 	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "task not found")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestCreateTaskHandler_UserSuccess() {
// 	task_id := primitive.NewObjectID()
// 	user_id := primitive.NewObjectID()
// 	fixedTime := time.Date(2024, time.August, 14, 10, 21, 23, 0, time.UTC)
// 	mockTask := domain.Task{
// 		ID: task_id,
// 		Title: "New Task",
// 		Description: "A new task to be completed",
// 		DueDate: fixedTime,
// 		Status: "Pending",
// 	}
// 	suite.mockTaskUseCase.On("CreateTask", mockTask, user_id.Hex(), "USER").Return(task_id, nil)
// 	suite.router.POST("/tasks", MockAuthMiddleware("USER", user_id.Hex()), suite.controller.CreateTaskHandler)
// 	body, _ := json.Marshal(mockTask)
// 	fmt.Println(bytes.NewBuffer(body))
// 	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	suite.router.ServeHTTP(w, req)
// 	suite.Equal(http.StatusCreated, w.Code)
// 	suite.Contains(w.Body.String(), task_id.Hex())
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestCreateTaskHandler_UserFailure() {
// 	task_id := primitive.NewObjectID()
// 	user_id := primitive.NewObjectID()
// 	fixedTime := time.Date(2024, time.August, 14, 10, 21, 23, 0, time.UTC)
// 	mockTask := domain.Task{
// 		ID: task_id,
// 		Title: "New Task",
// 		Description: "A new task to be completed",
// 		DueDate: fixedTime,
// 		Status: "Pending",
// 	}
// 	suite.mockTaskUseCase.On("CreateTask", mockTask, user_id.Hex(), "USER").Return(nil, errors.New("Error"))
// 	suite.router.POST("/tasks", MockAuthMiddleware("USER", user_id.Hex()), suite.controller.CreateTaskHandler)
// 	body, _ := json.Marshal(mockTask)
// 	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	suite.router.ServeHTTP(w, req)
// 	suite.Equal(http.StatusNotFound, w.Code)
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestUpdateTaskHandler_UserSuccess() {
// 	task_id := primitive.NewObjectID()
// 	user_id := primitive.NewObjectID()
// 	fixedTime := time.Date(2024, time.August, 14, 10, 21, 23, 0, time.UTC)
// 	mockTask := domain.Task{
// 		ID: task_id,
// 		Title: "New Task",
// 		Description: "A new task to be completed",
// 		DueDate: fixedTime,
// 		Status: "Pending",
// 		UserId: user_id,
// 	}
// 	suite.router.PUT("/tasks/:id", MockAuthMiddleware("USER", user_id.Hex()), suite.controller.UpdateTaskHandler)
// 	suite.mockTaskUseCase.On("UpdateTask", task_id.Hex(), user_id.Hex(), "USER", mockTask).Return(true, nil)
// 	body, _ := json.Marshal(mockTask)
// 	fmt.Println(bytes.NewBuffer(body))
// 	req, _ := http.NewRequest("PUT", "/tasks/"+task_id.Hex(), bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	suite.router.ServeHTTP(w, req)
// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.Contains(w.Body.String(), "Task updated successfully")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestCreateTaskHandler_Failure() {
// 	task_id := primitive.NewObjectID()
// 	user_id := primitive.NewObjectID()
// 	suite.router.PUT("/tasks/:id", MockAuthMiddleware("USER", user_id.Hex()), suite.controller.UpdateTaskHandler)
// 	suite.mockTaskUseCase.On("UpdateTask", task_id.Hex(), user_id.Hex(), "USER", domain.Task{}).Return(false, errors.New("Error"))
// 	body, _ := json.Marshal(domain.Task{})
// 	fmt.Println(bytes.NewBuffer(body))
// 	req, _ := http.NewRequest("PUT", "/tasks/"+task_id.Hex(), bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	suite.router.ServeHTTP(w, req)
// 	suite.Equal(http.StatusNotFound, w.Code)
// 	suite.Contains(w.Body.String(), "Error")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestDeleteTaskHandler_UserSuccess() {
// 	suite.mockTaskUseCase.On("DeleteTaskById", "id", "user123", "USER").Return(true, nil)
// 	suite.router.DELETE("/tasks/:id", MockAuthMiddleware("USER", "user123"), suite.controller.DeleteTaskHandler)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/tasks/id", nil)
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.Contains(w.Body.String(), "Task deleted")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestDeleteTaskHandler_UserNotFound() {
// 	suite.mockTaskUseCase.On("DeleteTaskById", "id", "user123", "USER").Return(false, nil)
// 	suite.router.DELETE("/tasks/:id", MockAuthMiddleware("USER", "user123"), suite.controller.DeleteTaskHandler)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/tasks/id", nil)
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusNotFound, w.Code)
// 	suite.Contains(w.Body.String(), "Task not found")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestDeleteTaskHandler_AdminSuccess() {
// 	suite.mockTaskUseCase.On("DeleteTaskById", "id", "", "ADMIN").Return(true, nil)
// 	suite.router.DELETE("/tasks/:id", MockAuthMiddleware("ADMIN", "admin123"), suite.controller.DeleteTaskHandler)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/tasks/id", nil)
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.Contains(w.Body.String(), "Task deleted")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestDeleteTaskHandler_AdminNotFound() {

// 	suite.mockTaskUseCase.On("DeleteTaskById", "id", "", "ADMIN").Return(false, nil)
// 	suite.router.DELETE("/tasks/:id", MockAuthMiddleware("ADMIN", "admin123"), suite.controller.DeleteTaskHandler)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/tasks/id", nil)
// 	suite.router.ServeHTTP(w, req)

// 	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "Task not found")
// 	suite.mockTaskUseCase.AssertExpectations(suite.T())
// 	suite.mockTaskUseCase.On("DeleteTaskById", "someid", "admin123", "ADMIN").Return(false, nil)
// }

// func (suite *ControllerTestSuite) TestUserRegisterHandler_UserSuccess() {
// 	mockUser := domain.User{
// 		UserName: "user123",
// 		Password: "password",
// 		Role:     "USER",
// 	}
// 	suite.mockUserUseCase.On("Register", mockUser).Return("user_id", nil)
// 	suite.router.POST("/register", suite.controller.UserRegisterHandler)
// 	w := httptest.NewRecorder()
// 	body, err := json.Marshal(mockUser)
// 	suite.Nil(err)

// 	req, _ := http.NewRequest("POST", "/register", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.Contains(w.Body.String(), "user_id")
// 	suite.Contains(w.Body.String(), "SignedUp Successfully.")
// 	suite.mockUserUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestUserRegisterHandler_UserInvalidCredentials() {
// 	suite.mockUserUseCase.On("Register", domain.User{}).Return(nil, errors.New("invalid credentials"))

// 	suite.router.POST("/register", suite.controller.UserRegisterHandler)
// 	w := httptest.NewRecorder()
// 	reqBody := `{"userName":"","password":"","role":""}`
// 	req, err := http.NewRequest("POST", "/register", bytes.NewReader([]byte(reqBody)))
// 	suite.Nil(err)
// 	req.Header.Set("Content-Type", "application/json")
// 	suite.router.ServeHTTP(w, req)
// 	suite.Equal(http.StatusBadRequest, w.Code)
// 	suite.Contains(w.Body.String(), "invalid credential")
// }

// func (suite *ControllerTestSuite) TestUserRegisterHandler_EmptyCredentials() {
// 	suite.mockUserUseCase.On("Register", domain.User{}).Return(nil, errors.New("invalid credentials"))
// 	suite.router.POST("/register", suite.controller.UserRegisterHandler)
// 	w := httptest.NewRecorder()
// 	req, err := http.NewRequest("POST", "/register", nil)
// 	suite.Nil(err)
// 	req.Header.Set("Content-Type", "application/json")
// 	suite.router.ServeHTTP(w, req)
// 	suite.Equal(http.StatusBadRequest, w.Code)
// 	suite.Contains(w.Body.String(), "Invalid credential")
// }

// func (suite *ControllerTestSuite) TestUserRegisterHandler_UserUsernameInUse() {
// 	mockUser := domain.User{
// 		UserName: "user123",
// 		Password: "password",
// 		Role:     "USER",
// 	}
// 	suite.mockUserUseCase.On("Register", mockUser).Return(nil, errors.New("username already in use"))
// 	suite.router.POST("/register", suite.controller.UserRegisterHandler)
// 	w := httptest.NewRecorder()
// 	reqBody := `{"userName":"user123","password":"password","role":"USER"}`
// 	req, err := http.NewRequest("POST", "/register", bytes.NewReader([]byte(reqBody)))
// 	suite.Nil(err)
// 	req.Header.Set("Content-Type", "application/json")
// 	suite.router.ServeHTTP(w, req)
// 	suite.Equal(http.StatusBadRequest, w.Code)
// 	suite.Contains(w.Body.String(), "username already in use")
// }

// func (suite *ControllerTestSuite) TestUserLoginHandler_UserSuccess() {
// 	mockUser := domain.User{
// 		UserName: "user123",
// 		Password: "password",
// 	}
// 	suite.mockUserUseCase.On("Login", mockUser).Return("valid token", nil)
// 	suite.router.POST("/login", suite.controller.UserLoginHandler)
// 	w := httptest.NewRecorder()
// 	body, err := json.Marshal(mockUser)
// 	suite.Nil(err)

// 	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.Contains(w.Body.String(), "valid token")
// 	suite.Contains(w.Body.String(), "Logged successfully")
// 	suite.mockUserUseCase.AssertExpectations(suite.T())
// }

// func (suite *ControllerTestSuite) TestUserLoginHandler_UserInvalidCredentials() {
// 	suite.mockUserUseCase.On("Login", domain.User{}).Return("", errors.New("invalid credentials"))
// 	suite.router.POST("/login", suite.controller.UserLoginHandler)
// 	w := httptest.NewRecorder()
// 	req, err := http.NewRequest("POST", "/login", nil)
// 	suite.Nil(err)
// 	req.Header.Set("Content-Type", "application/json")
// 	suite.router.ServeHTTP(w, req)
// 	suite.Equal(http.StatusBadRequest, w.Code)
// 	suite.Contains(w.Body.String(), "Invalid credential")
// }

// func TestControllerTestSuite(t *testing.T) {
// 	suite.Run(t, new(ControllerTestSuite))
// }
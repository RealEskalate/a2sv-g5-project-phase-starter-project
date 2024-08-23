package controllers

// import (
// 	"bytes"
// 	"errors"
// 	"group3-blogApi/domain"
// 	"group3-blogApi/mocks"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )
// func setupRouterUser(userUsecase domain.UserUsecase) *gin.Engine {
// 	r := gin.Default()
// 	userController := NewUserController(userUsecase)

// 	// Middleware to simulate setting user_id and role in the context
// 	r.Use(func(c *gin.Context) {
// 		objID := primitive.NewObjectID()
// 		c.Set("user_id", objID.Hex())
// 		c.Set("role", "admin") // Simulate an admin role
// 		c.Next()
// 	})

// 	r.GET("/users/me", userController.GetMyProfile)
// 	r.GET("/users", userController.GetUsers)
// 	r.DELETE("/users/:id", userController.DeleteUser)
// 	r.GET("/users/:id", userController.GetUser)
// 	r.PUT("/users/:id/role", userController.UpdateUserRole)

// 	r.POST("/login", userController.Login)
// 	r.POST("/logoutAll", userController.LogoutAll)
// 	r.POST("/logoutDevice", userController.LogoutDevice)
// 	r.POST("/register", userController.Register)

// 	return r
// }

// func TestLogoutAll(t *testing.T) {
// 	mockUsecase := new(mocks.UserUsecase)
// 	r := setupRouterUser(mockUsecase)

// 	t.Run("Logout Error", func(t *testing.T) {
// 		// Simulate an error during logout from all devices
// 		mockUsecase.On("LogoutAllDevices", "12345").Return(errors.New("Logout failed")).Once()

// 		req, _ := http.NewRequest(http.MethodPost, "/logoutAll?userID=12345", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 		assert.Contains(t, w.Body.String(), "Logout failed")
// 	})

// 	t.Run("Successful Logout", func(t *testing.T) {
// 		// Simulate successful logout from all devices
// 		mockUsecase.On("LogoutAllDevices", "12345").Return(nil).Once()

// 		req, _ := http.NewRequest(http.MethodPost, "/logoutAll?userID=12345", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusOK, w.Code)
// 		assert.Contains(t, w.Body.String(), "Logged out from all devices")
// 	})
// }

// func TestLogoutDevice(t *testing.T) {
// 	mockUsecase := new(mocks.UserUsecase)
// 	r := setupRouterUser(mockUsecase)

// 	t.Run("Logout Error", func(t *testing.T) {
// 		// Simulate an error during logout from a specific device
// 		mockUsecase.On("LogoutDevice", "12345", "98765").Return(errors.New("Logout failed")).Once()

// 		req, _ := http.NewRequest(http.MethodPost, "/logoutDevice?userID=12345&deviceID=98765", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 		assert.Contains(t, w.Body.String(), "Logout failed")
// 	})

// 	t.Run("Successful Logout", func(t *testing.T) {
// 		// Simulate successful logout from a specific device
// 		mockUsecase.On("LogoutDevice", "12345", "98765").Return(nil).Once()

// 		req, _ := http.NewRequest(http.MethodPost, "/logoutDevice?userID=12345&deviceID=98765", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusOK, w.Code)
// 		assert.Contains(t, w.Body.String(), "Logged out successfully")
// 	})
// }
// func TestLogin(t *testing.T) {
// 	mockUsecase := new(mocks.UserUsecase)
// 	r := setupRouterUser(mockUsecase)

// 	t.Run("Invalid Request Body", func(t *testing.T) {
// 		// Simulate invalid JSON request body
// 		req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(`{invalid json}`))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 		assert.Contains(t, w.Body.String(), "invalid character")
// 	})

// 	t.Run("Login Error", func(t *testing.T) {
// 		// Simulate valid user login request with incorrect credentials
// 		user := domain.User{Email: "test@example.com", Password: "wrongpassword"}
// 		mockUsecase.On("Login", &user, mock.AnythingOfType("string")).Return(nil, errors.New("Invalid credentials")).Once()

// 		reqBody := `{"email": "test@example.com", "password": "wrongpassword"}`
// 		req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 500, w.Code)
// 		assert.Contains(t, w.Body.String(), "")
// 	})

// 	t.Run("Successful Login", func(t *testing.T) {
// 		// Simulate valid user login request with correct credentials
// 		user := domain.User{Email: "test@example.com", Password: "correctpassword"}
// 		loginResponse := map[string]string{
// 			"accessToken":  "someAccessToken",
// 			"refreshToken": "someRefreshToken",
// 		}
// 		mockUsecase.On("Login", &user, mock.AnythingOfType("string")).Return(loginResponse, nil).Once()

// 		reqBody := `{"email": "test@example.com", "password": "correctpassword"}`
// 		req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 500, w.Code)

// 	})
// }

// func TestRegister(t *testing.T) {
// 	mockUsecase := new(mocks.UserUsecase)
// 	r := setupRouterUser(mockUsecase)

// 	t.Run("Invalid Input", func(t *testing.T) {
// 		// Simulate invalid input
// 		reqBody := `{"email": ""}` // invalid email (empty)
// 		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte(reqBody)))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 500, w.Code)
// 		assert.Contains(t, w.Body.String(), "")
// 	})

// 	t.Run("Registration Error", func(t *testing.T) {
// 		// Simulate an error during registration
// 		user := domain.User{Email: "user@example.com", Password: "password"}
// 		mockUsecase.On("Register", user).Return(errors.New("Failed to register user")).Once()

// 		reqBody := `{"email": "user@example.com", "password": "password"}`
// 		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte(reqBody)))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 		assert.Contains(t, w.Body.String(), "Failed to register user")
// 	})

// 	t.Run("Successful Registration", func(t *testing.T) {
// 		// Simulate successful registration
// 		user := domain.User{Email: "user@example.com", Password: "password"}
// 		mockUsecase.On("Register", user).Return(nil).Once()

// 		reqBody := `{"email": "user@example.com", "password": "password"}`
// 		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte(reqBody)))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusOK, w.Code)
// 		assert.Contains(t, w.Body.String(), "Your Account created successfully please check your email")
// 	})
// }

// func TestGetMyProfile(t *testing.T) {
//     mockUserUsecase := new(mocks.UserUsecase)

//     // Use a consistent ObjectID
//     objID := "66c65e453009f61066250647" // Explicitly set this ID to ensure consistency
// 	num, _ := primitive.ObjectIDFromHex(objID)

//     // Mock user data
//     user := domain.User{
//         ID:     num , // Create ObjectID from the hex string
//         Username: "testuser",
//         Email:    "testuser@example.com",
//         Bio:      "This is a test user",
//         Role:     "user",
//     }

//     // Mocking the use case method to return the user
//     mockUserUsecase.On("GetMyProfile", objID).Return(user, nil)

//     // Setup the router
//     r := gin.Default()

//     // Middleware to simulate setting user_id in the context
//     r.Use(func(c *gin.Context) {
//         c.Set("user_id", objID) // Use the same consistent user ID
//         c.Next()
//     })

//     userController := NewUserController(mockUserUsecase)
//     r.GET("/users/me", userController.GetMyProfile)

//     // Create a new HTTP GET request
//     req, _ := http.NewRequest("GET", "/users/me", nil)
//     req.Header.Set("Content-Type", "application/json")

//     // Record the response
//     w := httptest.NewRecorder()
//     r.ServeHTTP(w, req) // Serve the request

//     // Check the response
//     assert.Equal(t, http.StatusOK, w.Code)
//     assert.Contains(t, w.Body.String(), "Welcome to your profile")
//     assert.Contains(t, w.Body.String(), "testuser")

//     // Ensure that all expectations were met
//     mockUserUsecase.AssertExpectations(t)
// }

// func TestGetUsers(t *testing.T) {
//     mockUserUsecase := new(mocks.UserUsecase)
//     r := setupRouterUser(mockUserUsecase)

//     // Prepare mock data
//     users := []domain.User{
//         {ID: primitive.NewObjectID(), Username: "user1", Email: "user1@example.com"},
//         {ID: primitive.NewObjectID(), Username: "user2", Email: "user2@example.com"},
//     }

//     // Mocking the use case method to return the list of users
//     mockUserUsecase.On("GetUsers").Return(users, nil)

//     req, _ := http.NewRequest("GET", "/users", nil)
//     req.Header.Set("Content-Type", "application/json")

//     w := httptest.NewRecorder()
//     c, _ := gin.CreateTestContext(w)
//     c.Request = req
//     c.Set("role", "admin") // Ensure the role is set to "admin"

//     r.ServeHTTP(w, req)

//     // Check response status and body
//     assert.Equal(t, 200, w.Code)
//     assert.Contains(t, w.Body.String(), "user1")
//     assert.Contains(t, w.Body.String(), "user2")

//     mockUserUsecase.AssertExpectations(t)
// }

// func TestGetUser(t *testing.T) {
// 	mockUsecase := new(mocks.UserUsecase)
// 	r := setupRouterUser(mockUsecase)

// 	t.Run("Unauthorized Access", func(t *testing.T) {
// 		// Set role to non-admin
// 		r.Use(func(c *gin.Context) {
// 			c.Set("role", "user")
// 			c.Next()
// 		})

// 		req, _ := http.NewRequest(http.MethodGet, "/users/12345", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 500, w.Code)
// 		assert.Contains(t, w.Body.String(), "")
// 	})

// 	t.Run("User Not Found", func(t *testing.T) {
// 		mockUsecase.On("GetMyProfile", "12345").Return(domain.User{}, errors.New("User not found")).Once()

// 		req, _ := http.NewRequest(http.MethodGet, "/users/12345", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 		assert.Contains(t, w.Body.String(), "User not found")
// 	})

// 	t.Run("User ID Mismatch", func(t *testing.T) {
// 		user := domain.User{ID: primitive.NewObjectID()}
// 		mockUsecase.On("GetMyProfile", "12345").Return(user, nil).Once()

// 		req, _ := http.NewRequest(http.MethodGet, "/users/12345", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusUnauthorized, w.Code)
// 		assert.Contains(t, w.Body.String(), "")
// 	})

// 	t.Run("Successful Retrieval", func(t *testing.T) {
// 		userID := primitive.NewObjectID().Hex()
// 		user := domain.User{ID: primitive.NewObjectID()}
// 		mockUsecase.On("GetMyProfile", userID).Return(user, nil).Once()

// 		req, _ := http.NewRequest(http.MethodGet, "/users/"+userID, nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 401, w.Code)
// 		// assert.Contains(t, w.Body.String(), userID)
// 	})
// }

// func TestDeleteUser(t *testing.T) {
// 	mockUsecase := new(mocks.UserUsecase)
// 	r := setupRouterUser(mockUsecase)

// 	t.Run("Unauthorized Access", func(t *testing.T) {
// 		// Set role to non-admin
// 		r.Use(func(c *gin.Context) {
// 			c.Set("role", "user")
// 			c.Next()
// 		})

// 		req, _ := http.NewRequest(http.MethodDelete, "/users/12345", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 500, w.Code)
// 		assert.Contains(t, w.Body.String(), "")
// 	})

// 	t.Run("User Not Found", func(t *testing.T) {
// 		mockUsecase.On("GetMyProfile", "12345").Return(domain.User{}, errors.New("User not found")).Once()

// 		req, _ := http.NewRequest(http.MethodDelete, "/users/12345", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 		assert.Contains(t, w.Body.String(), "User not found")
// 	})

// 	t.Run("User ID Mismatch", func(t *testing.T) {
// 		user := domain.User{ID: primitive.NewObjectID()}
// 		mockUsecase.On("GetMyProfile", "12345").Return(user, nil).Once()

// 		req, _ := http.NewRequest(http.MethodDelete, "/users/12345", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusUnauthorized, w.Code)
// 		assert.Contains(t, w.Body.String(), "Unauthorized")
// 	})

// 	t.Run("Successful Deletion", func(t *testing.T) {
// 		userID := primitive.NewObjectID().Hex()
// 		user := domain.User{ID: primitive.NewObjectID()}
// 		mockUsecase.On("GetMyProfile", userID).Return(user, nil).Once()
// 		mockUsecase.On("DeleteUser", userID).Return(user, nil).Once()

// 		req, _ := http.NewRequest(http.MethodDelete, "/users/"+userID, nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t,401 , w.Code)

// 	})
// }

// func TestUpdateUserRole(t *testing.T) {
// 	mockUsecase := new(mocks.UserUsecase)
// 	r := setupRouterUser(mockUsecase)

// 	t.Run("Unauthorized Access", func(t *testing.T) {
// 		// Set role to non-admin
// 		r.Use(func(c *gin.Context) {
// 			c.Set("role", "user")
// 			c.Next()
// 		})

// 		body := `{"role": "admin"}`
// 		req, _ := http.NewRequest(http.MethodPut, "/users/12345/role", strings.NewReader(body))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 500, w.Code)
// 		assert.Contains(t, w.Body.String(), "")
// 	})

// 	t.Run("Invalid Role Input", func(t *testing.T) {
// 		// Set role to admin
// 		r.Use(func(c *gin.Context) {
// 			c.Set("role", "admin")
// 			c.Next()
// 		})

// 		// Missing role field
// 		body := `{}`
// 		req, _ := http.NewRequest(http.MethodPut, "/users/12345/role", strings.NewReader(body))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 		assert.Contains(t, w.Body.String(), "Key: 'Role' Error:Field validation for 'Role' failed on the 'required' tag")
// 	})

// 	t.Run("Failed Role Update", func(t *testing.T) {
// 		// Set role to admin
// 		r.Use(func(c *gin.Context) {
// 			c.Set("role", "admin")
// 			c.Next()
// 		})

// 		mockUsecase.On("UpdateUserRole", "12345", "admin").Return(domain.User{}, errors.New("Update failed")).Once()

// 		body := `{"role": "admin"}`
// 		req, _ := http.NewRequest(http.MethodPut, "/users/12345/role", strings.NewReader(body))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 		assert.Contains(t, w.Body.String(), "Update failed")
// 	})

// 	t.Run("Successful Role Update", func(t *testing.T) {
// 		// Set role to admin
// 		r.Use(func(c *gin.Context) {
// 			c.Set("role", "admin")
// 			c.Next()
// 		})

// 		user := domain.User{ID: primitive.NewObjectID(), Role: "admin"}
// 		mockUsecase.On("UpdateUserRole", "12345", "admin").Return(user, nil).Once()

// 		body := `{"role": "admin"}`
// 		req, _ := http.NewRequest(http.MethodPut, "/users/12345/role", strings.NewReader(body))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusOK, w.Code)
// 		assert.Contains(t, w.Body.String(), "User role updated successfully")
// 		assert.Contains(t, w.Body.String(), "admin")
// 	})
// }

// func TestDeleteMyAccount(t *testing.T) {
// 	mockUsecase := new(mocks.UserUsecase)
// 	r := setupRouterUser(mockUsecase)

// 	t.Run("User Not Found", func(t *testing.T) {
// 		// Simulate the user_id context
// 		r.Use(func(c *gin.Context) {
// 			c.Set("user_id", "12345")
// 			c.Next()
// 		})

// 		// Simulate GetMyProfile returning an error
// 		mockUsecase.On("GetMyProfile", "12345").Return(domain.User{}, errors.New("User not found")).Once()

// 		req, _ := http.NewRequest(http.MethodDelete, "/users/12345", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 		assert.Contains(t, w.Body.String(), "User not found")
// 	})

// 	t.Run("Unauthorized Access", func(t *testing.T) {
// 		// Simulate the user_id context
// 		r.Use(func(c *gin.Context) {
// 			c.Set("user_id", "67890")
// 			c.Next()
// 		})

// 		user := domain.User{ID: primitive.NewObjectID()} // user.ID.Hex() will be different from "67890"
// 		mockUsecase.On("GetMyProfile", "67890").Return(user, nil).Once()

// 		req, _ := http.NewRequest(http.MethodDelete, "/users/67890", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 401, w.Code)
// 		assert.Contains(t, w.Body.String(), "")
// 	})

// 	t.Run("Failed Account Deletion", func(t *testing.T) {
// 		// Simulate the user_id context
// 		r.Use(func(c *gin.Context) {
// 			c.Set("user_id", "12345")
// 			c.Next()
// 		})

// 		user := domain.User{ID: primitive.NewObjectID()} // user.ID.Hex() will match "12345"
// 		mockUsecase.On("GetMyProfile", "12345").Return(user, nil).Once()
// 		mockUsecase.On("DeleteMyAccount", "12345").Return(errors.New("Delete failed")).Once()

// 		req, _ := http.NewRequest(http.MethodDelete, "/users/12345", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 401, w.Code)
// 	})

// 	t.Run("Successful Account Deletion", func(t *testing.T) {
// 		// Simulate the user_id context
// 		r.Use(func(c *gin.Context) {
// 			c.Set("user_id", "12345")
// 			c.Next()
// 		})

// 		user := domain.User{ID: primitive.NewObjectID()} // user.ID.Hex() will match "12345"
// 		mockUsecase.On("GetMyProfile", "12345").Return(user, nil).Once()
// 		mockUsecase.On("DeleteMyAccount", "12345").Return(nil).Once()

// 		req, _ := http.NewRequest(http.MethodDelete, "/users/12345", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 401, w.Code)
// 	})
// }

// func TestUpdateMyProfile(t *testing.T) {
// 	mockUsecase := new(mocks.UserUsecase)
// 	r := setupRouterUser(mockUsecase)

// 	t.Run("Invalid Request Body", func(t *testing.T) {
// 		// Simulate the user_id context
// 		r.Use(func(c *gin.Context) {
// 			c.Set("user_id", "12345")
// 			c.Next()
// 		})

// 		req, _ := http.NewRequest(http.MethodPut, "/users/me", strings.NewReader(`{invalid json}`))
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 404, w.Code)
// 		assert.Contains(t, w.Body.String(), "404 page not found")
// 	})

// 	t.Run("Update Profile Error", func(t *testing.T) {
// 		// Simulate the user_id context
// 		r.Use(func(c *gin.Context) {
// 			c.Set("user_id", "12345")
// 			c.Next()
// 		})

// 		user := domain.User{Username: "John Doe"}
// 		mockUsecase.On("UpdateMyProfile", user, "12345").Return(errors.New("Update failed")).Once()

// 		reqBody := `{"name": "John Doe"}`
// 		req, _ := http.NewRequest(http.MethodPut, "/users/me", strings.NewReader(reqBody))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 404, w.Code)
// 		assert.Contains(t, w.Body.String(), "404 page not found")
// 	})

// 	t.Run("Successful Profile Update", func(t *testing.T) {
// 		// Simulate the user_id context
// 		r.Use(func(c *gin.Context) {
// 			c.Set("user_id", "12345")
// 			c.Next()
// 		})

// 		user := domain.User{Username: "John Doe"}
// 		mockUsecase.On("UpdateMyProfile", user, "12345").Return(nil).Once()

// 		reqBody := `{"name": "John Doe"}`
// 		req, _ := http.NewRequest(http.MethodPut, "/users/me", strings.NewReader(reqBody))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 404, w.Code)
// 	})
// }



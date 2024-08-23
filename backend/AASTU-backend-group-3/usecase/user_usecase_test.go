package usecase

// import (
// 	"errors"
// 	"group3-blogApi/domain"
// 	"group3-blogApi/mocks"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func TestGetMyProfile(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)
// 	userID := primitive.NewObjectID()

// 	mockUser := domain.User{
// 		ID:       userID,
// 		Username: "testuser",
// 		Email:    "test@example.com",
// 		Bio:      "Test Bio",
// 	}

// 	uc := UserUsecase{
// 		UserRepo: mockUserRepo,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("GetMyProfile", userID.Hex()).Return(mockUser, nil).Once()

// 		result, err := uc.GetMyProfile(mockUser.ID.Hex())

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockUser, result)
// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("user not found", func(t *testing.T) {
// 		mockUserRepo.On("GetMyProfile", userID.Hex()).Return(domain.User{}, errors.New("user not found")).Once()

// 		result, err := uc.GetMyProfile(mockUser.ID.Hex())

// 		assert.Error(t, err)
// 		assert.Equal(t, domain.User{}, result)
// 		mockUserRepo.AssertExpectations(t)
// 	})
// }

// func TestDeleteUser(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)
// 	userID := primitive.NewObjectID()
// 	mockUser := domain.User{
// 		ID:       userID,
// 		Username: "testuser",
// 		Email:    "test@example.com",
// 		Bio:      "Test Bio",
// 	}

// 	uc := UserUsecase{
// 		UserRepo: mockUserRepo,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("DeleteUser", userID.Hex()).Return(mockUser, nil).Once()

// 		result, err := uc.DeleteUser(mockUser.ID.Hex())

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockUser, result)
// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("user not found", func(t *testing.T) {
// 		mockUserRepo.On("DeleteUser",userID.Hex()).Return(domain.User{}, errors.New("user not found")).Once()

// 		result, err := uc.DeleteUser(mockUser.ID.Hex())

// 		assert.Error(t, err)
// 		assert.Equal(t, domain.User{}, result)
// 		mockUserRepo.AssertExpectations(t)
// 	})
// }

// func TestUpdateUserRole(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)
// 	userID := primitive.NewObjectID()
// 	mockUser := domain.User{
// 		ID:    userID,
// 		Role:  "user",
// 		Email: "test@example.com",
// 	}

// 	uc := UserUsecase{
// 		UserRepo: mockUserRepo,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("UpdateUserRole", userID.Hex(), "admin").Return(mockUser, nil).Once()

// 		result, err := uc.UpdateUserRole(mockUser.ID.Hex(), "admin")

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockUser, result)
// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("user not found", func(t *testing.T) {
// 		mockUserRepo.On("UpdateUserRole", userID.Hex(), "admin").Return(domain.User{}, errors.New("user not found")).Once()

// 		result, err := uc.UpdateUserRole(mockUser.ID.Hex(), "admin")

// 		assert.Error(t, err)
// 		assert.Equal(t, domain.User{}, result)
// 		mockUserRepo.AssertExpectations(t)
// 	})
// }

// func TestDeleteMyAccount(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)
// 	userID := primitive.NewObjectID()
// 	mockUser := domain.User{
// 		ID:    userID,
// 		Role:  "user",
// 		Email: "teklu@gmail.com",
// 	}

// 	uc := UserUsecase{
// 		UserRepo: mockUserRepo,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("DeleteMyAccount", userID.Hex()).Return(nil).Once()

// 		err := uc.DeleteMyAccount(mockUser.ID.Hex())

// 		assert.NoError(t, err)
// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("user not found", func(t *testing.T) {
// 		mockUserRepo.On("DeleteMyAccount", userID.Hex()).Return(errors.New("user not found")).Once()

// 		err := uc.DeleteMyAccount(mockUser.ID.Hex())

// 		assert.Error(t, err)
// 		mockUserRepo.AssertExpectations(t)
// 	})

// }

// func TestUploadImage(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)
// 	userID := primitive.NewObjectID()
// 	mockUser := domain.User{
// 		ID:    userID,
// 		Role:  "user",
// 		Email: "medina@gmail.com",
// 	}

// 	uc := UserUsecase{
// 		UserRepo: mockUserRepo,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("UploadImage", userID.Hex(), "imagePath").Return(nil).Once()

// 		err := uc.UploadImage(mockUser.ID.Hex(), "imagePath")

// 		assert.NoError(t, err)
// 		mockUserRepo.AssertExpectations(t)
// 	} )

// 	t.Run("user not found", func(t *testing.T) {
// 		mockUserRepo.On("UploadImage", userID.Hex(), "imagePath").Return(errors.New("user not found")).Once()

// 		err := uc.UploadImage(mockUser.ID.Hex(), "imagePath")

// 		assert.Error(t, err)
// 		mockUserRepo.AssertExpectations(t)
// 	} )

// }

// func TestUpdateMyProfile(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)
// 	userID := primitive.NewObjectID()
// 	mockUser := domain.User{
// 		ID:    userID,
// 		Role:  "user",
// 		Email: "beki@gmai;.com",
// 		Username: "beki",
// 		Bio: "bio",
// 	}

// 	uc := UserUsecase{
// 		UserRepo: mockUserRepo,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("UpdateMyProfile", mockUser, userID.Hex()).Return(nil).Once()

// 		err := uc.UpdateMyProfile(mockUser, mockUser.ID.Hex())

// 		assert.NoError(t, err)
// 		mockUserRepo.AssertExpectations(t)
// 	} )

// 	t.Run("user not found", func(t *testing.T) {
// 		mockUserRepo.On("UpdateMyProfile", mockUser, userID.Hex()).Return(errors.New("user not found")).Once()

// 		err := uc.UpdateMyProfile(mockUser, mockUser.ID.Hex())

// 		assert.Error(t, err)
// 		mockUserRepo.AssertExpectations(t)
// 	} )

// 	t.Run("bio and username are required", func(t *testing.T) {
// 		mockUser.Bio = ""
// 		mockUser.Username = ""
// 		err := uc.UpdateMyProfile(mockUser, mockUser.ID.Hex())
// 		assert.Error(t, err)
// 		mockUserRepo.AssertExpectations(t)
// 	} )

// }

// func TestGetUsers(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)
// 	mockUsers := []domain.User{
// 		{
// 			ID:       primitive.NewObjectID(),
// 			Username: "testuser",
// 			Email:    "kenean@gmail.com",
// 			Bio:      "Test Bio",
// 		},
// 		{
// 			ID:       primitive.NewObjectID(),
// 			Username: "testuser2",
// 			Email:    "j@gmail.com",
// 			Bio:      "Test Bio",
// 		},
// 	}

// 	uc := UserUsecase{
// 		UserRepo: mockUserRepo,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("GetUsers").Return(mockUsers, nil).Once()

// 		result, err := uc.GetUsers()

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockUsers, result)
// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("not User found", func(t *testing.T) {
// 		mockUserRepo.On("GetUsers").Return([]domain.User{}, errors.New("user not found")).Once()

// 		result, err := uc.GetUsers()

// 		assert.Error(t, err)
// 		assert.Equal(t, []domain.User{}, result)
// 		mockUserRepo.AssertExpectations(t)
// 	})
// }

// func TestLogin(t *testing.T) {
//     // Create mocks
//     mockUserRepo := new(mocks.UserRepository)
//     mockPasswordService := new(mocks.PasswordService)
//     mockTokenGenerator := new(mocks.TokenGenerator)

//     // Mock user data
//     userID := primitive.NewObjectID()
//     mockUser := &domain.User{
//         ID:       userID,
//         Email:    "teklu@gmain.com",
//         Password: "hashedPassword",
//         Role:     "user",
//         RefreshTokens: []domain.RefreshToken{
//             {Token: "mockedRefreshToken", DeviceID: "deviceID", CreatedAt: time.Now()},
//             {Token: "mockedRefreshToken", DeviceID: "deviceID", CreatedAt: time.Now()},
//         },
//     }

//     // Fake mock expectations
//     mockUserRepo.On("Login", mock.Anything).Return(mockUser, nil).Once()
//     mockPasswordService.On("CheckPasswordHash", mock.Anything, mock.Anything).Return(true).Once()
//     mockTokenGenerator.On("GenerateRefreshToken", mock.Anything).Return("mockedRefreshToken", nil).Once()
//     mockTokenGenerator.On("GenerateToken", mock.Anything).Return("mockedAccessToken", nil).Once()
//     mockUserRepo.On("UpdateUser", mock.Anything).Return(nil).Once()

//     // Create UserUsecase with mocks
//     userUsecase := &UserUsecase{
//         UserRepo:    mockUserRepo,
//         PasswordSvc: mockPasswordService,
//         TokenGen:    mockTokenGenerator,
//     }

//     // Call the Login method
//     response, err := userUsecase.Login(mockUser, "deviceID")

//     // Assertions
//     assert.NoError(t, err)
//     assert.Equal(t, "mockedAccessToken", response.AccessToken)
//     assert.Equal(t, "mockedRefreshToken", response.RefreshToken)

//     // Assert expectations
//     mockUserRepo.AssertExpectations(t)
//     mockPasswordService.AssertExpectations(t)
//     mockTokenGenerator.AssertExpectations(t)
// }

// func TestLogout(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)

// 	// Sample user data
// 	userID := primitive.NewObjectID()
// 	mockUser := domain.User{
// 		ID: userID,
// 		RefreshTokens: []domain.RefreshToken{
// 			{
// 				Token:     "validToken",
// 				DeviceID:  "validDeviceID",
// 				CreatedAt: time.Now(),
// 			},
// 			{
// 				Token:     "anotherToken",
// 				DeviceID:  "anotherDeviceID",
// 				CreatedAt: time.Now(),
// 			},
// 		},
// 	}

// 	uc := UserUsecase{
// 		UserRepo: mockUserRepo,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		// Set up expectations
// 		mockUserRepo.On("GetUserByID", userID.Hex()).Return(mockUser, nil).Once()
// 		mockUserRepo.On("UpdateUser", &domain.User{
// 			ID: userID,
// 			RefreshTokens: []domain.RefreshToken{
// 				{
// 					Token:     "anotherToken",
// 					DeviceID:  "anotherDeviceID",
// 					CreatedAt: mockUser.RefreshTokens[1].CreatedAt,
// 				},
// 			},
// 		}).Return(nil).Once()

// 		// Call the Logout method
// 		err := uc.Logout(userID.Hex(), "validDeviceID", "validToken")

// 		// Assertions
// 		assert.NoError(t, err)
// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("user not found", func(t *testing.T) {
// 		// Set up expectations
// 		mockUserRepo.On("GetUserByID", userID.Hex()).Return(domain.User{}, errors.New("user not found")).Once()

// 		// Call the Logout method
// 		err := uc.Logout(userID.Hex(), "validDeviceID", "validToken")

// 		// Assertions
// 		assert.Error(t, err)
// 		assert.Equal(t, "user not found", err.Error())
// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("invalid token", func(t *testing.T) {
// 		// Set up expectations
// 		mockUserRepo.On("GetUserByID", userID.Hex()).Return(mockUser, nil).Once()

// 		// Call the Logout method
// 		err := uc.Logout(userID.Hex(), "validDeviceID", "invalidToken")

// 		// Assertions
// 		assert.Error(t, err)
// 		assert.Equal(t, "invalid token", err.Error())
// 		mockUserRepo.AssertExpectations(t)
// 	})

	
// }


// func TestLogoutAllDevices(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)

// 	// Sample user data
// 	userID := primitive.NewObjectID()
// 	mockUser := domain.User{
// 		ID: userID,
// 		RefreshTokens: []domain.RefreshToken{
// 			{
// 				Token:     "validToken1",
// 				DeviceID:  "validDeviceID1",
// 				CreatedAt: time.Now(),
// 			},
// 			{
// 				Token:     "validToken2",
// 				DeviceID:  "validDeviceID2",
// 				CreatedAt: time.Now(),
// 			},
// 		},
// 	}

// 	uc := UserUsecase{
// 		UserRepo: mockUserRepo,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		// Set up expectations
// 		mockUserRepo.On("GetUserByID", userID.Hex()).Return(mockUser, nil).Once()
// 		mockUserRepo.On("UpdateUser", &domain.User{
// 			ID:            userID,
// 			RefreshTokens: []domain.RefreshToken{},
// 		}).Return(nil).Once()

// 		// Call the LogoutAllDevices method
// 		err := uc.LogoutAllDevices(userID.Hex())

// 		// Assertions
// 		assert.NoError(t, err)
// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("user not found", func(t *testing.T) {
// 		// Set up expectations
// 		mockUserRepo.On("GetUserByID", userID.Hex()).Return(domain.User{}, errors.New("user not found")).Once()

// 		// Call the LogoutAllDevices method
// 		err := uc.LogoutAllDevices(userID.Hex())

// 		// Assertions
// 		assert.Error(t, err)
// 		assert.Equal(t, "user not found", err.Error())
// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("update user error", func(t *testing.T) {
// 		// Set up expectations
// 		mockUserRepo.On("GetUserByID", userID.Hex()).Return(mockUser, nil).Once()
// 		mockUserRepo.On("UpdateUser", mock.Anything).Return(errors.New("update failed")).Once()

// 		// Call the LogoutAllDevices method
// 		err := uc.LogoutAllDevices(userID.Hex())

// 		// Assertions
// 		assert.Error(t, err)
// 		assert.Equal(t, "update failed", err.Error())
// 		mockUserRepo.AssertExpectations(t)
// 	})
// }



// func TestLogoutDevice(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)

// 	// Sample user data
// 	userID := primitive.NewObjectID()
// 	deviceID := "validDeviceID"
// 	mockUser := domain.User{
// 		ID: userID,
// 		RefreshTokens: []domain.RefreshToken{
// 			{
// 				Token:     "validToken1",
// 				DeviceID:  deviceID,
// 				CreatedAt: time.Now(),
// 			},
// 			{
// 				Token:     "validToken2",
// 				DeviceID:  "anotherDeviceID",
// 				CreatedAt: time.Now(),
// 			},
// 		},
// 	}

// 	uc := UserUsecase{
// 		UserRepo: mockUserRepo,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		// Set up expectations
// 		mockUserRepo.On("GetUserByID", userID.Hex()).Return(mockUser, nil).Once()
// 		mockUserRepo.On("UpdateUser", &domain.User{
// 			ID: userID,
// 			RefreshTokens: []domain.RefreshToken{
// 				{
// 					Token:     "validToken2",
// 					DeviceID:  "anotherDeviceID",
// 					CreatedAt: time.Now(),
// 				},
// 			},
// 		}).Return(nil).Once()

// 		// Call the LogoutDevice method
// 		err := uc.LogoutDevice(userID.Hex(), deviceID)

// 		// Assertions
// 		assert.NoError(t, err)
// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("user not found", func(t *testing.T) {
// 		// Set up expectations
// 		mockUserRepo.On("GetUserByID", userID.Hex()).Return(domain.User{}, errors.New("user not found")).Once()

// 		// Call the LogoutDevice method
// 		err := uc.LogoutDevice(userID.Hex(), deviceID)

// 		// Assertions
// 		assert.Error(t, err)
// 		assert.Equal(t, "user not found", err.Error())
// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("device not found", func(t *testing.T) {
// 		// Set up expectations
// 		mockUserRepo.On("GetUserByID", userID.Hex()).Return(mockUser, nil).Once()

// 		// Call the LogoutDevice method with a non-existent device ID
// 		err := uc.LogoutDevice(userID.Hex(), "nonExistentDeviceID")

// 		// Assertions
// 		assert.Error(t, err)
// 		assert.Equal(t, "device not found", err.Error())
// 		mockUserRepo.AssertExpectations(t)
// 	})

	
// }



// func TestRegister(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)
// 	mockPasswordSvc := new(mocks.PasswordService)
// 	mockTokenGenerator := new(mocks.TokenGenerator)

// 	userUsecase := &UserUsecase{
// 		UserRepo:       mockUserRepo,
// 		PasswordSvc:    mockPasswordSvc,
// 		TokenGen: mockTokenGenerator,
// 	}

// 	t.Run("Missing required fields", func(t *testing.T) {
// 		err := userUsecase.Register(domain.User{})
// 		assert.EqualError(t, err, "all fields are required")
// 	})

// 	t.Run("Invalid email format", func(t *testing.T) {
// 		user := domain.User{
// 			Username: "validUsername",
// 			Email:    "invalidEmail",
// 			Password: "ValidPassword1!",
// 		}
// 		err := userUsecase.Register(user)
// 		assert.EqualError(t, err, "invalid email format")
// 	})

// 	t.Run("Invalid password format", func(t *testing.T) {
// 		user := domain.User{
// 			Username: "validUsername",
// 			Email:    "valid@email.com",
// 			Password: "weakpassword",
// 		}
// 		err := userUsecase.Register(user)
// 		assert.EqualError(t, err, "password must contain at least one uppercase letter, one lowercase letter, one digit, one special character and minimum length of 8 characters")
// 	})

	
	



	
// }




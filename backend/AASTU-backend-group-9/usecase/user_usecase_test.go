package usecase

import (
    "blog/domain"
    "blog/domain/mocks"
    "context"
    "errors"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
	 "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser(t *testing.T) {
    mockUserRepo := new(mocks.UserRepository)
    mockUser := &domain.CreateUser{
        Email:    "test@example.com",
		Username: "testuser",
        Password: "password123",
    }
    mockClaims := &domain.JwtCustomClaims{
        Role: "admin",
    }

    t.Run("success", func(t *testing.T) {
        mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

        u := NewUserUsecase(mockUserRepo, time.Second*2)

        err := u.CreateUser(context.TODO(), mockUser, mockClaims)

        assert.NoError(t, err)
        mockUserRepo.AssertExpectations(t)
    })

    t.Run("error invalid email", func(t *testing.T) {
        invalidUser := &domain.CreateUser{
            Email:    "invalid-email",
            Username: "testuser",
            Password: "password123",
        }

        u := NewUserUsecase(mockUserRepo, time.Second*2)

        err := u.CreateUser(context.TODO(), invalidUser, mockClaims)

        assert.Error(t, err)
        assert.Equal(t, "invalid email", err.Error())
    })

    t.Run("error invalid password", func(t *testing.T) {
        invalidUser := &domain.CreateUser{
            Email:    "test@example.com",
            Username: "testuser",
            Password: "short",
        }

        u := NewUserUsecase(mockUserRepo, time.Second*2)

        err := u.CreateUser(context.TODO(), invalidUser, mockClaims)

        assert.Error(t, err)
        assert.Equal(t, "password must be at least 8 characters long", err.Error())
    })

    t.Run("error cannot manipulate user", func(t *testing.T) {
        mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(errors.New("cannot manipulate user")).Once()

        u := NewUserUsecase(mockUserRepo, time.Second*2)

        err := u.CreateUser(context.TODO(), mockUser, mockClaims)

        assert.Error(t, err)
        assert.Equal(t, "cannot manipulate user", err.Error())
        mockUserRepo.AssertExpectations(t)
    })
}

func TestGetUserByEmail(t *testing.T) {
    mockUserRepo := new(mocks.UserRepository)
    mockUser := &domain.User{
        ID:       primitive.NewObjectID(),
        Email:    "test@example.com",
        Username: "testuser",
        Password: "password123",
    }

    t.Run("success", func(t *testing.T) {
        mockUserRepo.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(mockUser, nil).Once()

        uc := NewUserUsecase(mockUserRepo, time.Second*2)
        user, err := uc.GetUserByEmail(context.TODO(), "test@example.com")

        assert.NoError(t, err)
        assert.Equal(t, mockUser, user)
        mockUserRepo.AssertExpectations(t)
    })

    t.Run("user not found", func(t *testing.T) {
        mockUserRepo.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("user not found")).Once()

        uc := NewUserUsecase(mockUserRepo, time.Second*2)
        user, err := uc.GetUserByEmail(context.TODO(), "nonexistent@example.com")

        assert.Error(t, err)
        assert.Nil(t, user)
        assert.Equal(t, "user not found", err.Error())
        mockUserRepo.AssertExpectations(t)
    })
}

func TestGetUserByUsername(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
    mockUser := &domain.User{
        ID:       primitive.NewObjectID(),
        Email:    "test@example.com",
        Username: "testuser",
        Password: "password123",
    }

    t.Run("success", func(t *testing.T) {
        mockUserRepo.On("GetUserByUsername", mock.Anything, mock.AnythingOfType("string")).Return(mockUser, nil).Once()

        uc := NewUserUsecase(mockUserRepo, time.Second*2)
        user, err := uc.GetUserByUsername(context.TODO(), "testuser")

        assert.NoError(t, err)
        assert.Equal(t, mockUser, user)
        mockUserRepo.AssertExpectations(t)
    })
	t.Run("user not found", func(t *testing.T) {
		mockUserRepo.On("GetUserByUsername", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("user not found")).Once()
		
        uc := NewUserUsecase(mockUserRepo, time.Second*2)
		user, err := uc.GetUserByUsername(context.TODO(), "nonexistentuser")
		
        assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "user not found", err.Error())
		mockUserRepo.AssertExpectations(t)
	})
}
func TestGetUserByID(t *testing.T) {
    mockUserRepo := new(mocks.UserRepository)
    mockUserID := primitive.NewObjectID()
    mockUser := domain.User{
        ID:       mockUserID,
        Email:    "test@example.com",
        Username: "testuser",
    }

    t.Run("success", func(t *testing.T) {
        mockUserRepo.On("GetUserByID", mock.Anything, mockUserID).Return(&mockUser, nil).Once()

        uc := NewUserUsecase(mockUserRepo, time.Second*2)
        user, err := uc.GetUserByID(context.TODO(), mockUserID)

        assert.NoError(t, err)
        assert.Equal(t, &mockUser, user)
        mockUserRepo.AssertExpectations(t)
    })

    t.Run("user not found", func(t *testing.T) {
        mockUserRepo.On("GetUserByID", mock.Anything, mockUserID).Return(nil, errors.New("user not found")).Once()

        uc := NewUserUsecase(mockUserRepo, time.Second*2)
        user, err := uc.GetUserByID(context.TODO(), mockUserID)

        assert.Error(t, err)
        assert.Nil(t, user)
        assert.Equal(t, "user not found", err.Error())
        mockUserRepo.AssertExpectations(t)
    })
}
func TestGetAllUsers(t *testing.T) {
    mockUserRepo := new(mocks.UserRepository)
    mockUsers := []*domain.User{
        {
            ID:       primitive.NewObjectID(),
            Email:    "test1@example.com",
            Username: "testuser1",
        },
        {
            ID:       primitive.NewObjectID(),
            Email:    "test2@example.com",
            Username: "testuser2",
        },
    }

    t.Run("success", func(t *testing.T) {
        mockUserRepo.On("GetAllUsers", mock.Anything).Return(mockUsers, nil).Once()

        u := NewUserUsecase(mockUserRepo, time.Second*2)
        users, err := u.GetAllUsers(context.TODO())

        assert.NoError(t, err)
        assert.Equal(t, mockUsers, users)
        mockUserRepo.AssertExpectations(t)
    })
	t.Run("no users found", func(t *testing.T) {
        mockUserRepo.On("GetAllUsers", mock.Anything).Return(nil, errors.New("no users found")).Once()

        u := NewUserUsecase(mockUserRepo, time.Second*2)
        users, err := u.GetAllUsers(context.TODO())

        assert.Error(t, err)
        assert.Nil(t, users)
        assert.Equal(t, "no users found", err.Error())
        mockUserRepo.AssertExpectations(t)
    })
}

func TestDeleteUser(t *testing.T) {
    mockUserRepo := new(mocks.UserRepository)
    mockUserID := primitive.NewObjectID()
    mockClaims := &domain.JwtCustomClaims{
        Role: "admin",
    }

    t.Run("success", func(t *testing.T) {
        mockUserRepo.On("DeleteUser", mock.Anything, mockUserID).Return(nil).Once()

        u := NewUserUsecase(mockUserRepo, time.Second*2)

        err := u.DeleteUser(context.TODO(), mockUserID, mockClaims)

        assert.NoError(t, err)
        mockUserRepo.AssertExpectations(t)
    })
	t.Run("error cannot manipulate user", func(t *testing.T) {
        mockUserRepo.On("DeleteUser", mock.Anything, mockUserID).Return(errors.New("cannot manipulate user")).Once()

        u := NewUserUsecase(mockUserRepo, time.Second*2)

        err := u.DeleteUser(context.TODO(), mockUserID, mockClaims)

        assert.Error(t, err)
        assert.Equal(t, "cannot manipulate user", err.Error())
        mockUserRepo.AssertExpectations(t)
    })
}
func TestPromoteUser(t *testing.T) {
    mockUserRepo := new(mocks.UserRepository)
    mockUserID := primitive.NewObjectID()
    mockUser := &domain.User{
        ID:    mockUserID,
        Role:  "user",
        Email: "test@example.com",
    }
    mockClaims := &domain.JwtCustomClaims{
        Role: "admin",
    }

    t.Run("success", func(t *testing.T) {
        mockUserRepo.On("GetUserByID", mock.Anything, mockUserID).Return(mockUser, nil).Once()
        mockUserRepo.On("PromoteUser", mock.Anything, mockUserID).Return(nil).Once()

        u := NewUserUsecase(mockUserRepo, time.Second*2)
        err := u.PromoteUser(context.TODO(), mockUserID, mockClaims)

        assert.NoError(t, err)
        mockUserRepo.AssertExpectations(t)
    })
	t.Run("error cannot promote root user", func(t *testing.T) {
        rootUser := &domain.User{
            ID:   mockUserID,
            Role: "root",
        }
        mockUserRepo.On("GetUserByID", mock.Anything, mockUserID).Return(rootUser, nil).Once()

        u := NewUserUsecase(mockUserRepo, time.Second*2)
        err := u.PromoteUser(context.TODO(), mockUserID, mockClaims)

        assert.Error(t, err)
        assert.Equal(t, "cannot promote root user", err.Error())
        mockUserRepo.AssertExpectations(t)
    })
	t.Run("error user is already an admin", func(t *testing.T) {
        adminUser := &domain.User{
            ID:   mockUserID,
            Role: "admin",
        }
        mockUserRepo.On("GetUserByID", mock.Anything, mockUserID).Return(adminUser, nil).Once()

        u := NewUserUsecase(mockUserRepo, time.Second*2)
        err := u.PromoteUser(context.TODO(), mockUserID, mockClaims)

        assert.Error(t, err)
        assert.Equal(t, "user is already an admin", err.Error())
        mockUserRepo.AssertExpectations(t)
    })
	t.Run("error insufficient permissions", func(t *testing.T) {
        mockUserRepo.On("GetUserByID", mock.Anything, mockUserID).Return(mockUser, nil).Once()

        nonAdminClaims := &domain.JwtCustomClaims{
            Role: "user",
        }
        u := NewUserUsecase(mockUserRepo, time.Second*2)
        err := u.PromoteUser(context.TODO(), mockUserID, nonAdminClaims)

        assert.Error(t, err)
        assert.Equal(t, "a user must be an admin or root to promote another user", err.Error())
        mockUserRepo.AssertExpectations(t)
    })
}
func TestDemoteUser(t *testing.T) {
    mockUserRepo := new(mocks.UserRepository)
    mockUserID := primitive.NewObjectID()
    mockUser := &domain.User{
        ID:    mockUserID,
        Email: "test@example.com",
        Role:  "admin",
    }

    t.Run("success", func(t *testing.T) {
        mockUserRepo.On("GetUserByID", mock.Anything, mockUserID).Return(mockUser, nil).Once()
        mockUserRepo.On("UpdateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

        adminClaims := &domain.JwtCustomClaims{
            Role: "admin",
        }
        u := NewUserUsecase(mockUserRepo, time.Second*2)
        err := u.DemoteUser(context.TODO(), mockUserID, adminClaims)

        assert.NoError(t, err)
        mockUserRepo.AssertExpectations(t)
    })
}
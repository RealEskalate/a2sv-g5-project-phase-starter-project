package usecase

import (
    "Blog_Starter/domain"
    "Blog_Starter/domain/mocks"
    "context"
	"time"
	"testing"
	"github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "golang.org/x/crypto/bcrypt"
)

type UserUsecaseTestSuite struct {
    suite.Suite
    userRepo   *mocks.UserRepository
    userUsecase domain.UserUsecase
}

func (suite *UserUsecaseTestSuite) SetupTest() {
    suite.userRepo = new(mocks.UserRepository)
    suite.userUsecase = NewUserUsecase(suite.userRepo, time.Second*2)
}

func (suite *UserUsecaseTestSuite) TestDeleteUser_Success() {
    userID := "some-user-id"
    password := "password"
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    user := &domain.User{
        UserID:   primitive.NewObjectID(),
        Password: string(hashedPassword),
    }

    suite.userRepo.On("GetUserByID", mock.Anything, userID).Return(user, nil)
    suite.userRepo.On("DeleteUser", mock.Anything, userID).Return(nil)

    err := suite.userUsecase.DeleteUser(context.Background(), userID, password)
    suite.NoError(err)
    suite.userRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestDeleteUser_WrongPassword() {
    userID := "some-user-id"
    password := "password"
    wrongPassword := "wrong-password"
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    user := &domain.User{
        UserID:   primitive.NewObjectID(),
        Password: string(hashedPassword),
    }

    suite.userRepo.On("GetUserByID", mock.Anything, userID).Return(user, nil)

    err := suite.userUsecase.DeleteUser(context.Background(), userID, wrongPassword)
    suite.EqualError(err, "password incorrect")
    suite.userRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestGetAllUser_Success() {
    users := []*domain.User{
        {
            UserID: primitive.NewObjectID(),
            Username: "user1",
            Email: "user1@example.com",
            Name: "User One",
        },
        {
            UserID: primitive.NewObjectID(),
            Username: "user2",
            Email: "user2@example.com",
            Name: "User Two",
        },
    }

    suite.userRepo.On("GetAllUser", mock.Anything).Return(users, nil)

    userResponses, err := suite.userUsecase.GetAllUser(context.Background())
    suite.NoError(err)
    suite.Len(userResponses, 2)
    suite.userRepo.AssertExpectations(suite.T())
}

func TestUserUsecaseTestSuite(t *testing.T) {
    suite.Run(t, new(UserUsecaseTestSuite))
}
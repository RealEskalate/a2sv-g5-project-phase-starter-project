package usecase_test

import (
	"blogApp/internal/domain"
	"blogApp/mocks/repository"
	"blogApp/internal/usecase/user"
	"context"
	"time"
	"errors"
	"testing"
	"blogApp/pkg/hash"
	//"blogApp/pkg/jwt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	userUsecase *user.UserUsecase
	mockRepo    *mocks.UserRepository
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.UserRepository)
	suite.userUsecase = user.NewUserUsecase(suite.mockRepo)
}

func (suite *UserUsecaseTestSuite) TestRegisterUserSuccess() {
	newUser := &domain.User{
		Email:    "test@example.com",
		Password: "password123",
		UserName: "testuser",
		Profile: domain.UserProfile{
			ProfileUrl: "http://example.com/profile.jpg",
			FirstName:  "Test",
			LastName:   "User",
			Gender:     "Non-binary",
			Bio:        "Just a test user",
			Profession: "Software Engineer",
		},
	}

	suite.mockRepo.On("FindUserByEmail", context.Background(), newUser.Email).Return(nil, nil)
	suite.mockRepo.On("IsEmptyCollection", context.Background()).Return(false, nil)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	suite.mockRepo.On("CreateUser", context.Background(), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(1).(*domain.User)
		user.Password = string(hashedPassword) // Check if password is hashed
	})

	createdUser, err := suite.userUsecase.RegisterUser(newUser)

	suite.NoError(err)
	suite.Equal(newUser.Email, createdUser.Email)
	suite.Equal(newUser.UserName, createdUser.UserName)
	suite.Equal(newUser.Profile, createdUser.Profile)
	suite.NotEmpty(createdUser.Password) // Ensure the password was hashed
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRegisterUserAlreadyExists() {
	existingUser := &domain.User{
		Email: "test@example.com",
	}

	suite.mockRepo.On("FindUserByEmail", context.Background(), existingUser.Email).Return(existingUser, nil)

	createdUser, err := suite.userUsecase.RegisterUser(existingUser)

	suite.Error(err)
	suite.Nil(createdUser)
	suite.EqualError(err, "user already exists")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRegisterUserEmptyCollection() {
	newUser := &domain.User{
		Email:    "new@example.com",
		Password: "password123",
		UserName: "newuser",
		Profile: domain.UserProfile{
			ProfileUrl: "http://example.com/newprofile.jpg",
			FirstName:  "New",
			LastName:   "User",
			Gender:     "Female",
			Bio:        "New test user",
			Profession: "Product Manager",
		},
	}

	suite.mockRepo.On("FindUserByEmail", context.Background(), newUser.Email).Return(nil, nil)
	suite.mockRepo.On("IsEmptyCollection", context.Background()).Return(true, nil)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	suite.mockRepo.On("CreateUser", context.Background(), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(1).(*domain.User)
		user.Password = string(hashedPassword) // Check if password is hashed
	})

	createdUser, err := suite.userUsecase.RegisterUser(newUser)

	suite.NoError(err)
	suite.Equal("owner", createdUser.Role) // Check if role is set to "owner"
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestRegisterUserErrorIsEmptyCollection() {
	newUser := &domain.User{
		Email:    "new@example.com",
		Password: "password123",
	}

	suite.mockRepo.On("FindUserByEmail", context.Background(), newUser.Email).Return(nil, nil)
	suite.mockRepo.On("IsEmptyCollection", context.Background()).Return(false, errors.New("database error"))

	createdUser, err := suite.userUsecase.RegisterUser(newUser)

	suite.Error(err)
	suite.Nil(createdUser)
	suite.EqualError(err, "database error")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLoginSuccess() {
    email := "test@example.com"
    password := "password123"
    hashedPassword, err := hash.HashPassword(password)
    suite.NoError(err, "Hashing password should not return an error")

    // Setup the mock to return a user with the hashed password
    suite.mockRepo.On("FindUserByEmail", context.TODO(), email).Return(&domain.User{
        ID:       primitive.NewObjectID(),
        UserName: "testuser",
        Email:    email,
        Password: hashedPassword, // Ensure this matches the hashed password
        Profile:  domain.UserProfile{},
        Role:     "user",
        Created:  primitive.NewDateTimeFromTime(time.Now()),
        Updated:  primitive.NewDateTimeFromTime(time.Now()),
        Verified: true,
    }, nil).Once()

    // user := &domain.User{
    //     ID:       primitive.NewObjectID(),
    //     UserName: "testuser",
    //     Email:    email,
    //     Role:     "user",
    // }

    // expectedAccessToken, err := jwt.GenerateJWT(user.ID.Hex(), user.Email, user.UserName, user.Role)
    // suite.NoError(err, "Generating access token should not return an error")

    // expectedRefreshToken, err := jwt.GenerateRefreshToken(user.ID.Hex(), user.Email, user.Role, user.UserName)
    // suite.NoError(err, "Generating refresh token should not return an error")

    resultUser, _, err := suite.userUsecase.Login(email, password)

    suite.NoError(err, "Login should not return an error")
    suite.NotNil(resultUser, "Result user should not be nil")
    suite.Equal(email, resultUser.Email, "Emails should match")
    //suite.Equal(expectedAccessToken, tokens.AccessToken, "Access tokens should match")
    //suite.Equal(expectedRefreshToken, tokens.RefreshToken, "Refresh tokens should match")

    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLoginInvalidCredentials() {
	email := "test@example.com"
	password := "wrongpassword"

	suite.mockRepo.On("FindUserByEmail", context.TODO(), email).Return(nil, nil)

	resultUser, tokens, err := suite.userUsecase.Login(email, password)

	
	suite.Error(err)
	suite.Nil(resultUser)
	suite.Nil(tokens)
	suite.EqualError(err, "invalid credentials")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLoginUserNotFound() {
	email := "test@example.com"
	password := "password123"

	suite.mockRepo.On("FindUserByEmail", context.TODO(), email).Return(nil, nil)
	resultUser, tokens, err := suite.userUsecase.Login(email, password)

	suite.Error(err)
	suite.Nil(resultUser)
	suite.Nil(tokens)
	suite.EqualError(err, "invalid credentials")
	suite.mockRepo.AssertExpectations(suite.T())
}



func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

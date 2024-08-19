package usecases

// import (
// 	"astu-backend-g1/domain"
// 	"astu-backend-g1/mocks"
// 	"fmt"
// 	"log"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/suite"
// )

// type UserUsecaseTestSuite struct {
// 	suite.Suite
// 	userUsecase    domain.UserUsecase
// 	userRepository *mocks.UserRepository
// 	data           []domain.User
// }

// func (suite *UserUsecaseTestSuite) SetupSuite() {
// 	suite.userRepository = mocks.NewUserRepository(suite.T())
// 	userUsecase, err := NewUserUsecase(suite.userRepository)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	suite.userUsecase = userUsecase
// 	suite.data = []domain.User{
// 		{
// 			ID:        "1",
// 			Username:  "john_doe",
// 			Email:     "john.doe@example.com",
// 			FirstName: "John",
// 			LastName:  "Doe",
// 			Password:  "hashed_password_1",
// 			IsAdmin:   false,
// 			IsActive:  true,
// 		},
// 		{
// 			ID:        "2",
// 			Username:  "jane_smith",
// 			Email:     "jane.smith@example.com",
// 			FirstName: "Jane",
// 			LastName:  "Smith",
// 			Password:  "hashed_password_2",
// 			IsAdmin:   true,
// 			IsActive:  true,
// 		},
// 		{
// 			ID:        "3",
// 			Username:  "mike_jones",
// 			Email:     "mike.jones@example.com",
// 			FirstName: "Mike",
// 			LastName:  "Jones",
// 			Password:  "hashed_password_3",
// 			IsAdmin:   false,
// 			IsActive:  false,
// 		},
// 	}
// }

// func (suite *UserUsecaseTestSuite) TestCreate() {
// 	assert := assert.New(suite.T())
// 	user := suite.data[0]
// 	suite.userRepository.On("Create", &user).Return(user, nil)
// 	createdUser, err := suite.userUsecase.Create(&user)
// 	assert.Nil(err)
// 	assert.Equal(createdUser, user)
// }

// func (suite *UserUsecaseTestSuite) TestGet() {
// 	assert := assert.New(suite.T())
// 	suite.userRepository.On("Get", domain.UserFilterOption{Filter: domain.UserFilter{}}).Return(suite.data, nil)
// 	fetchedUsers, err := suite.userUsecase.Get()
// 	assert.Nil(err)
// 	assert.Equal(fetchedUsers, suite.data)
// }

// func (suite *UserUsecaseTestSuite) TestGetByID() {
// 	assert := assert.New(suite.T())
// 	suite.T().Parallel()
// 	suite.Run("success", func() {
// 		expectedUser := suite.data[1]
// 		suite.userRepository.On("Get", domain.UserFilterOption{Filter: domain.UserFilter{UserId: expectedUser.ID}}).Return([]domain.User{expectedUser}, nil)
// 		fetchedUser, err := suite.userUsecase.GetByID(expectedUser.ID)
// 		assert.Nil(err)
// 		assert.Equal(fetchedUser, expectedUser)
// 	})
// 	suite.Run("not found", func() {
// 		expectedError := fmt.Errorf("there is no user with the given id")
// 		suite.userRepository.On("Get", domain.UserFilterOption{Filter: domain.UserFilter{UserId: "5"}}).Return([]domain.User{suite.data[0]}, expectedError)
// 		_, err := suite.userUsecase.GetByID("5")
// 		assert.ErrorContains(err, expectedError.Error())
// 	})
// }

// func (suite *UserUsecaseTestSuite) TestGetByEmail() {
// 	assert := assert.New(suite.T())
// 	suite.T().Parallel()
// 	suite.Run("success", func() {
// 		expectedUser := suite.data[1]
// 		suite.userRepository.On("Get", domain.UserFilterOption{Filter: domain.UserFilter{Email: expectedUser.Email}}).Return([]domain.User{expectedUser}, nil)
// 		fetchedUser, err := suite.userUsecase.GetByEmail(expectedUser.Email)
// 		assert.Nil(err)
// 		assert.Equal(fetchedUser, expectedUser)
// 	})
// 	suite.Run("not found", func() {
// 		expectedError := fmt.Errorf("there is no user with the email")
// 		suite.userRepository.On("Get", domain.UserFilterOption{Filter: domain.UserFilter{Email: "email@don.exist"}}).Return([]domain.User{suite.data[0]}, expectedError)
// 		_, err := suite.userUsecase.GetByEmail("email@don.exist")
// 		assert.ErrorContains(err, expectedError.Error())
// 	})
// }

// func (suite *UserUsecaseTestSuite) TestGetByUsername() {
// 	assert := assert.New(suite.T())
// 	suite.Run("success", func() {
// 		expectedUser := suite.data[1]
// 		suite.userRepository.On("Get", domain.UserFilterOption{Filter: domain.UserFilter{Username: expectedUser.Username}}).Return([]domain.User{expectedUser}, nil)
// 		fetchedUser, err := suite.userUsecase.GetByUsername(expectedUser.Username)
// 		assert.Nil(err)
// 		assert.Equal(fetchedUser, expectedUser)
// 	})
// 	suite.Run("not found", func() {
// 		expecteduser := suite.data[1]
// 		expectedError := fmt.Errorf("there is no user with the given username")
// 		suite.userRepository.On("Get", domain.UserFilterOption{Filter: domain.UserFilter{Username: "donexist"}}).Return([]domain.User{expecteduser}, expectedError)
// 		_, err := suite.userUsecase.GetByUsername("donexist")
// 		assert.ErrorContains(err, expectedError.Error())
// 	})
// }

// func TestUserUsecase(t *testing.T) {
// 	suite.Run(t, new(UserUsecaseTestSuite))
// }

package tests

// import (
// 	"blog_project/domain"
// 	"blog_project/tests/mocks"
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/suite"
// )

// type UserUsecaseTestSuite struct {
// 	suite.Suite
// 	userRepo    *mocks.MockUserRepository
// 	userUsecase domain.IUserUsecase
// }

// func (suite *UserUsecaseTestSuite) SetupTest() {
// 	suite.userRepo = new(mocks.MockUserRepository)
// 	suite.userUsecase = UserUsecase
// }

// func (suite *UserUsecaseTestSuite) TearDownTest() {
// 	suite.userUsecase = nil
// 	suite.userRepo = nil
// }

// func (suite *UserUsecaseTestSuite) TestCreateUser_ExistingUsername() {
// 	suite.userUsecase.CreateUser(context.TODO(), domain.User{Username: "existing_user", Email: "myemail@gmail.com", Password: "Abcd1234."})
// 	suite.userRepo.On("CreateUser", domain.User{Username: "existing_user", Email: "myemail@gmail.com", Password: "Abcd1234."}).Return(domain.User{}, errors.New("username already in use"))

// 	_, err := suite.userUsecase.CreateUser(context.TODO(), domain.User{Username: "existing_user", Email: "myemail2@gmail.com", Password: "Abcd1234."})

// 	assert.NotNil(suite.T(), err)
// 	assert.EqualErrorf(suite.T(), err, "username already in use", "expected %v, got %v", "username already in use", err)

// }

// func (suite *UserUsecaseTestSuite) TestCreateUser_ExistingEmail() {
// 	suite.userUsecase.CreateUser(context.TODO(), domain.User{Email: "existing_user@gmail.com", Username: "myusername", Password: "Abcd1234."})
// 	suite.userRepo.On("CreateUser", domain.User{Email: "existing_user@gmail.com"}).Return(domain.User{}, errors.New("email already in use"))

// 	_, err := suite.userUsecase.CreateUser(context.TODO(), domain.User{Email: "existing_user@gmail.com", Username: "myusername2", Password: "Abcd1234."})

// 	assert.NotNil(suite.T(), err)
// 	assert.EqualErrorf(suite.T(), err, "email already in use", "expected %v, got %v", "email already in use", err)

// }

// func (suite *UserUsecaseTestSuite) TestCreateUser_InvalidEmail() {
// 	suite.userRepo.On("CreateUser", domain.User{Email: "not_even_an_email", Username: "myusername", Password: "Abcd1234."}).Return(domain.User{}, errors.New("invalid email"))

// 	_, err := suite.userUsecase.CreateUser(context.TODO(), domain.User{Email: "not_even_an_email", Username: "myusername3", Password: "Abcd1234."})

// 	assert.NotNil(suite.T(), err)
// 	assert.EqualErrorf(suite.T(), err, "invalid email", "expected %v, got %v", "invalid email", err)

// }

// func (suite *UserUsecaseTestSuite) TestCreateUser_InvalidPassword() {
// 	suite.userRepo.On("CreateUser", domain.User{Email: "myemail4@gmail.com", Username: "myusername4", Password: "1234"}).Return(domain.User{}, errors.New("invalid password, must contain at least one uppercase letter, one lowercase letter, one number, one special character, and minimum length of 8"))

// 	_, err := suite.userUsecase.CreateUser(context.TODO(), domain.User{Email: "myemail4@gmail.com", Username: "myusername4", Password: "1234"})

// 	assert.NotNil(suite.T(), err)
// 	assert.EqualErrorf(suite.T(), err, "invalid password, must contain at least one uppercase letter, one lowercase letter, one number, one special character, and minimum length of 8", "expected %v, got %v", "invalid password, must contain at least one uppercase letter, one lowercase letter, one number, one special character, and minimum length of 8", err)

// }

// func TestUserUsecaseTestSuite(t *testing.T) {
// 	suite.Run(t, new(UserUsecaseTestSuite))

// }

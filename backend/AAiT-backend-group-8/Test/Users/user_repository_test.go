package test

import (
	domain "AAiT-backend-group-8/Domain"
	mongodb "AAiT-backend-group-8/Infrastructure/mongodb"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestUserSuite struct {
	suite.Suite
	userRepo mongodb.UserRepositoryImpl
}

func (t *TestUserSuite) SetUpTest() {
	client := mongodb.InitMongoDB()
	testRepo := mongodb.CreateCollection(client, "unit-tests", "user")
	t.userRepo = *mongodb.NewUserRepository(
		testRepo,
		context.TODO(),
	)
}

var temp *domain.User = &domain.User{
	Name:     "test",
	Email:    "testacct@gmail.com",
	Password: "123abc",
}

func (t *TestUserSuite) TestCreateUser() {
	assert := assert.New(t.T())
	err := t.userRepo.CreateUser(temp)
	assert.Nil(err, "error should be nil")
}

func (t *TestUserSuite) TearDownTest() {
	t.userRepo.DropDataBase()
}

func (t *TestUserSuite) TestGetUserByEmail() {
	assert := assert.New(t.T())
	local, err := t.userRepo.GetUserByEmail(temp.Email)
	assert.Nil(err, "error should be nil ")
	assert.Equal(local.Email, temp.Email, "they should have the same email")
	assert.Equal(local.Name, temp.Name, "they should have the same name")
	assert.Equal(local.Password, temp.Password, "they should have the same password")

}

func (t *TestUserSuite) TestVerifyUser() {
	assert := assert.New(t.T())
	err := t.userRepo.VerifyUser(temp)

	assert.Nil(err, "error should be nil ")
}

func (t *TestUserSuite) TestGetUserCount() {
	assert := assert.New(t.T())
	number, err := t.userRepo.GetUserCount()
	assert.Nil(err, "error should be nil ")
	assert.GreaterOrEqual(number, int64(1), "there must be a user")
}

func (t *TestUserSuite) TestUpdatePasswordByEmail() {
	assert := assert.New(t.T())
	email := temp.Email
	err := t.userRepo.UpdatePasswordByEmail(email, "12345")
	assert.Nil(err, "error should be nil ")
	newUser, err := t.userRepo.GetUserByEmail(email)
	assert.Nil(err, "error should be nil ")
	assert.Equal(newUser.Password, "12345", "password must be changed")

}

func (t *TestUserSuite) TestPromoteUser() {
	assert := assert.New(t.T())
	email := temp.Email
	err := t.userRepo.PromoteUser(email)
	assert.Nil(err, "error should be nil ")
	newUser, err := t.userRepo.GetUserByEmail(email)
	assert.Nil(err, "error should be nil ")
	assert.Equal(newUser.Role, "admin", "must be promoted to admin")
}

func (t *TestUserSuite) TestDemoteUser() {
	assert := assert.New(t.T())
	email := temp.Email
	err := t.userRepo.DemoteUser(email)
	assert.Nil(err, "should be ")
	newUser, err := t.userRepo.GetUserByEmail(email)
	assert.Nil(err, "error should be nil ")
	assert.Equal(newUser.Role, "user", "must be demoted to user")
}
func (t *TestUserSuite) TestDeleteUser() {
	assert := assert.New(t.T())
	email := temp.Email
	err := t.userRepo.DeleteUser(email)
	assert.Nil(err, "error should be nil ")
	_, err = t.userRepo.GetUserByEmail(email)
	assert.NotNil(err, "should be error")
}
func TestUserRepositorySuite(t *testing.T) {

	suite := new(TestUserSuite)
	suite.SetT(t)
	suite.SetUpTest()
	suite.TestCreateUser()
	suite.TestGetUserByEmail()
	suite.TestVerifyUser()
	suite.TestGetUserCount()
	suite.TestUpdatePasswordByEmail()
	suite.TestPromoteUser()
	suite.TestDemoteUser()
	suite.TestDeleteUser()
	suite.TearDownTest()
}

package repository_test

import (
	"blog/database/mocks"
	"blog/domain"
	"blog/repository"
	"context"
	"errors"
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositorySuite struct {
	suite.Suite
	databaseHelper     *mocks.Database
	collectionHelper   *mocks.Collection
	cursorHelper       *mocks.Cursor
	singleResultHelper *mocks.SingleResult
}

func (suite *UserRepositorySuite) SetupTest() {
	suite.databaseHelper = &mocks.Database{}
	suite.collectionHelper = &mocks.Collection{}
	suite.cursorHelper = &mocks.Cursor{}
	suite.singleResultHelper = &mocks.SingleResult{}
}
func (suite *UserRepositorySuite) TearDownSuite() {
	suite.collectionHelper.AssertExpectations(suite.T())
	suite.databaseHelper.AssertExpectations(suite.T())
	suite.cursorHelper.AssertExpectations(suite.T())
	suite.singleResultHelper.AssertExpectations(suite.T())
}
func (suite *UserRepositorySuite) TestCreateUser() {
	suite.T().Run("CreateUser_Success", func(t *testing.T) {
		user := &domain.User{
			Username: "test",
			Email:    "",
			Password: "test",
		}
		suite.collectionHelper.On("InsertOne", mock.Anything, user).Return(user.ID, nil)
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		err := repo.CreateUser(context.Background(), user)
		suite.Nil(err)
	})
	suite.T().Run("CreateUser_Error", func(t *testing.T) {
		emptyUser := &domain.User{}
		suite.collectionHelper.On("InsertOne", mock.Anything, emptyUser).Return(nil, errors.New("error"))
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		err := repo.CreateUser(context.Background(), emptyUser)
		suite.NotNil(err)
	})
}
func (suite *UserRepositorySuite) TestDeleteUser() {
	suite.T().Run("DeleteUser_Success", func(t *testing.T) {
		id := primitive.NewObjectID()
		suite.collectionHelper.On("DeleteOne", mock.Anything, bson.M{"_id": id}).Return(int64(1), nil).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper).Once()
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		err := repo.DeleteUser(context.Background(), id)
		suite.Nil(err)
	})
	suite.T().Run("DeleteUser_Error", func(t *testing.T) {
		id := primitive.NewObjectID()
		suite.collectionHelper.On("DeleteOne", mock.Anything, bson.M{"_id": id}).Return(int64(0), errors.New("error")).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper).Once()
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		err := repo.DeleteUser(context.Background(), id)
		suite.NotNil(err)
	})
}
func (suite *UserRepositorySuite) TestGetUserByEmail() {
	suite.T().Run("GetUserByEmail_Success", func(t *testing.T) {
		email := "test"
		user := &domain.User{}
		suite.singleResultHelper.On("Decode", user).Return(nil).Once()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"email": email}).Return(suite.singleResultHelper).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper).Once()
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		resp, err := repo.GetUserByEmail(context.Background(), email)
		suite.Nil(err)
		suite.NotNil(resp)
	})

	suite.T().Run("GetUserByEmail_Error", func(t *testing.T) {
		email := "non-existing@gmail.com"
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"email": email}).Return(suite.singleResultHelper).Once()
		suite.singleResultHelper.On("Decode", mock.Anything).Return(mongo.ErrNoDocuments).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper).Once()
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		_, err := repo.GetUserByEmail(context.Background(), email)
		suite.NotNil(err)
	})
}

func (suite *UserRepositorySuite) TestGetUserByID() {
	suite.T().Run("GetUserByID_Success", func(t *testing.T) {
		id := primitive.NewObjectID()
		user := &domain.User{}
		suite.singleResultHelper.On("Decode", user).Return(nil).Once()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"_id": id}).Return(suite.singleResultHelper).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper).Once()
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		resp, err := repo.GetUserByID(context.Background(), id)
		suite.Nil(err)
		suite.NotNil(resp)
	})

	suite.T().Run("GetUserByID_Error", func(t *testing.T) {
		id := primitive.NewObjectID()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"_id": id}).Return(suite.singleResultHelper).Once()
		suite.singleResultHelper.On("Decode", mock.Anything).Return(mongo.ErrNoDocuments).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper).Once()
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		_, err := repo.GetUserByID(context.Background(), id)
		suite.NotNil(err)
	})
}

func (suite *UserRepositorySuite) TestGetUserByUsername() {
	suite.T().Run("GetUserByUsername_Success", func(t *testing.T) {
		username := "test"
		user := &domain.User{}
		suite.singleResultHelper.On("Decode", user).Return(nil).Once()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"username": username}).Return(suite.singleResultHelper).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper).Once()
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		resp, err := repo.GetUserByUsername(context.Background(), username)
		suite.Nil(err)
		suite.NotNil(resp)
	})

	suite.T().Run("GetUserByUsername_Error", func(t *testing.T) {
		username := "non-existing"
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"username": username}).Return(suite.singleResultHelper).Once()
		suite.singleResultHelper.On("Decode", mock.Anything).Return(mongo.ErrNoDocuments).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper).Once()
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		_, err := repo.GetUserByUsername(context.Background(), username)
		suite.NotNil(err)
	})
}

func (suite *UserRepositorySuite) TestGetAllUsers() {
	suite.T().Run("GetAllUsers_Success", func(t *testing.T) {
	suite.collectionHelper.On("Find", mock.Anything, mock.Anything).Return(suite.cursorHelper, nil).Once()
	suite.cursorHelper.On("Close", mock.Anything).Return(nil).Once()
	suite.cursorHelper.On("Next", mock.Anything).Return(false).Once()
	suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)

	repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
	_, err := repo.GetAllUsers(context.Background())
	suite.NoError(err)
	})
	suite.T().Run("GetAllUsers_Error", func(t *testing.T) {
		suite.collectionHelper.On("Find", mock.Anything, mock.Anything).Return(suite.cursorHelper, errors.New("Unexpected")).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		_, err := repo.GetAllUsers(context.Background())
		suite.Error(err)
	})
}

func (suite *UserRepositorySuite) TestUpdateUser() {
	suite.T().Run("UpdateUser_Success", func(t *testing.T) {
	mockUser := domain.User{First_Name: "test", Last_Name: "test", Email: "test", Password: "test", Role: "test"}
	updateresult := &mongo.UpdateResult{}
	suite.collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(updateresult, nil).Once()
	suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)
	repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
	err := repo.UpdateUser(context.Background(), &mockUser)
	suite.NoError(err)
	})
	suite.T().Run("UpdateUser_Error", func(t *testing.T) {
		emptyUser := &domain.User{
		}
		suite.collectionHelper.On("UpdateOne", mock.Anything,mock.Anything, mock.Anything).Return(nil,errors.New("error"))
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		err := repo.UpdateUser(context.Background(), emptyUser)
		suite.NotNil(err)
	})
}

func (suite *UserRepositorySuite) TestUpdatePassword() {
	suite.T().Run("UpdatePassword_Success", func(t *testing.T) {
		mockUser := domain.User{Password: "test", }
		updateresult := &mongo.UpdateResult{}
		suite.collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(updateresult, nil).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		err := repo.UpdatePassword(context.Background(), &mockUser)
		suite.NoError(err)
		})
		suite.T().Run("UpdateUser_Error", func(t *testing.T) {
			emptyUser := &domain.User{}
			suite.collectionHelper.On("UpdateOne", mock.Anything,mock.Anything, mock.Anything).Return(nil,errors.New("error"))
			suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)
			repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
			err := repo.UpdatePassword(context.Background(), emptyUser)
			suite.NotNil(err)
		})
}

func (suite *UserRepositorySuite) TestPromoteUser() {
	suite.T().Run("PromoteUser_Success", func(t *testing.T) {
		id := primitive.NewObjectID()
		updateresult := &mongo.UpdateResult{}
		suite.collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(updateresult, nil).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		err := repo.PromoteUser(context.Background(), id)
		suite.NoError(err)
	})
	suite.T().Run("PromoteUser_Error", func(t *testing.T) {
		id := primitive.NewObjectID()
			suite.collectionHelper.On("UpdateOne", mock.Anything,mock.Anything, mock.Anything).Return(nil,errors.New("error"))
			suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)
			repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
			err := repo.PromoteUser(context.Background(), id)
			suite.NotNil(err)
	})
}

func (suite *UserRepositorySuite) TestDemoteUser() {
	suite.T().Run("DemoteUser_Success", func(t *testing.T) {
		id := primitive.NewObjectID()
		updateresult := &mongo.UpdateResult{}
		suite.collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(updateresult, nil).Once()
		suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)
		repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
		err := repo.DemoteUser(context.Background(), id)
		suite.NoError(err)
	})
	suite.T().Run("DemoteUser_Error", func(t *testing.T) {
		id := primitive.NewObjectID()
			suite.collectionHelper.On("UpdateOne", mock.Anything,mock.Anything, mock.Anything).Return(nil,errors.New("error"))
			suite.databaseHelper.On("Collection", domain.CollectionUser).Return(suite.collectionHelper)
			repo := repository.NewUserRepository(suite.databaseHelper, domain.CollectionUser)
			err := repo.DemoteUser(context.Background(), id)
			suite.NotNil(err)
	})
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

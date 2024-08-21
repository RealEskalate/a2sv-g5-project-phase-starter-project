package repository_test

import (
	"blog/database/mocks"
	"blog/domain"
	"blog/repository"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type OTPRepositorySuite struct {
	suite.Suite
	databaseHelper     *mocks.Database
	collectionHelper   *mocks.Collection
	cursorHelper       *mocks.Cursor
	singleResultHelper *mocks.SingleResult
}

func (suite *OTPRepositorySuite) SetupTest() {
	suite.databaseHelper = &mocks.Database{}
	suite.collectionHelper = &mocks.Collection{}
	suite.cursorHelper = &mocks.Cursor{}
	suite.singleResultHelper = &mocks.SingleResult{}
}
func (suite *OTPRepositorySuite) TearDownSuite() {
	suite.collectionHelper.AssertExpectations(suite.T())
	suite.databaseHelper.AssertExpectations(suite.T())
	suite.cursorHelper.AssertExpectations(suite.T())
	suite.singleResultHelper.AssertExpectations(suite.T())
}

func (suite *OTPRepositorySuite) TestSaveOTP() {
	suite.Run("SaveOTP_succes", func() {
		otp := &domain.OTP{
			Value:     "value",
			Username:  "username",
			Email:     "email",
			Password:  "password",
			ExpiresAt: time.Now(),
			CreatedAt: time.Now(),
		}
		id := "id"
		suite.collectionHelper.On("InsertOne", mock.Anything, otp).Return(id, nil).Once()
		suite.databaseHelper.On("Collection", domain.CollectionOTP).Return(suite.collectionHelper).Once()
		repo := repository.NewOTPRepository(suite.databaseHelper, domain.CollectionOTP)
		err := repo.SaveOTP(context.Background(), otp)
		suite.Nil(err)
	})
	suite.Run("SaveOTP_error", func() {
		otp := &domain.OTP{}
		suite.collectionHelper.On("InsertOne", mock.Anything, otp).Return(nil, errors.New("error")).Once()
		suite.databaseHelper.On("Collection", domain.CollectionOTP).Return(suite.collectionHelper)
		repo := repository.NewOTPRepository(suite.databaseHelper, domain.CollectionOTP)
		err := repo.SaveOTP(context.Background(), otp)
		suite.NotNil(err)
	})
}

func (suite *OTPRepositorySuite) TestGetOTPByEmail() {
	suite.Run("GetOTPByEmail_succes", func() {
		otp := &domain.OTP{

		}
		suite.singleResultHelper.On("Decode", otp).Return(nil).Once()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"email": "email"}).Return(suite.singleResultHelper).Once()
		suite.databaseHelper.On("Collection", domain.CollectionOTP).Return(suite.collectionHelper).Once()
		repo := repository.NewOTPRepository(suite.databaseHelper, domain.CollectionOTP)
		_, err := repo.GetOTPByEmail(context.Background(), "email")
		suite.Nil(err)
	})
	suite.Run("GetOTPByEmail_error", func() {
		otp := &domain.OTP{}
		suite.singleResultHelper.On("Decode", otp).Return(errors.New("error")).Once()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"email": "email"}).Return(suite.singleResultHelper).Once()
		suite.databaseHelper.On("Collection", domain.CollectionOTP).Return(suite.collectionHelper).Once()
		repo := repository.NewOTPRepository(suite.databaseHelper, domain.CollectionOTP)
		_, err := repo.GetOTPByEmail(context.Background(), "email")
		suite.NotNil(err)
	})
}

func (suite *OTPRepositorySuite) TestDeleteOTP() {
	suite.Run("DeleteOTP_succes", func() {
		suite.collectionHelper.On("DeleteOne", mock.Anything, bson.M{"email": "email"}).Return(int64(1), nil).Once()
		suite.databaseHelper.On("Collection", domain.CollectionOTP).Return(suite.collectionHelper).Once()
		repo := repository.NewOTPRepository(suite.databaseHelper, domain.CollectionOTP)
		err := repo.DeleteOTP(context.Background(), "email")
		suite.Nil(err)
	})
	suite.Run("DeleteOTP_error", func() {
		suite.collectionHelper.On("DeleteOne", mock.Anything, bson.M{"email": "email"}).Return(int64(0), errors.New("error")).Once()
		suite.databaseHelper.On("Collection", domain.CollectionOTP).Return(suite.collectionHelper).Once()
		repo := repository.NewOTPRepository(suite.databaseHelper, domain.CollectionOTP)
		err := repo.DeleteOTP(context.Background(), "email")
		suite.NotNil(err)
	})
}

func TestOTPRepositorySuite(t *testing.T) {
	suite.Run(t, new(OTPRepositorySuite))
}

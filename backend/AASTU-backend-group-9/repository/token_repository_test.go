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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenRepositorySuite struct {
	suite.Suite
	databaseHelper     *mocks.Database
	collectionHelper   *mocks.Collection
	cursorHelper       *mocks.Cursor
	singleResultHelper *mocks.SingleResult
}

func (suite *TokenRepositorySuite) SetupTest() {
	suite.databaseHelper = &mocks.Database{}
	suite.collectionHelper = &mocks.Collection{}
	suite.cursorHelper = &mocks.Cursor{}
	suite.singleResultHelper = &mocks.SingleResult{}
}
func (suite *TokenRepositorySuite) TearDownSuite() {
	suite.collectionHelper.AssertExpectations(suite.T())
	suite.databaseHelper.AssertExpectations(suite.T())
	suite.cursorHelper.AssertExpectations(suite.T())
	suite.singleResultHelper.AssertExpectations(suite.T())
}

func (suite *TokenRepositorySuite) TestSaveToken() {
	suite.Run("SaveToke_succes", func() {
		token := &domain.Token{
			ID:           primitive.NewObjectID(),
			UserID:       primitive.NewObjectID(),
			RefreshToken: "refresh",
			ExpiresAt:    time.Now(),
			CreatedAt:    time.Now(),
		}
		suite.collectionHelper.On("InsertOne", mock.Anything, token).Return(token.ID, nil).Once()
		suite.databaseHelper.On("Collection", domain.TokenCollection).Return(suite.collectionHelper).Once()
		repo := repository.NewMongoTokenRepository(suite.databaseHelper, domain.TokenCollection)
		err := repo.SaveToken(context.Background(), token)
		suite.Nil(err)
	})
	suite.Run("SaveToken_error", func() {
		token := &domain.Token{}
		suite.collectionHelper.On("InsertOne", mock.Anything, token).Return(nil, errors.New("error")).Once()
		suite.databaseHelper.On("Collection", domain.TokenCollection).Return(suite.collectionHelper)
		repo := repository.NewMongoTokenRepository(suite.databaseHelper, domain.TokenCollection)
		err := repo.SaveToken(context.Background(), token)
		suite.NotNil(err)
	})
}

func (suite *TokenRepositorySuite) TestFindTokenByAccessToken() {
	suite.Run("FindTokenByAccessToken_success", func() {
		token := &domain.Token{
		}
		suite.singleResultHelper.On("Decode", token).Return(nil).Once()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"access_token": "access"}).Return(suite.singleResultHelper).Once()
		suite.databaseHelper.On("Collection", domain.TokenCollection).Return(suite.collectionHelper).Once()
		repo := repository.NewMongoTokenRepository(suite.databaseHelper, domain.TokenCollection)
		_, err := repo.FindTokenByAccessToken(context.Background(), "access")
		suite.Nil(err)
	})
	suite.Run("FindTokenByAccessToken_error", func() {
		suite.singleResultHelper.On("Decode", &domain.Token{}).Return(errors.New("error")).Once()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"access_token": "access"}).Return(suite.singleResultHelper).Once()
		suite.databaseHelper.On("Collection", domain.TokenCollection).Return(suite.collectionHelper).Once()
		repo := repository.NewMongoTokenRepository(suite.databaseHelper, domain.TokenCollection)
		_, err := repo.FindTokenByAccessToken(context.Background(), "access")
		suite.NotNil(err)
	})
}

func (suite *TokenRepositorySuite) TestFindTokenByRefreshToken() {
	suite.Run("FindTokenByRefreshToken_success", func() {
		token := &domain.Token{
		}
		suite.singleResultHelper.On("Decode", token).Return(nil).Once()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"refresh_token": "refresh"}).Return(suite.singleResultHelper).Once()
		suite.databaseHelper.On("Collection", domain.TokenCollection).Return(suite.collectionHelper).Once()
		repo := repository.NewMongoTokenRepository(suite.databaseHelper, domain.TokenCollection)
		_, err := repo.FindTokenByRefreshToken(context.Background(), "refresh")
		suite.Nil(err)
	})
	suite.Run("FindTokenByRefreshToken_error", func() {
		suite.singleResultHelper.On("Decode", &domain.Token{}).Return(errors.New("error")).Once()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"refresh_token": "refresh"}).Return(suite.singleResultHelper).Once()
		suite.databaseHelper.On("Collection", domain.TokenCollection).Return(suite.collectionHelper).Once()
		repo := repository.NewMongoTokenRepository(suite.databaseHelper, domain.TokenCollection)
		_, err := repo.FindTokenByRefreshToken(context.Background(), "refresh")
		suite.NotNil(err)
	})
}

func (suite *TokenRepositorySuite) TestDeleteToken() {
	suite.Run("DeleteToken_success", func() {
		tokenID := primitive.NewObjectID()
		suite.collectionHelper.On("DeleteOne", mock.Anything, bson.M{"_id": tokenID}).Return(int64(1), nil).Once()
		suite.databaseHelper.On("Collection", domain.TokenCollection).Return(suite.collectionHelper).Once()
		repo := repository.NewMongoTokenRepository(suite.databaseHelper, domain.TokenCollection)
		err := repo.DeleteToken(context.Background(), tokenID)
		suite.Nil(err)
	})
	suite.Run("DeleteToken_error", func() {
		tokenID := primitive.NewObjectID()
		suite.collectionHelper.On("DeleteOne", mock.Anything, bson.M{"_id": tokenID}).Return(int64(0), errors.New("error")).Once()
		suite.databaseHelper.On("Collection", domain.TokenCollection).Return(suite.collectionHelper).Once()
		repo := repository.NewMongoTokenRepository(suite.databaseHelper, domain.TokenCollection)
		err := repo.DeleteToken(context.Background(), tokenID)
		suite.NotNil(err)
	})
}


func TestTokenRepositorySuite(t *testing.T) {
	suite.Run(t, new(TokenRepositorySuite))
}

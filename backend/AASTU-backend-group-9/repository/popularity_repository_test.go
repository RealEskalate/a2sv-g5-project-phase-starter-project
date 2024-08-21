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
)

type PopularityRepositorySuite struct {
	suite.Suite
	databaseHelper     *mocks.Database
	collectionHelper   *mocks.Collection
	cursorHelper       *mocks.Cursor
	singleResultHelper *mocks.SingleResult
}

func (suite *PopularityRepositorySuite) SetupTest() {
	suite.databaseHelper = &mocks.Database{}
	suite.collectionHelper = &mocks.Collection{}
	suite.cursorHelper = &mocks.Cursor{}
	suite.singleResultHelper = &mocks.SingleResult{}
}
func (suite *PopularityRepositorySuite) TearDownSuite() {
	suite.collectionHelper.AssertExpectations(suite.T())
	suite.databaseHelper.AssertExpectations(suite.T())
	suite.cursorHelper.AssertExpectations(suite.T())
	suite.singleResultHelper.AssertExpectations(suite.T())
}

func (suite *PopularityRepositorySuite) TestHasUserLiked() {
	suite.Run("HasLiked_success", func() {
		id := primitive.NewObjectID()
		var count int64 = 1
		suite.collectionHelper.On("CountDocuments", mock.Anything,bson.M{"post_id": id, "user_id": id, "interaction_type": "Like"}).Return(count,nil).Once()
		suite.databaseHelper.On("Collection", "popularity").Return(suite.collectionHelper).Once()
		repo := repository.NewPopularityRepository(suite.databaseHelper, "popularity")
		_, err := repo.HasUserLiked(context.Background(), id, id)
		suite.Nil(err)
	})
	suite.Run("HasLiked_error", func() {
		id := primitive.NewObjectID()
		var count int64 = 1
		suite.collectionHelper.On("CountDocuments", mock.Anything,bson.M{"post_id": id, "user_id": id, "interaction_type": "Like"}).Return(count,errors.New("error")).Once()
		suite.databaseHelper.On("Collection", "popularity").Return(suite.collectionHelper).Once()
		repo := repository.NewPopularityRepository(suite.databaseHelper, "popularity")
		_, err := repo.HasUserLiked(context.Background(), id,id)
		suite.NotNil(err)
	})
}

func (suite *PopularityRepositorySuite) TestHasUserDisliked() {
	suite.Run("HasDisliked_success", func() {
		id := primitive.NewObjectID()
		var count int64 = 1
		suite.collectionHelper.On("CountDocuments", mock.Anything,bson.M{"post_id": id, "user_id": id, "interaction_type": "Dislike"}).Return(count,nil).Once()
		suite.databaseHelper.On("Collection", "popularity").Return(suite.collectionHelper).Once()
		repo := repository.NewPopularityRepository(suite.databaseHelper, "popularity")
		_, err := repo.HasUserDisliked(context.Background(), id, id)
		suite.Nil(err)
	})
	suite.Run("HasDisliked_error", func() {
		id := primitive.NewObjectID()
		var count int64 = 1
		suite.collectionHelper.On("CountDocuments", mock.Anything,bson.M{"post_id": id, "user_id": id, "interaction_type": "Dislike"}).Return(count,errors.New("error")).Once()
		suite.databaseHelper.On("Collection", "popularity").Return(suite.collectionHelper).Once()
		repo := repository.NewPopularityRepository(suite.databaseHelper, "popularity")
		_, err := repo.HasUserDisliked(context.Background(), id,id)
		suite.NotNil(err)
	})
}

func (suite *PopularityRepositorySuite) TestUserInteractionsAdder() {
	suite.Run("Adder_success", func() {
		user := domain.UserInteraction{}
		suite.collectionHelper.On("InsertOne", mock.Anything,user).Return(nil,nil).Once()
		suite.databaseHelper.On("Collection", "popularity").Return(suite.collectionHelper).Once()
		repo := repository.NewPopularityRepository(suite.databaseHelper, "popularity")
		err := repo.UserInteractionsAdder(context.Background(), user)
		suite.Nil(err)
	})
	suite.Run("Adder_error", func() {
		user := domain.UserInteraction{}
		suite.collectionHelper.On("InsertOne", mock.Anything,user).Return(nil,errors.New("error")).Once()
		suite.databaseHelper.On("Collection", "popularity").Return(suite.collectionHelper).Once()
		repo := repository.NewPopularityRepository(suite.databaseHelper, "popularity")
		err := repo.UserInteractionsAdder(context.Background(), user)
		suite.NotNil(err)
	})
}

func (suite *PopularityRepositorySuite) TestUserInteractionsDelete() {
	suite.Run("Delete_success", func() {
		user := domain.UserInteraction{}
		suite.collectionHelper.On("DeleteOne", mock.Anything,user).Return(int64(1),nil).Once()
		suite.databaseHelper.On("Collection", "popularity").Return(suite.collectionHelper).Once()
		repo := repository.NewPopularityRepository(suite.databaseHelper, "popularity")
		err := repo.UserInteractionsDelete(context.Background(), user)
		suite.Nil(err)
	})
	suite.Run("Delete_error", func() {
		user :=domain.UserInteraction{}
		suite.collectionHelper.On("DeleteOne", mock.Anything,user).Return(int64(0),errors.New("error")).Once()
		suite.databaseHelper.On("Collection", "popularity").Return(suite.collectionHelper).Once()
		repo := repository.NewPopularityRepository(suite.databaseHelper, "popularity")
		err := repo.UserInteractionsDelete(context.Background(), user)
		suite.NotNil(err)
	})
}

func TestPopularityRepositorySuite(t *testing.T) {
	suite.Run(t, new(PopularityRepositorySuite))
}

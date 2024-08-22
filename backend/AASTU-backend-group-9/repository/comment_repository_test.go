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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepositorySuite struct {
	suite.Suite
	databaseHelper     *mocks.Database
	collectionHelper   *mocks.Collection
	cursorHelper       *mocks.Cursor
	singleResultHelper *mocks.SingleResult
}

func (suite *CommentRepositorySuite) SetupTest() {
	suite.databaseHelper = &mocks.Database{}
	suite.collectionHelper = &mocks.Collection{}
	suite.cursorHelper = &mocks.Cursor{}
	suite.singleResultHelper = &mocks.SingleResult{}
}
func (suite *CommentRepositorySuite) TearDownSuite() {
	suite.collectionHelper.AssertExpectations(suite.T())
	suite.databaseHelper.AssertExpectations(suite.T())
	suite.cursorHelper.AssertExpectations(suite.T())
	suite.singleResultHelper.AssertExpectations(suite.T())
}

func (suite *CommentRepositorySuite) TestAddComment() {
	suite.Run("AddComment_success", func() {
		comment := &domain.Comment{
			Content: "content",
		}
		id := primitive.NewObjectID()
		suite.collectionHelper.On("InsertOne", mock.Anything, comment).Return(id, nil).Once()
		suite.databaseHelper.On("Collection", "comments").Return(suite.collectionHelper).Once()
		repo := repository.NewCommentRepository(suite.databaseHelper, "comments")
		err := repo.AddComment(context.Background(), id, id, comment)
		suite.Nil(err)
	})
	suite.Run("AddComment_error", func() {
		comment := &domain.Comment{
			Content: "content",
		}
		id := primitive.NewObjectID()
		suite.collectionHelper.On("InsertOne", mock.Anything, comment).Return(nil, errors.New("insert error")).Once()
		suite.databaseHelper.On("Collection", "comments").Return(suite.collectionHelper).Once()
		repo := repository.NewCommentRepository(suite.databaseHelper, "comments")
		err := repo.AddComment(context.Background(), id, id, comment)
		suite.NotNil(err)
	})
}
func TestCommentRepositorySuite(t *testing.T) {
	suite.Run(t, new(CommentRepositorySuite))
}

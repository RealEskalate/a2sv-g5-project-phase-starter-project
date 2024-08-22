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
func (suite *CommentRepositorySuite) TestGetComments() {
	suite.Run("GetComments_succes", func() {
		comment := &domain.Comment{
		}
		id := primitive.NewObjectID()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"blog_id": id}).Return(suite.singleResultHelper).Once()
		suite.singleResultHelper.On("Decode", comment).Return(nil).Once()
		suite.databaseHelper.On("Collection", "comments").Return(suite.collectionHelper)
		repo := repository.NewCommentRepository(suite.databaseHelper, "comments")
		_, err := repo.GetComments(context.Background(), id)
		suite.Nil(err)
	})
	suite.Run("GetComments_error", func() {
		comment := &domain.Comment{}
		id := primitive.NewObjectID()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"blog_id": id}).Return(suite.singleResultHelper).Once()
		suite.singleResultHelper.On("Decode", comment).Return(errors.New("error")).Once()
		suite.databaseHelper.On("Collection", "comments").Return(suite.collectionHelper)
		repo := repository.NewCommentRepository(suite.databaseHelper, "comments")
		_, err := repo.GetComments(context.Background(), id)
		suite.NotNil(err)
	})
}

func (suite *CommentRepositorySuite) TestDeleteComment() {
	suite.Run("DeleteComment_succes", func() {
		id := primitive.NewObjectID()
		comment := &domain.Comment{
		}
		suite.databaseHelper.On("Collection", "comments").Return(suite.collectionHelper).Once()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"_id": id}).Return(suite.singleResultHelper).Once()
		suite.singleResultHelper.On("Decode", comment).Return(nil).Once()
		suite.collectionHelper.On("DeleteOne", mock.Anything, bson.M{"_id": id}).Return(int64(1),nil).Once()
		repo := repository.NewCommentRepository(suite.databaseHelper, "comments")
		err := repo.DeleteComment(context.Background(), id, id, comment.AuthorID)
		suite.Nil(err)
	})
	suite.Run("DeleteComment_error", func() {
		id := primitive.NewObjectID()
		comment := &domain.Comment{}
		suite.databaseHelper.On("Collection", "comments").Return(suite.collectionHelper).Once()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"_id": id}).Return(suite.singleResultHelper).Once()
		suite.singleResultHelper.On("Decode", comment).Return(nil).Once()
		suite.collectionHelper.On("DeleteOne", mock.Anything, bson.M{"_id": id}).Return(int64(0),errors.New("error")).Once()
		
		repo := repository.NewCommentRepository(suite.databaseHelper, "comments")
		err := repo.DeleteComment(context.Background(), id, id, id)
		suite.NotNil(err)
	})
}

func (suite *CommentRepositorySuite) TestUpdateComment() {
	suite.Run("UpdateComment_succes", func() {
		comment := &domain.Comment{
		}
		id := primitive.NewObjectID()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"_id": id}).Return(suite.singleResultHelper).Once()
		suite.singleResultHelper.On("Decode", comment).Return(nil).Once()
		suite.collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()
		suite.databaseHelper.On("Collection", "comments").Return(suite.collectionHelper)
		repo := repository.NewCommentRepository(suite.databaseHelper, "comments")
		_ = id == comment.AuthorID
		err := repo.UpdateComment(context.Background(), id, id, comment.AuthorID, comment)
		suite.Nil(err)
	})
	suite.Run("UpdateComment_error", func() {
		comment := &domain.Comment{
		}
		id := primitive.NewObjectID()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"_id": id}).Return(suite.singleResultHelper).Once()
		suite.singleResultHelper.On("Decode", comment).Return(nil).Once()		
		suite.databaseHelper.On("Collection", "comments").Return(suite.collectionHelper)
		suite.collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()

		repo := repository.NewCommentRepository(suite.databaseHelper, "comments")
		err := repo.UpdateComment(context.Background(), id, id, id, comment)
		suite.NotNil(err)
	})
}

func TestCommentRepositorySuite(t *testing.T) {
	suite.Run(t, new(CommentRepositorySuite))
}

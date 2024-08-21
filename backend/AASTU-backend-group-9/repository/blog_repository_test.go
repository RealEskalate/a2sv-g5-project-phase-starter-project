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

type BlogRepositorySuite struct {
	suite.Suite
	databaseHelper     *mocks.Database
	collectionHelper   *mocks.Collection
	cursorHelper       *mocks.Cursor
	singleResultHelper *mocks.SingleResult
}

func (suite *BlogRepositorySuite) SetupTest() {
	suite.databaseHelper = &mocks.Database{}
	suite.collectionHelper = &mocks.Collection{}
	suite.cursorHelper = &mocks.Cursor{}
	suite.singleResultHelper = &mocks.SingleResult{}
}
func (suite *BlogRepositorySuite) TearDownSuite() {
	suite.collectionHelper.AssertExpectations(suite.T())
	suite.databaseHelper.AssertExpectations(suite.T())
	suite.cursorHelper.AssertExpectations(suite.T())
	suite.singleResultHelper.AssertExpectations(suite.T())
}

func (suite *BlogRepositorySuite) TestCreateBlog() {
	suite.Run("CreateBlog_succes", func() {
		blog := &domain.Blog{
			Title:   "title",
			Content: "content",
		}
		id := "id"
		suite.collectionHelper.On("InsertOne", mock.Anything, blog).Return(id, nil).Once()
		suite.databaseHelper.On("Collection", "blogs").Return(suite.collectionHelper).Once()
		repo := repository.NewBlogRepository(suite.databaseHelper, "blogs")
		err := repo.CreateBlog(context.Background(), blog)
		suite.Nil(err)
	})
	suite.Run("CreateBlog_error", func() {
		blog := &domain.Blog{}
		suite.collectionHelper.On("InsertOne", mock.Anything, blog).Return(nil, errors.New("error")).Once()
		suite.databaseHelper.On("Collection", "blogs").Return(suite.collectionHelper)
		repo := repository.NewBlogRepository(suite.databaseHelper, "blogs")
		err := repo.CreateBlog(context.Background(), blog)
		suite.NotNil(err)
	})
}

func (suite *BlogRepositorySuite) TestGetBlogByID() {
	suite.Run("GetBlogByID_succes", func() {
		blog := &domain.Blog{}
		id := primitive.NewObjectID()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"_id": id}).Return(suite.singleResultHelper).Once()
		suite.singleResultHelper.On("Decode", blog).Return(nil).Once()
		suite.databaseHelper.On("Collection", "blogs").Return(suite.collectionHelper).Once()
		repo := repository.NewBlogRepository(suite.databaseHelper, "blogs")
		_, err := repo.GetBlogByID(context.Background(), id)
		suite.Nil(err)
	})
	suite.Run("GetBlogByID_error", func() {
		id := primitive.NewObjectID()
		suite.collectionHelper.On("FindOne", mock.Anything, bson.M{"_id": id}).Return(suite.singleResultHelper).Once()
		suite.singleResultHelper.On("Decode", mock.Anything).Return(mongo.ErrNoDocuments).Once()
		suite.databaseHelper.On("Collection", "blogs").Return(suite.collectionHelper).Once()
		repo := repository.NewBlogRepository(suite.databaseHelper, "blogs")
		_, err := repo.GetBlogByID(context.Background(), id)
		suite.NotNil(err)
	})
}

func (suite *BlogRepositorySuite) TestUpdateBlog() {
	suite.Run("UpdateBlog_succes", func() {
		blog := &domain.Blog{
			Title:   "title",
			Content: "content",
		}
		updateresult := &mongo.UpdateResult{}
		suite.collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(updateresult, nil).Once()
		suite.databaseHelper.On("Collection", "blogs").Return(suite.collectionHelper).Once()
		repo := repository.NewBlogRepository(suite.databaseHelper, "blogs")
		err := repo.UpdateBlog(context.Background(), blog)
		suite.Nil(err)
	})
	suite.Run("UpdateBlog_error", func() {
		blog := &domain.Blog{
			Title:   "title",
			Content: "content",
		}
		suite.collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
		suite.databaseHelper.On("Collection", "blogs").Return(suite.collectionHelper).Once()
		repo := repository.NewBlogRepository(suite.databaseHelper, "blogs")
		err := repo.UpdateBlog(context.Background(), blog)
		suite.NotNil(err)
	})
}

func (suite *BlogRepositorySuite) TestDeleteBlog() {
	suite.Run("DeleteBlog_succes", func() {
		id := primitive.NewObjectID()
		suite.collectionHelper.On("DeleteOne", mock.Anything, bson.M{"_id": id}).Return(int64(1), nil).Once()
		suite.databaseHelper.On("Collection", "blogs").Return(suite.collectionHelper).Once()
		repo := repository.NewBlogRepository(suite.databaseHelper, "blogs")
		err := repo.DeleteBlog(context.Background(), id)
		suite.Nil(err)
	})
	suite.Run("DeleteBlog_error", func() {
		id := primitive.NewObjectID()
		suite.collectionHelper.On("DeleteOne", mock.Anything, bson.M{"_id": id}).Return(int64(0), errors.New("error")).Once()
		suite.databaseHelper.On("Collection", "blogs").Return(suite.collectionHelper).Once()
		repo := repository.NewBlogRepository(suite.databaseHelper, "blogs")
		err := repo.DeleteBlog(context.Background(), id)
		suite.NotNil(err)
	})
}

func (suite *BlogRepositorySuite) TestAddcomment() {
	suite.Run("Addcomment_succes", func() {
		id := primitive.NewObjectID()
		comment := &domain.Comment{}
		updat := &mongo.UpdateResult{}
		suite.collectionHelper.On("UpdateOne", mock.Anything, bson.M{"_id": id}, bson.M{"$push": bson.M{"comments": comment}}).Return(updat, nil).Once()
		suite.databaseHelper.On("Collection", "blogs").Return(suite.collectionHelper).Once()
		repo := repository.NewBlogRepository(suite.databaseHelper, "blogs")
		err := repo.AddComment(context.Background(), id, comment)
		suite.Nil(err)
	})
	suite.Run("Addcomment_error", func() {
		id := primitive.NewObjectID()
		comment := &domain.Comment{}
		suite.collectionHelper.On("UpdateOne", mock.Anything, bson.M{"_id": id}, bson.M{"$push": bson.M{"comments": comment}}).Return(nil, errors.New("error")).Once()
		suite.databaseHelper.On("Collection", "blogs").Return(suite.collectionHelper).Once()
		repo := repository.NewBlogRepository(suite.databaseHelper, "blogs")
		err := repo.AddComment(context.Background(), id, comment)
		suite.NotNil(err)
	})
}

func TestBlogRepositorySuite(t *testing.T) {
	suite.Run(t, new(BlogRepositorySuite))
}

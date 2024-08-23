package test

import (
	domain "AAiT-backend-group-8/Domain"

	mongodb "AAiT-backend-group-8/Infrastructure/mongodb"

	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestLikeSuite struct {
	suite.Suite
	likeRepo mongodb.LikeRepository
}

var likeObj *domain.Like = &domain.Like{
	UserID: primitive.NewObjectID(),
	BlogID: primitive.NewObjectID(),
}

func (t *TestLikeSuite) SetUpLikeTest() {
	client := mongodb.InitMongoDB()
	testRepo := mongodb.CreateCollection(client, "unit-tests", "like")
	t.likeRepo = *mongodb.NewLikeRepository(
		testRepo, context.TODO(),
	)
}

func (t *TestLikeSuite) TestLikeBlog() {
	assert := assert.New(t.T())
	err := t.likeRepo.LikeBlog(*likeObj)
	assert.Nil(err, "error should be nil ")
}

func (t *TestLikeSuite) TestUnlikeBlog() {
	assert := assert.New(t.T())
	// First like the blog to have something to unlike
	err := t.likeRepo.LikeBlog(*likeObj)
	assert.Nil(err, "error should be nil ")

	// Now unlike the blog
	err = t.likeRepo.UnlikeBlog(likeObj.Id)
	assert.Nil(err, "error should be nil ")
}

func (t *TestLikeSuite) TestGetLikes() {
	assert := assert.New(t.T())
	_, err := t.likeRepo.GetLikes(likeObj.BlogID)
	assert.Nil(err, "error should be nil ")
}

func (t *TestLikeSuite) TestCheckIfLiked() {
	assert := assert.New(t.T())
	// First like the blog to have something to check
	err := t.likeRepo.LikeBlog(*likeObj)
	assert.Nil(err, "error should be nil ")

	isLiked, like := t.likeRepo.CheckIfLiked(likeObj.UserID, likeObj.BlogID)
	assert.True(isLiked, "should be liked")
	assert.Equal(like.UserID, likeObj.UserID, "user ID should match")
	assert.Equal(like.BlogID, likeObj.BlogID, "blog ID should match")
}

func (t *TestLikeSuite) TestDeleteByBlogID() {
	assert := assert.New(t.T())
	// First like the blog to have something to delete
	err := t.likeRepo.LikeBlog(*likeObj)
	assert.Nil(err, "error should be nil ")

	err = t.likeRepo.DeleteByBLogID(likeObj.BlogID)
	assert.Nil(err, "error should be nil ")
}

func (t *TestLikeSuite) TearDownTest() {
	assert := assert.New(t.T())
	// err := t.likeRepo.collection.Drop(context.TODO())
	err := t.likeRepo.DropDataBase()
	assert.Nil(err, "error should be nil ")
}

func TestLikeRepositorySuite(t *testing.T) {
	suite := new(TestLikeSuite)
	suite.SetT(t)
	suite.SetUpLikeTest()
	// call the other functions
	suite.TestLikeBlog()
	suite.TestGetLikes()
	suite.TestCheckIfLiked()
	suite.TestDeleteByBlogID()
	// suite.TestUnlikeBlog()
	suite.TearDownTest()
}

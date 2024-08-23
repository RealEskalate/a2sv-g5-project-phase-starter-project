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

type TestCommentSuite struct {
	suite.Suite
	commentRepo mongodb.CommentRepository
}

var comment *domain.Comment = &domain.Comment{
	Body:       "A test comment",
	AuthorName: "test user",
	AuthorID:   primitive.NewObjectID(),
	BlogID:     primitive.NewObjectID(),
}

func (t *TestCommentSuite) SetUpCommentTest() {
	client := mongodb.InitMongoDB()
	testRepo := mongodb.CreateCollection(client, "unit-tests", "comment")
	t.commentRepo = *mongodb.NewCommentRepository(
		testRepo, context.TODO(),
	)
}

func (t *TestCommentSuite) TestCreateComment() {
	assert := assert.New(t.T())
	err := t.commentRepo.CreateComment(comment)
	assert.Nil(err, "error should be nil ")
}

func (t *TestCommentSuite) TestGetComments() {
	assert := assert.New(t.T())
	comments, err := t.commentRepo.GetComments(comment.BlogID)
	assert.Nil(err, "error should be nil ")
	assert.Equal(comments[0].Body, comment.Body, "their bodies should be equal")
	assert.Equal(comments[0].AuthorName, comment.AuthorName, "author name should be same")
}

func (t *TestCommentSuite) TestDeleteComment() {
	assert := assert.New(t.T())
	err := t.commentRepo.DeleteComment(primitive.NewObjectID())
	assert.Nil(err, "error should be nil ")
	_, err = t.commentRepo.GetCommentByID(primitive.NewObjectID())
	assert.NotNil(err, "there should be some error")
}

func (t *TestCommentSuite) TearDownTest() {
	assert := assert.New(t.T())
	err := t.commentRepo.DropDataBase()
	assert.Nil(err, "error should be nil ")

}

func TestTaskRepositorySuite(t *testing.T) {
	suite := new(TestCommentSuite)
	suite.SetT(t)
	suite.SetUpCommentTest()
	// call the other functions
	suite.TestCreateComment()
	suite.TestGetComments()
	// suite.TestDeleteComment()
	suite.TearDownTest()

}

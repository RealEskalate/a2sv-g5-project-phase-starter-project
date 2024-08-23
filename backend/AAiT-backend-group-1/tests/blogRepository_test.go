package tests

import (
    "fmt"
    "testing"
    "time"

    "github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
    "github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/mocks"
    "github.com/stretchr/testify/suite"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogTestSuite struct {
    suite.Suite
    mockRepo *mocks.MockBlogRepository
}

func (suite *BlogTestSuite) SetupTest() {
    suite.mockRepo = new(mocks.MockBlogRepository)
}

func (suite *BlogTestSuite) TearDownTest() {
    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogTestSuite) TestCreateBlog() {
    mockBlog := &domain.Blog{
        ID:        primitive.NewObjectID(),
        Title:     "Test Blog",
        Content:   "Test Content",
        AuthorID:  primitive.NewObjectID(),
        CreatedAt: time.Now(),
    }

    suite.mockRepo.On("Create", mockBlog).Return(mockBlog, nil)

    result, err := suite.mockRepo.Create(mockBlog)

    suite.Nil(err)
    suite.Equal(mockBlog, result)
}

func (suite *BlogTestSuite) TestFindById() {
    mockBlog := &domain.Blog{
        ID:        primitive.NewObjectID(),
        Title:     "Test Blog",
        Content:   "Test Content",
        AuthorID:  primitive.NewObjectID(),
        CreatedAt: time.Now(),
    }

    suite.mockRepo.On("FindById", "1").Return(mockBlog, nil)

    result, err := suite.mockRepo.FindById("1")

    suite.Nil(err)
    suite.Equal(mockBlog, result)
}

func (suite *BlogTestSuite) TestFindById_Error() {
    mockError := &domain.CustomError{
        Code:    404,
        Message: "Blog not found",
    }

    suite.mockRepo.On("FindById", "8").Return(nil, mockError)

    result, err := suite.mockRepo.FindById("8")
    fmt.Println(result, err)
    suite.Nil(result)
    suite.NotNil(err)
    suite.Equal(404, err.StatusCode())
    suite.Equal("Blog not found", err.Error())
}

func (suite *BlogTestSuite) TestUpdateBlog() {
    mockBlog := &domain.Blog{
        ID:        primitive.NewObjectID(),
        Title:     "Updated Blog",
        Content:   "Updated Content",
        AuthorID:  primitive.NewObjectID(),
        CreatedAt: time.Now(),
    }

    suite.mockRepo.On("Update", "1", mockBlog).Return(mockBlog, nil)

    result, err := suite.mockRepo.Update("1", mockBlog)

    suite.Nil(err)
    suite.Equal(mockBlog, result)
}

func (suite *BlogTestSuite) TestDeleteBlog() {
    suite.mockRepo.On("Delete", "1").Return(nil)

    err := suite.mockRepo.Delete("1")

    suite.Nil(err)
}

func (suite *BlogTestSuite) TestDeleteBlog_Error() {
    mockError := &domain.CustomError{
        Code:    500,
        Message: "Internal Server Error",
    }

    suite.mockRepo.On("Delete", "1").Return(mockError)

    err := suite.mockRepo.Delete("1")

    suite.NotNil(err)
    suite.Equal(500, err.StatusCode())
    suite.Equal("Internal Server Error", err.Error())
}

func TestBlogTestSuite(t *testing.T) {
    suite.Run(t, new(BlogTestSuite))
}

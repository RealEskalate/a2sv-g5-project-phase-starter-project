package Usecases

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/suite"

// 	"ASTU-backend-group-3/Blog_manager/Domain" // Adjust the import path
// 	"ASTU-backend-group-3/Blog_manager/mocks"  // Adjust the import path
// )

// type BlogUsecaseTestSuite struct {
// 	suite.Suite
// 	mockRepo *mocks.BlogRepository
// 	usecase  BlogUsecase
// }

// func (suite *BlogUsecaseTestSuite) SetupTest() {
// 	suite.mockRepo = new(mocks.BlogRepository)
// 	suite.usecase = NewBlogUsecase(suite.mockRepo)
// }

// func (suite *BlogUsecaseTestSuite) TestCreateBlog_Success() {
// 	blog := &Domain.Blog{Title: "Learning Go", Content: "Content", Tags: []string{"Go"}}
// 	createdBlog := *blog
// 	createdBlog.Id = "blog123"

// 	suite.mockRepo.On("Save", blog).Return(&createdBlog, nil)

// 	result, err := suite.usecase.CreateBlog(blog)

// 	assert.Nil(suite.T(), err)
// 	assert.Equal(suite.T(), &createdBlog, result)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestCreateBlog_Error() {
// 	blog := &Domain.Blog{Title: "Learning Go", Content: "Content", Tags: []string{"Go"}}

// 	suite.mockRepo.On("Save", blog).Return(nil, errors.New("failed to create blog"))

// 	result, err := suite.usecase.CreateBlog(blog)

// 	assert.NotNil(suite.T(), err)
// 	assert.Equal(suite.T(), "failed to create blog", err.Error())
// 	assert.Nil(suite.T(), result)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestUpdateBlog_Success() {
// 	blogID := "blog123"
// 	input := Domain.UpdateBlogInput{Title: "Updated Title", Content: "Updated Content", Tags: []string{"Updated"}}
// 	author := "author123"
// 	existingBlog := &Domain.Blog{Id: blogID, Author: author}
// 	updatedBlog := *existingBlog
// 	updatedBlog.Title = input.Title
// 	updatedBlog.Content = input.Content
// 	updatedBlog.Tags = input.Tags
// 	updatedBlog.UpdatedAt = time.Now().Format(time.RFC3339)

// 	suite.mockRepo.On("FindByID", blogID).Return(existingBlog, nil)
// 	suite.mockRepo.On("Save", &updatedBlog).Return(&updatedBlog, nil)

// 	result, err := suite.usecase.UpdateBlog(blogID, input, author)

// 	assert.Nil(suite.T(), err)
// 	assert.Equal(suite.T(), &updatedBlog, result)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestUpdateBlog_Error() {
// 	blogID := "blog123"
// 	input := Domain.UpdateBlogInput{Title: "Updated Title", Content: "Updated Content", Tags: []string{"Updated"}}
// 	author := "author123"

// 	suite.mockRepo.On("FindByID", blogID).Return(nil, errors.New("blog not found"))

// 	result, err := suite.usecase.UpdateBlog(blogID, input, author)

// 	assert.NotNil(suite.T(), err)
// 	assert.Equal(suite.T(), "blog not found", err.Error())
// 	assert.Nil(suite.T(), result)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestRetrieveBlogs_Success() {
// 	page, pageSize := 1, 10
// 	sortBy := "latest"
// 	blogs := []Domain.Blog{{Id: "blog123", Title: "Learning Go"}}
// 	total := int64(1)

// 	suite.mockRepo.On("RetrieveBlogs", page, pageSize, sortBy).Return(blogs, total, nil)

// 	result, count, err := suite.usecase.RetrieveBlogs(page, pageSize, sortBy)

// 	assert.Nil(suite.T(), err)
// 	assert.Equal(suite.T(), blogs, result)
// 	assert.Equal(suite.T(), total, count)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestRetrieveBlogs_Error() {
// 	page, pageSize := 1, 10
// 	sortBy := "latest"

// 	suite.mockRepo.On("RetrieveBlogs", page, pageSize, sortBy).Return(nil, int64(0), errors.New("failed to retrieve blogs"))

// 	result, count, err := suite.usecase.RetrieveBlogs(page, pageSize, sortBy)

// 	assert.NotNil(suite.T(), err)
// 	assert.Equal(suite.T(), "failed to retrieve blogs", err.Error())
// 	assert.Nil(suite.T(), result)
// 	assert.Equal(suite.T(), int64(0), count) // Ensure the expected value is int64
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestSearchBlogs_Success() {
// 	title := "Go"
// 	author := "author123"
// 	tags := []string{"Go"}
// 	blogs := []Domain.Blog{{Id: "blog123", Title: "Learning Go"}}

// 	suite.mockRepo.On("SearchBlogs", title, author, tags).Return(blogs, nil)

// 	result, err := suite.usecase.SearchBlogs(title, author, tags)

// 	assert.Nil(suite.T(), err)
// 	assert.Equal(suite.T(), blogs, result)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestSearchBlogs_Error() {
// 	title := "Go"
// 	author := "author123"
// 	tags := []string{"Go"}

// 	suite.mockRepo.On("SearchBlogs", title, author, tags).Return(nil, errors.New("failed to search blogs"))

// 	result, err := suite.usecase.SearchBlogs(title, author, tags)

// 	assert.NotNil(suite.T(), err)
// 	assert.Equal(suite.T(), "failed to search blogs", err.Error())
// 	assert.Nil(suite.T(), result)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestAddComment_Success() {
// 	blogID := "blog123"
// 	comment := Domain.Comment{Content: "Nice post!", UserID: "user123"}

// 	suite.mockRepo.On("AddComment", blogID, comment).Return(nil)

// 	err := suite.usecase.AddComment(blogID, comment)

// 	assert.Nil(suite.T(), err)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestAddComment_Error() {
// 	blogID := "blog123"
// 	comment := Domain.Comment{Content: "Nice post!", UserID: "user123"}

// 	suite.mockRepo.On("AddComment", blogID, comment).Return(errors.New("failed to add comment"))

// 	err := suite.usecase.AddComment(blogID, comment)

// 	assert.NotNil(suite.T(), err)
// 	assert.Equal(suite.T(), "failed to add comment", err.Error())
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestDeleteBlogByID_Success() {
// 	blogID := "blog123"

// 	suite.mockRepo.On("DeleteBlogByID", blogID).Return(nil)

// 	err := suite.usecase.DeleteBlogByID(blogID)

// 	assert.Nil(suite.T(), err)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestDeleteBlogByID_Error() {
// 	blogID := "blog123"

// 	suite.mockRepo.On("DeleteBlogByID", blogID).Return(errors.New("failed to delete blog"))

// 	err := suite.usecase.DeleteBlogByID(blogID)

// 	assert.NotNil(suite.T(), err)
// 	assert.Equal(suite.T(), "failed to delete blog", err.Error())
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestFilterBlogs_Success() {
// 	tags := []string{"technology", "programming", "golang"}
// 	startDate, _ := time.Parse(time.RFC3339, "2024-08-01T00:00:00Z")
// 	endDate, _ := time.Parse(time.RFC3339, "2024-08-20T23:59:59Z")
// 	sortBy := "popularity"
// 	suite.mockRepo.On("FilterBlogs", tags, startDate, endDate, sortBy).Return([]Domain.Blog{}, nil)
// 	res, err := suite.usecase.FilterBlogs(tags, startDate, endDate, sortBy)

// 	assert.Nil(suite.T(), err)
// 	assert.Equal(suite.T(), []Domain.Blog{}, res)
// 	suite.mockRepo.AssertExpectations(suite.T())

// }

// func (suite *BlogUsecaseTestSuite) TestFilterBlogs_Error() {
// 	tags := []string{"technology", "programming", "golang"}
// 	startDate, _ := time.Parse(time.RFC3339, "2024-08-01T00:00:00Z")
// 	endDate, _ := time.Parse(time.RFC3339, "2024-08-20T23:59:59Z")
// 	sortBy := "popularity"
// 	suite.mockRepo.On("FilterBlogs", tags, startDate, endDate, sortBy).Return(nil, errors.New("invalid filter"))
// 	res, err := suite.usecase.FilterBlogs(tags, startDate, endDate, sortBy)

// 	assert.NotNil(suite.T(), err)
// 	assert.Empty(suite.T(), res)
// 	assert.Equal(suite.T(), "invalid filter", err.Error())
// 	suite.mockRepo.AssertExpectations(suite.T())

// }

// func (suite *BlogUsecaseTestSuite) TestFindByID_Success() {
// 	blogID := "blog123"
// 	blog := &Domain.Blog{Id: blogID, Title: "Learning Go"}

// 	suite.mockRepo.On("FindByID", blogID).Return(blog, nil)

// 	result, err := suite.usecase.FindByID(blogID)

// 	assert.Nil(suite.T(), err)
// 	assert.Equal(suite.T(), blog, result)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestFindByID_Error() {
// 	blogID := "blog123"

// 	suite.mockRepo.On("FindByID", blogID).Return(&Domain.Blog{}, errors.New("blog not found"))

// 	result, err := suite.usecase.FindByID(blogID)

// 	assert.NotNil(suite.T(), err)
// 	assert.Equal(suite.T(), &Domain.Blog{}, result)
// 	assert.Equal(suite.T(), "blog not found", err.Error())
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestIncrementViewCount_Success() {
// 	blogID := "blog123"

// 	suite.mockRepo.On("IncrementViewCount", blogID).Return(nil)

// 	err := suite.usecase.IncrementViewCount(blogID)

// 	assert.Nil(suite.T(), err)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestIncrementViewCount_Error() {
// 	blogID := "blog123"

// 	suite.mockRepo.On("IncrementViewCount", blogID).Return(errors.New("failed to increment view count"))

// 	err := suite.usecase.IncrementViewCount(blogID)

// 	assert.NotNil(suite.T(), err)
// 	assert.Equal(suite.T(), "failed to increment view count", err.Error())
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestToggleLike_Success() {
// 	blogID := "blog123"
// 	username := "user123"

// 	suite.mockRepo.On("ToggleLike", blogID, username).Return(nil)

// 	err := suite.usecase.ToggleLike(blogID, username)

// 	assert.Nil(suite.T(), err)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestToggleLike_Error() {
// 	blogID := "blog123"
// 	username := "user123"

// 	suite.mockRepo.On("ToggleLike", blogID, username).Return(errors.New("failed to toggle like"))

// 	err := suite.usecase.ToggleLike(blogID, username)

// 	assert.NotNil(suite.T(), err)
// 	assert.Equal(suite.T(), "failed to toggle like", err.Error())
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestToggleDislike_Success() {
// 	blogID := "blog123"
// 	username := "user123"

// 	suite.mockRepo.On("ToggleDislike", blogID, username).Return(nil)

// 	err := suite.usecase.ToggleDislike(blogID, username)

// 	assert.Nil(suite.T(), err)
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func (suite *BlogUsecaseTestSuite) TestToggleDislike_Error() {
// 	blogID := "blog123"
// 	username := "user123"

// 	suite.mockRepo.On("ToggleDislike", blogID, username).Return(errors.New("failed to toggle dislike"))

// 	err := suite.usecase.ToggleDislike(blogID, username)

// 	assert.NotNil(suite.T(), err)
// 	assert.Equal(suite.T(), "failed to toggle dislike", err.Error())
// 	suite.mockRepo.AssertExpectations(suite.T())
// }

// func TestBlogUsecaseTestSuite(t *testing.T) {
// 	suite.Run(t, new(BlogUsecaseTestSuite))
// }

package tests

import (
	"blog_api/delivery/env"
	"blog_api/domain"
	"blog_api/domain/dtos"
	"blog_api/repository"
	initdb "blog_api/infrastructure/db"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Define the test suite
type BlogRepositorySuite struct {
	suite.Suite
	repository domain.BlogRepositoryInterface
	database   *mongo.Database
	cleanup    func()
}

// SetupSuite initializes the test database and repository
func (suite *BlogRepositorySuite) SetupSuite() {
	err := env.LoadEnvironmentVariables("../.env")
	suite.Require().NoError(err)

	// connect to mongodb
	client, err := initdb.ConnectDB(env.ENV.DB_ADDRESS, env.ENV.TEST_DB_NAME)
	suite.Require().NoError(err)

	db := client.Database(env.ENV.TEST_DB_NAME)
	suite.database = db

	repository := repository.NewBlogRepository(db.Collection(domain.CollectionBlogs))
	suite.repository = repository

	suite.cleanup = func() {
		db.Collection(domain.CollectionBlogs).Drop(context.TODO())
	}
}

// cleans up the database before each test
func (suite *BlogRepositorySuite) SetupTest() {
	if suite.cleanup != nil {
		suite.cleanup()
	}
}

// TestFetchBlogPostByID tests the FetchBlogPostByID method
func (suite *BlogRepositorySuite) TestFetchBlogPostByIDSuccess() {
	// Prepare a test blog post
	blogID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Blog",
		Content:   "This is a test blog.",
		Username:  "testuser",
		Tags:      []string{"go", "mongodb"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 10,
	}

	// Insert the test blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Success scenario with incrementView = true
	fetchedBlog, err := suite.repository.FetchBlogPostByID(context.TODO(), blogID.Hex(), true)
	suite.Require().NoError(err)
	suite.NotNil(fetchedBlog)
	suite.Equal(blog.Title, fetchedBlog.Title)
	suite.Equal(blog.ViewCount+1, fetchedBlog.ViewCount) // ViewCount should be incremented by 1
}

func (suite *BlogRepositorySuite) TestFetchBlogPostByIDNoIncrement() {
	// Prepare a test blog post
	blogID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Blog",
		Content:   "This is a test blog.",
		Username:  "testuser",
		Tags:      []string{"go", "mongodb"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 10,
	}

	// Insert the test blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Success scenario with incrementView = false
	fetchedBlog, err := suite.repository.FetchBlogPostByID(context.TODO(), blogID.Hex(), false)
	suite.Require().NoError(err)
	suite.NotNil(fetchedBlog)
	suite.Equal(blog.Title, fetchedBlog.Title)
	suite.Equal(blog.ViewCount, fetchedBlog.ViewCount) // ViewCount should not be incremented
}

func (suite *BlogRepositorySuite) TestFetchBlogPostByIDInvalidID() {
	// Failure scenario with invalid ID
	invalidID := "invalidID"
	fetchedBlog, err := suite.repository.FetchBlogPostByID(context.TODO(), invalidID, true)
	suite.Nil(fetchedBlog)
	suite.Equal(domain.ERR_BAD_REQUEST, err.GetCode())
	suite.Equal("Invalid blog ID", err.Error())
}

func (suite *BlogRepositorySuite) TestFetchBlogPostByIDNotFound() {
	// Failure scenario with a non-existent blog ID
	nonExistentID := primitive.NewObjectID().Hex()
	fetchedBlog, err := suite.repository.FetchBlogPostByID(context.TODO(), nonExistentID, true)
	suite.Nil(fetchedBlog)
	suite.Equal(domain.ERR_NOT_FOUND, err.GetCode())
	suite.Equal("Blog post not found", err.Error())
}

// TestDeleteBlogPost tests the DeleteBlogPost method
func (suite *BlogRepositorySuite) TestDeleteBlogPostSuccess() {
	blogID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Blog",
		Content:   "This is a test blog.",
		Username:  "testuser",
		Tags:      []string{"go", "mongodb"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 10,
	}

	// Insert a test blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	err = suite.repository.DeleteBlogPost(context.TODO(), blogID.Hex())
	suite.Require().NoError(err)

	var result dtos.BlogDTO
	err = suite.database.Collection(domain.CollectionBlogs).FindOne(context.TODO(), bson.M{"_id": blogID}).Decode(&result)
	suite.Equal(mongo.ErrNoDocuments, err)
}

func (suite *BlogRepositorySuite) TestDeleteBlogPostInvalidID() {
	invalidID := "invalidID"
	err := suite.repository.DeleteBlogPost(context.TODO(), invalidID)
	suite.Equal(domain.ERR_BAD_REQUEST, err.GetCode())
	suite.Equal("Invalid blog ID", err.Error())
}

// TestInsertBlogPost tests the InsertBlogPost method
func (suite *BlogRepositorySuite) TestInsertBlogPostSuccess() {
	blog := &domain.Blog{
		Title:    "New Blog",
		Content:  "This is a new blog.",
		Username: "testuser",
		Tags:     []string{"go", "mongodb"},
	}

	// Success scenario: Insert new blog post
	err := suite.repository.InsertBlogPost(context.TODO(), blog)
	suite.Require().NoError(err)

	// Verify that the blog post was inserted
	var result dtos.BlogDTO
	dbErr := suite.database.Collection(domain.CollectionBlogs).FindOne(context.TODO(), bson.M{"title": blog.Title}).Decode(&result)
	suite.Require().NoError(dbErr)
	suite.Equal(blog.Title, result.Title)
}

func (suite *BlogRepositorySuite) TestUpdateBlogPostSuccess() {
	// Prepare a blog post
	blogID := primitive.NewObjectID()
	originalBlog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Old Title",
		Content:   "Old content.",
		Username:  "testuser",
		Tags:      []string{"old"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 5,
	}

	// Insert the original blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), originalBlog)
	suite.Require().NoError(err)

	// Define the update data
	updateData := &domain.NewBlog{
		Title:   "New Title",
		Content: "Updated content.",
		Tags:    []string{"new"},
	}

	// Success scenario: Update existing blog post
	err = suite.repository.UpdateBlogPost(context.TODO(), blogID.Hex(), updateData)
	suite.Require().NoError(err)

	// Verify that the blog post was updated
	var updatedBlog dtos.BlogDTO
	err = suite.database.Collection(domain.CollectionBlogs).FindOne(context.TODO(), bson.M{"_id": blogID}).Decode(&updatedBlog)
	suite.Require().NoError(err)
	suite.Equal(updateData.Title, updatedBlog.Title)
	suite.Equal(updateData.Content, updatedBlog.Content)
	suite.ElementsMatch(updateData.Tags, updatedBlog.Tags)
}

func (suite *BlogRepositorySuite) TestUpdateBlogPostInvalidID() {
	invalidID := "invalidID"
	updateData := &domain.NewBlog{
		Title:   "New Title",
		Content: "Updated content.",
		Tags:    []string{"new"},
	}
	err := suite.repository.UpdateBlogPost(context.TODO(), invalidID, updateData)
	suite.Equal(domain.ERR_BAD_REQUEST, err.GetCode())
	suite.Equal("Invalid blog ID", err.Error())
}

// TestFetchBlogPosts tests the FetchBlogPosts method with various filters
func (suite *BlogRepositorySuite) TestFetchBlogPostsByTitle() {
	blog := dtos.BlogDTO{
		ID:        primitive.NewObjectID(),
		Title:     "Tech Blog",
		Content:   "Tech content",
		Username:  "user1",
		Tags:      []string{"tech"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 50,
		LikedBy: []string{"user2"},
		DislikedBy: []string{"user3"},
	}

	// Insert a test blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	filter := domain.BlogFilterOptions{Title: "Tech Blog"}
	blogs,total, err := suite.repository.FetchBlogPosts(context.TODO(), filter)
	suite.Require().NoError(err)
	suite.Equal(total, 1)
	suite.Equal(blog.Title, blogs[0].Title)
}

func (suite *BlogRepositorySuite) TestFetchBlogPostsByTag() {
	blog := dtos.BlogDTO{
		ID:        primitive.NewObjectID(),
		Title:     "Travel Blog",
		Content:   "Travel content",
		Username:  "user3",
		Tags:      []string{"travel"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 20,
		LikedBy: []string{"user2"},
		DislikedBy: []string{"user3"},
	}

	// Insert a test blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	filter := domain.BlogFilterOptions{Tags: []string{"travel"}}
	blogs,total, err := suite.repository.FetchBlogPosts(context.TODO(), filter)
	suite.Require().NoError(err)
	suite.Equal(total, 1)
	suite.Equal(blog.Title, blogs[0].Title)
}

func (suite *BlogRepositorySuite) TestFetchBlogPostsByMultipleFilters() {
	blog := dtos.BlogDTO{
		ID:        primitive.NewObjectID(),
		Title:     "Tech Blog",
		Content:   "Tech content",
		Username:  "user1",
		Tags:      []string{"tech"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 50,
		LikedBy: []string{"user2"},
		DislikedBy: []string{"user3"},
	}

	// Insert a test blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	filter := domain.BlogFilterOptions{
		Tags:        []string{"tech"},
		MinViewCount: 30,
	}
	blogs,total, err := suite.repository.FetchBlogPosts(context.TODO(), filter)
	suite.Require().NoError(err)
	suite.Equal(total, 1)
	suite.Equal(blog.Title, blogs[0].Title)
}

func (suite *BlogRepositorySuite) TestFetchBlogPostsNoResults() {
	filter := domain.BlogFilterOptions{Title: "Non-existent Blog"}
	_,total, err := suite.repository.FetchBlogPosts(context.TODO(), filter)
	suite.Require().NoError(err)
	suite.Equal(total, 0)
}

// Track blog popularity test
func (suite *BlogRepositorySuite) TestTrackBlogPopularitySuccessLike() {
	// Prepare a blog post
	blogID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Title",
		Content:   "Test content.",
		Username:  "testuser",
		Tags:      []string{"test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 5,
		LikedBy: []string{"user2"},
		DislikedBy: []string{"user3"},
	}

	// Insert the blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Define parameters for liking the blog
	action := "like"
	state := true
	username := "testuser"

	// Success scenario: Like the blog
	err = suite.repository.TrackBlogPopularity(context.TODO(), blogID.Hex(), action, state, username)
	suite.Require().NoError(err)

	// Verify that the blog post was updated
	var updatedBlog dtos.BlogDTO
	err = suite.database.Collection(domain.CollectionBlogs).FindOne(context.TODO(), bson.M{"_id": blogID}).Decode(&updatedBlog)
	suite.Require().NoError(err)
	suite.Contains(updatedBlog.LikedBy, username)
	suite.NotContains(updatedBlog.DislikedBy, username)
}

func (suite *BlogRepositorySuite) TestTrackBlogPopularitySuccessDislike() {
	// Prepare a blog post with a user who has already liked it
	blogID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Title",
		Content:   "Test content.",
		Username:  "testuser",
		Tags:      []string{"test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 5,
		LikedBy:   []string{"testuser"},
		DislikedBy: []string{},
	}

	// Insert the blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Define parameters for disliking the blog
	action := "dislike"
	state := true
	username := "testuser"

	// Success scenario: Dislike the blog
	err = suite.repository.TrackBlogPopularity(context.TODO(), blogID.Hex(), action, state, username)
	suite.Require().NoError(err)

	// Verify that the blog post was updated
	var updatedBlog dtos.BlogDTO
	err = suite.database.Collection(domain.CollectionBlogs).FindOne(context.TODO(), bson.M{"_id": blogID}).Decode(&updatedBlog)
	suite.Require().NoError(err)
	suite.Contains(updatedBlog.DislikedBy, username)
	suite.NotContains(updatedBlog.LikedBy, username)
}

func (suite *BlogRepositorySuite) TestTrackBlogPopularityInvalidBlogID() {
	// Failure scenario: Invalid blog ID
	action := "like"
	state := true
	username := "testuser"
	invalidID := "invalidID"

	err := suite.repository.TrackBlogPopularity(context.TODO(), invalidID, action, state, username)
	suite.Require().Error(err)
	suite.Equal(domain.ERR_BAD_REQUEST, err.GetCode())
}

func (suite *BlogRepositorySuite) TestTrackBlogPopularityInvalidAction() {
	// Prepare a blog post
	blogID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Title",
		Content:   "Test content.",
		Username:  "testuser",
		Tags:      []string{"test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 5,
	}

	// Insert the blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Failure scenario: Invalid action
	action := "invalid_action"
	state := true
	username := "testuser"

	newErr := suite.repository.TrackBlogPopularity(context.TODO(), blogID.Hex(), action, state, username)
	suite.Require().Error(newErr)
	suite.Equal(domain.ERR_BAD_REQUEST, newErr.GetCode())
}


// Comment tests
// create a comment tests
func (suite *BlogRepositorySuite) TestCreateCommentSuccess() {
	// Prepare a blog post
	blogID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Title",
		Content:   "Test content.",
		Username:  "testuser",
		Tags:      []string{"test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 5,
	}

	// Insert the blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Define parameters for the comment
	comment := &domain.Comment{
		Content: "This is a test comment.",
	}
	createdBy := "commenter"

	// Success scenario: Add a comment
	err = suite.repository.CreateComment(context.TODO(), comment, blogID.Hex(), createdBy)
	suite.Require().NoError(err)

	// Verify that the comment was added
	var updatedBlog dtos.BlogDTO
	err = suite.database.Collection(domain.CollectionBlogs).FindOne(context.TODO(), bson.M{"_id": blogID}).Decode(&updatedBlog)
	suite.Require().NoError(err)
	suite.Len(updatedBlog.Comments, 1)
	suite.Equal("This is a test comment.", updatedBlog.Comments[0].Content)
	suite.Equal(createdBy, updatedBlog.Comments[0].Username)
}

func (suite *BlogRepositorySuite) TestCreateCommentInvalidBlogID() {
	// Failure scenario: Invalid blog ID
	comment := &domain.Comment{
		Content: "This is a test comment.",
	}
	createdBy := "commenter"
	invalidID := "invalidID"

	err := suite.repository.CreateComment(context.TODO(), comment, invalidID, createdBy)
	suite.Require().Error(err)
	suite.Equal(domain.ERR_BAD_REQUEST, err.GetCode())
}

func (suite *BlogRepositorySuite) TestCreateCommentBlogNotFound() {
	// Failure scenario: Blog not found
	nonexistentID := primitive.NewObjectID().Hex()
	comment := &domain.Comment{
		Content: "This is a test comment.",
	}
	createdBy := "commenter"

	err := suite.repository.CreateComment(context.TODO(), comment, nonexistentID, createdBy)
	suite.Require().Error(err)
	suite.Equal(domain.ERR_NOT_FOUND, err.GetCode())
}

// Delete a comment tests
func (suite *BlogRepositorySuite) TestDeleteCommentSuccess() {
	// Prepare a blog post with a comment
	blogID := primitive.NewObjectID()
	commentID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Title",
		Content:   "Test content.",
		Username:  "testuser",
		Tags:      []string{"test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 5,
		Comments: []dtos.CommentDTO{
			{
				ID:        commentID,
				Content:   "This is a test comment.",
				Username:  "commenter",
				CreatedAt: time.Now(),
			},
		},
	}

	// Insert the blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Success scenario: Delete a comment
	err = suite.repository.DeleteComment(context.TODO(), commentID.Hex(), blogID.Hex(), "commenter")
	suite.Require().NoError(err)

	// Verify that the comment was deleted
	var updatedBlog dtos.BlogDTO
	err = suite.database.Collection(domain.CollectionBlogs).FindOne(context.TODO(), bson.M{"_id": blogID}).Decode(&updatedBlog)
	suite.Require().NoError(err)
	suite.Len(updatedBlog.Comments, 0)
}

func (suite *BlogRepositorySuite) TestDeleteCommentInvalidBlogID() {
	// Failure scenario: Invalid blog ID
	invalidID := "invalidID"
	commentID := primitive.NewObjectID().Hex()
	userName := "commenter"

	err := suite.repository.DeleteComment(context.TODO(), commentID, invalidID, userName)
	suite.Require().Error(err)
	suite.Equal(domain.ERR_BAD_REQUEST, err.GetCode())
}

func (suite *BlogRepositorySuite) TestDeleteCommentBlogNotFound() {
	// Failure scenario: Blog not found
	nonexistentID := primitive.NewObjectID().Hex()
	commentID := primitive.NewObjectID().Hex()
	userName := "commenter"

	err := suite.repository.DeleteComment(context.TODO(), commentID, nonexistentID, userName)
	suite.Require().Error(err)
	suite.Equal(domain.ERR_NOT_FOUND, err.GetCode())
}

func (suite *BlogRepositorySuite) TestDeleteCommentUnauthorized() {
	// Prepare a blog post with a comment
	blogID := primitive.NewObjectID()
	commentID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Title",
		Content:   "Test content.",
		Username:  "testuser",
		Tags:      []string{"test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 5,
		Comments: []dtos.CommentDTO{
			{
				ID:        commentID,
				Content:   "This is a test comment.",
				Username:  "commenter",
				CreatedAt: time.Now(),
			},
		},
	}

	// Insert the blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Failure scenario: Unauthorized user
	newErr := suite.repository.DeleteComment(context.TODO(), commentID.Hex(), blogID.Hex(), "anotheruser")
	suite.Require().Error(newErr)
	suite.Equal(domain.ERR_FORBIDDEN, newErr.GetCode())
}

func (suite *BlogRepositorySuite) TestDeleteCommentCommentNotFound() {
	// Prepare a blog post with a comment
	blogID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Title",
		Content:   "Test content.",
		Username:  "testuser",
		Tags:      []string{"test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 5,
		Comments: []dtos.CommentDTO{
			{
				ID:        primitive.NewObjectID(),
				Content:   "This is a test comment.",
				Username:  "commenter",
				CreatedAt: time.Now(),
			},
		},
	}

	// Insert the blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Failure scenario: Comment not found
	nonexistentCommentID := primitive.NewObjectID().Hex()
	newErr := suite.repository.DeleteComment(context.TODO(), nonexistentCommentID, blogID.Hex(), "commenter")
	suite.Require().Error(newErr)
	suite.Equal(domain.ERR_NOT_FOUND, newErr.GetCode())
}

func (suite *BlogRepositorySuite) TestUpdateCommentSuccess() {
	// Prepare a blog post with a comment
	blogID := primitive.NewObjectID()
	commentID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Title",
		Content:   "Test content.",
		Username:  "testuser",
		Tags:      []string{"test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 5,
		Comments: []dtos.CommentDTO{
			{
				ID:        commentID,
				Content:   "Old comment content.",
				Username:  "commenter",
				CreatedAt: time.Now(),
			},
		},
	}

	// Insert the blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Success scenario: Update the comment
	updateComment := &domain.NewComment{
		Content: "Updated comment content.",
	}
	err = suite.repository.UpdateComment(context.TODO(), updateComment, commentID.Hex(), blogID.Hex(), "commenter")
	suite.Require().NoError(err)

	// Verify that the comment was updated
	var updatedBlog dtos.BlogDTO
	err = suite.database.Collection(domain.CollectionBlogs).FindOne(context.TODO(), bson.M{"_id": blogID}).Decode(&updatedBlog)
	suite.Require().NoError(err)
	suite.Len(updatedBlog.Comments, 1)
	suite.Equal("Updated comment content.", updatedBlog.Comments[0].Content)
}

func (suite *BlogRepositorySuite) TestUpdateCommentInvalidBlogID() {
	// Failure scenario: Invalid blog ID
	invalidID := "invalidID"
	commentID := primitive.NewObjectID().Hex()
	updateComment := &domain.NewComment{Content: "Some content"}
	userName := "commenter"

	err := suite.repository.UpdateComment(context.TODO(), updateComment, commentID, invalidID, userName)
	suite.Require().Error(err)
	suite.Equal(domain.ERR_BAD_REQUEST, err.GetCode())
}

func (suite *BlogRepositorySuite) TestUpdateCommentBlogNotFound() {
	// Failure scenario: Blog not found
	nonexistentID := primitive.NewObjectID().Hex()
	commentID := primitive.NewObjectID().Hex()
	updateComment := &domain.NewComment{Content: "Some content"}
	userName := "commenter"

	err := suite.repository.UpdateComment(context.TODO(), updateComment, commentID, nonexistentID, userName)
	suite.Require().Error(err)
	suite.Equal(domain.ERR_NOT_FOUND, err.GetCode())
}

func (suite *BlogRepositorySuite) TestUpdateCommentUnauthorized() {
	// Prepare a blog post with a comment
	blogID := primitive.NewObjectID()
	commentID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Title",
		Content:   "Test content.",
		Username:  "testuser",
		Tags:      []string{"test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 5,
		Comments: []dtos.CommentDTO{
			{
				ID:        commentID,
				Content:   "This is a test comment.",
				Username:  "commenter",
				CreatedAt: time.Now(),
			},
		},
	}

	// Insert the blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Failure scenario: Unauthorized user
	updateComment := &domain.NewComment{Content: "Updated comment content."}
	newErr := suite.repository.UpdateComment(context.TODO(), updateComment, commentID.Hex(), blogID.Hex(), "anotheruser")
	suite.Require().Error(newErr)
	suite.Equal(domain.ERR_FORBIDDEN, newErr.GetCode())
}

func (suite *BlogRepositorySuite) TestUpdateCommentCommentNotFound() {
	// Prepare a blog post with a comment
	blogID := primitive.NewObjectID()
	blog := dtos.BlogDTO{
		ID:        blogID,
		Title:     "Test Title",
		Content:   "Test content.",
		Username:  "testuser",
		Tags:      []string{"test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ViewCount: 5,
		Comments: []dtos.CommentDTO{
			{
				ID:        primitive.NewObjectID(),
				Content:   "This is a test comment.",
				Username:  "commenter",
				CreatedAt: time.Now(),
			},
		},
	}

	// Insert the blog post
	_, err := suite.database.Collection(domain.CollectionBlogs).InsertOne(context.TODO(), blog)
	suite.Require().NoError(err)

	// Failure scenario: Comment not found
	nonexistentCommentID := primitive.NewObjectID().Hex()
	updateComment := &domain.NewComment{Content: "Updated comment content."}
	newErr := suite.repository.UpdateComment(context.TODO(), updateComment, nonexistentCommentID, blogID.Hex(), "commenter")
	suite.Require().Error(newErr)
	suite.Equal(domain.ERR_NOT_FOUND, newErr.GetCode())
}

// Run the tests
func TestBlogRepositorySuite(t *testing.T) {
	suite.Run(t, new(BlogRepositorySuite))
}

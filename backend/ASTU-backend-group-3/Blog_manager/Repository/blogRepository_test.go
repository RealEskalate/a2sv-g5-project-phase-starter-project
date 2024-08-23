package Repository

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"context"
	"strconv"

	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	// "go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepositoryTestSuite struct {
	suite.Suite
	client     *mongo.Client
	collection *mongo.Collection
	repo       *blogRepository
}

func (suite *BlogRepositoryTestSuite) SetupSuite() {
	// Connect to the MongoDB database
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	suite.NoError(err)

	// Verify connection
	err = client.Ping(context.TODO(), nil)
	suite.NoError(err)

	suite.client = client
	suite.collection = client.Database("testdb").Collection("blogs")
	suite.repo = NewBlogRepository(suite.collection)
}

func (suite *BlogRepositoryTestSuite) TearDownSuite() {
	// Disconnect from MongoDB
	err := suite.client.Disconnect(context.TODO())
	suite.NoError(err)
}

func (suite *BlogRepositoryTestSuite) TearDownTest() {
	// Clear the collection after each test
	err := suite.collection.Drop(context.TODO())
	suite.NoError(err)
}

func (suite *BlogRepositoryTestSuite) TestSave() {
	blog := &Domain.Blog{
		Title:     "New Blog",
		Author:    "Author Name",
		Tags:      []string{"tag1", "tag2"},
		CreatedAt: time.Now().Format(time.RFC3339), // Converts time.Time to string
	}

	savedBlog, err := suite.repo.Save(blog)
	suite.NoError(err)
	suite.NotEmpty(savedBlog.Id) // Ensure ID was generated
}

func (suite *BlogRepositoryTestSuite) TestFindByID() {
	// First save a blog
	blog := &Domain.Blog{
		Title:     "Find Me",
		Author:    "Author Name",
		Tags:      []string{"tag1"},
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	savedBlog, err := suite.repo.Save(blog)
	suite.NoError(err)

	// Now try to find it by ID
	foundBlog, err := suite.repo.FindByID(savedBlog.Id)
	suite.NoError(err)
	suite.Equal(savedBlog.Title, foundBlog.Title)
	suite.Equal(savedBlog.Author, foundBlog.Author)
}

func (suite *BlogRepositoryTestSuite) TestDeleteBlogByID() {
	// First save a blog
	blog := &Domain.Blog{
		Title:     "To Be Deleted",
		Author:    "Author Name",
		Tags:      []string{"tag1"},
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	savedBlog, err := suite.repo.Save(blog)
	suite.NoError(err)

	// Now delete it by ID
	err = suite.repo.DeleteBlogByID(savedBlog.Id)
	suite.NoError(err)

	// Try to find it again
	foundBlog, err := suite.repo.FindByID(savedBlog.Id)
	suite.NoError(err)
	suite.Nil(foundBlog) // Should be nil since it's deleted
}

func (suite *BlogRepositoryTestSuite) TestRetrieveBlogs() {
	for i := 1; i <= 5; i++ {
		blog := &Domain.Blog{
			Title:     "Blog " + strconv.Itoa(i),
			Author:    "Author " + strconv.Itoa(i),
			Tags:      []string{"tag1", "tag" + strconv.Itoa(i)},
			CreatedAt: time.Now().Format(time.RFC3339),
		}
		_, err := suite.repo.Save(blog)
		suite.NoError(err)
	}

	blogs, total, err := suite.repo.RetrieveBlogs(1, 2, "created_at")
	suite.NoError(err)
	suite.Equal(2, len(blogs))
	suite.Equal(int64(5), total)
}

func (suite *BlogRepositoryTestSuite) TestSearchBlogs() {
	blog := &Domain.Blog{
		Title:     "Unique Blog",
		Author:    "Special Author",
		Tags:      []string{"unique"},
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	_, err := suite.repo.Save(blog)
	suite.NoError(err)

	blogs, err := suite.repo.SearchBlogs("Unique", "", []string{"unique"})
	suite.NoError(err)
	suite.Equal(1, len(blogs))
	suite.Equal("Unique Blog", blogs[0].Title)
}

func (suite *BlogRepositoryTestSuite) TestIncrementViewCount() {
	blog := &Domain.Blog{
		Title:     "View Blog",
		Author:    "Viewer",
		CreatedAt: time.Now().Format(time.RFC3339),
		ViewCount: 0,
	}
	_, err := suite.repo.Save(blog)
	suite.NoError(err)

	err = suite.repo.IncrementViewCount(blog.Id)
	suite.NoError(err)

	updatedBlog, err := suite.repo.FindByID(blog.Id)
	suite.NoError(err)
	suite.Equal(int64(1), updatedBlog.ViewCount)
}

func (suite *BlogRepositoryTestSuite) TestAddComment() {
	// Create and save a new blog entry
	blog := &Domain.Blog{
		Title:     "Comment Blog",
		Author:    "Commenter",
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	_, err := suite.repo.Save(blog)
	suite.NoError(err)

	// Create a new comment
	comment := Domain.Comment{
		Id:        primitive.NewObjectID().Hex(), // Ensure Id is set
		Content:   "This is a comment",
		PostID:    blog.Id,
		UserID:    "user1",
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	// Add the comment to the blog
	err = suite.repo.AddComment(blog.Id, comment)
	suite.NoError(err)

	// Retrieve the blog to check if the comment was added
	updatedBlog, err := suite.repo.FindByID(blog.Id)
	suite.NoError(err)
	suite.Equal(1, len(updatedBlog.Comments))
	suite.Equal("This is a comment", updatedBlog.Comments[0].Content)
}

func (suite *BlogRepositoryTestSuite) TestToggleLikeAndDislike() {
	blog := &Domain.Blog{
		Title:     "Like Blog",
		Author:    "Liker",
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	_, err := suite.repo.Save(blog)
	suite.NoError(err)

	// Toggle Like
	err = suite.repo.ToggleLike(blog.Id, "user1")
	suite.NoError(err)

	updatedBlog, err := suite.repo.FindByID(blog.Id)
	suite.NoError(err)
	suite.Contains(updatedBlog.Likes, "user1")

	// Toggle Dislike
	err = suite.repo.ToggleDislike(blog.Id, "user1")
	suite.NoError(err)

	updatedBlog, err = suite.repo.FindByID(blog.Id)
	suite.NoError(err)
	suite.Contains(updatedBlog.Dislikes, "user1")
	suite.NotContains(updatedBlog.Likes, "user1")
}

func TestBlogRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BlogRepositoryTestSuite))
}

package test

import (
	domain "AAiT-backend-group-8/Domain"
	mongodb "AAiT-backend-group-8/Infrastructure/mongodb"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestBlogSuite struct {
	suite.Suite
	blogRepo mongodb.BlogRepository
}

var blogs []domain.Blog

func (t *TestBlogSuite) SetUpTest() {
	client := mongodb.InitMongoDB()
	testRepo := mongodb.CreateCollection(client, "unit-tests", "blog")
	t.blogRepo = *mongodb.NewBlogRepository(
		testRepo,
	)
}
func (t *TestBlogSuite) TearDownTest() {
	err := t.blogRepo.DropDB()
	if err != nil {
		t.T().Fatal("Failed to clear the database after the test:", err)
	}
}

func (suite *TestBlogSuite) TestCreateBlogPositive() {
	assert := assert.New(suite.T())

	temp := &domain.Blog{
		Title: "test title",
		Body:  "test test test test test test test test testtest test testtest test test",
		Tags:  []string{"test1", "Test2", "TEST 3"},
	}
	err := suite.blogRepo.Create(temp)
	assert.Nil(err, "error should be nil ")
}

func (suite *TestBlogSuite) TestFindAllPositive() {
	assert := assert.New(suite.T())

	// Seed the database with some test blogs
	blog1 := &domain.Blog{
		Title: "test title 1",
		Body:  "test body 1",
		Tags:  []string{"test1", "Test2"},
	}
	blog2 := &domain.Blog{
		Title: "test title 2",
		Body:  "test body 2",
		Tags:  []string{"test3", "Test4"},
	}

	err := suite.blogRepo.Create(blog1)
	assert.Nil(err, "error should be nil when creating blog1")

	err = suite.blogRepo.Create(blog2)
	assert.Nil(err, "error should be nil when creating blog2")

	local, err := suite.blogRepo.FindAll(1, 10, "created_at")
	assert.Nil(err, "error should be nil when finding all blogs")
	assert.GreaterOrEqual(len(local), 2, "there should be 2 blogs found")
	blogs = local

}

func (suite *TestBlogSuite) TestFindAllNegative() {
	assert := assert.New(suite.T())

	// Try to retrieve blogs from an empty collection
	blogs, err := suite.blogRepo.FindAll(1, 10, "created_at")
	assert.Nil(err, "error should be nil when finding all blogs")
	assert.Equal(0, len(blogs), "there should be 0 blogs found")
}

func (suite *TestBlogSuite) TestFindByIDPositive() {
	assert := assert.New(suite.T())
	temp := blogs[0]
	stringID := temp.Id.Hex()
	fmt.Println(stringID)
	got, err := suite.blogRepo.FindByID(stringID)
	gotder := *got
	assert.Equal(gotder, temp, "they should be equal")
	assert.Nil(err, "Error should be nil ")
}
func (suite *TestBlogSuite) TestFindByIDNegative() {
	assert := assert.New(suite.T())

	// Try to find a blog with a non-existing ID
	_, err := suite.blogRepo.FindByID("60d5f2c3c2b7f088c8e6b85a")
	assert.NotNil(err, "error should not be nil when finding blog by a non-existing ID")
}

func (suite *TestBlogSuite) TestUpdatePositive() {
	assert := assert.New(suite.T())

	temp := blogs[0]
	err := suite.blogRepo.Create(&temp)
	assert.Nil(err, "error should be nil when creating blog")

	// Update the blog
	temp.Title = "updated title"
	err = suite.blogRepo.Update(&temp)
	assert.Nil(err, "error should be nil when updating blog")

	updatedBlog, err := suite.blogRepo.FindByID(temp.Id.Hex())
	assert.Nil(err, "error should be nil when finding updated blog")
	assert.Equal("updated title", updatedBlog.Title, "the title should be updated")
}

func (suite *TestBlogSuite) TestUpdateNegative() {
	assert := assert.New(suite.T())

	// Try to update a blog with a non-existing ID
	temp := &domain.Blog{
		Title: "non-existing title",
		Body:  "non-existing body",
	}
	err := suite.blogRepo.Update(temp)
	assert.NotNil(err, "error should not be nil when updating a non-existing blog")
}
func (suite *TestBlogSuite) TestDeletePositive() {
	assert := assert.New(suite.T())

	temp := &domain.Blog{
		Title: "test title",
		Body:  "test body",
		Tags:  []string{"test1", "Test2"},
	}
	err := suite.blogRepo.Create(temp)
	assert.Nil(err, "error should be nil when creating blog")

	err = suite.blogRepo.Delete(temp.Id.Hex())
	assert.Nil(err, "error should be nil when deleting blog")

	_, err = suite.blogRepo.FindByID(temp.Id.Hex())
	assert.NotNil(err, "error should not be nil when finding a deleted blog")
}

func (suite *TestBlogSuite) TestDeleteNegative() {
	assert := assert.New(suite.T())

	// Try to delete a blog with a non-existing ID
	err := suite.blogRepo.Delete("60d5f2c3c2b7f088c8edds6b85a")
	assert.NotNil(err, "error should not be nil when deleting a non-existing blog")
}

func TestTaskRepositorySuite(t *testing.T) {
	suite := new(TestBlogSuite)
	suite.SetT(t)
	suite.SetUpTest()
	// suite.TestFindAllNegative()
	suite.TestCreateBlogPositive()
	suite.TestFindAllPositive()
	suite.TestFindByIDPositive()
	suite.TestFindByIDNegative()
	suite.TestDeletePositive()
	suite.TestDeleteNegative()

	suite.TearDownTest()
	suite.TestFindAllNegative()

}

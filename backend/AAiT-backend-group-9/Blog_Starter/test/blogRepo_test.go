package test

import (
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"context"
	"log"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepositorySuit struct {
	suite.Suite
	// the funcionalities we need to test
	repository domain.BlogRepository
	db         *mongo.Database
}

func (suite *BlogRepositorySuit) SetupSuite() {
	// this function runs once before all tests in the suite

	// some initialization setup
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("testdb")
	repository := repository.NewBlogRepository(db, "blogs")

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.repository = repository
	suite.db = db
}

func (suite *BlogRepositorySuit) TearDownSuite() {
	// we need to drop the table we used in the tests
	defer suite.db.Drop(context.Background())
}
func (suite *BlogRepositorySuit) TestCreateBlog() {
	userID := primitive.NewObjectID()
	blog := domain.Blog{
		UserID:  userID,
		Title:   "GERD",
		Content: "The Grand Ethiopian Renaissance Dam, formerly known as the Millennium Dam and sometimes referred to as the Hidase Dam, is a gravity dam on the Blue Nile River in Ethiopia. The dam is in the",
		Tags:    []string{"Ethiopia", "Egypt", "Sudan"},
		Author:  "John Doe",
	}

	createdBlog, err := suite.repository.CreateBlog(context.Background(), &blog)

	suite.NoError(err)
	suite.NotEmpty(createdBlog.BlogID)
	suite.Equal(blog.UserID, createdBlog.UserID)

}

func (suite *BlogRepositorySuit) TestGetBlog() {
	userID := primitive.NewObjectID()
	blog := domain.Blog{
		UserID:  userID,
		Title:   "GERD",
		Content: "The Grand Ethiopian Renaissance Dam, formerly known as the Millennium Dam and sometimes referred to as the Hidase Dam, is a gravity dam on the Blue Nile River in Ethiopia. The dam is in the",
		Tags:    []string{"Ethiopia", "Egypt", "Sudan"},
	}

	// create a blog
	createdBlog, err := suite.repository.CreateBlog(context.Background(), &blog)

	suite.NoError(err)
	Newblog, err := suite.repository.GetBlogByID(context.Background(), createdBlog.BlogID.String())
	suite.NoError(err)
	suite.Equal(createdBlog.BlogID, Newblog.BlogID)

}

func (suite *BlogRepositorySuit) TestIncrementViewCount() {
	userID := primitive.NewObjectID()
	blog := domain.Blog{
		UserID:  userID,
		Title:   "GERD",
		Content: "The Grand Ethiopian Renaissance Dam, formerly known as the Millennium Dam and sometimes referred to as the Hidase Dam, is a gravity dam on the Blue Nile River in Ethiopia. The dam is in the",
		Tags:    []string{"Ethiopia", "Egypt", "Sudan"},
	}

	// create a blog
	createdBlog, err := suite.repository.CreateBlog(context.Background(), &blog)

	suite.NoError(err)
	Newblog, err := suite.repository.IncrementViewCount(context.Background(), createdBlog.BlogID.String())
	suite.NoError(err)
	suite.Equal(createdBlog.ViewCount+1, Newblog.ViewCount)

}

func (suite *BlogRepositorySuit) TestGetAllBlog() {
	userID := primitive.NewObjectID()
	blog := domain.Blog{
		UserID:  userID,
		Title:   "GERD",
		Content: "The Grand Ethiopian Renaissance Dam, formerly known as the Millennium Dam and sometimes referred to as the Hidase Dam, is a gravity dam on the Blue Nile River in Ethiopia. The dam is in the",
		Tags:    []string{"Ethiopia", "Egypt", "Sudan"},
	}

	// create a blog
	createdBlog, err := suite.repository.CreateBlog(context.Background(), &blog)

	suite.NoError(err)
	blogs, _, err := suite.repository.GetAllBlog(context.Background(), 0, 10, "createtimestamp")
	suite.NoError(err)
	suite.NotEmpty(blogs)
	suite.Equal(1, len(blogs))
	suite.Equal(createdBlog.BlogID, blogs[0].BlogID)

}

func (suite *BlogRepositorySuit) TestUpdateBlog() {
	userID := primitive.NewObjectID()
	blog := domain.Blog{
		UserID:  userID,
		Title:   "GERD",
		Content: "The Grand Ethiopian Renaissance Dam, formerly known as the Millennium Dam and sometimes referred to as the Hidase Dam, is a gravity dam on the Blue Nile River in Ethiopia. The dam is in the",
		Tags:    []string{"Ethiopia", "Egypt", "Sudan"},
	}

	// create a blog
	// update Blog
	// compare blog

	createdBlog, err := suite.repository.CreateBlog(context.Background(), &blog)

	suite.NoError(err)

	BlogId := createdBlog.BlogID.String()

	updateBlog := domain.BlogUpdate{
		UserID:  userID.String(),
		Title:   "GERED!",
		Content: "new nsaf fjklsdjfkl",
		Tags:    []string{"Ethiopia", "Egypt", "Sudan"},
	}
	updatedBlog, err := suite.repository.UpdateBlog(context.Background(), &updateBlog, BlogId)
	suite.NoError(err)
	suite.Equal(updateBlog.Title, updatedBlog.Title)
	suite.NotEqual(blog.Title, updatedBlog.Title)
	suite.NotEqual(blog.Content, updatedBlog.Content)
	suite.NotNil(updatedBlog.UpdatedAt)

}

func (suite *BlogRepositorySuit) TestDeleteBlog() {
	userID := primitive.NewObjectID()
	blog := domain.Blog{
		UserID:  userID,
		Title:   "GERD",
		Content: "The Grand Ethiopian Renaissance Dam, formerly known as the Millennium Dam and sometimes referred to as the Hidase Dam, is a gravity dam on the Blue Nile River in Ethiopia. The dam is in the",
		Tags:    []string{"Ethiopia", "Egypt", "Sudan"},
	}

	// create a blog
	createdBlog, err := suite.repository.CreateBlog(context.Background(), &blog)

	suite.NoError(err)
	err = suite.repository.DeleteBlog(context.Background(), createdBlog.BlogID.String())
	suite.NoError(err)
	_, err = suite.repository.GetBlogByID(context.Background(), createdBlog.BlogID.String())
	suite.Error(err)

}

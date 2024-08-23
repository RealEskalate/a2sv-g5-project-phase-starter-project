package test

import (
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
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

// I want clear my database after every test run
func (suite *BlogRepositorySuit) SetupTest() {
	// this function runs before every test in the suite
	// we need to clear the table before every test
	_, err := suite.db.Collection("blogs").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
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

func (suite *BlogRepositorySuit) TestGetBlogByIDPositive() {
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
	Newblog, err := suite.repository.GetBlogByID(context.Background(), createdBlog.BlogID.Hex())
	suite.NoError(err)
	suite.Equal(createdBlog.BlogID, Newblog.BlogID)

}
func (suite *BlogRepositorySuit) TestGetBlogByIDNegative() {

	_, err := suite.repository.GetBlogByID(context.Background(), "60b1b2e0b4c4e9b6f4e3c3d3")
	suite.Error(err)
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
	Newblog, err := suite.repository.IncrementViewCount(context.Background(), createdBlog.BlogID.Hex())
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

	BlogId := createdBlog.BlogID.Hex()

	updateBlog := domain.BlogUpdate{
		UserID:  userID.Hex(),
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
	err = suite.repository.DeleteBlog(context.Background(), createdBlog.BlogID.Hex())
	suite.NoError(err)
	_, err = suite.repository.GetBlogByID(context.Background(), createdBlog.BlogID.Hex())
	suite.Error(err)

}

func (suite *BlogRepositorySuit) TestFilterBlogs() {
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
	blogRequest := domain.BlogFilterRequest{
		LikeLowerRange: 0,
		ViewLowerRange: 0,
		Date:           nil,
		Tags:           []string{"Ethiopia", "Egypt", "Sudan"},
	}
	blogs, err := suite.repository.FilterBlogs(context.Background(), &blogRequest)
	suite.NoError(err)
	suite.NotEmpty(blogs)
	suite.Equal(1, len(blogs))
	suite.Equal(createdBlog.BlogID, blogs[0].BlogID)

}

func (suite *BlogRepositorySuit) TestSearchBlogs() {
	userID := primitive.NewObjectID()
	blog := domain.Blog{
		UserID:  userID,
		Title:   "GERD",
		Content: "The Grand Ethiopian Renaissance Dam, formerly known as the Millennium Dam and sometimes referred to as the Hidase Dam, is a gravity dam on the Blue Nile River in Ethiopia. The dam is in the",
		Tags:    []string{"Ethiopia", "Egypt", "Sudan"},
		Author:  "John Doe",
	}

	// create a blog
	createdBlog, err := suite.repository.CreateBlog(context.Background(), &blog)

	suite.NoError(err)
	searchRequest := domain.BlogSearchRequest{
		Title:  "GERD",
		Author: "John Doe",
	}
	blogs, err := suite.repository.SearchBlogs(context.Background(), &searchRequest)

	suite.NoError(err)
	suite.NotEmpty(blogs)
	suite.Equal(1, len(blogs))
	suite.Equal(createdBlog.BlogID, blogs[0].BlogID)

}

func (suite *BlogRepositorySuit) TestInsertRating() {
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
	rating := domain.BlogRating{
		BlogID: createdBlog.BlogID.Hex(),
		Rating: 5,
	}
	err = suite.repository.InsertRating(context.Background(), &rating)
	suite.NoError(err)
	Newblog, err := suite.repository.GetBlogByID(context.Background(), createdBlog.BlogID.Hex())
	suite.NoError(err)
	suite.Equal(float64(5), Newblog.AverageRating)

}

func (suite *BlogRepositorySuit) TestUpdateRating() {
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
	rating := domain.BlogRating{
		BlogID: createdBlog.BlogID.Hex(),
		Rating: 5,
	}
	err = suite.repository.InsertRating(context.Background(), &rating)
	suite.NoError(err)
	rating.Rating = 4
	err = suite.repository.UpdateRating(context.Background(), &rating, 5)
	suite.NoError(err)
	Newblog, err := suite.repository.GetBlogByID(context.Background(), createdBlog.BlogID.Hex())
	suite.NoError(err)
	suite.Equal(float64(4), Newblog.AverageRating)

}

func (suite *BlogRepositorySuit) TestDeleteRating() {
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
	rating := domain.BlogRating{
		BlogID: createdBlog.BlogID.Hex(),
		Rating: 5,
	}
	err = suite.repository.InsertRating(context.Background(), &rating)
	suite.NoError(err)

	err = suite.repository.DeleteRating(context.Background(), &rating)

	suite.NoError(err)
	Newblog, err := suite.repository.GetBlogByID(context.Background(), createdBlog.BlogID.Hex())
	suite.NoError(err)
	suite.Equal(float64(0), Newblog.AverageRating)

}

func (suite *BlogRepositorySuit) TestUpdateCommentCount() {
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
	err = suite.repository.UpdateCommentCount(context.Background(), createdBlog.BlogID.Hex(), true)
	suite.NoError(err)
	Newblog, err := suite.repository.GetBlogByID(context.Background(), createdBlog.BlogID.Hex())
	suite.NoError(err)
	suite.Equal(1, Newblog.CommentCount)

}

func (suite *BlogRepositorySuit) TestUpdateLikeCount() {
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
	err = suite.repository.UpdateLikeCount(context.Background(), createdBlog.BlogID.Hex(), true)
	suite.NoError(err)
	Newblog, err := suite.repository.GetBlogByID(context.Background(), createdBlog.BlogID.Hex())
	suite.NoError(err)
	suite.Equal(1, Newblog.LikeCount)

}

// we need this to run all tests in our suite
func Test_blogRepositorySuite(t *testing.T) {
	/// we still need this to run all tests in our suite
	suite.Run(t, &BlogRepositorySuit{})
}

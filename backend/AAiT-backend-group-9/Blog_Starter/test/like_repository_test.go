package test

import (
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogLikeRepositorySuit struct {
	suite.Suite
	// the funcionalities we need to test
	repository domain.LikeRepository
	db         *mongo.Database
}

func (suite *BlogLikeRepositorySuit) SetupSuite() {
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
	repository := repository.NewLikeRepository(db, "like")

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.repository = repository
	suite.db = db
}

func (suite *BlogLikeRepositorySuit) TearDownSuite() {
	// we need to drop the table we used in the tests
	defer suite.db.Drop(context.Background())
}

func (suite *BlogLikeRepositorySuit) SetupTest() {
	// this function runs before every test in the suite
	// we need to clear the table before every test
	_, err := suite.db.Collection("like").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *BlogLikeRepositorySuit) TestLikeBlog() {
	// testing the like blog functionality
	like := &domain.Like{
		UserID: "123",
		BlogID: "456",
	}

	like, err := suite.repository.LikeBlog(context.Background(), like)
	suite.NoError(err)
	suite.NotNil(like)
}

func (suite *BlogLikeRepositorySuit) TestGetByID() {
	// testing the get by id functionality
	like := &domain.Like{
		UserID: "123",
		BlogID: "456",
	}

	like, err := suite.repository.LikeBlog(context.Background(), like)
	suite.NoError(err)
	suite.NotNil(like)

	like, err = suite.repository.GetByID(context.Background(), like.UserID, like.BlogID)
	suite.NoError(err)
	suite.NotNil(like)
}

func (suite *BlogLikeRepositorySuit) TestUnlikeBlog() {
	// testing the unlike blog functionality
	like := &domain.Like{
		UserID: "123",
		BlogID: "456",
	}

	like, err := suite.repository.LikeBlog(context.Background(), like)
	suite.NoError(err)
	suite.NotNil(like)

	like, err = suite.repository.UnlikeBlog(context.Background(), like.LikeID)
	suite.NoError(err)
	suite.NotNil(like)
}

func Test_BlogLikeRepositorySuit(t *testing.T) {
	// this is what actually runs our suite
	suite.Run(t, &BlogLikeRepositorySuit{})
}

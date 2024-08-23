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

type BlogRatingRepositorySuit struct {
	suite.Suite
	// the funcionalities we need to test
	repository domain.BlogRatingRepository
	db         *mongo.Database
}

func (suite *BlogRatingRepositorySuit) SetupSuite() {
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
	repository := repository.NewBlogRatingRepository(db, "rating", nil)

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.repository = repository
	suite.db = db
}

func (suite *BlogRatingRepositorySuit) TearDownSuite() {
	// we need to drop the table we used in the tests
	defer suite.db.Drop(context.Background())
}

func (suite *BlogRatingRepositorySuit) SetupTest() {
	// this function runs before every test in the suite
	// we need to clear the table before every test
	_, err := suite.db.Collection("rating").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *BlogRatingRepositorySuit) TestGetRatingByBlogID() {
	// we need to insert some data first
	rating := domain.BlogRating{
		RatingID: primitive.NewObjectID(),
		BlogID:   "1",
		UserID:   "1",
		Rating:   5,
	}

	_, err := suite.repository.InsertRating(context.Background(), &rating)

	if err != nil {
		suite.T().Fatal(err)
	}

	// now we can test the GetRatingByBlogID function
	ratings, err := suite.repository.GetRatingByBlogID(context.Background(), "1")

	if err != nil {
		suite.T().Fatal(err)
	}

	suite.Equal(rating.BlogID, ratings[0].BlogID)
}

func (suite *BlogRatingRepositorySuit) TestGetRatingByID() {
	// we need to insert some data first
	rating := domain.BlogRating{
		RatingID: primitive.NewObjectID(),
		BlogID:   "1",
		UserID:   "1",
		Rating:   5,
	}

	_, err := suite.repository.InsertRating(context.Background(), &rating)

	if err != nil {
		suite.T().Fatal(err)
	}

	// now we can test the GetRatingByID function
	foundRating, err := suite.repository.GetRatingByID(context.Background(), rating.RatingID.Hex())

	if err != nil {
		suite.T().Fatal(err)
	}

	suite.Equal(rating.BlogID, foundRating.BlogID)
}

func (suite *BlogRatingRepositorySuit) TestInsertRating() {
	// we need to insert some data first
	rating := domain.BlogRating{
		RatingID: primitive.NewObjectID(),
		BlogID:   "1",
		UserID:   "1",
		Rating:   5,
	}

	_, err := suite.repository.InsertRating(context.Background(), &rating)

	if err != nil {
		suite.T().Fatal(err)
	}

	// now we can test the InsertRating function
	insertedRating, err := suite.repository.GetRatingByID(context.Background(), rating.RatingID.Hex())

	if err != nil {
		suite.T().Fatal(err)
	}

	suite.Equal(rating.BlogID, insertedRating.BlogID)
}

func (suite *BlogRatingRepositorySuit) TestUpdateRating() {
	// we need to insert some data first
	rating := domain.BlogRating{
		RatingID: primitive.NewObjectID(),
		BlogID:   "1",
		UserID:   "1",
		Rating:   5,
	}

	_, err := suite.repository.InsertRating(context.Background(), &rating)

	if err != nil {
		suite.T().Fatal(err)
	}

	// now we can test the UpdateRating function
	updatedRating, prevRating, err := suite.repository.UpdateRating(context.Background(), 4, rating.RatingID.Hex())

	if err != nil {
		suite.T().Fatal(err)
	}

	suite.Equal(5, prevRating)
	suite.Equal(4, updatedRating.Rating)
}

func (suite *BlogRatingRepositorySuit) TestDeleteRating() {
	// we need to insert some data first
	rating := domain.BlogRating{
		RatingID: primitive.NewObjectID(),
		BlogID:   "1",
		UserID:   "1",
		Rating:   5,
	}

	_, err := suite.repository.InsertRating(context.Background(), &rating)

	if err != nil {
		suite.T().Fatal(err)
	}

	// now we can test the DeleteRating function
	deletedRating, err := suite.repository.DeleteRating(context.Background(), rating.RatingID.Hex())

	if err != nil {
		suite.T().Fatal(err)
	}

	suite.Equal(rating.BlogID, deletedRating.BlogID)
}

func (suite *BlogRatingRepositorySuit) TestGetRatingByBlogIDFail() {
	// we need to insert some data first
	rating := domain.BlogRating{
		RatingID: primitive.NewObjectID(),
		BlogID:   "1",
		UserID:   "1",
		Rating:   5,
	}

	_, err := suite.repository.InsertRating(context.Background(), &rating)

	if err != nil {
		suite.T().Fatal(err)
	}

	// now we can test the GetRatingByBlogID function
	blogRate, err := suite.repository.GetRatingByBlogID(context.Background(), "2")

	suite.Nil(err)
	suite.Empty(blogRate)

}

func (suite *BlogRatingRepositorySuit) TestGetRatingByIDFail() {
	// we need to insert some data first
	rating := domain.BlogRating{
		RatingID: primitive.NewObjectID(),
		BlogID:   "1",
		UserID:   "1",
		Rating:   5,
	}

	_, err := suite.repository.InsertRating(context.Background(), &rating)

	if err != nil {
		suite.T().Fatal(err)
	}

	// now we can test the GetRatingByID function
	_, err = suite.repository.GetRatingByID(context.Background(), "5")

	suite.NotNil(err)
}

func (suite *BlogRatingRepositorySuit) TestInsertRatingFail() {
	// we need to insert some data first
	rating := domain.BlogRating{
		RatingID: primitive.NewObjectID(),
		BlogID:   "1",
		UserID:   "1",
		Rating:   5,
	}

	_, err := suite.repository.InsertRating(context.Background(), &rating)

	if err != nil {
		suite.T().Fatal(err)
	}

	// now we can test the InsertRating function
	_, err = suite.repository.InsertRating(context.Background(), &rating)

	suite.Nil(err) //TODO: It must be nil because the rating is already inserted
}

func Test_RunBlogRatingRepositorySuit(t *testing.T) {
	suite.Run(t, &BlogRatingRepositorySuit{})
}

package tests

import (
	"blog_api/delivery/env"
	"blog_api/domain"
	cron_jobs "blog_api/infrastructure/cron"
	initdb "blog_api/infrastructure/db"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const LIFESPAN = 12

var UserData = []domain.User{
	{Username: "test1", Email: "test1@gmail.com", CreatedAt: time.Now().Add(-LIFESPAN * time.Hour), IsVerified: false},
	{Username: "test2", Email: "test2@gmail.com", CreatedAt: time.Now().Add(-(LIFESPAN + 1) * time.Hour), IsVerified: false},
	{Username: "test3", Email: "test3@gmail.com", CreatedAt: time.Now(), IsVerified: false},
	{Username: "test4", Email: "test4@gmail.com", CreatedAt: time.Now().Add(-LIFESPAN * time.Hour), IsVerified: true},
}

type CronJobsTestSuite struct {
	suite.Suite
	client     *mongo.Client
	collection *mongo.Collection
}

func (suite *CronJobsTestSuite) SetupSuite() {
	err := env.LoadEnvironmentVariables("../.env")
	if err != nil {
		suite.T().Fatal(err)
		return
	}

	client, err := initdb.ConnectDB(env.ENV.DB_ADDRESS, env.ENV.TEST_DB_NAME)
	if err != nil {
		suite.T().Fatal(err)
		return
	}

	suite.client = client
	suite.collection = client.Database(env.ENV.TEST_DB_NAME).Collection(domain.CollectionUsers)
}

func (suite *CronJobsTestSuite) SetupTest() {
	suite.collection.DeleteMany(context.Background(), bson.D{})
}

func (suite *CronJobsTestSuite) TearDownSuite() {
	initdb.DisconnectDB(suite.client)
}

func (suite *CronJobsTestSuite) TestDeleteUnverifiedUsers() {
	for _, user := range UserData {
		_, err := suite.collection.InsertOne(context.Background(), user)
		suite.Nil(err)
	}

	cron_jobs.DeleteUnverifiedUsers(suite.collection, time.Duration(LIFESPAN)*time.Hour)()

	count, err := suite.collection.CountDocuments(context.Background(), bson.D{})
	suite.Nil(err)
	suite.Equal(count, int64(2))

	res := suite.collection.FindOne(context.Background(), bson.D{{Key: "username", Value: UserData[0].Username}})
	suite.NotNil(res.Err())
	suite.Equal(res.Err(), mongo.ErrNoDocuments)

	res = suite.collection.FindOne(context.Background(), bson.D{{Key: "username", Value: UserData[1].Username}})
	suite.NotNil(res.Err())
	suite.Equal(res.Err(), mongo.ErrNoDocuments)

	res = suite.collection.FindOne(context.Background(), bson.D{{Key: "username", Value: UserData[2].Username}})
	suite.Nil(res.Err())

	res = suite.collection.FindOne(context.Background(), bson.D{{Key: "username", Value: UserData[3].Username}})
	suite.Nil(res.Err())
}

func TestCronJobs(t *testing.T) {
	suite.Run(t, new(CronJobsTestSuite))
}

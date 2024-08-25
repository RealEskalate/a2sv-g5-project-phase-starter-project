package tests

import (
	"blog_api/delivery/env"
	"blog_api/domain"
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
	{Username: "test1", CreatedAt: time.Now().Add(-LIFESPAN * time.Hour), IsVerified: false},
	{Username: "test2", CreatedAt: time.Now().Add(-(LIFESPAN + 1) * time.Hour), IsVerified: false},
	{Username: "test3", CreatedAt: time.Now(), IsVerified: false},
	{Username: "test4", CreatedAt: time.Now().Add(-LIFESPAN * time.Hour), IsVerified: true},
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

}

func TestCronJobs(t *testing.T) {
	suite.Run(t, new(CronJobsTestSuite))
}

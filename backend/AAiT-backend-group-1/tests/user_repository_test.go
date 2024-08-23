package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/repository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	user_repo domain.UserRepository
	client    *mongo.Client
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	connectionString := os.Getenv("DATABASE_URI")
	testUserCollectionName := os.Getenv("TEST_USER_COLLECTION")
	databaseName := os.Getenv("TEST_DB_NAME")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, errConnect := mongo.Connect(context.TODO(), clientOptions)

	if errConnect != nil {
		fmt.Println("Connection error : " + errConnect.Error())
	}

	errPing := client.Ping(context.TODO(), nil)
	if errPing != nil {
		fmt.Println("Ping error : " + errPing.Error())
	}

	suite.client = client
	database := client.Database(databaseName)
	testUserCollection := database.Collection(testUserCollectionName)
	infrastructure.EstablisUniqueUsernameIndex(testUserCollection, "username")
	infrastructure.EstablisUniqueUsernameIndex(testUserCollection, "email")
	suite.user_repo = repository.NewUserRespository(testUserCollection)

}

func (suite *UserRepositoryTestSuite) SetupTest() {
	testDatabaseName := os.Getenv("TEST_DB_NAME")
	testUserCollectionName := os.Getenv("TEST_USER_COLLECTION")
	if _, err := suite.client.Database(testDatabaseName).Collection(testUserCollectionName).DeleteMany(context.TODO(), bson.D{}); err != nil {
		fmt.Println("Error clearing out the test collection" + err.Error())
	}
}

func (suite *UserRepositoryTestSuite) TestCreateUser_Positive() {
	user := domain.User{
		Username: "testOne",
		Password: "testOne",
		Role:     "user",
	}

	fetchedUser, err := suite.user_repo.Create(context.TODO(), &user)
	suite.Nil(err, "Error should be nil")
	suite.Equal(user.Username, fetchedUser.Username, "Username should be the same")
	suite.Equal(user.Password, fetchedUser.Password, "Password should be the same")
	suite.Equal(user.Role, fetchedUser.Role, "Role should be the same")
}

func (suite *UserRepositoryTestSuite) TestCreateUser_Duplicate() {
	user := domain.User{
		Username: "testOne",
		Password: "testOne",
		Role:     "user",
	}

	userTwo := domain.User{
		Username: "testOne",
		Password: "testOne",
		Role:     "user",
	}

	suite.user_repo.Create(context.TODO(), &user)
	_, err := suite.user_repo.Create(context.TODO(), &userTwo)

	suite.Error(err, "Error should be nil")
	suite.Equal(http.StatusConflict, err.StatusCode(), "Status code should be 409")
}

func TestUserSuite(t *testing.T) {
	errLoad := godotenv.Load("../.env")
	if errLoad != nil {
		fmt.Println("error loading the env file")
	}
	suite.Run(t, &UserRepositoryTestSuite{})
}

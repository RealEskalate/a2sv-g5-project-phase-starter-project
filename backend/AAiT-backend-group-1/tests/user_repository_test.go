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
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		Role:     "admin",
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

func (suite *UserRepositoryTestSuite) TestCountByUsername_Positive() {
	user := domain.User{
		Username: "testOne",
		Password: "testOne",
		Role:     "admin",
	}
	count, err := suite.user_repo.CountByUsername(context.TODO(), user.Username)

	suite.NoError(err, "Error should be nil")
	suite.Equal(0, count, "Count should be 0")

	suite.user_repo.Create(context.TODO(), &user)
	count, err = suite.user_repo.CountByUsername(context.TODO(), user.Username)
	suite.NoError(err, "Error should be nil")
	suite.Equal(1, count, "Count should be 1")

}

func (suite *UserRepositoryTestSuite) TestCountByEmail_Positive() {
	user := domain.User{
		Username: "testOne",
		Email:    "example@example.com",
	}
	count, err := suite.user_repo.CountByEmail(context.TODO(), user.Email)

	suite.NoError(err, "Error should be nil")
	suite.Equal(0, count, "Count should be 0")

	suite.user_repo.Create(context.TODO(), &user)
	count, err = suite.user_repo.CountByEmail(context.TODO(), user.Email)
	suite.NoError(err, "Error should be nil")
	suite.Equal(1, count, "Count should be 1")
}

func (suite *UserRepositoryTestSuite) TestCheckExistence_Positive() {
	user := domain.User{
		Username: "testOne",
		Email:    "example@example.com",
	}

	count, err := suite.user_repo.CheckExistence(context.TODO(), primitive.NewObjectID().Hex())
	suite.NoError(err, "Error should be nil")
	suite.Equal(0, count, "Count should be 0")

	var fetchedUser domain.User
	suite.user_repo.Create(context.TODO(), &user)
	testUserCollectionName := os.Getenv("TEST_USER_COLLECTION")
	databaseName := os.Getenv("TEST_DB_NAME")
	errFetch := suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}}).Decode(&fetchedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")

	count, err = suite.user_repo.CheckExistence(context.TODO(), fetchedUser.ID.Hex())
	suite.Equal(1, count, "Count should be 1")
}

func (suite *UserRepositoryTestSuite) TestFindByID_Positive() {
	user := domain.User{
		Username: "testOne",
		Email:    "example@example.com",
	}

	var fetchedUser domain.User
	suite.user_repo.Create(context.TODO(), &user)
	testUserCollectionName := os.Getenv("TEST_USER_COLLECTION")
	databaseName := os.Getenv("TEST_DB_NAME")
	errFetch := suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}}).Decode(&fetchedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")

	fetchedUserByID, err := suite.user_repo.FindById(context.TODO(), fetchedUser.ID.Hex())
	suite.NoError(err, "Error should be nil")
	suite.Equal(fetchedUser.Username, fetchedUserByID.Username, "Username should be the same")
	suite.Equal(fetchedUser.Email, fetchedUserByID.Email, "Email should be the same")
}

func (suite *UserRepositoryTestSuite) TestFindByEmail_Positive() {
	user := domain.User{
		Username: "testOne",
		Email:    "example@example.com",
	}

	suite.user_repo.Create(context.TODO(), &user)

	fetchedUser, err := suite.user_repo.FindByEmail(context.TODO(), user.Email)
	suite.NoError(err, "Error should be nil")
	suite.Equal(fetchedUser.Username, user.Username, "Username should be the same")
	suite.Equal(fetchedUser.Email, user.Email, "Email should be the same")
}

func (suite *UserRepositoryTestSuite) TestFindByUsername_Positive() {
	user := domain.User{
		Username: "testOne",
		Email:    "example@example.com",
	}

	suite.user_repo.Create(context.TODO(), &user)

	fetchedUser, err := suite.user_repo.FindByUsername(context.TODO(), user.Username)
	suite.NoError(err, "Error should be nil")
	suite.Equal(fetchedUser.Username, user.Username, "Username should be the same")
	suite.Equal(fetchedUser.Email, user.Email, "Email should be the same")
}

func (suite *UserRepositoryTestSuite) TestFindAll_Positive() {
	user := domain.User{
		Username: "testOne",
		Email:    "emailOne",
		Password: "testOne",
		Role:     "user",
	}

	userTwo := domain.User{
		Username: "testTwo",
		Email:    "emailTwo",
		Password: "testOne",
		Role:     "user",
	}

	suite.user_repo.Create(context.TODO(), &user)

	fetchedUser, err := suite.user_repo.FindAll(context.TODO())
	suite.NoError(err, "Error should be nil")
	suite.Equal(1, len(fetchedUser), "Length should be 1")
	suite.Equal(fetchedUser[0].Username, user.Username, "Email should be the same")

	_, errCreate := suite.user_repo.Create(context.TODO(), &userTwo)
	suite.NoError(errCreate, "Error should be nil")
	fetchedUser, err = suite.user_repo.FindAll(context.TODO())
	suite.NoError(err, "Error should be nil")
	suite.Equal(2, len(fetchedUser), "Length should be 2")
}

func (suite *UserRepositoryTestSuite) TestUpdateProfile_Positive() {
	user := domain.User{
		Username: "testOne",
		Email:    "emailOne",
		Password: "testOne",
		Role:     "user",
	}

	suite.user_repo.Create(context.TODO(), &user)

	var fetchedUser domain.User
	testUserCollectionName := os.Getenv("TEST_USER_COLLECTION")
	databaseName := os.Getenv("TEST_DB_NAME")
	errFetch := suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}}).Decode(&fetchedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")

	fetchedUser.Username = "testTwo"
	fetchedUser.Email = "emailTwo"
	fetchedUser.Bio = "some new bio"
	fetchedUser.Role = "admin"

	err := suite.user_repo.UpdateProfile(context.TODO(), fetchedUser.ID.Hex(), map[string]interface{}{
		"username": fetchedUser.Username,
		"email":    fetchedUser.Email,
		"bio":      fetchedUser.Bio,
		"role":     fetchedUser.Role,
	})

	var updatedUser domain.User
	errFetch = suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: fetchedUser.Username}}).Decode(&updatedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")

	suite.NoError(err, "Error should be nil")
	suite.Equal(updatedUser.Username, fetchedUser.Username, "Username should be the same")
	suite.Equal(updatedUser.Email, fetchedUser.Email, "Email should be the same")
	suite.Equal(updatedUser.Bio, fetchedUser.Bio, "Bio should be the same")
	suite.Equal(user.Role, updatedUser.Role, "Role should be the same")
}

func (suite *UserRepositoryTestSuite) TestUpdateProfile_Negative() {
	err := suite.user_repo.UpdateProfile(context.TODO(), primitive.NewObjectID().Hex(), map[string]interface{}{
		"username": "not important",
		"email":    "not important",
		"bio":      "not important",
		"role":     "not important",
	})

	suite.Error(err, "Error should not be nil")
	suite.Equal(http.StatusNotFound, err.StatusCode(), "Status code should be 404")
}

func (suite *UserRepositoryTestSuite) TestUpdateUser_Empty() {
	user := domain.User{
		Username: "testOne",
		Email:    "emailOne",
		Password: "testOne",
		Role:     "user",
	}

	suite.user_repo.Create(context.TODO(), &user)

	var fetchedUser domain.User
	testUserCollectionName := os.Getenv("TEST_USER_COLLECTION")
	databaseName := os.Getenv("TEST_DB_NAME")
	errFetch := suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}}).Decode(&fetchedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")

	err := suite.user_repo.UpdateProfile(context.TODO(), fetchedUser.ID.Hex(), map[string]interface{}{})

	var updatedUser domain.User
	errFetch = suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: fetchedUser.Username}}).Decode(&updatedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")

	suite.NoError(err, "Error should be nil")
	suite.Equal(updatedUser.Username, fetchedUser.Username, "Username should be the same")
	suite.Equal(updatedUser.Email, fetchedUser.Email, "Email should be the same")
	suite.Equal(updatedUser.Bio, fetchedUser.Bio, "Bio should be the same")
	suite.Equal(user.Role, updatedUser.Role, "Role should be the same")
}

func (suite *UserRepositoryTestSuite) TestUpdatePassword_Positive() {
	user := domain.User{
		Username: "testOne",
		Email:    "emailOne",
		Password: "testOne",
		Role:     "user",
	}

	suite.user_repo.Create(context.TODO(), &user)

	var fetchedUser domain.User
	testUserCollectionName := os.Getenv("TEST_USER_COLLECTION")
	databaseName := os.Getenv("TEST_DB_NAME")
	errFetch := suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}}).Decode(&fetchedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")

	err := suite.user_repo.UpdatePassword(context.TODO(), fetchedUser.ID.Hex(), "newPassword")
	suite.NoError(err, "Error should be nil")

	var updatedUser domain.User
	errFetch = suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: fetchedUser.Username}}).Decode(&updatedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")
	suite.NotEqual(updatedUser.Password, fetchedUser.Password, "Password should not be the same")

}

func (suite *UserRepositoryTestSuite) TestUpdatePassword_Negative() {
	err := suite.user_repo.UpdatePassword(context.TODO(), primitive.NewObjectID().Hex(), "newPassword")
	suite.Error(err, "Error should not be nil")
	suite.Equal(http.StatusNotFound, err.StatusCode(), "Status code should be 404")
}

func (suite *UserRepositoryTestSuite) TestUpdateRole_Positive() {
	user := domain.User{
		Username: "testOne",
		Email:    "emailOne",
		Password: "testOne",
		Role:     "user",
	}

	suite.user_repo.Create(context.TODO(), &user)

	var fetchedUser domain.User
	testUserCollectionName := os.Getenv("TEST_USER_COLLECTION")
	databaseName := os.Getenv("TEST_DB_NAME")
	errFetch := suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}}).Decode(&fetchedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")

	err := suite.user_repo.UpdateRole(context.TODO(), fetchedUser.ID.Hex(), "admin")
	suite.NoError(err, "Error should be nil")

	var updatedUser domain.User
	errFetch = suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: fetchedUser.Username}}).Decode(&updatedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")
	suite.Equal("admin", updatedUser.Role, "Role should be admin")
}

func (suite *UserRepositoryTestSuite) TestUpdateRole_Negative() {
	err := suite.user_repo.UpdateRole(context.TODO(), primitive.NewObjectID().Hex(), "admin")
	suite.Error(err, "Error should not be nil")
	suite.Equal(http.StatusNotFound, err.StatusCode(), "Status code should be 404")
}

func (suite *UserRepositoryTestSuite) TestDeleteUser_Positive() {
	user := domain.User{
		Username: "testOne",
		Email:    "emailOne",
		Password: "testOne",
		Role:     "user",
	}

	suite.user_repo.Create(context.TODO(), &user)

	var fetchedUser domain.User
	testUserCollectionName := os.Getenv("TEST_USER_COLLECTION")
	databaseName := os.Getenv("TEST_DB_NAME")
	errFetch := suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}}).Decode(&fetchedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")

	err := suite.user_repo.Delete(context.TODO(), fetchedUser.ID.Hex())
	suite.NoError(err, "Error should be nil")

	var deletedUser domain.User
	errFetch = suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: fetchedUser.Username}}).Decode(&deletedUser)
	suite.Error(errFetch, "Error should not be nil")
	suite.Equal(mongo.ErrNoDocuments, errFetch, "Error should be mongo.ErrNoDocuments")
}

func (suite *UserRepositoryTestSuite) TestDeleteUser_Negative() {
	err := suite.user_repo.Delete(context.TODO(), primitive.NewObjectID().Hex())
	suite.Error(err, "Error should not be nil")
	suite.Equal(http.StatusNotFound, err.StatusCode(), "Status code should be 404")
}

func (suite *UserRepositoryTestSuite) TestUploadProfilePicture() {
	user := domain.User{
		Username: "testOne",
		Email:    "emailOne",
		Password: "testOne",
		Role:     "user",
	}

	photo := domain.Photo{
		Filename:  "test.jpg",
		FilePath:  "test.jpg",
		Public_id: "test.jpg",
	}

	suite.user_repo.Create(context.TODO(), &user)

	var fetchedUser domain.User
	testUserCollectionName := os.Getenv("TEST_USER_COLLECTION")
	databaseName := os.Getenv("TEST_DB_NAME")
	errFetch := suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}}).Decode(&fetchedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")

	err := suite.user_repo.UploadProfilePicture(context.TODO(), photo, fetchedUser.ID.Hex())
	suite.NoError(err, "Error should be nil")

	var updatedUser domain.User
	errFetch = suite.client.Database(databaseName).Collection(testUserCollectionName).FindOne(context.TODO(), bson.D{{Key: "username", Value: fetchedUser.Username}}).Decode(&updatedUser)
	suite.NoError(errFetch, "error when trying to search the database, most out of the scope of this test")

	suite.Equal(photo.Filename, updatedUser.ProfilePicture.Filename, "Filename should be the same")
	suite.Equal(photo.FilePath, updatedUser.ProfilePicture.FilePath, "FilePath should be the same")
	suite.Equal(photo.Public_id, updatedUser.ProfilePicture.Public_id, "Public_id should be the same")

}

func (suite *UserRepositoryTestSuite) TestUploadProfilePicture_Negative() {
	photo := domain.Photo{
		Filename:  "test.jpg",
		FilePath:  "test.jpg",
		Public_id: "test.jpg",
	}

	err := suite.user_repo.UploadProfilePicture(context.TODO(), photo, primitive.NewObjectID().Hex())
	suite.Error(err, "Error should not be nil")
	suite.Equal(http.StatusNotFound, err.StatusCode(), "Status code should be 404")

}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	err := suite.client.Disconnect(context.Background())
	if err != nil {
		fmt.Println("Error disconnecting the database")
	}
}

func TestUserSuite(t *testing.T) {
	errLoad := godotenv.Load("../.env")
	if errLoad != nil {
		fmt.Println("error loading the env file")
	}
	suite.Run(t, &UserRepositoryTestSuite{})
}

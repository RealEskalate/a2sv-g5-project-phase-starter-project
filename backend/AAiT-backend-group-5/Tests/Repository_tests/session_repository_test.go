package tests

// import (
// 	"context"
// 	"log"
// 	"testing"

// 	"github.com/stretchr/testify/suite"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"

// 	"github.com/aait.backend.g5.main/backend/Domain/Models"
// 	"github.com/aait.backend.g5.main/backend/Repository"
// )

// type SessionRepoTestSuite struct {
// 	suite.Suite
// 	Client     *mongo.Client
// 	Collection *mongo.Collection
// 	Repo       *repository.SessionRepo
// }

// func (suite *SessionRepoTestSuite) SetupTest() {
// 	// Connect to the test MongoDB instance
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = client.Ping(context.TODO(), readpref.Primary())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	suite.Client = client
// 	db := suite.Client.Database("test-db")
// 	suite.Collection = db.Collection("session-collection")
// 	suite.Repo = repository.NewSessionRepository(db).(*repository.SessionRepo)
// }

// func (suite *SessionRepoTestSuite) TearDownSuite() {
// 	err := suite.Client.Disconnect(context.TODO())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }


// func (suite *SessionRepoTestSuite) TestSaveToken() {
// 	session := &models.Session{
// 		UserID:       primitive.NewObjectID().Hex(),
// 		AccessToken:  "access_token_value",
// 		RefreshToken: "refresh_token_value",
// 	}

// 	err := suite.Repo.SaveToken(context.TODO(), session)
// 	suite.Empty(err)

// 	// Verify the session was saved correctly
// 	var foundSession models.Session
// 	ferr := suite.Collection.FindOne(context.TODO(), bson.M{"user_id": session.UserID}).Decode(&foundSession)
// 	suite.Nil(ferr)
// 	suite.Equal(session.AccessToken, foundSession.AccessToken)
// 	suite.Equal(session.RefreshToken, foundSession.RefreshToken)
// }

// func (suite *SessionRepoTestSuite) TestUpdateToken() {
// 	session := &models.Session{
// 		UserID:       primitive.NewObjectID().Hex(),
// 		AccessToken:  "old_access_token",
// 		RefreshToken: "old_refresh_token",
// 	}

// 	// Insert initial session
// 	_, err := suite.Collection.InsertOne(context.TODO(), session)
// 	suite.Nil(err)

// 	// Update tokens
// 	session.AccessToken = "new_access_token"
// 	session.RefreshToken = "new_refresh_token"
// 	err = suite.Repo.UpdateToken(context.TODO(), session)
// 	suite.Empty(err)

// 	// Verify the tokens were updated
// 	var updatedSession models.Session
// 	err = suite.Collection.FindOne(context.TODO(), bson.M{"user_id": session.UserID}).Decode(&updatedSession)
// 	suite.Nil(err)
// 	suite.Equal("new_access_token", updatedSession.AccessToken)
// 	suite.Equal("new_refresh_token", updatedSession.RefreshToken)
// }

// func (suite *SessionRepoTestSuite) TestRemoveToken() {
// 	session := &models.Session{
// 		UserID:       primitive.NewObjectID().Hex(),
// 		AccessToken:  "access_token_value",
// 		RefreshToken: "refresh_token_value",
// 	}

// 	// Insert session
// 	_, err := suite.Collection.InsertOne(context.TODO(), session)
// 	suite.Empty(err)

// 	// Remove session token
// 	err = suite.Repo.RemoveToken(context.TODO(), session.UserID)
// 	suite.Empty(err)
// }

// func (suite *SessionRepoTestSuite) TestGetToken() {
// 	session := &models.Session{
// 		UserID:       primitive.NewObjectID().Hex(),
// 		AccessToken:  "access_token_value",
// 		RefreshToken: "refresh_token_value",
// 	}

// 	// Insert session
// 	_, err := suite.Collection.InsertOne(context.TODO(), session)
// 	suite.Empty(err)

	
// 	// Retrieve the session
// 	foundSession, errResp := suite.Repo.GetToken(context.TODO(), session.UserID)
// 	suite.Empty(errResp)
// 	suite.NotNil(foundSession)
// 	suite.Equal(session.AccessToken, foundSession.AccessToken)
// 	suite.Equal(session.RefreshToken, foundSession.RefreshToken)
// }

// func TestSessionRepoTestSuite(t *testing.T) {
// 	suite.Run(t, new(SessionRepoTestSuite))
// }

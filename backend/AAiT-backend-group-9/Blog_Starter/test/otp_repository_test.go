/*

type otpRepository struct {
	database   *mongo.Database
	collection string
}

func NewOtpRepository(db *mongo.Database, collection string) domain.OtpRepository {
	return &otpRepository{
		database:   db,
		collection: collection,
	}
}

func (or *otpRepository) SaveOtp(c context.Context, otp *domain.Otp) error {
	collection := or.database.Collection(or.collection)

	// Define an update operation
	update := bson.M{
		"$set": otp,
	}

	// Define options for the update operation (e.g., to perform an upsert)
	options := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(c, bson.M{"_id": otp.ID}, update, options)

	return err
}

func (or *otpRepository) InvalidateOtp(c context.Context, otp *domain.Otp) error {
	collection := or.database.Collection(or.collection)

	// Define an update operation
	update := bson.M{}

	// Define options for the update operation (e.g., to perform an upsert)
	options := options.Update().SetUpsert(false)

	// Perform the update operation
	_, err := collection.UpdateOne(c, bson.M{"_id": otp.ID}, update, options)

	return err
}

func (or *otpRepository) GetOtpByEmail(c context.Context, email string) (domain.Otp, error) {
	collection := or.database.Collection(or.collection)
	var otp domain.Otp
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&otp)
	return otp, err
}

func (or *otpRepository) GetByID(c context.Context, id string) (domain.Otp, error) {
	collection := or.database.Collection(or.collection)
	var otp domain.Otp
	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&otp)
	return otp, err
}


*/

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

type OtpRepositorySuit struct {
	suite.Suite
	// the funcionalities we need to test
	repository domain.OtpRepository
	db         *mongo.Database
}

func (suite *OtpRepositorySuit) SetupSuite() {
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
	repository := repository.NewOtpRepository(db, "otp")

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.repository = repository
	suite.db = db
}

func (suite *OtpRepositorySuit) TearDownSuite() {
	// we need to drop the table we used in the tests
	defer suite.db.Drop(context.Background())
}

func (suite *OtpRepositorySuit) SetupTest() {
	// this function runs before every test in the suite
	// we need to clear the table before every test
	_, err := suite.db.Collection("otp").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}

func TestOtpRepositorySuit(t *testing.T) {
	suite.Run(t, new(OtpRepositorySuit))
}

func (suite *OtpRepositorySuit) TestSaveOtp() {
	otp := domain.Otp{
		ID:    primitive.NewObjectID(),
		Email: "eyob@gmail.com",
		Otp:   "1234",
	}

	err := suite.repository.SaveOtp(context.Background(), &otp)
	suite.Nil(err)

	otp2, err := suite.repository.GetByID(context.Background(), otp.ID.Hex())
	suite.Nil(err)
	suite.Equal(otp.Otp, otp2.Otp)
}

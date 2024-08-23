package repositories

import (
	domain "aait-backend-group4/Domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type otpRepository struct {
	database   mongo.Database
	collection string
}

// NewOTPRepository creates and initializes a new instance of otpRepository.
// It takes a mongo.Database instance and a collection name as parameters.
// Returns an instance of domain.OTPRepository, which is the interface implemented by otpRepository.
func NewOTPRepository(database mongo.Database, collection string) domain.OTPRepository {
	return &otpRepository{
		database:   database,
		collection: collection,
	}
}

// CreateOTP inserts a new OTP (One-Time Password) record into the database.
// It takes a context.Context and a pointer to a domain.UserOTPVerification struct as parameters.
// The struct represents the OTP information associated with a user.
// Returns an error if the insert operation fails.
func (or *otpRepository) CreateOTP(c context.Context, otp *domain.UserOTPVerification) error {
	collection := or.database.Collection(or.collection)

	// Attempt to insert the OTP record into the collection
	_, err := collection.InsertOne(c, otp)
	return err
}

// GetOTPByEmail retrieves an OTP record from the database based on the provided email address.
// It takes a context.Context and an email string as parameters.
// The email string is used to query the database for the corresponding OTP record.
// Returns the retrieved domain.UserOTPVerification struct and an error if any issues occur.
// If no OTP is found for the given email, an error with a descriptive message is returned.
func (or *otpRepository) GetOTPByEmail(c context.Context, email string) (otp domain.UserOTPVerification, err error) {
	collection := or.database.Collection(or.collection)
	var otpResponse domain.UserOTPVerification

	// Query the collection for an OTP record with the specified email
	err = collection.FindOne(c, bson.D{{Key: "email", Value: email}}).Decode(&otpResponse)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Specific error message if no document is found
			return otpResponse, fmt.Errorf("otp with this email not found")
		}
		// For other errors, return a general error message
		return otpResponse, fmt.Errorf("error retrieving otp: %v", err)
	}

	return otpResponse, nil
}

// DeleteOTPByEmail removes an OTP record from the database based on the provided email address.
// It takes a context.Context and an email string as parameters.
// The email string is used to identify which OTP record to delete.
// Returns an error if the deletion operation fails.
func (or *otpRepository) DeleteOTPByEmail(c context.Context, email string) error {
	collection := or.database.Collection(or.collection)

	// Attempt to delete the OTP record from the collection with the specified email
	_, err := collection.DeleteOne(c, bson.D{{Key: "email", Value: email}})
	return err
}

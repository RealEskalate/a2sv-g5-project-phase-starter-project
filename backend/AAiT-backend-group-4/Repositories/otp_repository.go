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

func NewOTPRepository(database mongo.Database, collection string) domain.OTPRepository {
	return &otpRepository{
		database,
		collection,
	}
}

func (or *otpRepository) CreateOTP(c context.Context, otp *domain.UserOTPVerification) error {
	collection := or.database.Collection(or.collection)
	_, err := collection.InsertOne(c, otp)
	return err
}

func (or *otpRepository) GetOTPByEmail(c context.Context, email string) (otp domain.UserOTPVerification, err error) {
	collection := or.database.Collection(or.collection)
	var otpResponse domain.UserOTPVerification
	err = collection.FindOne(c, bson.D{{Key: "email", Value: email}}).Decode(&otpResponse)
	if err != nil {
		return otpResponse, fmt.Errorf("otp with this email not found")
	}
	return otpResponse, nil
}

func (or *otpRepository) DeleteOTPByEmail(c context.Context, email string) error {
	collection := or.database.Collection(or.collection)
	_, err := collection.DeleteOne(c, bson.D{{Key: "email", Value: email}})
	return err
}

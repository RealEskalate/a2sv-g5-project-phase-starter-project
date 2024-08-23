package repository

import (
	"Blog_Starter/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Otp{}, err
	}
	collection := or.database.Collection(or.collection)
	var otp domain.Otp
	err = collection.FindOne(c, bson.M{"_id": objID}).Decode(&otp)
	return otp, err
}

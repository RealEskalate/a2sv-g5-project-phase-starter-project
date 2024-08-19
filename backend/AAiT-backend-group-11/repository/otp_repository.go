package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OtpRepository struct {
	collection *mongo.Collection
}

// GetByID implements interfaces.OTPRepository.
func (o *OtpRepository) GetByID(id string) (entities.OTP, error) {
	colllection := o.collection.Database().Collection("otp")
	ctx := context.Background()

	var otp entities.OTP

	err := colllection.FindOne(ctx, bson.M{"_id": id}).Decode(&otp)

	return otp, err
}

// GetOtpByEmail implements interfaces.OTPRepository.
func (o *OtpRepository) GetOtpByEmail(email string) (entities.OTP, error) {
	collection := o.collection.Database().Collection("otp")
	ctx := context.Background()

	var otp entities.OTP

	err := collection.FindOne(ctx, bson.M {"email": email}).Decode(&otp)

	return otp, err
}

// InvalidateOtp implements interfaces.OTPRepository.
func (o *OtpRepository) InvalidateOtp(otp *entities.OTP) error {
	collection := o.collection.Database().Collection("otp")
	ctx := context.Background()

	update := bson.M{}

	// Define options for the update operation (e.g., to perform an upsert)
	options := options.Update().SetUpsert(false)

	// Perform the update operation
	_, err := collection.UpdateOne(ctx, bson.M{"_id": otp.ID}, update, options)

	return err
}

// SaveOtp implements interfaces.OTPRepository.
func (o *OtpRepository) SaveOtp(otp *entities.OTP) error {
	collection := o.collection.Database().Collection("otp")
	ctx := context.Background()

	update := bson.M{
		"$set": otp,
	}

	options := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, bson.M{"_id": otp.ID}, update, options)

	return err

}

func NewOtpRepository(collection *mongo.Collection) interfaces.OTPRepository {
	return &OtpRepository{
		collection: collection,
	}
}

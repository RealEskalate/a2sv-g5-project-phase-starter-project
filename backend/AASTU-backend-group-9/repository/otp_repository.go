package repository

import (
	"blog/database"
	"blog/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type OTPRepository struct {
	database   database.Database
	collection string
}

func (u *OTPRepository) SaveOTP(c context.Context, otp *domain.OTP) error {
	collection := u.database.Collection(domain.CollectionOTP)
	_, err := collection.InsertOne(c, otp)
	return err
}
func (u *OTPRepository) GetOTPByEmail(c context.Context, email string) (*domain.OTP, error) {
	collection := u.database.Collection(domain.CollectionOTP)
	filter := bson.M{"email": email}
	otp := &domain.OTP{}
	err := collection.FindOne(c, filter).Decode(otp)
	return otp, err
}

func (u *OTPRepository) DeleteOTP(c context.Context, email string) error {
	collection := u.database.Collection(domain.CollectionOTP)
	filter := bson.M{"email": email}
	_, err := collection.DeleteOne(c, filter)
	return err
}
func NewOTPRepository(db database.Database, collection string) domain.OTPRepository {
	return &OTPRepository{
		database:   db,
		collection: collection,
	}
}

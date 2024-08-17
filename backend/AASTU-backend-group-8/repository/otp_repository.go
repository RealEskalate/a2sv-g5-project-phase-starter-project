package repository

import (
	"context"
	"meleket/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

type OtpRepository struct{
	collection domain.Collection
}

func NewOtpRepository(col domain.Collection) *OtpRepository {
	return &OtpRepository{ collection : col } 
}

func (r *UserRepository) StoreOTP(userID primitive.ObjectID, otp string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	otpEntry := domain.OTP{
		UserID:    userID,
		OTP:       otp,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}
	_, err := r.collection.InsertOne(ctx, otpEntry)
	return err
}

func (r *UserRepository) GetOTP(userID primitive.ObjectID) (*domain.OTP, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var otpEntry domain.OTP
	err := r.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&otpEntry)
	return &otpEntry, err
}

func (r *UserRepository) DeleteOTP(userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.collection.DeleteOne(ctx, bson.M{"user_id": userID})
	return err
}
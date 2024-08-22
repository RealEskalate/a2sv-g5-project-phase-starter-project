package repository

import (
	"context"
	"meleket/domain"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type OTPRepository struct {
	collection domain.Collection
	mutex      sync.RWMutex
}

func NewOtpRepository(col domain.Collection) *OTPRepository {
	return &OTPRepository{
		collection: col,
		mutex:      sync.RWMutex{},
	}
}

func (r *OTPRepository) StoreOTP(otp *domain.OTP) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	_, err := r.collection.InsertOne(ctx, otp)
	return err
}

func (r *OTPRepository) GetOTPByEmail(email string) (*domain.OTP, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var otpEntry domain.OTP
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&otpEntry)
	return &otpEntry, err
}

func (r *OTPRepository) DeleteOTPByEmail(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	_, err := r.collection.DeleteOne(ctx, bson.M{"email": email})
	return err
}

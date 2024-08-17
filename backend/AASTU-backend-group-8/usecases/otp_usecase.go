package usecases

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *UserUsecase) VerifyOTP(userID primitive.ObjectID, inputOTP string) error {
	otpEntry, err := u.userRepo.GetOTP(userID)
	if err != nil {
		return err
	}

	if otpEntry.ExpiresAt.Before(time.Now()) {
		return errors.New("OTP expired")
	}

	if otpEntry.OTP != inputOTP {
		return errors.New("Invalid OTP")
	}

	// Clean up the used OTP
	return u.userRepo.DeleteOTP(userID)
}
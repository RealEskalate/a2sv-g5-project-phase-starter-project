package repository

import (
	"Blog_Starter/domain"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginRepository struct {
	database   *mongo.Database
	collection string
}


func NewLoginRepository(db *mongo.Database, collection string) domain.LoginRepository {
	return &LoginRepository{
		database:   db,
		collection: collection,
	}
}


// Login implements domain.LoginRepository.
func (l *LoginRepository) Login(c context.Context, user *domain.UserLogin) (*domain.LoginResponse, error) {
	var loginUser domain.LoginResponse
	collection := l.database.Collection(l.collection)
	err := collection.FindOne(c, bson.M{"email": user.Email}).Decode(&loginUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &loginUser, nil
}
// UpdatePassword implements domain.LoginRepository.
func (l *LoginRepository) UpdatePassword(c context.Context, req domain.ChangePasswordRequest, userID string) error {
    collection := l.database.Collection(domain.CollectionOTP)
    
    // Find the OTP document
    var otpDoc bson.M
    err := collection.FindOne(c, bson.M{"email": req.Email, "otp": req.OTP}).Decode(&otpDoc)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return fmt.Errorf("otp not found")
        }
        return err
    }

    // Update the user's password
    userCollection := l.database.Collection(l.collection)
    _, err = userCollection.UpdateOne(c, bson.M{"_id": userID}, bson.M{"$set": bson.M{"password": req.Password}})
    if err != nil {
        return err
    }
    
    return nil
}
package repository

import (
	"context"
	"errors"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type resetPasswordRepository struct {
	database        mongo.Database
	usersCollection string
	resetCollection string
}

func NewResetPasswordRepository(db mongo.Database, userCollection string, resetCollection string) entities.ResetPasswordRepository {
	return &resetPasswordRepository{
		database:        db,
		usersCollection: userCollection,
		resetCollection: resetCollection,
	}
}

func (rp *resetPasswordRepository) GetUserByEmail(c context.Context, email string) (*entities.User, error) {
	collection := rp.database.Collection(rp.usersCollection)
	var user entities.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, custom_error.ErrUserNotFound
	}
	return &user, err
}
func (rp *resetPasswordRepository) ResetPassword(c context.Context, userID string, resetPassword *entities.ResetPasswordRequest) error {

	collection := rp.database.Collection(rp.usersCollection)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"password": resetPassword.NewPassword}})
	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}
	if res.MatchedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	return nil
}

func (rp *resetPasswordRepository) SaveOtp(c context.Context, otp *entities.OtpSave) error {
	collection := rp.database.Collection(rp.resetCollection)

	_, err := collection.InsertOne(c, otp)

	if err != nil {
		return custom_error.ErrErrorSavingOtp
	}

	return err
}

func (rp *resetPasswordRepository) GetOTPByEmail(c context.Context, email string) (*entities.OtpSave, error) {

	collection := rp.database.Collection(rp.resetCollection)
	var otp entities.OtpSave

	err := collection.FindOne(c, bson.M{"email": email}).Decode(&otp)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, custom_error.ErrUserNotFound
	}

	if err != nil {
		return nil, custom_error.ErrErrorGettingOtp
	}

	return &otp, err
}

func (rp *resetPasswordRepository) DeleteOtp(c context.Context, email string) error {

	collection := rp.database.Collection(rp.resetCollection)

	_, err := collection.DeleteOne(c, bson.M{"email": email})

	if err != nil {
		return custom_error.ErrErrorDeletingOtp
	}

	return err
}

package repositories

import (
	"context"
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/repositories/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOtpRepository struct {
	collection *mongo.Collection
}

func NewMongoOtpRepository(db *mongo.Database, collectionName string) repository_interface.IOtpRepository {
	return &MongoOtpRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoOtpRepository) SaveOtp(ctx context.Context, otp models.OtpEntry) error {
	_, err := r.collection.InsertOne(ctx, otp)
	return err
}

func (r *MongoOtpRepository) FindByOtp(ctx context.Context, otp string) (*models.OtpEntry, error) {
	var otpEntry models.OtpEntry
	err := r.collection.FindOne(ctx, bson.M{"otp": otp}).Decode(&otpEntry)
	if err != nil {
		return nil, err
	}
	return &otpEntry, nil
}
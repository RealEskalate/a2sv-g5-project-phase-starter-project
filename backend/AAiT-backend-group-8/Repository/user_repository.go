package repository

import (
	"AAiT-backend-group-8/Domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	db  *mongo.Collection
	ctx context.Context
}

func NewUserRepository(db *mongo.Collection, cont context.Context) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db, ctx: cont}
}

func (r *UserRepositoryImpl) CreateUser(user *Domain.User) error {
	_, err := r.db.InsertOne(context.Background(), user)
	return err
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*Domain.User, error) {
	var user Domain.User
	err := r.db.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (r *UserRepositoryImpl) GetUserByVerificationToken(token string) (*Domain.User, error) {
	var user Domain.User
	err := r.db.FindOne(r.ctx, bson.M{"verification_token": token}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserRepositoryImpl) VerifyUser(user *Domain.User) error {
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{
		"verified":           true,
		"verification_token": user.VerificationToken,
	}}
	_, err := r.db.UpdateOne(r.ctx, filter, update)
	fmt.Println(err)
	return err
}
func (r *UserRepositoryImpl) GetUserCount() (int64, error) {
	count, err := r.db.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}
	return count, nil
}

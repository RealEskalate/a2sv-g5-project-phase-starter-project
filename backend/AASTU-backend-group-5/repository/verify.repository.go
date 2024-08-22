package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailVRepo struct{
	UserRepository
}

func NewEmailVRepo(user_repo UserRepository)*EmailVRepo {
	return &EmailVRepo{
		UserRepository: user_repo,
	}
}

func (repo *EmailVRepo) VerifyUser(id string) error {
	objID,_ := primitive.ObjectIDFromHex(id) 
	filter := bson.D{{Key: "_id" , Value: objID}}
	setter := bson.D{{Key:"$set" , Value: bson.D{{Key:"is_verified" , Value: true}}}}

	_,err := repo.Collection.UpdateOne(context.TODO() , filter , setter)

	return err
}
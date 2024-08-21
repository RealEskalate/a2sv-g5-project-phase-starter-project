package repository

import (
	"context"

	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UploadRepo struct {
	UserRepository
}

func NewUploadRepository(user_repo UserRepository) *UploadRepo {
	return &UploadRepo{
		UserRepository: user_repo,
	}
}

func (repo *UploadRepo) AddProfile(media domain.Media , id string) error {
	objID,_ := primitive.ObjectIDFromHex(id)
	media.ID = primitive.NewObjectID()
	filter := bson.D{{Key : "_id" , Value: objID}}
	data := bson.D{{Key: "profile_picture" , Value: media}}
	setter := bson.D{{Key: "$set" , Value: data}}

	_,err := repo.UserRepository.Collection.UpdateOne(context.TODO() , filter , setter)

	return err
}
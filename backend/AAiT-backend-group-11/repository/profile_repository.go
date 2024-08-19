package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type profileRepository struct {
	db *mongo.Database
	collection *mongo.Collection
	context context.Context
}

func NewProfileRepository(ctx context.Context, db *mongo.Database) interfaces.ProfileRepository{
	return profileRepository{db: db, collection: db.Collection("profile"),context: ctx}
} 
func (repo profileRepository) GetUserProfile(user_id string) (*entities.Profile,error){
	filter:=bson.D{{"userId",user_id}}
	user:=repo.collection.FindOne(context.TODO(),filter)
	if user.Err() ==nil{
		return &entities.Profile{},errors.New("couldn't find the user")
	}
	var profile entities.Profile
	user.Decode(&profile)
	return &profile,nil

}

func (repo profileRepository) CreateUserProfile(profile *entities.Profile) (*entities.Profile, error){
	_,err:=repo.collection.InsertOne(repo.context,profile) 
	if err!=nil{
		return nil,err
	}
	return profile,nil
}

func (repo profileRepository) UpdateUserProfile(profile *entities.Profile)(*entities.Profile,error){
	user_id:=profile.UserID
	filter:=bson.D{{"userId",user_id}}
	_,err:=repo.collection.UpdateOne(context.TODO(),filter,profile)
	if err!=nil{
		return nil,err
	}
	
	return profile,nil
}

func (repo profileRepository) DeleteUserProfile(user_id string)error{
	filter:=bson.D{{"userId",user_id}}
	_,err:=repo.collection.DeleteOne(repo.context,filter)
	if err!=nil{
		return err
	}

	return nil
}

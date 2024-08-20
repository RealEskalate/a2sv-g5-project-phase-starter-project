package repository

import (
	"context"
	"fmt"
	"group3-blogApi/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *UserRepositoryImpl) GetMyProfile(userID string) (domain.User, error) {
	var user domain.User
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid user id")
	}

	err = ur.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil	
}


func (ur *UserRepositoryImpl) GetUsers() ([]domain.User, error) {
	var users []domain.User
	cursor, err := ur.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return []domain.User{}, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user domain.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}


func (ur *UserRepositoryImpl) DeleteUser(userID string) (domain.User, error){
	var user domain.User
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid user id")
	}

	err = ur.collection.FindOneAndDelete(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ur *UserRepositoryImpl) UpdateUserRole(userID, role string) (domain.User, error) {
	var user domain.User
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid user id")
	}

	err = ur.collection.FindOneAndUpdate(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": bson.M{"role": role}}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func(ur *UserRepositoryImpl) DeleteMyAccount(userID string) error{
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user id")
	}

	_, err = ur.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}

func(ur *UserRepositoryImpl) UploadImage(userID string, imagePath string) error{
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user id")
	}

	_, err = ur.collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": bson.M{"image": imagePath}})
	if err != nil {
		return err
	}
	return nil
}

func(ur *UserRepositoryImpl) UpdateMyProfile(user domain.User, UserID string) error{
	objectID, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		return fmt.Errorf("invalid user id")
	}

	_, err = ur.collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": bson.M{"name": user.Username, "bio": user.Bio}})
	if err != nil {
		return err
	}
	return nil
}

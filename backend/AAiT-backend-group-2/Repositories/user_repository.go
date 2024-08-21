package repositories

import (
	domain "AAiT-backend-group-2/Domain"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	userCollection *mongo.Collection
	tokenCollection *mongo.Collection
}

func NewUserRepository(db *mongo.Database)  domain.UserRepository{
	return &userRepository{
			userCollection: db.Collection("users"),
			tokenCollection: db.Collection("resetTokens"),
	}
}

func (ur *userRepository) FindAll(c context.Context) ([]domain.User, error) {
	cursor, err := ur.userCollection.Find(c, bson.M{})

	if err != nil {
		return nil, err
	}

	var users []domain.User

	
	if err := cursor.All(c, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) FindByID(c context.Context, id string) (*domain.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("invalid id")
	}

	var user domain.User

	filter := bson.M{"_id": objectID}

	err = ur.userCollection.FindOne(c, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (ur *userRepository) FindByEmailOrUsername(c context.Context, emailOrUsername string) (*domain.User, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"email": emailOrUsername},
			{"username": emailOrUsername},
		},
	}

	var user domain.User
	err := ur.userCollection.FindOne(c, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}


func (ur *userRepository) Save(c context.Context, user domain.User) error {
	_, err := ur.userCollection.InsertOne(c, user)

	if err != nil {
		return err
	}

	return nil
}
	

func (ur *userRepository) Update(c context.Context, id string, updateData domain.UpdateData) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}

	filter := bson.M{"_id": objectID}


	update := bson.M{
		"$set": updateData,
	}

	_, err = ur.userCollection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Delete(c context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}
	filter := bson.M{"_id": objectID}

	_, err = ur.userCollection.DeleteOne(c, filter)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) CountDocuments(c context.Context) (int64, error) {

	count, err := ur.userCollection.CountDocuments(c, bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ur *userRepository) PromoteUser(c context.Context, id string, updateData domain.UpdateData) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": updateData,
	}

	_, err = ur.userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func (ur *userRepository) DemoteAdmin(c context.Context, id string, updateData domain.UpdateData) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": updateData,
	}

	_, err = ur.userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func (ur *userRepository) ForgotPassword(c context.Context, email string, token string) (string, error) {
	user, err := ur.FindByEmailOrUsername(c, email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	resetToken := domain.ResetToken{
		Token:     token,
		UserID:    user.ID.Hex(),
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	_, err = ur.tokenCollection.InsertOne(c, resetToken)
	if err != nil {
		return "", err
	}

	return resetToken.Token, nil
}



func (ur *userRepository) ValidateResetToken(c context.Context, userID, token string) error {
	var resetToken domain.ResetToken

	filter := bson.M{
		"userid": userID,
	}

	err := ur.tokenCollection.FindOne(c, filter).Decode(&resetToken)

	if err != nil {
		return errors.New("invalid token")
	}

	fmt.Println("refreshToken", resetToken.Token, userID)
	fmt.Println("token", token)

	if resetToken.Token != token {
		return errors.New("invalid token")
	}

	if time.Now().After(resetToken.ExpiresAt) {
		return errors.New("token expired")
	}

	return nil
	
}

func (ur *userRepository) InvalidateResetToken(c context.Context, userID string) error {
	filter := bson.M{
		"userid": userID,
	}

	_, err := ur.tokenCollection.DeleteMany(c,filter)
	
	if err != nil {
		return err
	}

	return nil
}
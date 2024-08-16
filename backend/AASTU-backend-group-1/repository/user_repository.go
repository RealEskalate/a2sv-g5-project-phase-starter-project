package repository

import (
	"blogs/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	userCollection  *mongo.Collection
	tokenCollection *mongo.Collection
}

func NewUserRepository(db mongo.Database) domain.UserRepository {
	return &UserRepository{
		userCollection:  db.Collection("users"),
		tokenCollection: db.Collection("tokens"),
	}
}

func filterUser(usernameoremail string) bson.M {
	return bson.M{
		"$or": []bson.M{
			{"username": usernameoremail},
			{"email": usernameoremail},
		},
	}
}

func (ur *UserRepository) CheckUsernameAndEmail(username, email string) error {
	var user domain.User
	filter := bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}

	err := ur.userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err == nil {
		return errors.New("username or email already exists")
	}

	if mongo.ErrNoDocuments != err {
		return err
	}

	return nil

}

func (ur *UserRepository) RegisterUser(user *domain.User) error {
	_, err := ur.userCollection.InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByUsernameorEmail(usernameoremail string) (*domain.User, error) {
	var user domain.User
	filter := filterUser(usernameoremail)

	err := ur.userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateProfile(usernameoremail string, user *domain.User) error {
	filter := filterUser(usernameoremail)

	update := bson.M{
		"$set": bson.M{
			"firstname":  user.FirstName,
			"lastname":   user.LastName,
			"bio":        user.Bio,
			"avatar":     user.Avatar,
			"username":   user.Username,
			"email":      user.Email,
			"role":       user.Role,
			"address":    user.Address,
			"joinedDate": user.JoinedDate,
		},
	}

	_, err := ur.userCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}

	return nil

}

func (ur UserRepository) Resetpassword(usernameoremail string, password string) error {
	filter := filterUser(usernameoremail)

	update := bson.M{
		"$set": bson.M{
			"password": password,
		},
	}

	_, err := ur.userCollection.UpdateOne(context.TODO(), filter, update)

	if err == mongo.ErrNoDocuments {
		return errors.New("user not found")
	}

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) InsertToken(token *domain.Token) error {
	_, err := ur.tokenCollection.InsertOne(context.TODO(), token)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetTokenByUsername(username string) (*domain.Token, error) {
	var token domain.Token
	filter := bson.M{
		"username": username,
	}

	err := ur.tokenCollection.FindOne(context.TODO(), filter).Decode(&token)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("token not found")
	}

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (ur *UserRepository) DeleteToken(username string) error {
	filter := bson.M{
		"username": username,
	}

	_, err := ur.tokenCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}

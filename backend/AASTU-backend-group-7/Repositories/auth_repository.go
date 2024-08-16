package Repositories

import (
	"blogapp/Domain"
	"blogapp/Infrastructure"
	jwtservice "blogapp/Infrastructure/jwt_service"
	"context"
	"errors"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type authRepository struct {
	collection Domain.Collection
}

func NewAuthRepository(_collection Domain.Collection) *authRepository {
	return &authRepository{

		collection: _collection,
	}

}

// login
func (au *authRepository) Login(ctx context.Context, user *Domain.User) (string, error, int) {
	filter := bson.D{{"email", user.Email}}
	var existingUser Domain.User
	err := au.collection.FindOne(ctx, filter).Decode(&existingUser)

	if err != nil || !Infrastructure.CompareHashAndPasswordCustom(existingUser.Password, user.Password) {
		fmt.Printf("Login Called:%v, %v", existingUser.Password, user.Password)
		
		// cpmpare the hashed password
		hashedPassword, _ := Infrastructure.GenerateFromPasswordCustom(user.Password)
		fmt.Print(existingUser.Password == hashedPassword)
		return "", errors.New("Invalid credentials"), http.StatusBadRequest
	}

	// Generate JWT
	jwtToken, err := jwtservice.SignJwt(existingUser)
	if err != nil {
		return "", err, 500
	}

	return jwtToken, nil, 200
}

// register
func (au *authRepository) Register(ctx context.Context, user *Domain.User) (*Domain.OmitedUser, error, int) {

	// Check if the email is already taken
	existingUserFilter := bson.D{{"email", user.Email}}
	existingUserCount, err := au.collection.CountDocuments(ctx, existingUserFilter)
	if err != nil {
		fmt.Println("error at count", err)
		return &Domain.OmitedUser{}, err, 500
	}
	if existingUserCount > 0 {
		return &Domain.OmitedUser{}, errors.New("Email is already taken"), http.StatusBadRequest
	}

	// User registration logic
	hashedPassword, err := Infrastructure.GenerateFromPasswordCustom(user.Password)
	if err != nil {
		fmt.Println("error at hashing", err)
		return &Domain.OmitedUser{}, err, 500
	}

	user.Password = string(hashedPassword)
	InsertedID, err := au.collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("error at insert", err)
		return &Domain.OmitedUser{}, err, 500
	}

	// Fetch the inserted task
	var fetched Domain.OmitedUser

	// Access the InsertedID field from the InsertOneResult struct
	insertedID := InsertedID.InsertedID.(primitive.ObjectID)
	
	err = au.collection.FindOne(context.TODO(), bson.D{{"_id", insertedID}}).Decode(&fetched)
	if err != nil {
		fmt.Println(err)
		return &Domain.OmitedUser{}, errors.New("User Not Created"), 500
	}
	if fetched.Email != user.Email {
		return &Domain.OmitedUser{}, errors.New("User Not Created"), 500
	}
	fetched.Password = ""
	return &fetched, nil, 200
}

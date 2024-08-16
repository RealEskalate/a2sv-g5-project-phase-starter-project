package Repositories

import (
	"blogapp/Domain"
	jwtservice "blogapp/Infrastructure/jwt_service"
	"blogapp/Infrastructure/password_services"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

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
func (au *authRepository) Login(ctx context.Context, user *Domain.User) (Domain.Tokens, error, int) {
	filter := bson.D{{"email", user.Email}}
	var existingUser Domain.User
	err := au.collection.FindOne(ctx, filter).Decode(&existingUser)

	if err != nil || !password_services.CompareHashAndPasswordCustom(existingUser.Password, user.Password) {
		fmt.Printf("Login Called:%v, %v", existingUser.Password, user.Password)

		// cpmpare the hashed password
		hashedPassword, _ := password_services.GenerateFromPasswordCustom(user.Password)
		fmt.Print(existingUser.Password == hashedPassword)
		return Domain.Tokens{}, errors.New("Invalid credentials"), http.StatusBadRequest
	}

	// Generate JWT access
	jwtAccessToken, err := jwtservice.CreateAccessToken(existingUser)
	if err != nil {
		return Domain.Tokens{}, err, 500
	}

	jwtRefreshToken, err := jwtservice.CreateRefreshToken(existingUser)
	if err != nil {
		return Domain.Tokens{}, err, 500
	}

	tokens := Domain.Tokens{
		AccessToken:  jwtAccessToken,
		RefreshToken: jwtRefreshToken,
	}
	
	return tokens, nil, 200
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
	hashedPassword, err := password_services.GenerateFromPasswordCustom(user.Password)
	if err != nil {
		fmt.Println("error at hashing", err)
		return &Domain.OmitedUser{}, err, 500
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
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

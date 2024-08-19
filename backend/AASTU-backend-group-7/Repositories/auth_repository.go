package Repositories

import (
	"blogapp/Domain"
	"blogapp/Dtos"
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
	UserCollection  Domain.Collection
	TokenRepository Domain.RefreshRepository
}

func NewAuthRepository(user_collection Domain.Collection, token_collection Domain.Collection) *authRepository {
	return &authRepository{

		UserCollection:  user_collection,
		TokenRepository: NewRefreshRepository(token_collection),
	}

}

// login
func (au *authRepository) Login(ctx context.Context, user *Domain.User) (Domain.Tokens, error, int) {
	filter := bson.D{{"email", user.Email}}
	var existingUser Domain.User
	err := au.UserCollection.FindOne(ctx, filter).Decode(&existingUser)

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

	//check if the refresh token is already stored
	filter = primitive.D{{"_id", existingUser.ID}}
	existingTokenCount, err := au.UserCollection.CountDocuments(ctx, filter)
	fmt.Println("existingTokenCount", existingTokenCount)
	if err != nil {
		fmt.Println("error at count", err)
		return Domain.Tokens{}, err, 500
	}

	if existingTokenCount > 0 {
		// update the refresh token
		err, statusCode := au.TokenRepository.UpdateToken(ctx, jwtRefreshToken, existingUser.ID)
		if err != nil {
			return Domain.Tokens{}, err, statusCode
		}

	} else {
		err, statusCode := au.TokenRepository.StoreToken(ctx, existingUser.ID, jwtRefreshToken)
		if err != nil {
			return Domain.Tokens{}, err, statusCode
		}

	}

	tokens := Domain.Tokens{
		AccessToken:  jwtAccessToken,
		RefreshToken: jwtRefreshToken,
	}
	return tokens, nil, 200
}

// register
func (au *authRepository) Register(ctx context.Context, user *Dtos.RegisterUserDto) (*Domain.OmitedUser, error, int) {

	// Check if the email is already taken
	existingUserFilter := bson.D{}
	if user.UserName != "" {
		existingUserFilter = bson.D{
			{"$or", bson.A{
bson.D{{Key: "email", Value: user.Email}},
				bson.D{{Key: "username", Value: user.UserName}},
			}},
		}
	} else {
		existingUserFilter = bson.D{
			{Key: "email", Value: user.Email},
		}
	}
	existingUserCount, err := au.UserCollection.CountDocuments(ctx, existingUserFilter)
	if err != nil {
		fmt.Println("error at count", err)
		return &Domain.OmitedUser{}, err, 500
	}
	if existingUserCount > 0 {
		return &Domain.OmitedUser{}, errors.New("Email is already taken"), http.StatusBadRequest
	}
	// check if password is following the rules
	err = password_services.CheckPasswordStrength(user.Password)
	if err != nil {
		return &Domain.OmitedUser{}, err, http.StatusBadRequest
	}
	// User registration logic
	hashedPassword, err := password_services.GenerateFromPasswordCustom(user.Password)
	if err != nil {
		fmt.Println("error at hashing", err)
		return &Domain.OmitedUser{}, err, 500
	}

	user.Password = string(hashedPassword)
	user.Role = "user"
	user.CreatedAt = time.Now()

	InsertedID, err := au.UserCollection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("error at insert", err)
		return &Domain.OmitedUser{}, err, 500
	}

	// Fetch the inserted task
	var fetched Domain.OmitedUser

	// Access the InsertedID field from the InsertOneResult struct
	insertedID := InsertedID.InsertedID.(primitive.ObjectID)

	err = au.UserCollection.FindOne(context.TODO(), bson.D{{"_id", insertedID}}).Decode(&fetched)
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

// logout
func (au *authRepository) Logout(ctx context.Context, user_id primitive.ObjectID) (error, int) {
	// delete the refresh token
	err, statusCode := au.TokenRepository.DeleteToken(ctx, user_id)
	if err != nil {
		return err, statusCode
	}
	return nil, 200
}

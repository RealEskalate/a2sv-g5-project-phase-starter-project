package Repositories

import (
	"blogapp/Domain"
	"blogapp/Dtos"
	jwtservice "blogapp/Infrastructure/jwt_service"
	"blogapp/Infrastructure/password_services"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GOOGLE_KEY = "1078174695401-jcvnclinh131aqj2saod6hj35706e7vb.apps.googleusercontent.com"
var GOOGLE_SECRET = "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"
var GoogleCallback = "http://localhost:3000/callback"

type authRepository struct {
	UserCollection  Domain.Collection
	TokenRepository Domain.RefreshRepository
	oauth2Config    oauth2.Config
}

func NewAuthRepository(user_collection Domain.Collection, token_collection Domain.Collection) *authRepository {

	oauth_config := &oauth2.Config{
		ClientID:     GOOGLE_KEY,
		ClientSecret: GOOGLE_SECRET,
		RedirectURL:  GoogleCallback,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	return &authRepository{

		UserCollection:  user_collection,
		TokenRepository: NewRefreshRepository(token_collection),
		oauth2Config:    *oauth_config,
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

	return au.GenerateTokenFromUser(ctx, existingUser)
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
	user.EmailVerified = false
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

// google login
func (au *authRepository) GoogleLogin(ctx context.Context) string {
	url := au.oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return url

}

func (au *authRepository) CallbackHandler(ctx context.Context, code string) (Domain.Tokens, error, int) {
	token, err := au.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return Domain.Tokens{}, errors.New("Couldn't exchange token: "), http.StatusInternalServerError

	}

	client := au.oauth2Config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return Domain.Tokens{}, errors.New("Failed to get user info: "), http.StatusInternalServerError
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return Domain.Tokens{}, errors.New("Failed to decode user info"), http.StatusInternalServerError

	}
	log.Println(userInfo)
	// check if the user is already registered
	filter := bson.D{{"email", userInfo["email"]}}
	var existingUser Domain.User
	err = au.UserCollection.FindOne(ctx, filter).Decode(&existingUser)
	if err != nil {
		// register the user
		user := Dtos.RegisterUserDto{
			Email:          userInfo["email"].(string),
			UserName:       userInfo["name"].(string),
			ProfilePicture: userInfo["picture"].(string),
			EmailVerified:  userInfo["email_verified"].(bool),
		}
		_, err, _ := au.Register(ctx, &user)
		if err != nil {
			return Domain.Tokens{}, err, 500
		}
		err = au.UserCollection.FindOne(ctx, filter).Decode(&existingUser)
		if err != nil {
			return Domain.Tokens{}, err, 500
		}

	}
	return au.GenerateTokenFromUser(ctx, existingUser)

}

func (au *authRepository) GenerateTokenFromUser(ctx context.Context, existingUser Domain.User) (Domain.Tokens, error, int) {
	filter := bson.D{{Key: "email", Value: existingUser.Email}}
	// Generate JWT access
	jwtAccessToken, err := jwtservice.CreateAccessToken(existingUser)
	if err != nil {
		return Domain.Tokens{}, err, 500
	}
	refreshToken, err := jwtservice.CreateRefreshToken(existingUser)
	if err != nil {
		return Domain.Tokens{}, err, 500
	}

	filter = primitive.D{{"_id", existingUser.ID}}
	existingTokenCount, err := au.UserCollection.CountDocuments(ctx, filter)
	fmt.Println("existingTokenCount", existingTokenCount)
	if err != nil {
		fmt.Println("error at count", err)
		return Domain.Tokens{}, err, 500
	}

	if existingTokenCount > 0 {
		// update the refresh token
		err, statusCode := au.TokenRepository.UpdateToken(ctx, refreshToken, existingUser.ID)
		if err != nil {
			return Domain.Tokens{}, err, statusCode
		}

	} else {
		err, statusCode := au.TokenRepository.StoreToken(ctx, existingUser.ID, refreshToken)
		if err != nil {
			return Domain.Tokens{}, err, statusCode
		}
	}

	return Domain.Tokens{
		AccessToken:  jwtAccessToken,
		RefreshToken: refreshToken,
	}, nil, 200
}

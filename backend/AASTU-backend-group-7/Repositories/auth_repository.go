package Repositories

import (
	"blogapp/Config"
	"blogapp/Domain"
	"blogapp/Dtos"
	emailservice "blogapp/Infrastructure/email_service"
	jwtservice "blogapp/Infrastructure/jwt_service"
	"blogapp/Infrastructure/password_services"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type authRepository struct {
	validator       *validator.Validate
	UserCollection  Domain.Collection
	TokenRepository Domain.RefreshRepository
	oauth2Config    oauth2.Config
	userRepository  Domain.UserRepository
	emailservice    emailservice.MailTrapService
	mu              sync.RWMutex
}

func NewAuthRepository(user_collection Domain.Collection, token_collection Domain.Collection, userRepository Domain.UserRepository) *authRepository {

	oauth_config := &oauth2.Config{
		ClientID:     Config.GOOGLE_KEY,
		ClientSecret: Config.GOOGLE_SECRET,
		RedirectURL:  Config.Google_Callback,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	return &authRepository{
		validator:       validator.New(),
		UserCollection:  user_collection,
		TokenRepository: NewRefreshRepository(token_collection),
		oauth2Config:    *oauth_config,
		emailservice:    emailservice.NewMailService(),
		userRepository:  userRepository,
		mu:              sync.RWMutex{},
	}
}

// login
func (ar *authRepository) Login(ctx context.Context, user *Domain.User) (Domain.Tokens, error, int) {
	ar.mu.RLock()
	defer ar.mu.RUnlock()

	filter := bson.D{{"email", user.Email}}
	var existingUser Domain.User
	err := ar.UserCollection.FindOne(ctx, filter).Decode(&existingUser)

	if err != nil || !password_services.CompareHashAndPasswordCustom(existingUser.Password, user.Password) {
		fmt.Printf("Login Called:%v, %v", existingUser.Password, user.Password)

		// cpmpare the hashed password
		hashedPassword, _ := password_services.GenerateFromPasswordCustom(user.Password)
		fmt.Print(existingUser.Password == hashedPassword)
		return Domain.Tokens{}, errors.New("Invalid credentials"), http.StatusBadRequest
	}
	fmt.Println("emailverified :", existingUser.EmailVerified, "email", existingUser.Email)
	// if existingUser.EmailVerified == false {
	// 	err, statusCode := ar.SendActivationEmail(user.Email)
	// 	if err != nil {
	// 		fmt.Println("error at sending email", err)
	// 		return Domain.Tokens{}, err, statusCode
	// 	}
	// 	return Domain.Tokens{}, errors.New("email is not activated , an activation email has been sent"), http.StatusUnauthorized
	// }
	return ar.GenerateTokenFromUser(ctx, existingUser)

}

// register
func (ar *authRepository) Register(ctx context.Context, user *Dtos.RegisterUserDto) (*Domain.OmitedUser, error, int) {
	ar.mu.RLock()
	defer ar.mu.RUnlock()

	// Validate the user input
	err := ar.validator.Struct(user)
	if err != nil {
		return &Domain.OmitedUser{}, err, http.StatusBadRequest
	}
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
	existingUserCount, err := ar.UserCollection.CountDocuments(ctx, existingUserFilter)
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
	user.UpdatedAt = time.Now()

	InsertedID, err := ar.UserCollection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("error at insert", err)
		return &Domain.OmitedUser{}, err, 500
	}

	// Fetch the inserted task
	var fetched Domain.OmitedUser

	// Access the InsertedID field from the InsertOneResult struct
	insertedID := InsertedID.InsertedID.(primitive.ObjectID)

	err = ar.UserCollection.FindOne(context.TODO(), bson.D{{"_id", insertedID}}).Decode(&fetched)
	if err != nil {
		fmt.Println(err)
		return &Domain.OmitedUser{}, errors.New("User Not Created"), 500
	}
	if fetched.Email != user.Email {
		return &Domain.OmitedUser{}, errors.New("User Not Created"), 500
	}
	fetched.Password = ""
	// err, statusCode := ar.SendActivationEmail(fetched.Email)
	// if err != nil {
	// 	return &fetched, err, statusCode
	// }
	return &fetched, err, 200
}

// logout
func (ar *authRepository) Logout(ctx context.Context, user_id primitive.ObjectID) (error, int) {
	ar.mu.RLock()
	defer ar.mu.RUnlock()

	// delete the refresh token
	err, statusCode := ar.TokenRepository.DeleteToken(ctx, user_id)
	if err != nil {
		return err, statusCode
	}
	return nil, 200
}

func (ar *authRepository) ForgetPassword(ctx context.Context, email string) (error, int) {
	ar.mu.RLock()
	defer ar.mu.RUnlock()
	_, err, status := ar.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return err, status
	}
	resetToken, err := jwtservice.GenerateToken(email)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = ar.emailservice.SendEmail(email, "Reset Password", `Click "http://localhost:8080/auth/forget-password/`+resetToken+`">hereto reset your password.
`, "reset")
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusOK

}

func (ar *authRepository) ResetPassword(ctx context.Context, email string, password string, resetToken string) (error, int) {
	ar.mu.RLock()
	defer ar.mu.RUnlock()
	_, err := jwtservice.VerifyToken(resetToken)
	if err != nil {
		return err, http.StatusBadRequest
	}
	if password == "" {
		return errors.New("password is required"), http.StatusBadRequest
	}
	err = password_services.CheckPasswordStrength(password)
	if err != nil {
		return err, http.StatusBadRequest
	}
	hashed, err := password_services.GenerateFromPasswordCustom(password)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	_, err, _ = ar.userRepository.ChangePassByEmail(ctx, email, hashed)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	fmt.Println("password:", password, "reset_token", resetToken)
	return nil, http.StatusOK

}

// google login
func (ar *authRepository) GoogleLogin(ctx context.Context) string {
	ar.mu.RLock()
	defer ar.mu.RUnlock()
	url := ar.oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return url

}

func (ar *authRepository) CallbackHandler(ctx context.Context, code string) (Domain.Tokens, error, int) {
	ar.mu.RLock()
	defer ar.mu.RUnlock()

	token, err := ar.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return Domain.Tokens{}, errors.New("Couldn't exchange token: "), http.StatusInternalServerError

	}

	client := ar.oauth2Config.Client(ctx, token)
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
	err = ar.UserCollection.FindOne(ctx, filter).Decode(&existingUser)
	if err != nil {
		// register the user
		user := Dtos.RegisterUserDto{
			Name:           userInfo["name"].(string),
			Email:          userInfo["email"].(string),
			UserName:       userInfo["name"].(string),
			ProfilePicture: userInfo["picture"].(string),
			EmailVerified:  userInfo["email_verified"].(bool),
			Password:       "Test!2" + userInfo["sub"].(string),
		}
		_, err, _ := ar.Register(ctx, &user)
		if err != nil {
			return Domain.Tokens{}, err, 500
		}
		err = ar.UserCollection.FindOne(ctx, filter).Decode(&existingUser)
		if err != nil {
			return Domain.Tokens{}, err, 500
		}

	}
	return ar.GenerateTokenFromUser(ctx, existingUser)

}

func (ar *authRepository) GenerateTokenFromUser(ctx context.Context, existingUser Domain.User) (Domain.Tokens, error, int) {
	ar.mu.RLock()
	defer ar.mu.RUnlock()

	// filter := bson.D{{Key: "email", Value: existingUser.Email}}
	// Generate JWT access
	jwtAccessToken, err := jwtservice.CreateAccessToken(existingUser)
	if err != nil {
		return Domain.Tokens{}, err, 500
	}
	refreshToken, err := jwtservice.CreateRefreshToken(existingUser)
	if err != nil {
		return Domain.Tokens{}, err, 500
	}

	// filter := primitive.D{{"_id", existingUser.ID}}
	existingToken, err, statusCode := ar.TokenRepository.FindToken(ctx, existingUser.ID)
	if err != nil && err.Error() != "mongo: no documents in result" {
		fmt.Println("error at count", err)
		return Domain.Tokens{}, err, statusCode
	}

	if existingToken != "" {
		// update the refresh token
		err, statusCode := ar.TokenRepository.UpdateToken(ctx, refreshToken, existingUser.ID)
		if err != nil {
			return Domain.Tokens{}, err, statusCode
		}

	} else {
		err, statusCode := ar.TokenRepository.StoreToken(ctx, existingUser.ID, refreshToken)
		if err != nil {
			return Domain.Tokens{}, err, statusCode
		}
	}

	return Domain.Tokens{
		AccessToken:  jwtAccessToken,
		RefreshToken: refreshToken,
	}, nil, 200
}

func (ar *authRepository) ActivateAccount(ctx context.Context, token string) (error, int) {
	ar.mu.RLock()
	defer ar.mu.RUnlock()
	email, err := jwtservice.VerifyToken(token)
	if err != nil {
		return err, http.StatusBadRequest
	}
	fmt.Println("email:", email, "token:", token)

	filter := bson.D{{"email", email}}

	update := bson.D{
		{"$set", bson.D{
			{"email_verified", true},
			{"created_at", time.Now()},
		}},
	}
	UpdatedResult, err := ar.UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if UpdatedResult.ModifiedCount == 0 {
		return errors.New("user does not exist"), 400
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", UpdatedResult.MatchedCount, UpdatedResult.ModifiedCount)

	return nil, http.StatusOK

}

func (ar *authRepository) SendActivationEmail(email string) (error, int) {
	ar.mu.RLock()
	defer ar.mu.RUnlock()

	activationToken, err := jwtservice.GenerateToken(email)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = ar.emailservice.SendEmail(email, "Verify Email", `Click "`+Config.BASE_URL+`/auth/activate/`+activationToken+`"here to verify email.
`, "reset")
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

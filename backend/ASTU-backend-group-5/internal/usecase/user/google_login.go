package user

import (
	"blogApp/internal/config"
	"blogApp/internal/domain"
	"blogApp/pkg/jwt"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GoogleCallback handles Google OAuth callback
func (u *UserUsecase) GoogleCallback(code string) (*domain.User, *domain.Token, error) {
	googleConfig := config.GoogleConfig()

	// Exchange the authorization code for an access token
	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error exchanging token:", err)
		return nil, nil, errors.New("invalid credentials")
	}

	// Fetch user info from Google
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("Error fetching user info:", err)
		return nil, nil, errors.New("something went wrong")
	}
	defer resp.Body.Close()

	// Parse user info
	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading user data:", err)
		return nil, nil, errors.New("JSON Parsing Failed")
	}

	googleUser := domain.GoogleUser{}
	err = json.Unmarshal(userData, &googleUser)
	if err != nil {
		fmt.Println("Error unmarshaling user data:", err)
		return nil, nil, errors.New("JSON Parsing Failed")
	}

	if googleUser.Email == "" {
		return nil, nil, errors.New("invalid credentials")
	}

	// Prepare user data
	user := &domain.User{
		Email: googleUser.Email,
		Role:  "user",
		Profile: domain.UserProfile{
			FirstName:  googleUser.GivenName,
			LastName:   googleUser.FamilyName,
			ProfileUrl: googleUser.Picture,
		},
	}

	// Check if user exists in the database
	dbUser, err := u.repo.FindUserByEmail(context.TODO(), googleUser.Email)
	if err != nil {
		return nil, nil, err
	}

	// Create a new user if not found
	if dbUser == nil {
		user.ID = primitive.NewObjectID()
		randomPassword, err := generateRandomPassword(30)
		if err != nil {
			return nil, nil, err
		}
		user.Password = randomPassword
		err = u.repo.CreateUser(context.TODO(), user)
		if err != nil {
			return nil, nil, err
		}
		dbUser = user
	}

	// Generate JWT tokens
	accessToken, err := jwt.GenerateJWT(dbUser.ID.Hex(), dbUser.UserName, dbUser.Email, dbUser.Role)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(dbUser.ID.Hex(), dbUser.UserName, dbUser.Email, dbUser.Role)
	if err != nil {
		return nil, nil, err
	}

	return dbUser, &domain.Token{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

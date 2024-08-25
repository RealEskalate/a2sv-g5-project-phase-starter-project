package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const googleTokenInfoURL = "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%s"

// UserInfo struct represents the user information obtained from the federated identity provider
type UserInfo struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	ProfilePicUrl string `json:"profilePicUrl"`
}

func VerifyFederatedToken(token, googleClientID string) (*UserInfo, error) {
	// Verify the token with the Google API
	userInfo, err := VerifyGoogleToken(token, googleClientID)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func VerifyGoogleToken(idToken, googleClientID string) (*UserInfo, error) {
	// TokenInfo struct represents the response from Google's tokeninfo endpoint
	type TokenInfo struct {
		Audience string `json:"aud"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Picture  string `json:"picture"`
		// Add other relevant fields as needed
	}
	// Send a request to Google's tokeninfo endpoint to verify the token
	resp, err := http.Get(fmt.Sprintf(googleTokenInfoURL, idToken))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to verify Google token")
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var tokenInfo TokenInfo
	err = json.Unmarshal(body, &tokenInfo)
	if err != nil {
		return nil, err
	}

	// Verify audience and other claims as needed
	if tokenInfo.Audience != googleClientID {
		return nil, errors.New("invalid audience in Google token")
	}

	// Extract user information from the token
	userInfo := &UserInfo{
		Name:          tokenInfo.Name,
		Email:         tokenInfo.Email,
		ProfilePicUrl: tokenInfo.Picture,
	}
	return userInfo, nil
}

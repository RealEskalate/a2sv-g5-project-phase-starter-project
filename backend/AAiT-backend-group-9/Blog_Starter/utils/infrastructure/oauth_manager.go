package infrastructure

import (
    "Blog_Starter/utils"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
)

const googleTokenInfoURL = "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%s"

// NewOAuthManager is the implementation of the OAuthManager interface
type NewOAuthManager struct{}

// VerifyFederatedToken verifies the federated token with the Google API
func (o *NewOAuthManager) VerifyFederatedToken(token, googleClientID string) (*utils.UserInfo, error) {
    userInfo, err := o.VerifyGoogleToken(token, googleClientID)
    if err != nil {
        return nil, err
    }

    return userInfo, nil
}

// VerifyGoogleToken verifies the Google token and extracts user information
func (o *NewOAuthManager) VerifyGoogleToken(idToken, googleClientID string) (*utils.UserInfo, error) {
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
    userInfo := &utils.UserInfo{
        Name:          tokenInfo.Name,
        Email:         tokenInfo.Email,
        ProfilePicUrl: tokenInfo.Picture,
    }
    return userInfo, nil
}
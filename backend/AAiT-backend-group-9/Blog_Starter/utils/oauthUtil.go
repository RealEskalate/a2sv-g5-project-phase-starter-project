package utils

// UserInfo struct represents the user information obtained from the federated identity provider
type UserInfo struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	ProfilePicUrl string `json:"profilePicUrl"`
}

// OAuthManager defines the methods for verifying federated tokens
type OAuthManager interface {
	VerifyFederatedToken(token, googleClientID string) (*UserInfo, error)
	VerifyGoogleToken(idToken, googleClientID string) (*UserInfo, error)
}

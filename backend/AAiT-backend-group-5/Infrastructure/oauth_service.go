package infrastructure

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	config "github.com/aait.backend.g5.main/backend/Config"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"

	"golang.org/x/oauth2/google"
)

type OAuthService struct {
	env      config.Env
	userRepo interfaces.UserRepository
}

func NewOAuthService(env config.Env, userRepo interfaces.UserRepository) interfaces.OAuthService {
	return &OAuthService{
		env:      env,
		userRepo: userRepo,
	}
}

func (os *OAuthService) OAuthTokenValidator(token string, ctx context.Context) (*models.JWTCustome, *models.ErrorResponse) {
	tokenInfo, err := os.validateGoogleToken(token)
	if err != nil {
		return nil, err
	}

	email, ok := tokenInfo["email"].(string)
	if !ok {
		return nil, models.Unauthorized("Email not found in token info")
	}

	user, uErr := os.getUser(email, ctx)
	if uErr != nil {
		return nil, uErr
	}

	return &models.JWTCustome{
		ID:    user.ID,
		Email: user.Email,
		Role:  string(user.Role),
	}, nil
}

func (os *OAuthService) RefreshTokenValidator(refreshToken string, ctx context.Context) (*models.JWTCustome, *models.ErrorResponse) {
	tokenInfo, err := os.validateGoogleRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	access_token, ok := tokenInfo["access_token"].(string)

	if !ok {
		return nil, models.Unauthorized("Something went wrong")
	}

	user, uErr := os.OAuthTokenValidator(access_token, ctx)

	if uErr != nil {
		return nil, uErr
	}
	return user, nil
}

func (os *OAuthService) validateGoogleToken(token string) (map[string]interface{}, *models.ErrorResponse) {
	googleTokenInfoURL := os.env.GOOGLE_TOKEN_INFO_URL

	resp, err := http.Get(googleTokenInfoURL + token)
	if err != nil {
		return nil, models.InternalServerError("Error validating token: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, models.Unauthorized("Invalid token: status code " + http.StatusText(resp.StatusCode))
	}

	var tokenInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, models.InternalServerError("Error decoding token info: " + err.Error())
	}

	return tokenInfo, nil
}

func (os *OAuthService) validateGoogleRefreshToken(refreshToken string) (map[string]interface{}, *models.ErrorResponse) {
	tokenURL := google.Endpoint.TokenURL
	// "https://oauth2.googleapis.com/token"

	form := url.Values{}
	form.Set("client_id", os.env.OAUTH_CLIENT_ID)
	form.Set("client_secret", os.env.OAUTH_CLIENT_SECRET)
	form.Set("refresh_token", refreshToken)
	form.Set("grant_type", "refresh_token")

	resp, err := http.PostForm(tokenURL, form)
	if err != nil {
		return nil, models.InternalServerError("Error validating refresh token: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, models.Unauthorized("Invalid refresh token: status code " + http.StatusText(resp.StatusCode))
	}

	var tokenInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, models.InternalServerError("Error decoding token info: " + err.Error())
	}

	return tokenInfo, nil
}

func (os *OAuthService) getUser(email string, ctx context.Context) (*models.User, *models.ErrorResponse) {
	user, err := os.userRepo.GetUserByEmailOrUsername(ctx, email, email)
	if err != nil {
		return nil, models.InternalServerError("Error retrieving user: " + err.Error())
	}
	return user, nil
}

func (os *OAuthService) GenerateAccessToken(ctx context.Context, refreshToken string) (string, *models.ErrorResponse) {
	clientID := os.env.GOOGEL_CLIENT_ID
	clientSecret := os.env.OAUTH_CLIENT_SECRET
	tokenURL := google.Endpoint.TokenURL

	resp, err := http.PostForm(tokenURL, url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"refresh_token": {refreshToken},
		"grant_type":    {"refresh_token"},
	})
	if err != nil {
		return "", models.InternalServerError("Error generating access token: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", models.Unauthorized("Failed to generate access token: " + resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", models.InternalServerError("Error decoding token info: " + err.Error())
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", models.InternalServerError("Access token not found in response")
	}

	return accessToken, nil
}

//

// map[
// access_token:ya29.a0AcM612z8wwfjfMRSjoKStM7js7p7H49sZ5vsI-YQR1v644xi56XJBgcPwBFdLgQhVMkRU-PlJEbOV8QpmKpVJzXN2b15E2M_VHNVGawPVikGKDGtrwg3V0JLbhDfp2dMlg9GTYnwzeeSpVAhazw4odue2XYXsC4CFenYna1DaCgYKAZcSARASFQHGX2Mitv-H1fRiZjGgANd35E9rjQ0175
// expires_in:3599
// id_token:eyJhbGciOiJSUzI1NiIsImtpZCI6ImE0OTM5MWJmNTJiNThjMWQ1NjAyNTVjMmYyYTA0ZTU5ZTIyYTdiNjUiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI4NzYwMDc1OTQ3MDEtbWNtdDMwbmhsbTUzOHA5bzM2N25rZzE1aDdlMWc0cGwuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI4NzYwMDc1OTQ3MDEtbWNtdDMwbmhsbTUzOHA5bzM2N25rZzE1aDdlMWc0cGwuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDQxODcwNDQ4MTgxNzA4OTQ3MTIiLCJoZCI6ImFhdS5lZHUuZXQiLCJlbWFpbCI6ImRhd2l0LmFicmFoYW1AYWF1LmVkdS5ldCIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJhdF9oYXNoIjoiTncyQXIzUWduaktrMmY0RVUzeUVvUSIsIm5hbWUiOiJEYXdpdCBBYnJhaGFtIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0ppamgxMFROWjBmRWJvbUItOUF3blQ3NUtkSUI4U0Joc21ZanZvWENPNFFIWVQ4aWc9czk2LWMiLCJnaXZlbl9uYW1lIjoiRGF3aXQiLCJmYW1pbHlfbmFtZSI6IkFicmFoYW0iLCJpYXQiOjE3MjQ0MDk0ODUsImV4cCI6MTcyNDQxMzA4NX0.J_AolhNxhFkgzJhh3aWlGrlu_anp-hkdR5LJFiPtqQB1G5FN1RyfieZWsBWImo4ZEUPZUcdKQJdvJGmT00r23E0EG410dvMWc-oB1ut2jp5JHPqT2o0CiRLWaowo5RTf4PCkYzIQPebsauM8I6tkoo7PaD3Ghk4_EqGDVDXOi_8pOI2x40M_rfZ_eWvgOFrQIwiqQBkK_B_24LZNWPRmr-Xjv_WQv2zWPbAoElxtpGbH0tLJ28xoEM-jia2fnzvD_JB8__p-pax-L0UaMhQqsP4R0qekc3h5isVervbjx0M-MnNpd-PwWGYSAkSU-QeITt6NYDX40t26d0FaMqZWmA
// scope:https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile openid token_type:Bearer]
//
//

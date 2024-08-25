package controller

import (
	"Blog_Starter/domain"
	"Blog_Starter/config"
	EmailUtil "Blog_Starter/utils"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpController struct {
	signUpUsecase domain.SignupUsecase
	otpUsecase    domain.OtpUsecase
	Env           *config.Env
}

func NewSignUpController(signUpUsecase domain.SignupUsecase, otpUsecase domain.OtpUsecase, env *config.Env) *SignUpController {
	return &SignUpController{
		signUpUsecase: signUpUsecase,
		otpUsecase:    otpUsecase,
		Env:           env, 
	}
}

func (s *SignUpController) SignUp(c *gin.Context) {
    var UserSignUp domain.UserSignUp

    if err := c.ShouldBindJSON(&UserSignUp); err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: "Invalid request",
        })
        return
    }

    err := UserSignUp.Validate()
    if err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    user, err := s.signUpUsecase.CreateUser(c, &UserSignUp)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    // Generate a random number between 0 and 9999 (inclusive).
    randNumber := rand.Intn(10000)

    // Format the code as a 4-digit string with leading zeros.
    code := fmt.Sprintf("%04d", randNumber)

    otp := domain.Otp{
        Email:      user.Email,
        Otp:        code,
        Expiration: time.Now().Add(5 * time.Minute),
    }

    oldOtp, err := s.otpUsecase.GetOtpByEmail(c, UserSignUp.Email)
    if err == nil {
        otp.ID = oldOtp.ID
    } else {
        otp.ID = primitive.NewObjectID()
    }

    // Save OTP to database
    err = s.otpUsecase.SaveOtp(c, &otp)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    emailContent := `
    <p>Thank you for signing up. To verify your account and complete the signup process, please use the following verification code:</p>
    <h3>` + code + `</h3>
    <p><strong>This verification code is valid for 5 minutes.</strong> Please enter it on the verification page to proceed.</p>
    <p>If you did not sign up for the BlogApp, please ignore this email.</p>`

    // Create the email subject
    emailSubject := "Verify Your Email"

    // Generate the email body using the template function
    emailBody := EmailUtil.GenerateEmailTemplate("Account Verification", emailContent)

    // Create the email template
    emailTemplate := domain.EmailTemplate{
        Subject: emailSubject,
        Body:    emailBody,
    }

    err = EmailUtil.SendTestEmail(UserSignUp.Email, emailTemplate.Subject, emailTemplate.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, domain.Response{
        Success: true,
        Message: "Successfully sent OTP",
    })
}

func (s *SignUpController) VerifyEmail(c *gin.Context) {
    var VerifyEmailRequest domain.VerifyEmailRequest

    if err := c.ShouldBindJSON(&VerifyEmailRequest); err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: "Invalid request",
        })
        return
    }
    VerifyEmailRequest.Email = strings.ToLower(VerifyEmailRequest.Email)

    otp, err := s.otpUsecase.GetOtpByEmail(c, VerifyEmailRequest.Email)
    if err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: "No OTP requested with the given email",
        })
        return
    }

    if otp.Otp != VerifyEmailRequest.OTP {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: "Invalid OTP",
        })
        return
    } else if time.Now().After(otp.Expiration) {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: "OTP expired",
        })
        return
    }

    user, err := s.signUpUsecase.VerifyEmail(c, &VerifyEmailRequest)
    if err != nil {
        if err.Error() == "mongo: no documents in result" {
            c.JSON(http.StatusNotFound, domain.Response{
                Success: false,
                Message: "User not found",
            })
            return
        }
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, domain.Response{
        Success: true,
        Message: "Email verified",
        Data:    user,
    })
}

func (s *SignUpController) ResendOTP(c *gin.Context) {
    var ResendOTPRequest domain.ResendOTPRequest

    if err := c.ShouldBindJSON(&ResendOTPRequest); err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: "Invalid request",
        })
        return
    }

    ResendOTPRequest.Email = strings.ToLower(ResendOTPRequest.Email)

    err := s.signUpUsecase.ResendOTP(c, &ResendOTPRequest)
    if err != nil {
        if err.Error() == "mongo: no documents in result" {
            c.JSON(http.StatusNotFound, domain.Response{
                Success: false,
                Message: "User not found",
            })
            return
        }
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    otp, err := s.otpUsecase.GetOtpByEmail(c, ResendOTPRequest.Email)
    if err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: "No OTP requested with the given email",
        })
        return
    }

    // Generate a random number between 0 and 9999 (inclusive).
    randNumber := rand.Intn(10000)

    // Format the code as a 4-digit string with leading zeros.
    code := fmt.Sprintf("%04d", randNumber)

    otp.Otp = code
    otp.Expiration = time.Now().Add(5 * time.Minute)

    // Save OTP to database
    err = s.otpUsecase.SaveOtp(c, &otp)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    emailContent := `
    <p>Thank you for signing up. To verify your account and complete the signup process, please use the following verification code:</p>
    <h3>` + code + `</h3>
    <p><strong>This verification code is valid for 5 minutes.</strong> Please enter it on the verification page to proceed.</p>
    <p>If you did not sign up for the BlogApp, please ignore this email.</p>`

    // Create the email subject
    emailSubject := "Verify Your Email"

    // Generate the email body using the template function
    emailBody := EmailUtil.GenerateEmailTemplate("Account Verification", emailContent)

    // Create the email template
    emailTemplate := domain.EmailTemplate{
        Subject: emailSubject,
        Body:    emailBody,
    }

    err = EmailUtil.SendTestEmail(ResendOTPRequest.Email, emailTemplate.Subject, emailTemplate.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, domain.Response{
        Success: true,
        Message: "Successfully sent OTP",
    })
}



func (s *SignUpController) FederatedSignup(c *gin.Context) {
	var request domain.FederatedSignupRequest

	// Bind the request body to the FederatedSignupRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: errorsutil.MapErrors(err)})
		return
	}

	// Validate and handle the federated authentication token
	if request.Provider == "google" {
		// Verify the Google authentication token (you might use a library or make a request to Google's API for verification)
		// ...

		// Assuming verification is successful, store or update the user in your database
		user, err := s.handleFederatedSignup(c, request.Token, request.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.Response{Message: errorsutil.MapErrors(err)})
			return
		}

		//Set user object
		c.Set("user", user)

		accessToken, err := s.SignupUsecase.CreateAccessToken(&user, s.Env.AccessTokenSecret, s.Env.AccessTokenExpiryHour)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.Response{Message: errorsutil.MapErrors(err)})
			return
		}

		refreshToken, err := s.SignupUsecase.CreateRefreshToken(&user, s.Env.RefreshTokenSecret, s.Env.RefreshTokenExpiryHour)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.Response{Message: errorsutil.MapErrors(err)})
			return
		}

		signupResponse := domain.LoginResponse{
			AccessToken:     accessToken,
			RefreshToken:    refreshToken,
			ID:              user.ID.Hex(),
			Name:            user.FullName,
			Email:           user.Email,
			ProfilePicUrl:   user.ProfilePicUrl,
			Role:            user.Role,
			ProfileComplete: user.ProfileComplete,
			ProfileStatus:   user.ProfileStatus,
		}

		c.JSON(http.StatusOK, domain.Response{
			Success: true,
			Data:    signupResponse,
		})

	} else {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "Invalid federated provider"})
	}
}

func (s *SignUpController) handleFederatedSignup(c *gin.Context, token, role string) (domain.User, error) {
	// Verify the token with the federated identity provider (e.g., Google)
	userInfo, err := verifyFederatedToken(token, s.Env.GoogleClientID)
	if err != nil {
		return domain.User{}, err
	}

	userInfo.Email = strings.ToLower(userInfo.Email)

	// Check if the user already exists in the database
	existingUser, err := s.SignupUsecase.GetUserByEmail(c, userInfo.Email)
	if err != nil {
		// User doesn't exist, create a new user
		newUser := domain.User{
			ID:            primitive.NewObjectID(),
			FullName:      userInfo.Name,
			Email:         userInfo.Email,
			Password:      "",   // You may leave this empty or handle it differently for federated signup
			Role:          role, // Set a default role or customize based on your application
			CreatedAt:     time.Now(),
			IsActivated:   true,                   // Assuming the user is activated upon federated signup
			ProfileStatus: enumutil.Incomplete,    // Set a default profile status or customize based on your application
			ProfilePicUrl: userInfo.ProfilePicUrl, // Assuming the federated provider provides a profile picture URL
		}

		// Save the new user to the database
		err := s.SignupUsecase.Create(c, &newUser)
		if err != nil {
			return domain.User{}, err
		}

		return newUser, nil
	}

	return existingUser, nil
}

// Function to verify the federated authentication token (e.g., Google)

// UserInfo struct represents the user information obtained from the federated identity provider
type UserInfo struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	ProfilePicUrl string `json:"profilePicUrl"`
}

func verifyFederatedToken(token, googleClientID string) (*UserInfo, error) {
	// Verify the token with the Google API
	userInfo, err := verifyGoogleToken(token, googleClientID)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func verifyGoogleToken(idToken, googleClientID string) (*UserInfo, error) {
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


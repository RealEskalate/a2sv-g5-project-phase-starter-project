package usecases

import (
	domain "blogs/Domain"
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	oauth2Service "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)
type OauthUseCase struct {
	signupRepository domain.SignupRepository
	contextTimeout   time.Duration
	oauthService     domain.OauthConfig
}


func NewOauthUsecase(signupRepository domain.SignupRepository, timeout time.Duration , oauth domain.OauthConfig) domain.OauthUsecase {
	return &OauthUseCase{
		signupRepository: signupRepository,
		contextTimeout:   timeout,
		oauthService:     oauth,}
}


func (u *OauthUseCase) OauthService() (interface{}) {
	

	// Call InitialConfig and handle potential errors
	credentials, err := u.oauthService.InitialConfig()
	if err != nil {
		return &domain.ErrorResponse{Message: "Error generating OAuth configuration", Status: 500}
	}

	// Generate the AuthCode URL
	url := credentials.AuthCodeURL("oauthStateString")
	return &domain.URL{URL: url} 
}

func (u *OauthUseCase)  OauthCallback(c context.Context , query string) (interface{}) {
	credentials, err := u.oauthService.InitialConfig()

	if err != nil {
		return &domain.ErrorResponse{Message: "Error generating OAuth configuration", Status: 500}
	}


	ctx , cancel := context.WithTimeout(c , u.contextTimeout)
	defer cancel()
	token, err := credentials.Exchange(ctx, query)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error exchanging token", Status: 500}
	}


	client := credentials.Client(ctx, token)

	oauth2Service, err := oauth2Service.NewService(ctx, option.WithHTTPClient(client))

	if err != nil {	
		return &domain.ErrorResponse{Message: "Error creating OAuth2 service", Status: 500}
	}

	userinfo , err := oauth2Service.Userinfo.V2.Me.Get().Do()

	if err != nil { 
		return &domain.ErrorResponse{Message: "Error getting user info", Status: 500}
	}

	

	// check if the email already exists
	existingUser , err := u.signupRepository.FindUserByEmail(ctx, userinfo.Email)


	if err != nil {
		var userData domain.User

		userData.Email = userinfo.Email
		userData.Full_Name = userinfo.Name
		userData.Profile_image_url = userinfo.Picture
		userData.GoogleID = userinfo.Id
		userData.Role = "user"
		userData.Verified = true
		userData.ID = primitive.NewObjectID()
		// spliting with @
		userData.Username = userinfo.Email[:strings.Index(userinfo.Email, "@")]


		createdUser, err := u.signupRepository.Create(ctx, userData)

		if err != nil {
			return &domain.ErrorResponse{Message: "Error creating user", Status: 500}
			
		} 
		
		return &domain.UserResponse{User: createdUser}
		
		
	}



	return &domain.UserResponse{User: existingUser}
	

}
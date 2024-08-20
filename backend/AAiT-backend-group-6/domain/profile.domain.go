package domain

import "context"

type Profile struct {
	Name  	 string `json:"name"`
	Email 	 string `json:"email"`
	Username string `json:"username"`
	Bio 	 string `json:"bio"`

}

type ProfileRepository interface{
	GetProfileByID(c context.Context, userID string) (*Profile, error)
	UpdateProfile(c context.Context, userID string, profile *Profile) error
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*Profile, error)
	UpdateProfile(c context.Context, userID string, profile *Profile) error
}
package interfaces

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/entities"
	"mime/multipart"
)

type ProfileRepository interface {
	GetAllProfiles() ([]*entities.Profile, error)
	GetUserProfile(userId string) (*entities.Profile, error)
	UpdateUserProfile(profile *entities.Profile) (*entities.Profile, error)
	CreateUserProfile(profile *entities.Profile) (*entities.Profile, error)
	DeleteUserProfile(user_id string) error
	UpdateProfilePicture(user_id,path string) error
	GetProfilePicture(user_id string) (string,error)
	DeleteProfilePicture(user_id string) error
}

type ProfileService interface {
	GetAllProfiles() ([]*entities.Profile, error)
	GetUserProfile(userId string) (*dto.ProfileResponse, error)
	UpdateUserProfile(profile *dto.UpdateProfileDto) (*dto.ProfileResponse, error)
	CreateUserProfile(profile *dto.CreateProfileDto) (*dto.ProfileResponse, error)
	DeleteUserProfile(user_id string) error
	UpdateProfilePicture(user_id string,file *multipart.FileHeader) (string,error)
	GetProfilePicture(user_id string) (string,error)
	DeleteProfilePicture(user_id string) error
}

type ImageService interface {
	UploadImage(file *multipart.FileHeader) (string, error)
	DeleteImage(path string) error
}
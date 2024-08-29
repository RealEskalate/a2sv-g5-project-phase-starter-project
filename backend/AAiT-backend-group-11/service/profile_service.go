package service

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type profileService struct {
	Profile_repo interfaces.ProfileRepository
	imageService interfaces.ImageService
}

func NewProfileService(repo interfaces.ProfileRepository,imageS interfaces.ImageService) interfaces.ProfileService {
	return &profileService{Profile_repo: repo,imageService: imageS}
}
func (profileService *profileService) GetAllProfiles() ([]*entities.Profile, error) {
	profiles,err:=profileService.Profile_repo.GetAllProfiles()
	if err!=nil{
		return nil,err
	}
	return profiles,nil
} 

func (profileService *profileService) GetUserProfile(userId string) (*dto.ProfileResponse, error) {
	profile, err := profileService.Profile_repo.GetUserProfile(userId)
	if err != nil {
		return nil, err
	}
	res := &dto.ProfileResponse{
		UserID:         profile.UserID.Hex(),
		Bio:            profile.Bio,
		ProfilePicture: profile.ProfilePicture,
		ContactInfo:    profile.ContactInfo,
	}
	return res, nil

}
func (profileService *profileService) UpdateUserProfile(profile_dto *dto.UpdateProfileDto) (*dto.ProfileResponse, error) {
	userID, err := primitive.ObjectIDFromHex(profile_dto.UserID)
	profile := &entities.Profile{
		UserID:         userID,
		Bio:            profile_dto.Bio,
		ProfilePicture: profile_dto.ProfilePicture,
		ContactInfo: entities.ContactInfo{
			Address: profile_dto.Address,
		},
	}
	updated, err := profileService.Profile_repo.UpdateUserProfile(profile)
	if err != nil {
		return nil, err
	}

	res := &dto.ProfileResponse{
		UserID:         updated.UserID.Hex(),
		Bio:            updated.Bio,
		ProfilePicture: updated.ProfilePicture,
		ContactInfo:    updated.ContactInfo,
	}
	return res, nil
}
func (profileService *profileService) CreateUserProfile(profile_dto *dto.CreateProfileDto) (*dto.ProfileResponse, error) {
	userID, err := primitive.ObjectIDFromHex(profile_dto.UserID)
	if err != nil {
		return nil, err
	}

	profile := &entities.Profile{
		UserID:         userID,
		Bio:            profile_dto.Bio,
		ProfilePicture: profile_dto.ProfilePicture,
		ContactInfo: entities.ContactInfo{
			PhoneNumber: profile_dto.PhoneNumber,
			Email:       profile_dto.Email,
			Address:     profile_dto.Address,
		},
	}
	result, err := profileService.Profile_repo.CreateUserProfile(profile)
	if err != nil {
		return nil, err
	}
	response := &dto.ProfileResponse{
		UserID:         result.UserID.Hex(),
		Bio:            result.Bio,
		ProfilePicture: result.ProfilePicture,
		ContactInfo:    result.ContactInfo,
	}
	return response, nil
}
func (profileService *profileService) DeleteUserProfile(user_id string) error {
	err := profileService.Profile_repo.DeleteUserProfile(user_id)
	if err != nil {
		return err
	}
	return nil
}

func (profileService *profileService) UpdateProfilePicture(user_id string, file *multipart.FileHeader) (string, error) {
	url, _ := profileService.GetProfilePicture(user_id)
	if url != "" {
		err := profileService.Profile_repo.UpdateProfilePicture(user_id, url)
		if err != nil {
			return "", err
		}
	}
	url, err := profileService.imageService.UploadImage(file)
	if err != nil {
		return "", err
	}

	err = profileService.Profile_repo.UpdateProfilePicture(user_id, url)
	if err != nil {
		return "", err
	}
	
	return url, nil

}

func (profileService *profileService) GetProfilePicture(user_id string) (string, error) {
	url, err := profileService.Profile_repo.GetProfilePicture(user_id)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (profileService *profileService) DeleteProfilePicture(user_id string) error {
	url, err := profileService.Profile_repo.GetProfilePicture(user_id)
	if err != nil {
		return err
	}
	err = profileService.Profile_repo.DeleteProfilePicture(user_id)
	if err != nil {
		return err
	}
	err = profileService.imageService.DeleteImage(url)
	if err != nil {
		return err
	}
	return nil
}
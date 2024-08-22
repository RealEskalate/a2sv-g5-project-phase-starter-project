package service

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type profileService struct {
	Profile_repo interfaces.ProfileRepository
}

func NewProfileService(repo interfaces.ProfileRepository) interfaces.ProfileService {
	return &profileService{Profile_repo: repo}
}

func (profileService *profileService) GetUserProfile(userId string) (*entities.Profile, error) {
	profile, err := profileService.Profile_repo.GetUserProfile(userId)
	if err != nil {
		return nil, err
	}
	return profile, nil

}
func (profileService *profileService) UpdateUserProfile(profile_dto *dto.UpdateProfileDto) (*entities.Profile, error) {
	userID, err := primitive.ObjectIDFromHex(profile_dto.UserID)
	profile := &entities.Profile{
		UserID:         userID,
		Bio:            profile_dto.Bio,
		ProfilePicture: profile_dto.ProfilePicture,
		ContactInfo: entities.ContactInfo{
			Address:     profile_dto.Address,
		},
	}
	updated, err := profileService.Profile_repo.UpdateUserProfile(profile)
	if err != nil {
		return nil, err
	}
	return updated, nil
}
func (profileService *profileService) CreateUserProfile(profile_dto *dto.CreateProfileDto) (*entities.Profile, error) {
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
	return result, nil
}
func (profileService *profileService) DeleteUserProfile(user_id string) error {
	err := profileService.Profile_repo.DeleteUserProfile(user_id)
	if err != nil {
		return err
	}
	return nil
}

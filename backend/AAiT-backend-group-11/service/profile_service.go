package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
)

type profileService struct {
	Profile_repo interfaces.ProfileRepository
}

func NewProfileService(repo interfaces.ProfileRepository) interfaces.ProfileService{
	return &profileService{Profile_repo: repo}
}

func (profileService *profileService) GetUserProfile(userId string) (*entities.Profile, error){
	profile,err:=profileService.Profile_repo.GetUserProfile(userId)
	if err!=nil{
		return nil,err
	}
	return profile,nil

}
func (profileService *profileService) UpdateUserProfile(profile *entities.Profile) (*entities.Profile, error){
	updated,err:=profileService.UpdateUserProfile(profile)
	if err!=nil{
		return nil,err
	}
	return updated,nil
}
func (profileService *profileService) CreateUserProfile(profile *entities.Profile) (*entities.Profile, error){
	result,err:=profileService.CreateUserProfile(profile)
	if err!=nil{
		return nil,err
	}
	return result,nil
}
func (profileService *profileService) DeleteUserProfile(user_id string) error{
	err:=profileService.DeleteUserProfile(user_id)
	if err!=nil{
		return err
	}
	return nil
}


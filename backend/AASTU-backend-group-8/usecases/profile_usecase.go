package usecases

import (
	"fmt"
	"meleket/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileUsecase struct {
	profileRepo domain.ProfileRepository
}

// func NewTokenRepository(tr repository.UserRepositoryInterface, js domain.JWTService) *UserUsecase {
func NewProfileUsecase(tr domain.ProfileRepository) domain.ProfileUsecase {
	return &ProfileUsecase{
		profileRepo: tr,
	}
}

func (r *ProfileUsecase) SaveProfile(profile *domain.Profile) error {
	objid, err := primitive.ObjectIDFromHex(profile.UserID)
	fmt.Println(objid)
	if err != nil {
		return err
	}
	return r.profileRepo.SaveProfile(profile)
}

func (r *ProfileUsecase) FindProfile(userID string) (*domain.Profile, error) {
	return r.profileRepo.FindProfile(userID)
}

func (r *ProfileUsecase) DeleteProfile(userID string) error {
	return r.profileRepo.DeleteProfile(userID)
}

func (r *ProfileUsecase) UpdateProfile(profile *domain.Profile) error {
	objid, err := primitive.ObjectIDFromHex(profile.UserID)
	fmt.Println(objid)
	if err != nil {
		return err
	}
	return r.profileRepo.UpdateProfile(profile)
}

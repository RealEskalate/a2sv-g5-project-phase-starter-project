package usecase

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/infrastructure"
	"AAiT-backend-group-6/repository"
	"context"
	"errors"
	"time"
)

type profileUsecase struct {
	userRepository domain.UserRepository
	profileRepository repository.UserProfileRepository
	contextTimeout time.Duration
	imageService  infrastructure.ImageService
}

func NewProfileUsecase(userRepository domain.UserRepository,profileRepository repository.UserProfileRepository, timeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
		imageService: *infrastructure.NewImageService("cloudinary://247973594783441:SKqGqfJ96lwdjR8PXhhyaqL8_fM@dfqfhf3zf"),
		profileRepository: profileRepository,
	}
}

func (uu *profileUsecase) UpdateProfile(c context.Context, userId string, updateProfileDto *domain.UpdateProfileDto) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	// if err := infrastructure.ValidateStruct(uu.validator, updateProfileDto); err != nil {
	// 	return fmt.Errorf("validation error: %v", err.Error())
	// }

	// existingUser, err := uu.userRepository.FindByEmailOrUsername(c, updateProfileDto.UserProfile.Username)
	// if err == nil && existingUser != nil && existingUser.ID.Hex() != userId {
	// 	return errors.New("username already exists")
	// }

	user, err := uu.userRepository.GetByID(c, userId)
	if err != nil {
		return errors.New("user not found")
	}

	if user.ID.Hex() != userId {
		return errors.New("unatuhorized")
	}

	var imageUrl string

	if updateProfileDto.Avatar != nil {
		_, err := uu.imageService.SaveProfileImage(updateProfileDto)
		if err != nil {
			return errors.New("internal server error")
		}

		url, err := uu.imageService.UploadImage(c, updateProfileDto.Avatar)
		imageUrl = url
		if err != nil {
			return errors.New("internal server error")
		}
	}

	updateData := map[string]interface{}{
		"username": user.Username,
		"profile": domain.Profile{
			Bio: updateProfileDto.UserProfile.Bio,
			Url: imageUrl,
		},
	}

	if err := uu.profileRepository.Update(ctx, user.ID.Hex(), updateData); err != nil {
		return err
	}

	return nil
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.Profile{Name: user.Name, Username: user.Username, Email: user.Email}, nil
}

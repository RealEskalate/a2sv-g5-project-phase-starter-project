package usecase

import (
	"blog/domain"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &ProfileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *ProfileUsecase) UpdateProfile(c context.Context, profile *domain.Profile,userid primitive.ObjectID) (*domain.ProfileResponse,error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	user := &domain.User{
		ID: userid,
		First_Name: profile.First_Name,
		Last_Name: profile.Last_Name,
		Bio: profile.Bio,
		Profile_Picture: profile.Profile_Picture,
		Contact_Info: profile.Contact_Info,
	 }

	err := pu.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil,errors.New("failed to update profile")
	}
	return &domain.ProfileResponse{First_Name: user.First_Name,
		Last_Name: user.Last_Name,
		Bio : user.Bio,
		Profile_Picture: user.Profile_Picture,
		Contact_Info: user.Contact_Info,
		},nil
}


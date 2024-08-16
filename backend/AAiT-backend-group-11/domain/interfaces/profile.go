package interfaces

import  "backend-starter-project/domain/entities"

type ProfileRepository interface {
	GetUserProfile(userId string) (*entities.Profile, error)
	UpdateUserProfile(profile *entities.Profile) (*entities.Profile, error)
}

type ProfileService interface {
	GetUserProfile(userId string) (*entities.Profile, error)
	UpdateUserProfile(profile *entities.Profile) (*entities.Profile, error)
}

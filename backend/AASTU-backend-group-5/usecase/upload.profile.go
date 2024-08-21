package usecase

import (
	"time"
	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/repository"
)

type UploadProfileUsecase struct {
	Repo repository.UploadRepo
}

func NewUploadUsecase(repo repository.UploadRepo) *UploadProfileUsecase {
	return &UploadProfileUsecase{
		Repo: repo,
	}
}

func (uploaduc UploadProfileUsecase)UploadPicture(path string , id string) error {
	profile_picture := domain.Media{
		Path: path,
		Uplaoded_date: time.Now(),
	}
	err := uploaduc.Repo.AddProfile(profile_picture , id)
	return err
}
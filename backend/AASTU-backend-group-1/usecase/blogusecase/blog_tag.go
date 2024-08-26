package blogusecase

import (
	"blogs/domain"
	"errors"
)

func (b *BlogUsecase) GetTags() ([]*domain.Tag, error) {
	return b.TagRepo.GetTags()
}

func (b *BlogUsecase) InsertTag(tag *domain.Tag, claim *domain.LoginClaims) error {
	if claim.Role != "admin" {
		return errors.New("Only admin can insert tags")
		
	}
	return b.TagRepo.InsertTag(tag)
}

func (b *BlogUsecase) DeleteTag(tag *domain.Tag , claim *domain.LoginClaims) error {
	if claim.Role != "admin" {
		return errors.New("Only admin can insert tags")
		
	}

	return b.TagRepo.DeleteTag(tag)
}

func (b *BlogUsecase) CheckTag(tags []string) error {
	return b.TagRepo.CheckTag(tags)
}
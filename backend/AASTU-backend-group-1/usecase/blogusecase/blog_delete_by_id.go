package blogusecase

import (
	"blogs/domain"
	"errors"
)

// DeleteBlogByID implements domain.BlogUsecase.
func (b *BlogUsecase) DeleteBlogByID(id string, claim *domain.LoginClaims) error {

	blog, err := b.BlogRepo.GetBlogByID(id)
	if err != nil {
		return err
	}
	if blog.Author.Username != claim.Username && claim.Role != "admin" {
		return errors.New("you are not the author of this blog")
	}

	err = b.BlogRepo.DeleteBlogByID(id)
	if err != nil {
		return err
	}

	return nil
}

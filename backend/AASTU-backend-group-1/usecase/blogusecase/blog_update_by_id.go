package blogusecase

import (
	"blogs/domain"
	"errors"
)

// UpdateBlogByID implements domain.BlogUsecase.
func (b *BlogUsecase) UpdateBlogByID(id string, newblog *domain.Blog, claim *domain.LoginClaims) error {

	blog, err := b.BlogRepo.GetBlogByID(id)
	if err != nil {
		return err
	}
	if blog.Author.Username != claim.Username {
		return errors.New("you are not the author of this blog")
	}

	err = b.BlogRepo.UpdateBlogByID(id, newblog)
	if err != nil {
		return err
	}

	return nil
}
